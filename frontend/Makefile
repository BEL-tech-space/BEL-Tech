format:
	@templ fmt .

scripts:
	@pnpm sw
	@pnpm build

templ:
	@templ generate --watch -watch-pattern .+\.templ$ --proxy="http://localhost:8090" --open-browser=false

generate:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --minify
	@templ generate && templ fmt . 

dev: scripts templ

tailwind:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --watch


.PHONY: format generate scripts templ dev tailwind
