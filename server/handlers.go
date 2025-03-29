package server

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/syumai/syum.ai/server/pages/indexpage"
	"github.com/syumai/syumaigen"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ascii", asciiHandler)
	mux.HandleFunc("/image", nocacheImageHandler)
	mux.HandleFunc("/og", ogImageHandler)
	mux.HandleFunc("/image/random", randomImageHandler)
	mux.HandleFunc("/favicon.ico", cachedImageHandler)
	mux.HandleFunc("/robots.txt", robotsHandler)
	mux.HandleFunc("/{$}", indexHandler)
	mux.HandleFunc("/", assetsHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'unsafe-inline'; img-src 'self' blob:;")
	w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	colorCode := r.URL.Query().Get("colorCode")
	initialColorCode := generateRandomColorCode()
	indexpage.Index(initialColorCode, colorCode).Render(r.Context(), w)
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, strings.NewReader(SyumaiASCIIArt)); err != nil {
		log.Fatal(err)
	}
}

func ogImageHandler(w http.ResponseWriter, r *http.Request) {
	const width, height = 1200, 630

	baseImg := image.NewRGBA(image.Rect(0, 0, width, height))
	bgColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	draw.Draw(baseImg, baseImg.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	colorCode := r.URL.Query().Get("colorCode")
	drawOGAvatar(baseImg, colorCode)
	drawOGLabel(baseImg, colorCode)

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, baseImg); err != nil {
		log.Fatal(err)
	}
}

func drawOGAvatar(baseImg *image.RGBA, colorCode string) {
	var cMap syumaigen.ColorMap
	if isValidColorCode(colorCode) {
		cMap = syumaigen.GenerateColorMapByColorCode(colorCode)
	} else {
		cMap = syumaigen.DefaultColorMap
	}
	syumaiImg, _ := syumaigen.GenerateImage(
		syumaigen.Pattern,
		cMap,
		20,
	)
	baseImgBounds := baseImg.Bounds()
	syumaiImgBounds := syumaiImg.Bounds()
	pos := image.Point{
		X: baseImgBounds.Dx()/2 - syumaiImgBounds.Dx()/2,
		Y: baseImgBounds.Dy()/2 - syumaiImgBounds.Dy()/2 - 50,
	}
	rect := image.Rectangle{Min: pos, Max: pos.Add(syumaiImg.Bounds().Size())}
	draw.Draw(baseImg, rect, syumaiImg, syumaiImg.Bounds().Min, draw.Over)
}

func drawOGLabel(baseImg *image.RGBA, colorCode string) {
	label := "syum.ai"
	if isValidColorCode(colorCode) {
		label = "#" + strings.ToUpper(colorCode)
	}

	face := basicfont.Face7x13
	d := &font.Drawer{
		Dst:  baseImg,
		Src:  image.NewUniform(color.Black),
		Face: face,
	}
	textWidth := int(d.MeasureString(label) >> 6)
	textHeight := face.Height

	tmp := image.NewRGBA(image.Rect(0, 0, textWidth, textHeight))
	d.Dst = tmp
	d.Dot = fixed.P(0, face.Ascent)
	d.DrawString(label)

	const scale = 7.5
	newWidth := int(float64(textWidth) * scale)
	newHeight := int(float64(textHeight) * scale)
	scaledText := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.NearestNeighbor.Scale(scaledText, scaledText.Bounds(), tmp, tmp.Bounds(), draw.Over, nil)

	baseImgBounds := baseImg.Bounds()

	x := baseImgBounds.Dx()/2 - newWidth/2
	y := baseImgBounds.Dy()/2 + newHeight/2 + 100

	rect := image.Rect(x, y, x+newWidth, y+newHeight)
	draw.Draw(baseImg, rect, scaledText, image.Point{}, draw.Over)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	var cMap syumaigen.ColorMap
	code := r.URL.Query().Get("code")
	if isValidColorCode(code) {
		cMap = syumaigen.GenerateColorMapByColorCode(code)
	} else {
		cMap = syumaigen.DefaultColorMap
	}
	imgType := r.URL.Query().Get("type")
	if imgType == "svg" {
		writeSVG(w, cMap)
		return
	}
	scale := 10
	scaleStr := r.URL.Query().Get("scale")
	if scaleStr != "" {
		var err error
		scale, err = strconv.Atoi(scaleStr)
		if err != nil {
			http.Error(w, "Invalid scale", http.StatusBadRequest)
			return
		}
	}
	writePNG(w, cMap, scale)
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
