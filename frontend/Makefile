format:
	@templ fmt .

generate:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css
	@templ generate && templ fmt . 

scripts:
	@pnpm sw
	@pnpm build

dev:
	# @pnpm sw
	@pnpm build
	@templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

tailwind:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --watch


.PHONY: generate dev tailwind
