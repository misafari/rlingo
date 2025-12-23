.PHONY:migrate-up migrate-down migrate-create migrate-force

DB_HOST=localhost
DB_PORT=5455
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=rlingo
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Migration configuration
MIGRATIONS_PATH=./internal/repository/postgres/migrations
MIGRATE=migrate

migrate-create: ## Create new migration (usage: make migrate-create NAME=add_users_table)
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required"; \
		echo "Usage: make migrate-create NAME=your_migration_name"; \
		exit 1; \
	fi
	@echo "Creating migration: $(NAME)"
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_PATH) -seq $(NAME)
	@echo "$(COLOR_GREEN)✓ Migration files created in $(MIGRATIONS_PATH)/"

migrate-up: ## Run all up migrations
	@echo "Running migrations..."
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up
	@echo "$(COLOR_GREEN)✓ Migrations applied"

migrate-up-1: ## Run next up migration
	@echo "Running next migration..."
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up 1
	@echo "$(COLOR_GREEN)✓ Migration applied"

migrate-down: ## Rollback last migration
	@echo "⚠️  Rolling back last migration..."
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1
	@echo "$(COLOR_GREEN)✓ Migration rolled back"

migrate-down-all: ## Rollback all migrations
	@echo "⚠️  This will rollback ALL migrations!"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down -all; \
		echo "$(COLOR_GREEN)✓ All migrations rolled back"; \
	fi

migrate-force: ## Force migration version (usage: make migrate-force VERSION=1)
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required"; \
		echo "Usage: make migrate-force VERSION=version_number"; \
		exit 1; \
	fi
	@echo "⚠️  Forcing migration version to $(VERSION)..."
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(VERSION)
	@echo "$(COLOR_GREEN)✓ Migration version forced"

migrate-version: ## Show current migration version
	@echo "Current migration version:"
	@$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

migrate-status: migrate-version
