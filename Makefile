# =========================
# PROJECT CONFIG
# =========================
APP_NAME=go-graphql
COMPOSE=docker compose
GO_REQUIRED=go1.24.0

# =========================
# DEFAULT
# =========================
.PHONY: help
help:
	@echo ""
	@echo "Available commands:"
	@echo ""
	@echo "  make dev            Run app in DEV mode (with override + hot reload)"
	@echo "  make dev-build      Build & run DEV containers"
	@echo "  make prod           Run app in PROD mode (no override)"
	@echo "  make prod-build     Build & run PROD containers"
	@echo ""
	@echo "  make test             Run all tests"
	@echo "  make test-unit        Run unit tests only"
	@echo "  make test-integration Run integration tests (Redis required)"
	@echo ""
	@echo "  make down           Stop all containers"
	@echo "  make logs           Tail logs"
	@echo "  make ps             Show running containers"
	@echo ""
	@echo "  make migrate        Run DB migrations"
	@echo "  make clean          Remove containers, volumes, images"
	@echo ""


# =========================
# GO VERSION GUARD
# =========================
.PHONY: check-go
check-go:
	@echo "🔍 Checking Go version..."
	@go version | grep -q "$(GO_REQUIRED)" || ( \
		echo "❌ Go version must be $(GO_REQUIRED)"; \
		go version; \
		exit 1 \
	)
	@echo "✅ Go version OK ($(GO_REQUIRED))"

# =========================
# DEV
# =========================
.PHONY: dev
dev:
	$(COMPOSE) up

.PHONY: dev-build
dev-build:
	$(COMPOSE) up --build

# =========================
# PROD
# =========================
.PHONY: prod
prod:
	$(COMPOSE) -f docker-compose.yml up -d

.PHONY: prod-build
prod-build:
	$(COMPOSE) -f docker-compose.yml up -d --build

# =========================
# TESTING
# =========================
.PHONY: test
test: check-go
	go test ./...

.PHONY: test-unit
test-unit: check-go
	go test ./... -run Test

.PHONY: test-integration
test-integration: check-go
	go test ./... -tags=integration

# =========================
# UTILS
# =========================
.PHONY: down
down:
	$(COMPOSE) down

.PHONY: logs
logs:
	$(COMPOSE) logs -f

.PHONY: ps
ps:
	$(COMPOSE) ps

# =========================
# DATABASE
# =========================
.PHONY: migrate
migrate:
	$(COMPOSE) exec gateway go run internal/infra/db/migrations/main.go

# =========================
# CLEANUP
# =========================
.PHONY: clean
clean:
	$(COMPOSE) down -v --rmi all --remove-orphans
