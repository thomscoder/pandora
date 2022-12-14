package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"pandora/engine/css"
	"pandora/engine/html"
	"pandora/engine/render"
)

// A browser engine, also known as a layout engine or rendering engine,
// is the part of a web browser that is responsible for interpreting and rendering HTML, CSS, and other web content.

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	htmlFile := flag.String("html", "example/example.html", "the name of the html to search")
	cssFile := flag.String("css", "example/example.css", "the name of the html to search")

	// Parse the command-line flags
	flag.Parse()

	// Print the value of the flag
	_html, err := ioutil.ReadFile(*htmlFile)
	check(err)

	_css, err := ioutil.ReadFile(*cssFile)
	check(err)

	rootNode, err := html.ParseHTML(string(_html))
	check(err)

	stylesheet, err := css.ParseCSS(string(_css))
	check(err)

	renderTree := render.NewRenderTree(rootNode, stylesheet)

	layoutTree := render.NewLayoutTree(renderTree)
	displayList := layoutTree.BuildDisplayList()

	file, err := os.Create("image.png")
	check(err)
	defer file.Close()

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	for _, item := range displayList {
		render.PaintNode(img, item)
	}

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}
