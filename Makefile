# go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
DB_URL=postgres://tms:tms@localhost:5432/tms?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-force:
	migrate -path migrations -database "$(DB_URL)" force

migrate-new:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name
