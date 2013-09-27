package main

import (
	"fmt"
	"net/http"
	"image"
	"image/png"
	"image/color"
)

const (
	defaultAddress = "localhost:8888"
	defaultPlotSize = 600
)

type Root struct{}

func (this Root) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w,"<h1>hello</h1><img src=\"/test.png\"/>")
}


type Plot struct {
	PlotSize int
	Plot image.Image
}

func (this *Plot) initGray() {
	this.PlotSize = defaultPlotSize
	p := image.NewGray(image.Rectangle{image.ZP,
		image.Point{this.PlotSize,this.PlotSize}})
	maxval := this.PlotSize*this.PlotSize
	for x := 0; x < this.PlotSize; x++ {
		for y := 0; y < this.PlotSize; y++ {
			p.Set(x,y,color.Gray{uint8(255*x*y/maxval)})
		}
	}
	this.Plot = p
}

func (this Plot) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	png.Encode(w, this.Plot)
}

func main() {
	fmt.Println("Listening on",defaultAddress)
	var r Root
	http.Handle("/",r)
	var p Plot
	p.initGray()
	http.Handle("/test.png",p)
	http.ListenAndServe(defaultAddress, nil)
}

