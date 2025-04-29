package server

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/syumai/syumaigen"
)

var colorCodeRegex = regexp.MustCompile(`^(?:[0-9a-fA-F]{3}){1,2}$`)

func isValidColorCode(code string) bool {
	return colorCodeRegex.MatchString(code)
}

func writePNG(w http.ResponseWriter, cMap syumaigen.ColorMap, scale int) {
	w.Header().Set("Content-Type", "image/png")
	img, err := syumaigen.GenerateImage(
		syumaigen.Pattern,
		cMap,
		scale,
	)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
	var buf bytes.Buffer
	var e png.Encoder
	e.CompressionLevel = png.BestSpeed
	err = e.Encode(&buf, img)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Fatal(err)
	}
}

func writeSVG(w http.ResponseWriter, cMap syumaigen.ColorMap) {
	w.Header().Set("Content-Type", "image/svg+xml")
	img, err := syumaigen.GenerateSVG(
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
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, img); err != nil {
		log.Fatal(err)
	}
}
