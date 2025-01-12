run:
	@go run .

PHONY: run

templ:
	@templ generate --watch --cmd="go run ." --proxy="http://localhost:8090" --open-browser=false -v

generate:
	@templ generate && templ fmt .