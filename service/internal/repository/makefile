
sqlc:
	@sqlc generate
up:
	@migrate -path postgres/migration -database "postgresql://postgres:spellit@localhost:5432/beltech?sslmode=disable" -verbose up
down:
	@migrate -path postgres/migration -database "postgresql://postgres:spellit@localhost:5432/beltech?sslmode=disable" -verbose down