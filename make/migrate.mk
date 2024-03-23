.PHONY: migrate/up
migrate/up:
	migrate -path ./schema -database 'postgres://postgres:password@localhost:5433/storages?sslmode=disable' up

.PHONY: migrate/down
migrate/down:
	migrate -path ./schema -database 'postgres://postgres:password@localhost:5433/storages?sslmode=disable' down