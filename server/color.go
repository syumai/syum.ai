package server

import (
	"math/rand/v2"
	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func generateRandomColorCode() string {
	h := rand.Float64() * 360.0
	c := 0.4 + rand.Float64()*0.6
	l := 0.5
	return strings.TrimLeft(colorful.Hcl(h, c, l).Hex(), "#")
}
