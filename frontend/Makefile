format:
	@templ fmt .

generate:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css
	@templ generate && templ fmt . 

dev:
	@templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

tailwind:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --watch


.PHONY: generate dev tailwind
