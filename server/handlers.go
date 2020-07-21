package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/syumai/syumaigen"
)

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
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

func cachedImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=864000, public")
	w.Header().Set("Expires", time.Now().AddDate(0, 0, 10).Format(time.RFC1123))
	imageHandler(w, r)
}
