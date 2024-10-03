package frontend

import "embed"

//go:generate tailwindcss -i assets/styles/input.css -o assets/styles/output.css
//go:generate templ generate

//go:embed assets
var AssetsDir embed.FS
