format:
	@templ fmt .

scripts:
	@pnpm sw
	@pnpm notification

pnpm:
	@pnpm run dev

templ:
	@templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

generate:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --minify
	@templ generate && templ fmt . 

dev: scripts templ

tailwind:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --watch


.PHONY: format generate scripts templ dev tailwind
