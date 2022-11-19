package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ascii", asciiHandler)
	r.HandleFunc("/host", hostHandler)
	r.HandleFunc("/image", imageHandler)
	r.HandleFunc("/image/random", randomImageHandler)
	r.HandleFunc("/favicon.ico", cachedImageHandler)
	r.PathPrefix("/").HandlerFunc(assetsHandler)
	return r
}
