package http

import (
	frontend "beltech/frontend"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

func (a *App) sw(w http.ResponseWriter, r *http.Request) {

	filePath := strings.TrimPrefix(r.URL.Path, "/public/")
	staticFiles, err := fs.Sub(frontend.AssetsDir, "public")
	if err != nil {
		a.serverError(w, fmt.Errorf("server public dir inaccessible: %w", err))
		return
	}

	file, err := staticFiles.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		a.serverError(w, fmt.Errorf("cannot open sw.js: %w", err))
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		a.serverError(w, fmt.Errorf("failed to read sw.js: %w", err))
		return
	}

	info, err := file.Stat()
	if err != nil {
		a.serverError(w, fmt.Errorf("failed to get sw.js file information: %w", err))
		return
	}

	w.Header().Set("Service-Worker-Allowed", "/")
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeContent(w, r, filePath, info.ModTime(), bytes.NewReader(content))
}
