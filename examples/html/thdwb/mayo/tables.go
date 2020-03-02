package mayo

import (
	"regexp"

	"github.com/p9c/learngio/examples/html/thdwb/structs"
)

var rgba = regexp.MustCompile(`rgba?\([\.?\d?\.?\d?%?\s?,?]+\)`)
var rgbaParams = regexp.MustCompile(`\([\.?\d?\.?\d?%?\s?,?]+\)`)

var colorTable = map[string]*structs.ColorRGBA{
	"maroon":  &structs.ColorRGBA{R: 0.5, G: 0.0, B: 0.0, A: 1.0},
	"red":     &structs.ColorRGBA{R: 1.0, G: 0.0, B: 0.0, A: 1.0},
	"orange":  &structs.ColorRGBA{R: 1.0, G: 0.6, B: 0.0, A: 1.0},
	"yellow":  &structs.ColorRGBA{R: 1.0, G: 1.0, B: 0.0, A: 1.0},
	"olive":   &structs.ColorRGBA{R: 0.5, G: 0.5, B: 0.0, A: 1.0},
	"green":   &structs.ColorRGBA{R: 0.0, G: 1.0, B: 0.0, A: 1.0},
	"purple":  &structs.ColorRGBA{R: 0.5, G: 0.0, B: 0.5, A: 1.0},
	"fuchsia": &structs.ColorRGBA{R: 1.0, G: 0.0, B: 1.0, A: 1.0},
	"lime":    &structs.ColorRGBA{R: 0.0, G: 1.0, B: 0.0, A: 1.0},
	"teal":    &structs.ColorRGBA{R: 0.0, G: 0.5, B: 0.5, A: 1.0},
	"aqua":    &structs.ColorRGBA{R: 0.0, G: 1.0, B: 1.0, A: 1.0},
	"blue":    &structs.ColorRGBA{R: 0.0, G: 0.0, B: 1.0, A: 1.0},
	"navy":    &structs.ColorRGBA{R: 0.0, G: 0.0, B: 0.5, A: 1.0},
	"black":   &structs.ColorRGBA{R: 0.0, G: 0.0, B: 0.0, A: 1.0},
	"gray":    &structs.ColorRGBA{R: 0.5, G: 0.5, B: 0.5, A: 1.0},
	"silver":  &structs.ColorRGBA{R: 0.7, G: 0.7, B: 0.7, A: 1.0},
	"white":   &structs.ColorRGBA{R: 1.0, G: 1.0, B: 1.0, A: 1.0},
}

var elementFontTable = map[string]float64{
	"h1": float64(32),
	"h2": float64(28),
	"p":  float64(14),
}
