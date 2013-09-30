package main

import (
	"fmt"
	"math"
	"net/http"
)

const (
	defaultAddress  = "localhost:8888"
	defaultPlotSize = 600
)

type Root struct{}

func (this Root) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "<h1>hello</h1><img src=\"/test.png\"/>")
}

func main() {
	fmt.Println("Listening on", defaultAddress)
	var r Root
	http.Handle("/", r)
	var p Plot
	p.setSize(defaultPlotSize, defaultPlotSize)
	maxval := float64(defaultPlotSize * defaultPlotSize)
	p.setFunc(func(x, y int) float64 {
		return float64(x) * float64(y) / maxval
	})
	p.setFunc(func(x, y int) float64 {
		return math.Cos(float64(x)/25.0)*math.Sin(float64(y)/25.0)/2.0 + 0.5
	})
	p.plotGray()
	http.Handle("/test.png", p)
	http.ListenAndServe(defaultAddress, nil)
}
