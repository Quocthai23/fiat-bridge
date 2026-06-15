package blockchain

import (
	"context"
	"math/big"

	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/contract"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/kms"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/nonce"

	"github.com/ethereum/go-ethereum/common"
	coreTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Emitter struct {
	client       *ethclient.Client
	signer       kms.Signer
	contractAddr common.Address
	nonceManager nonce.Manager
	instance     *contract.Contract
}

// NewEmitter creates a new Blockchain Emitter
func NewEmitter(client *ethclient.Client, signer kms.Signer, contractAddrHex string, nm nonce.Manager) (*Emitter, error) {
	contractAddr := common.HexToAddress(contractAddrHex)
	instance, err := contract.NewContract(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &Emitter{
		client:       client,
		signer:       signer,
		contractAddr: contractAddr,
		nonceManager: nm,
		instance:     instance,
	}, nil
}

// MintTokens signs a mint transaction but does not send it to the network (NoSend pattern)
func (e *Emitter) MintTokens(ctx context.Context, coreTxId string, userAddress common.Address, amount *big.Int) (*coreTypes.Transaction, uint64, error) {
	fromAddress := e.signer.Address()

	// Get and increment nonce atomically from Manager
	n, err := e.nonceManager.GetAndIncrementNonce(ctx, fromAddress)
	if err != nil {
		return nil, 0, err
	}

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		e.nonceManager.ReleaseNonce(fromAddress, n)
		return nil, 0, err
	}

	chainID, err := e.client.ChainID(ctx)
	if err != nil {
		e.nonceManager.ReleaseNonce(fromAddress, n)
		return nil, 0, err
	}

	auth, err := e.signer.GetTransactor(chainID)
	if err != nil {
		e.nonceManager.ReleaseNonce(fromAddress, n)
		return nil, 0, err
	}
	auth.Nonce = big.NewInt(int64(n))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.NoSend = true // IMPORTANT: Sign only, do not send

	tx, err := e.instance.Mint(auth, coreTxId, userAddress, amount)
	if err != nil {
		// Roll back nonce since transaction signing failed
		e.nonceManager.ReleaseNonce(fromAddress, n)
		return nil, 0, err
	}

	return tx, n, nil
}

// BurnTokens signs a burn transaction but does not send it to the network (NoSend pattern)
func (e *Emitter) BurnTokens(ctx context.Context, coreTxId string, userAddress common.Address, amount *big.Int) (*coreTypes.Transaction, uint64, error) {
	fromAddress := e.signer.Address()

	// Get and increment nonce atomically from Manager
	n, err := e.nonceManager.GetAndIncrementNonce(ctx, fromAddress)
	if err != nil {
		return nil, 0, err
	}

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		e.nonceManager.ReleaseNonce(fromAddress, n)
		return nil, 0, err
	}

	chainID, err := e.client.ChainID(ctx)
	if err != nil {
		e.nonceManager.ReleaseNonce(fromAddress, n)
		return nil, 0, err
	}

	auth, err := e.signer.GetTransactor(chainID)
	if err != nil {
		e.nonceManager.ReleaseNonce(fromAddress, n)
		return nil, 0, err
	}
	auth.Nonce = big.NewInt(int64(n))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.NoSend = true // IMPORTANT: Sign only, do not send

	tx, err := e.instance.Burn(auth, coreTxId, userAddress, amount)
	if err != nil {
		// Roll back nonce since transaction signing failed
		e.nonceManager.ReleaseNonce(fromAddress, n)
		return nil, 0, err
	}

	return tx, n, nil
}

// BroadcastTransaction sends a previously signed transaction to the network
func (e *Emitter) BroadcastTransaction(ctx context.Context, tx *coreTypes.Transaction) error {
	return e.client.SendTransaction(ctx, tx)
}

// ReleaseNonce exposes the nonce manager's ReleaseNonce using the signer's address
func (e *Emitter) ReleaseNonce(nonce uint64) {
	fromAddress := e.signer.Address()
	e.nonceManager.ReleaseNonce(fromAddress, nonce)
}
