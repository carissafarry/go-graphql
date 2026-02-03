# =========================
# PROJECT CONFIG
# =========================
APP_NAME=go-graphql
COMPOSE=docker compose

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
	@echo "  make down           Stop all containers"
	@echo "  make logs           Tail logs"
	@echo "  make ps             Show running containers"
	@echo ""
	@echo "  make migrate        Run DB migrations"
	@echo "  make clean          Remove containers, volumes, images"
	@echo ""

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
