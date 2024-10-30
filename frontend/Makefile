format:
	@templ fmt .

generate:
	@tailwindcss -i assets/styles/input.css  -o assets/styles/output.css
	@templ generate && templ fmt . 

dev:
	@go generate && templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v	@

tailwind:
	@tailwindcss -i assets/styles/input.css  -o assets/styles/output.css --watch


.PHONY: generate dev tailwind
