.PHONY: build run test clean docker-up docker-down

# Build the go binary
build:
	@echo "Building bridge-app..."
	@go build -o bin/bridge-app ./cmd/bridge/main.go

# Run locally
run:
	@echo "Running bridge-app locally..."
	@go run ./cmd/bridge/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf bin/

# Start all docker containers (Postgres, RabbitMQ, Redis, Bridge-app)
docker-up:
	@echo "Starting docker containers..."
	@docker-compose up --build -d

# Stop all docker containers
docker-down:
	@echo "Stopping docker containers..."
	@docker-compose down

# Generate Smart Contract Bindings
abigen:
	@echo "Generating Go bindings for Smart Contract..."
	@go run github.com/ethereum/go-ethereum/cmd/abigen@latest --abi EnterpriseFiatToken.abi --bin EnterpriseFiatToken.bin --pkg contract --out internal/infrastructure/contract/fiat_token.go
