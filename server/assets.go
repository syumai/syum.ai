//go:build !js

package server

import (
	"io/fs"
	"net/http"
	"os"
)

var (
	assetsFS          fs.FS
	assetsHandlerBase http.Handler
)

func init() {
	assetsRoot, err := os.OpenRoot("./public")
	if err != nil {
		panic(err)
	}
	assetsFS = assetsRoot.FS()
	assetsHandlerBase = http.FileServerFS(assetsFS)
}

// assetsHandler serves static assets for dev environment.
func assetsHandler(w http.ResponseWriter, r *http.Request) {
	assetsHandlerBase.ServeHTTP(w, r)
}
