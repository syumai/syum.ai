//go:build !js

package server

import (
	"embed"
	"io/fs"
	"net/http"
)

var (
	//go:embed static
	assetsFS      embed.FS
	assetsHandler http.Handler
)

func init() {
	f, err := fs.Sub(assetsFS, "static")
	if err != nil {
		panic(err)
	}
	assetsHandler = http.FileServer(http.FS(f))
}
