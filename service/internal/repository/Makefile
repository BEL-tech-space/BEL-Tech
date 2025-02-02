db := ""
sqlc:
	@sqlc generate

migrate:
	@migrate create -ext=sql -dir=postgres/migrations -seq init

up:
	@migrate -path postgres/migrations -database ${db} -verbose up
	
down:
	@migrate -path postgres/migrations -database ${db} -verbose down