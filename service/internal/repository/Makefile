
sqlc:
	@sqlc generate
up:
	@migrate -path postgres/migrations -database "" -verbose up
down:
	@migrate -path postgres/migrations -database "" -verbose down 