package server

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"log"
	"net/http"
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
	w.Header().Set("Content-Type", "image/png")
	var cMap syumaigen.ColorMap
	code := r.URL.Query().Get("code")
	if code != "" {
		cMap = syumaigen.GenerateColorMapByColorCode(code)
	} else {
		cMap = syumaigen.DefaultColorMap
	}
	img, err := syumaigen.GenerateImage(
		syumaigen.Pattern,
		cMap,
		10,
	)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Fatal(err)
	}
}

func randomImageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("/image?code=%s", generateRandomColorCode()), http.StatusTemporaryRedirect)
}

func cachedImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=864000, public")
	w.Header().Set("Expires", time.Now().AddDate(0, 0, 10).Format(time.RFC1123))
	imageHandler(w, r)
}
