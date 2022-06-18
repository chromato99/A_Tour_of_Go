package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// set rectangle size 100 * 100
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

// set green color
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(0), uint8(250), uint8(0), uint8(255)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
