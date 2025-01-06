run:
	@go run .

PHONY: run

templ:
	@templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

generate:
	@tailwindcss -i frontend/public/assets/styles/input.css  -o frontend/public/assets/styles/output.css --minify
	@templ generate && templ fmt .