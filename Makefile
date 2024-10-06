# Database connection string
DB_URL=postgres://rishikeshsuvarna@localhost/go-reddit?sslmode=disable
MIGRATIONS_DIR=backend/migrations

.PHONY: migrate-up migrate-down migrate-create migrate-version migrate-force

# Run all available migrations
migrate-up:
	@echo "Running migrations up..."
	migrate -source file://$(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Revert all migrations
migrate-down:
	@echo "Running migrations down..."
	migrate -source file://$(MIGRATIONS_DIR) -database "$(DB_URL)" down

# Create a new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $${name}

# Show current migration version
migrate-version:
	@echo "Current migration version:"
	migrate -source file://$(MIGRATIONS_DIR) -database "$(DB_URL)" version

# Force set a specific version
migrate-force:
	@read -p "Enter version to force: " version; \
	migrate -source file://$(MIGRATIONS_DIR) -database "$(DB_URL)" force $$version

# Run exactly one migration up
migrate-up-1:
	@echo "Running one migration up..."
	migrate -source file://$(MIGRATIONS_DIR) -database "$(DB_URL)" up 1

# Revert exactly one migration
migrate-down-1:
	@echo "Running one migration down..."
	migrate -source file://$(MIGRATIONS_DIR) -database "$(DB_URL)" down 1