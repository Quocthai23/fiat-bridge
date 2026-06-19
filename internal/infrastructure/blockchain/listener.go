package blockchain

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/contract"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Listener struct {
	client       *ethclient.Client
	contractAddr common.Address
	instance     *contract.Contract
}

// NewListener creates a new Blockchain Listener
func NewListener(client *ethclient.Client, contractAddrHex string) (*Listener, error) {
	contractAddr := common.HexToAddress(contractAddrHex)
	instance, err := contract.NewContract(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &Listener{
		client:       client,
		contractAddr: contractAddr,
		instance:     instance,
	}, nil
}

// StartListening starts the listener for both FiatMinted and FiatBurned events with 12 blocks confirmation
func (l *Listener) StartListening(ctx context.Context, defaultStartBlock uint64) {
	log.Printf("Starting listener for contract events from default block %d with 12 blocks confirmation...\n", defaultStartBlock)

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				latestBlock, err := l.client.BlockNumber(ctx)
				if err != nil {
					log.Printf("Failed to get latest block number: %v\n", err)
					continue
				}

				if latestBlock < 12 {
					continue
				}

				safeBlock := latestBlock - 12

				// Fetch SyncState from DB
				var syncState domain.SyncState
				if err := db.DB.FirstOrCreate(&syncState, domain.SyncState{ID: 1}).Error; err != nil {
					log.Printf("Failed to fetch SyncState: %v\n", err)
					continue
				}

				mintStartBlock := syncState.LastMintBlock + 1
				if syncState.LastMintBlock == 0 {
					mintStartBlock = defaultStartBlock
				}

				burnStartBlock := syncState.LastBurnBlock + 1
				if syncState.LastBurnBlock == 0 {
					burnStartBlock = defaultStartBlock
				}

				chunkSize := uint64(2000)

				// 1. Process Mints
				currentMintStart := mintStartBlock
				for currentMintStart <= safeBlock {
					endBlock := currentMintStart + chunkSize
					if endBlock > safeBlock {
						endBlock = safeBlock
					}

					filterOpts := &bind.FilterOpts{
						Start:   currentMintStart,
						End:     &endBlock,
						Context: ctx,
					}

					mintIter, err := l.instance.FilterFiatMinted(filterOpts, nil)
					if err != nil {
						log.Printf("Failed to filter FiatMinted: %v\n", err)
						break // breaking loop for this tick, will retry next tick
					}

					for mintIter.Next() {
						event := mintIter.Event
						log.Printf("Confirmed FiatMinted Event Received: CoreTxID=%s, Amount=%v, To=%s\n", event.CoreTxId, event.Amount, event.To.Hex())

						// Find existing transaction and update status to Completed
						var dbTx domain.Transaction
						if err := db.DB.Where("core_tx_id = ?", event.CoreTxId).First(&dbTx).Error; err == nil {
							// Only update if not already completed
							if dbTx.Status != domain.StatusCompleted {
								dbTx.Status = domain.StatusCompleted
								dbTx.TxHash = event.Raw.TxHash.Hex()
								dbTx.BlockNumber = event.Raw.BlockNumber
								if err := db.UpdateTransaction(&dbTx); err != nil {
									log.Printf("Failed to update status to completed for TxID %s: %v\n", event.CoreTxId, err)
								} else {
									log.Printf("Successfully marked mint transaction %s as completed\n", event.CoreTxId)

									var fiatOrder domain.FiatOrder
									if err := db.DB.Where("core_tx_id = ?", event.CoreTxId).First(&fiatOrder).Error; err == nil {
										if fiatOrder.WebhookUrl != "" {
											payloadMap := map[string]interface{}{
												"core_tx_id":  event.CoreTxId,
												"tx_hash":     event.Raw.TxHash.Hex(),
												"status":      "SUCCESS",
												"amount":      event.Amount.String(),
												"webhook_url": fiatOrder.WebhookUrl,
											}
											payloadBytes, _ := json.Marshal(payloadMap)

											outboxMsg := domain.OutboxEvent{
												EventType: "WEBHOOK",
												Payload:   string(payloadBytes),
												Status:    "PENDING",
											}
											if err := db.DB.Create(&outboxMsg).Error; err != nil {
												log.Printf("Failed to create webhook outbox event for TxID %s: %v\n", event.CoreTxId, err)
											} else {
												log.Printf("Created WEBHOOK outbox event for TxID %s\n", event.CoreTxId)
											}
										}
									}
								}
							}
						} else {
							log.Printf("Mint transaction %s not found in DB or error fetching: %v\n", event.CoreTxId, err)
						}
					}
					if err := mintIter.Error(); err != nil {
						log.Printf("Error iterating mint events: %v\n", err)
						break
					}

					// Successfully processed this chunk, update SyncState
					syncState.LastMintBlock = endBlock
					if err := db.DB.Save(&syncState).Error; err != nil {
						log.Printf("Failed to save LastMintBlock: %v\n", err)
					}

					currentMintStart = endBlock + 1
				}

				// 2. Process Burns
				currentBurnStart := burnStartBlock
				for currentBurnStart <= safeBlock {
					endBlock := currentBurnStart + chunkSize
					if endBlock > safeBlock {
						endBlock = safeBlock
					}

					filterOpts := &bind.FilterOpts{
						Start:   currentBurnStart,
						End:     &endBlock,
						Context: ctx,
					}

					burnIter, err := l.instance.FilterFiatBurned(filterOpts, nil)
					if err != nil {
						log.Printf("Failed to filter FiatBurned: %v\n", err)
						break
					}

					for burnIter.Next() {
						event := burnIter.Event
						log.Printf("Confirmed FiatBurned Event Received: CoreTxID=%s, Amount=%v, From=%s\n", event.CoreTxId, event.Amount, event.From.Hex())

						var existingTx domain.Transaction
						if err := db.DB.Where("core_tx_id = ?", event.CoreTxId).First(&existingTx).Error; err == nil {
							if existingTx.Status != domain.StatusCompleted {
								existingTx.Status = domain.StatusCompleted
								existingTx.TxHash = event.Raw.TxHash.Hex()
								existingTx.BlockNumber = event.Raw.BlockNumber
								if err := db.UpdateTransaction(&existingTx); err != nil {
									log.Printf("Failed to update status to completed for Burn TxID %s: %v\n", event.CoreTxId, err)
								} else {
									log.Printf("Successfully marked burn transaction %s as completed\n", event.CoreTxId)
								}
							}
						} else {
							tx := &domain.Transaction{
								CoreTxId:    event.CoreTxId,
								UserAddress: event.From.Hex(),
								Amount:      event.Amount.String(), // Using String() since Amount is now string
								Type:        domain.TxTypeBurn,
								Status:      domain.StatusCompleted,
								TxHash:      event.Raw.TxHash.Hex(),
								BlockNumber: event.Raw.BlockNumber,
							}

							if err := db.SaveTransaction(tx); err != nil {
								log.Printf("Failed to save missing burn transaction %s: %v\n", event.CoreTxId, err)
							} else {
								log.Printf("Successfully processed and saved missing burn transaction: %s\n", event.CoreTxId)
							}
						}
					}
					if err := burnIter.Error(); err != nil {
						log.Printf("Error iterating burn events: %v\n", err)
						break
					}

					// Successfully processed this chunk, update SyncState
					syncState.LastBurnBlock = endBlock
					if err := db.DB.Save(&syncState).Error; err != nil {
						log.Printf("Failed to save LastBurnBlock: %v\n", err)
					}

					currentBurnStart = endBlock + 1
				}
			}
		}
	}()
}
