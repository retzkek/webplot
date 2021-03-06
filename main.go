package main

import (
	"flag"
	"fmt"
	"html/template"
	"math"
	"net/http"
)

var (
	address  = flag.String("a", "localhost:8888", "Address and port to listen on")
	plotSize = flag.Int("s", 600, "Plot image size")
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
	flag.Parse()

	var r Root
	http.Handle("/", r)
	var p Plot
	p.setSize(*plotSize, *plotSize)
	//maxval := float64(defaultPlotSize * defaultPlotSize)
	//p.setFunc(func(x, y int) float64 {
	//	return float64(x) * float64(y) / maxval
	//})
	p.setFunc(func(x, y int) float64 {
		return math.Cos(float64(x)/25.0)*math.Sin(float64(y)/25.0)/2.0 + 0.5
	})
	p.plotGray()
	http.Handle("/test.png", p)

	fmt.Println("Listening on", *address)
	http.ListenAndServe(*address, nil)
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
