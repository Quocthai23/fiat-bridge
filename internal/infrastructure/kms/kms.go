package kms

import (
	"context"
	"crypto/ecdsa"
	"encoding/asn1"
	"fmt"
	"math/big"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	coreTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Signer is an interface for signing transactions without exposing the private key directly
type Signer interface {
	Address() common.Address
	GetTransactor(chainID *big.Int) (*bind.TransactOpts, error)
}

// MockKMSSigner simulates a KMS backend using a local private key for testing purposes
type MockKMSSigner struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

// NewMockKMSSigner initializes the mock signer
func NewMockKMSSigner(privKeyHex string) (*MockKMSSigner, error) {
	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &MockKMSSigner{
		privateKey: privateKey,
		address:    address,
	}, nil
}

func (s *MockKMSSigner) Address() common.Address {
	return s.address
}

// GetTransactor returns a TransactOpts that uses the local private key (mock)
func (s *MockKMSSigner) GetTransactor(chainID *big.Int) (*bind.TransactOpts, error) {
	keyAddr := s.address
	signer := coreTypes.LatestSignerForChainID(chainID)

	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *coreTypes.Transaction) (*coreTypes.Transaction, error) {
			if address != keyAddr {
				return nil, bind.ErrNotAuthorized
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), s.privateKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}, nil
}

// AWSKMSSigner implements the Signer interface using AWS KMS
type AWSKMSSigner struct {
	client  *kms.Client
	keyID   string
	address common.Address
}

// NewAWSKMSSigner creates a new AWS KMS signer
func NewAWSKMSSigner(ctx context.Context, cfg aws.Config, keyID string) (*AWSKMSSigner, error) {
	client := kms.NewFromConfig(cfg)

	pubKeyOutput, err := client.GetPublicKey(ctx, &kms.GetPublicKeyInput{
		KeyId: aws.String(keyID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get public key from KMS: %w", err)
	}

	// AWS KMS returns SPKI format for secp256k1
	var spki struct {
		Algorithm        asn1.RawValue
		SubjectPublicKey asn1.BitString
	}
	_, err = asn1.Unmarshal(pubKeyOutput.PublicKey, &spki)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal SPKI from KMS: %w", err)
	}

	pubKey, err := crypto.UnmarshalPubkey(spki.SubjectPublicKey.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse secp256k1 public key: %w", err)
	}

	address := crypto.PubkeyToAddress(*pubKey)

	return &AWSKMSSigner{
		client:  client,
		keyID:   keyID,
		address: address,
	}, nil
}

func (s *AWSKMSSigner) Address() common.Address {
	return s.address
}

// GetTransactor returns a TransactOpts that uses AWS KMS to sign
func (s *AWSKMSSigner) GetTransactor(chainID *big.Int) (*bind.TransactOpts, error) {
	keyAddr := s.address
	signer := coreTypes.LatestSignerForChainID(chainID)

	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *coreTypes.Transaction) (*coreTypes.Transaction, error) {
			if address != keyAddr {
				return nil, bind.ErrNotAuthorized
			}

			// 1. Get the hash of the transaction to sign
			txHash := signer.Hash(tx).Bytes()

			// 2. Send the hash to AWS KMS for signing
			signOutput, err := s.client.Sign(context.Background(), &kms.SignInput{
				KeyId:            aws.String(s.keyID),
				Message:          txHash,
				MessageType:      types.MessageTypeDigest,
				SigningAlgorithm: types.SigningAlgorithmSpecEcdsaSha256,
			})
			if err != nil {
				return nil, fmt.Errorf("KMS signing failed: %w", err)
			}

			// 3. Decode ASN.1 DER signature from KMS
			var sig struct {
				R *big.Int
				S *big.Int
			}
			_, err = asn1.Unmarshal(signOutput.Signature, &sig)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal KMS signature: %w", err)
			}

			// EIP-2 / SECP256K1 standards: S must be in the lower half of the curve to prevent malleability
			secp256k1N := crypto.S256().Params().N
			halfN := new(big.Int).Div(secp256k1N, big.NewInt(2))
			if sig.S.Cmp(halfN) > 0 {
				sig.S.Sub(secp256k1N, sig.S)
			}

			// AWS KMS doesn't return the recovery ID (V). We must calculate it.
			rBytes := sig.R.Bytes()
			sBytes := sig.S.Bytes()

			// Pad R and S to 32 bytes
			rPadded := common.LeftPadBytes(rBytes, 32)
			sPadded := common.LeftPadBytes(sBytes, 32)

			var finalSig []byte
			var found bool

			for v := 0; v < 2; v++ {
				candidateSig := append(rPadded, sPadded...)
				candidateSig = append(candidateSig, byte(v))

				recoveredPub, err := crypto.SigToPub(txHash, candidateSig)
				if err != nil {
					continue
				}

				recoveredAddr := crypto.PubkeyToAddress(*recoveredPub)
				if recoveredAddr == s.address {
					finalSig = candidateSig
					found = true
					break
				}
			}

			if !found {
				return nil, fmt.Errorf("failed to calculate recovery ID (V)")
			}

			// 4. Attach signature to transaction
			return tx.WithSignature(signer, finalSig)
		},
	}, nil
}
