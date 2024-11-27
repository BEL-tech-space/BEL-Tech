format:
	@templ fmt .

generate:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css
	@pnpm sw
	@pnpm build
	@templ generate && templ fmt . 

dev:
	@pnpm sw
	@pnpm build
	@templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

tailwind:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --watch


.PHONY: generate dev tailwind
