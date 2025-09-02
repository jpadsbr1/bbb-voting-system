include .env

# Caminho para as migrations
MIGRATIONS_DIR=./internal/infrastructure/storage/migrations

# URL de conexão
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Comando base do migrate
MIGRATE_CMD=migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)"

## ---- Targets ----

.PHONY: migrate-up migrate-down migrate-force migrate-new migrate-drop migrate-version

# Aplica todas as migrations pendentes
migrate-up:
	$(MIGRATE_CMD) up

# Volta 1 migration
migrate-down:
	$(MIGRATE_CMD) down 1

# Força o schema para uma versão específica (corrigir erros)
# Exemplo: make migrate-force VERSION=1
migrate-force:
	$(MIGRATE_CMD) force $(VERSION)

# Cria nova migration (usar: make migrate-new NAME=add_users_table)
migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

# Dropa todo o schema (cuidado!)
migrate-drop:
	$(MIGRATE_CMD) drop -f

# Mostra a versão atual das migrations
migrate-version:
	$(MIGRATE_CMD) version