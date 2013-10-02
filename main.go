package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
)

const (
	defaultAddress  = "cobra:8888"
	defaultPlotSize = 600
)

type Root struct{}

func (this Root) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	t, _ := template.New("root").Parse(rootTemplate)
	d := map[string]string{
		"Title":     "webplot example",
		"ImagePath": "/test.png"}
	t.ExecuteTemplate(w, "root", d)
}

func main() {
	fmt.Println("Listening on", defaultAddress)
	var r Root
	http.Handle("/", r)
	var p Plot
	p.setSize(defaultPlotSize, defaultPlotSize)
	//maxval := float64(defaultPlotSize * defaultPlotSize)
	//p.setFunc(func(x, y int) float64 {
	//	return float64(x) * float64(y) / maxval
	//})
	p.setFunc(func(x, y int) float64 {
		return math.Cos(float64(x)/25.0)*math.Sin(float64(y)/25.0)/2.0 + 0.5
	})
	p.plotGray()
	http.Handle("/test.png", p)
	http.ListenAndServe(defaultAddress, nil)
}

var rootTemplate = `
<html>
<head>
<title>{{.Title}}</title>
</head>
<body>
<h1>{{.Title}}</h1>
<img src="{{.ImagePath}}"/>
</body>
</html>
`
