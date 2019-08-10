package app

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, strings.NewReader(syumaiASCIIArt)); err != nil {
		log.Fatal(err)
	}
}
