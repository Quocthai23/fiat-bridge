package blockchain

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"

	"github.com/ethereum/go-ethereum/common"
	coreTypes "github.com/ethereum/go-ethereum/core/types"
)

type GasBumper struct {
	emitter *Emitter
}

// NewGasBumper creates a new Gas Bumper Worker
func NewGasBumper(emitter *Emitter) *GasBumper {
	return &GasBumper{
		emitter: emitter,
	}
}

// Start starts the gas bumper cronjob
func (gb *GasBumper) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				gb.bumpGasForStuckTransactions()
			}
		}
	}()

	log.Printf("Gas Bumper Worker started. Interval: %v\n", interval)
}

func (gb *GasBumper) bumpGasForStuckTransactions() {
	log.Println("[GasBumper] Scanning for stuck transactions...")

	var stuckTxs []domain.Transaction
	fiveMinsAgo := time.Now().Add(-5 * time.Minute)

	// Find transactions that have been pending for more than 5 minutes
	if err := db.DB.Where("status = ? AND created_at < ?", domain.StatusPendingOnChain, fiveMinsAgo).Find(&stuckTxs).Error; err != nil {
		log.Printf("[GasBumper] Error fetching stuck txs: %v\n", err)
		return
	}

	for _, tx := range stuckTxs {
		log.Printf("[GasBumper] Bumping gas for CoreTxId: %s (Old Hash: %s)\n", tx.CoreTxId, tx.TxHash)

		ctx := context.Background()
		userAddr := common.HexToAddress(tx.UserAddress)
		amount, ok := new(big.Int).SetString(tx.Amount, 10)
		if !ok {
			log.Printf("[GasBumper] Invalid amount string for tx %s: %v\n", tx.CoreTxId, tx.Amount)
			continue
		}

		// Retrieve the old transaction from blockchain to check status and old gas price
		oldTx, isPending, err := gb.emitter.client.TransactionByHash(ctx, common.HexToHash(tx.TxHash))
		if err == nil && !isPending {
			// Transaction was already mined, no need to bump. Just update status.
			log.Printf("[GasBumper] Transaction %s was already mined. Updating status to completed.\n", tx.CoreTxId)
			tx.Status = domain.StatusCompleted
			db.UpdateTransaction(&tx)
			continue
		}

		var oldFeeCap, oldTipCap *big.Int
		if err == nil && oldTx != nil {
			oldFeeCap = oldTx.GasFeeCap()
			oldTipCap = oldTx.GasTipCap()
		} else {
			log.Printf("[GasBumper] Warning: Could not fetch old transaction %s from pool: %v. Using default fallback.\n", tx.TxHash, err)
			oldFeeCap = big.NewInt(0)
			oldTipCap = big.NewInt(0)
		}

		// Calculate the minimum required gas price for RBF (must be at least 10% higher than old tx)
		// We'll calculate oldCap * 1.15 for safety
		minBumpedFeeCap := new(big.Int).Mul(oldFeeCap, big.NewInt(115))
		minBumpedFeeCap.Div(minBumpedFeeCap, big.NewInt(100))
		
		minBumpedTipCap := new(big.Int).Mul(oldTipCap, big.NewInt(115))
		minBumpedTipCap.Div(minBumpedTipCap, big.NewInt(100))

		// Fetch current suggested tip cap from the network
		suggestedTipCap, err := gb.emitter.client.SuggestGasTipCap(ctx)
		if err != nil {
			log.Printf("[GasBumper] Failed to get suggested tip cap: %v\n", err)
			continue
		}

		// Get current base fee from latest block header
		header, err := gb.emitter.client.HeaderByNumber(ctx, nil)
		if err != nil || header.BaseFee == nil {
			log.Printf("[GasBumper] Failed to get base fee: %v\n", err)
			continue
		}

		// suggestedFeeCap = (baseFee * 2) + suggestedTipCap
		baseFeeX2 := new(big.Int).Mul(header.BaseFee, big.NewInt(2))
		suggestedFeeCap := new(big.Int).Add(baseFeeX2, suggestedTipCap)

		// Choose the higher price between minBumped and suggested
		bumpedFeeCap := suggestedFeeCap
		if minBumpedFeeCap.Cmp(suggestedFeeCap) > 0 {
			bumpedFeeCap = minBumpedFeeCap
			log.Printf("[GasBumper] Using RBF minimum bumped FeeCap: %s\n", bumpedFeeCap.String())
		}
		
		bumpedTipCap := suggestedTipCap
		if minBumpedTipCap.Cmp(suggestedTipCap) > 0 {
			bumpedTipCap = minBumpedTipCap
			log.Printf("[GasBumper] Using RBF minimum bumped TipCap: %s\n", bumpedTipCap.String())
		}

		chainID, err := gb.emitter.client.ChainID(ctx)
		if err != nil {
			log.Printf("[GasBumper] Failed to get chain ID: %v\n", err)
			continue
		}

		auth, err := gb.emitter.signer.GetTransactor(chainID)
		if err != nil {
			log.Printf("[GasBumper] Failed to get transactor: %v\n", err)
			continue
		}

		// USE ORIGINAL NONCE for Replace-By-Fee
		auth.Nonce = big.NewInt(int64(tx.Nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasFeeCap = bumpedFeeCap
		auth.GasTipCap = bumpedTipCap
		auth.NoSend = true // Only sign, do not broadcast synchronously

		var newTx *coreTypes.Transaction
		var errMintBurn error

		if tx.Type == domain.TxTypeBurn {
			newTx, errMintBurn = gb.emitter.instance.Burn(auth, tx.CoreTxId, userAddr, amount)
		} else {
			newTx, errMintBurn = gb.emitter.instance.Mint(auth, tx.CoreTxId, userAddr, amount)
		}

		if errMintBurn != nil {
			log.Printf("[GasBumper] Failed to sign bumped tx %s: %v\n", tx.CoreTxId, errMintBurn)
			continue
		}

		if err := gb.emitter.BroadcastTransaction(ctx, newTx); err != nil {
			log.Printf("[GasBumper] Failed to broadcast bumped tx %s: %v\n", tx.CoreTxId, err)
			continue
		}

		newHash := newTx.Hash().Hex()
		log.Printf("[GasBumper] Successfully bumped tx %s. New Hash: %s\n", tx.CoreTxId, newHash)

		// Update DB only if it's still pending
		tx.TxHash = newHash
		tx.CreatedAt = time.Now() // Reset timer
		
		result := db.DB.Model(&tx).Where("core_tx_id = ? AND status = ?", tx.CoreTxId, domain.StatusPendingOnChain).Updates(map[string]interface{}{
			"tx_hash":    newHash,
			"created_at": tx.CreatedAt,
		})

		if result.Error != nil {
			log.Printf("[GasBumper] Failed to update DB for tx %s: %v\n", tx.CoreTxId, result.Error)
		} else if result.RowsAffected == 0 {
			log.Printf("[GasBumper] Tx %s was already completed by Listener. Ignored DB update.\n", tx.CoreTxId)
		}
	}
}
