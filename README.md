# Fiat Bridge Enterprise 🌉

![Docker Pulls](https://img.shields.io/docker/pulls/quocthai23/fiat-bridge?style=flat-square)
![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go)
![Architecture](https://img.shields.io/badge/Architecture-Event%20Driven-success?style=flat-square)

Fiat Bridge is a robust, production-ready Multi-tenant White-label middleware that bridges traditional Web2 Fiat transactions (VND) with Web3 On-chain smart contracts. 

It acts as a **Banking-as-a-Service (BaaS)**, allowing DApps to seamlessly integrate fiat On-ramp and Off-ramp features without directly interacting with banking infrastructure.

---

## 🚀 Quick Start (Production via Docker Hub)

The easiest way to run Fiat Bridge is using our pre-built Docker image. You **do not** need to clone this repository or install Go.

**1. Create a `docker-compose.yml` file:**
```yaml
version: '3.8'
services:
  app:
    image: quocthai23/fiat-bridge:v1.0  # Or :latest
    container_name: fiat-bridge-app
    ports:
      - "8080:8080"
    env_file:
      - .env
    depends_on:
      - postgres
      - redis
      - rabbitmq
    restart: always

  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: secretpassword
      POSTGRES_DB: bridge
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: always

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    restart: always

  rabbitmq:
    image: rabbitmq:3-management-alpine
    ports:
      - "5672:5672"
      - "15672:15672"
    volumes:
      - rabbitmq_data:/var/lib/rabbitmq
    restart: always

volumes:
  postgres_data:
  redis_data:
  rabbitmq_data:
```

**2. Create a `.env` file in the same directory:**
```env
DATABASE_URL=postgres://root:secretpassword@postgres:5432/bridge?sslmode=disable
REDIS_ADDR=redis:6379
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
CONTRACT_ADDRESS=0xYourSmartContractAddress
PRIVATE_KEY_HEX=YourAdminPrivateKey
```

**3. Start the system:**
```bash
docker-compose up -d
```
Your gateway is now running at `http://localhost:8080`.

## 🛠 Local Development (From Source)
If you want to contribute to the codebase or build it locally:

**Requirements:**
- Go 1.20+
- Docker & docker-compose

```bash
# 1. Clone the repository
git clone https://github.com/Quocthai23/fiat-bridge.git
cd fiat-bridge

# 2. Start the infrastructure (DB, Redis, RabbitMQ)
docker-compose -f docker-compose.infra.yml up -d

# 3. Run the Go application
make run
```

## ✨ Core Features
- **Zero-Knowledge PII Storage:** User's sensitive bank account details are never stored permanently; utilizing short-lived Redis TTL cache for utmost privacy.
- **Race-Condition Proof:** Leverages PostgreSQL `FOR UPDATE` row-level locking to prevent double-minting during network spikes.
- **Dynamic VietQR On-ramp:** Auto-generates QR links embedded with `core_tx_id` for automated banking reconciliation.
- **Smart Gas Bumper:** Uses EIP-1559 Replace-By-Fee (RBF) to automatically untangle stuck transactions in the EVM mempool.
- **Webhook Outbox Pattern:** Ensures guaranteed delivery of Webhooks to DApps with HMAC-SHA256 signatures and exponential backoff retry.
- **API Rate Limiting:** Built-in Redis pipeline-based Token Bucket limiter per API Key to prevent DDoS.

## 📊 Dashboard & Monitoring
When running via Docker, monitoring tools are accessible at:
- **RabbitMQ Dashboard:** `http://localhost:15672` (guest/guest)
- **(Optional) Grafana:** `http://localhost:3000` (admin/secret)
- **(Optional) Prometheus:** `http://localhost:9090`
