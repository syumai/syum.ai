package server

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/syumai/syumaigen"
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

//go:embed static
var assetsFS embed.FS

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, strings.NewReader(SyumaiASCIIArt)); err != nil {
		log.Fatal(err)
	}
}

func hostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, strings.NewReader(r.Host)); err != nil {
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
	w.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate,max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", time.Now().Format(time.RFC1123))
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

func cachedImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=864000, public")
	w.Header().Set("Expires", time.Now().AddDate(0, 0, 10).Format(time.RFC1123))
	imageHandler(w, r)
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=3600, public")
	w.Header().Set("Expires", time.Now().Add(time.Hour).Format(time.RFC1123))
	path := "static" + strings.TrimSuffix(r.URL.Path, "/")
	f := openAssetsFile(w, path)
	if f == nil {
		return
	}
	stat, _ := f.Stat() // this always succeeds
	if stat.IsDir() {
		path = filepath.Join(path, "index.html")
		f = openAssetsFile(w, path)
		if f == nil {
			return
		}
	}
	ext := filepath.Ext(path)
	if ext != "" {
		w.Header().Set("Content-Type", mime.TypeByExtension(ext))
	}
	if _, err := io.Copy(w, f); err != nil {
		log.Fatalf("assets file copy: %v", err)
	}
}

func openAssetsFile(w http.ResponseWriter, path string) fs.File {
	f, err := assetsFS.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
			return nil
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatalf("assets handler fs open: %v", err)
	}
	return f
}
