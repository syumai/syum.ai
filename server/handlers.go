package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/a-h/templ"
	"github.com/syumai/syum.ai/server/pages/indexpage"
	"github.com/syumai/syumaigen"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ascii", asciiHandler)
	mux.HandleFunc("/image", nocacheImageHandler)
	mux.HandleFunc("/image/random", randomImageHandler)
	mux.HandleFunc("/favicon.ico", cachedImageHandler)
	mux.Handle("/", templ.Handler(indexpage.Index()))
	return mux
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, strings.NewReader(SyumaiASCIIArt)); err != nil {
		log.Fatal(err)
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	var cMap syumaigen.ColorMap
	code := r.URL.Query().Get("code")
	if code != "" {
		cMap = syumaigen.GenerateColorMapByColorCode(code)
	} else {
		cMap = syumaigen.DefaultColorMap
	}
	imgType := r.URL.Query().Get("type")
	if imgType == "svg" {
		writeSVG(w, cMap)
		return
	}
	writePNG(w, cMap)
}

func randomImageHandler(w http.ResponseWriter, r *http.Request) {
	v := url.Values{}
	v.Set("code", generateRandomColorCode())
	imgType := r.URL.Query().Get("type")
	if imgType != "" {
		v.Set("type", imgType)
	}
	http.Redirect(w, r, fmt.Sprintf("/image?%s", v.Encode()), http.StatusTemporaryRedirect)
}

func nocacheImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store,max-age=0")
	imageHandler(w, r)
}

func cachedImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=864000")
	imageHandler(w, r)
}
