package server

import (
	"bytes"
	"fmt"
	"image"
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
	if _, err := io.Copy(w, strings.NewReader(syumaiASCIIArt)); err != nil {
		log.Fatal(err)
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	var (
		img image.Image
		err error
	)
	code := r.URL.Query().Get("code")
	if code != "" {
		img, err = syumaigen.GenerateImage(
			syumaigen.Pattern,
			syumaigen.GenerateColorMapByColorCode(code),
			10,
		)
	} else {
		img, err = syumaigen.GenerateImage(
			syumaigen.Pattern,
			syumaigen.GenerateRandomColorMap(),
			10,
		)
	}
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

func cachedImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=864000, public")
	w.Header().Set("Expires", time.Now().AddDate(0, 0, 10).Format(time.RFC1123))
	imageHandler(w, r)
}
