package reconciliation

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/contract"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Engine struct {
	client       *ethclient.Client
	contractAddr common.Address
	instance     *contract.Contract
}

// NewEngine creates a new Reconciliation Engine
func NewEngine(client *ethclient.Client, contractAddrHex string) (*Engine, error) {
	contractAddr := common.HexToAddress(contractAddrHex)
	instance, err := contract.NewContract(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &Engine{
		client:       client,
		contractAddr: contractAddr,
		instance:     instance,
	}, nil
}

// Start starts the reconciliation cronjob
func (e *Engine) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				e.runReconciliation()
			}
		}
	}()

	log.Printf("Reconciliation Engine started. Interval: %v\n", interval)
}

func (e *Engine) runReconciliation() {
	log.Println("[Reconciliation] Running reconciliation check...")

	header, err := e.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Printf("[Reconciliation Error] Failed to get latest block header: %v\n", err)
		return
	}
	blockNum := header.Number.Uint64()
	log.Printf("[Reconciliation] Syncing at Block Number: %d\n", blockNum)
	totalSupply, err := e.instance.TotalSupply(&bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(blockNum),
	})
	if err != nil {
		log.Printf("[Reconciliation Error] Failed to get totalSupply: %v\n", err)
		return
	}

	var totalMintedStr, totalBurnedStr string

	db.DB.Model(&domain.Transaction{}).
		Select("COALESCE(SUM(CAST(amount AS NUMERIC)), 0)").
		Where("type = ? AND status = ? AND block_number <= ?", domain.TxTypeMint, domain.StatusCompleted, blockNum).
		Scan(&totalMintedStr)

	totalMinted, ok := new(big.Int).SetString(totalMintedStr, 10)
	if !ok {
		totalMinted = big.NewInt(0)
	}

	db.DB.Model(&domain.Transaction{}).
		Select("COALESCE(SUM(CAST(amount AS NUMERIC)), 0)").
		Where("type = ? AND status = ? AND block_number <= ?", domain.TxTypeBurn, domain.StatusCompleted, blockNum).
		Scan(&totalBurnedStr)

	totalBurned, ok := new(big.Int).SetString(totalBurnedStr, 10)
	if !ok {
		totalBurned = big.NewInt(0)
	}

	// dbTotal = totalMinted - totalBurned
	dbTotal := new(big.Int).Sub(totalMinted, totalBurned)

	log.Printf("[Reconciliation] Blockchain TotalSupply: %s, DB TotalMinted: %s, DB TotalBurned: %s, DB Calculated Supply: %s\n", 
		totalSupply.String(), totalMinted.String(), totalBurned.String(), dbTotal.String())

	if totalSupply.Cmp(dbTotal) != 0 {
		log.Printf("[Reconciliation WARNING] Mismatch detected! SC: %s vs DB: %s\n", totalSupply.String(), dbTotal.String())
	} else {
		log.Println("[Reconciliation] Ledger matches SC perfectly.")
	}
}
