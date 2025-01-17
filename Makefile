run:
	@go run .

PHONY: run

templ:
	@templ generate --watch --cmd="go run ." --proxy="http://localhost:8090" --open-browser=false -v

generate:
	@templ generate && templ fmt .

dockerHost:
	# this allows you to bypass docker network isolation and access host machine. note image has been named "beltech"
	@docker run --rm --network host -e DB_PORT=5432 -e DB_HOST=127.0.0.1 --env-file dev.env beltech

docker:
	# this allows you to access host machine from docker container. note image has been named "beltech"
	# host.docker.internal is a special DNS name which resolves to the internal IP address used by the host.
	# this has to be discovered by running `docker run --rm alpine nslookup host.docker.internal` or
	# `docker run --rm alpine ping host.docker.internal`. This value then has to be added to the pg.hba.conf file 
	# e.g host all all ***/32 trust. Also, the postgressql.conf file needs to change `listen_address='*'`
	@docker run --rm -p 127.0.0.1:8090:8090 -e DB_PORT=5432 -e DB_HOST=host.docker.internal --env-file dev.env beltech