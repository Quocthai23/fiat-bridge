# Fiat Bridge Enterprise

Fiat Bridge is a robust, production-ready Multi-tenant White-label middleware for bridging traditional Web2 Fiat transactions (VND) with Web3 On-chain smart contracts.

This project enables DApps to seamlessly integrate fiat on-ramp (deposits) and off-ramp (withdrawals) features without directly interacting with banking infrastructure.

## Core Features
* **Multi-tenant White-label**: Each DApp registers its own bank details, allowing Fiat money to flow directly into their bank accounts.
* **Dynamic VietQR On-ramp**: Automatically generates `vietqr.io` payment links customized for each DApp with specific `core_tx_id` identifiers for automatic reconciliation.
* **Web2 to Web3 Off-ramp Mapping**: Securely handles Off-ramp withdrawals via `PayoutOrder` without exposing sensitive user bank details to the public blockchain.
* **Callback Security (Webhook Dispatcher)**: Reliable Outbox-pattern webhook delivery with HMAC-SHA256 signatures and Exponential Backoff retry mechanics (5s -> 1m -> 5m -> 1h).
* **API Rate Limiting**: Built-in Redis Token Bucket / Sliding Window rate limiter mapped to API keys.
* **Blockchain Resiliency**:
  - `GasBumper`: Cron job that scans the mempool for stuck transactions and auto-replaces them using Replace-By-Fee (RBF).
  - `Listener`: 12-block confirmation listener for `FiatMinted` and `FiatBurned` events.
  - `Reconciliation Engine`: Hourly cron job checking for mismatches between Database and Blockchain state.

---

## High-Level Architecture

### On-ramp (Fiat Deposit -> Token Mint)
1. **DApp Request**: User wants to deposit 100k VND. DApp calls `POST /api/v1/fiat/orders` with DApp API Key.
2. **VietQR Gen**: Bridge checks `DappConfig` for that API Key, gets the DApp's Vietcombank/Techcombank details, and generates a dynamic VietQR image link. Returns `core_tx_id` and `qr_url`.
3. **User Pays**: User scans QR and transfers VND.
4. **Bank Webhook**: Bank API (e.g. SePay/Casso) pushes webhook to `POST /api/v1/webhooks/bank`.
5. **Reconciliation**: Bridge extracts `core_tx_id` from transfer description, marks order as `PAID`.
6. **Mint Queue**: Bridge pushes `MINT` event to RabbitMQ.
7. **Worker & Blockchain**: Worker consumes from RabbitMQ, calls `Mint()` on Smart Contract using KMS / Private Key signing.
8. **DApp Callback**: Once mined, `Outbox Relay` sends HMAC-SHA256 signed webhook back to DApp.

### Off-ramp (Token Burn -> Fiat Payout)
1. **DApp Request (Web2)**: User enters their bank account info on DApp. DApp calls `POST /api/v1/fiat/payout-orders` with `user_address`, `bank_account`, `bank_bin`.
2. **Secure Mapping**: Bridge saves a `PayoutOrder` to DB (Status: `WAITING_FOR_BURN`) and returns `core_tx_id`. **PII Compliance:** The user's sensitive `bank_account` and `bank_bin` are NEVER stored in the database. Instead, they are temporarily cached in Redis with a 24-hour TTL, ensuring full data privacy.
3. **On-chain Burn**: User signs Web3 transaction calling `burn(core_tx_id, amount)` on Smart Contract.
4. **Listener**: Bridge Listener detects `FiatBurned` event after 12 block confirmations.
5. **Payout Dispatch**: Listener inserts `PAYOUT` event into Outbox -> RabbitMQ `payout_queue`.
6. **Payout Worker**: Reads `core_tx_id`, fetches the user's bank details from the ephemeral Redis cache, triggers internal `core-banking` API to transfer VND, and then immediately purges the cache.

---

## Setup & Running

### Requirements
* Docker & docker-compose
* Go 1.20+

### Run with Docker
Start PostgreSQL, RabbitMQ, Redis, Prometheus, Grafana, and the Bridge App:
```bash
make docker-up
```

### Run Locally (Development)
```bash
# Requires .env variables (DATABASE_URL, REDIS_ADDR, RABBITMQ_URL)
make build
make run
```

## Dashboard & Monitoring
* **RabbitMQ**: `http://localhost:15673` (guest/guest)
* **Grafana**: `http://localhost:3000` (admin/secret)
* **Prometheus**: `http://localhost:9090`
