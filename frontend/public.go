package frontend

import "embed"

//go:generate tailwindcss -i public/assets/styles/input.css -o public/assets/styles/output.css
//go:generate templ generate

//go:embed public
var AssetsDir embed.FS
