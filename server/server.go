package server

import (
	"fmt"
	"net/http"
	"time"
)

func NewServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ascii", asciiHandler)
	mux.HandleFunc("/image", imageHandler)
	mux.HandleFunc("/favicon.ico", cachedImageHandler)
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
