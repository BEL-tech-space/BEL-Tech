templ:
	@templ generate && templ fmt . 

dev: live

live:
	@go generate && templ fmt . && templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

tailwind:
	@tailwindcss -i assets/styles/input.css  -o assets/styles/output.css --watch

.PHONY: dev live tailwind
