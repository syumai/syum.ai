package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewServer(port string) *http.Server {
	r := mux.NewRouter()
	setupRedirectHandlers(r)
	mainRouter := r.Host("syum.ai").Subrouter()
	mainRouter.HandleFunc("/", indexHandler)
	mainRouter.HandleFunc("/ascii", asciiHandler)
	mainRouter.HandleFunc("/image", imageHandler)
	mainRouter.HandleFunc("/image/random", randomImageHandler)
	mainRouter.HandleFunc("/favicon.ico", cachedImageHandler)
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func setupRedirectHandlers(r *mux.Router) {
	redirectMap := map[string]string{
		"tw.syum.ai": "https://twitter.com/__syumai",
		"gh.syum.ai": "https://github.com/syumai",
	}
	for host, url := range redirectMap {
		r.Handle("/", http.RedirectHandler(url, http.StatusMovedPermanently)).
			Host(host)
	}
}
