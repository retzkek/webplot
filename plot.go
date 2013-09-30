package main

import (
	"image"
	"image/color"
	"image/png"
	"net/http"
)

// Plot generates and stores a 2D plot described by a function
// F(x,y) that is defined on (0,1).
type Plot struct {
	Width  int
	Height int
	F      func(int, int) float64

	plot image.Image
}

func (this *Plot) setSize(width, height int) {
	this.Width = width
	this.Height = height
}

func (this *Plot) setFunc(f func(int, int) float64) {
	this.F = f
}

func (this *Plot) plotGray() image.Image {
	p := image.NewGray(image.Rectangle{image.ZP,
		image.Point{this.Width, this.Height}})
	for x := 0; x < this.Width; x++ {
		for y := 0; y < this.Height; y++ {
			p.Set(x, y, color.Gray{uint8(this.F(x, y) * 255)})
		}
	}
	this.plot = p
	return p
}

func (this Plot) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	png.Encode(w, this.plot)
}
