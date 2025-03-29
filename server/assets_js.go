//go:build js

package server

import (
	"net/http"
)

// assetsHandler redirects to indexHandler.
// - Static assets are served using Cloudflare Workers' Static Assets feature.
// - TinyGo's mux doesn't appear to support {$} as a path pattern.
func assetsHandler(w http.ResponseWriter, r *http.Request) {
	indexHandler(w, r)
}
