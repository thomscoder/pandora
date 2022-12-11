package main

import (
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
	_html, err := ioutil.ReadFile("example.html")
	check(err)

	_css, err := ioutil.ReadFile("example.css")
	check(err)

	rootNode, err := html.ParseHTML(string(_html))
	check(err)

	stylesheet, err := css.ParseCSS(string(_css))
	check(err)

	renderTree := render.NewRenderTree(rootNode, stylesheet)

	layoutTree := render.NewLayoutTree(renderTree)
	displayList := layoutTree.BuildDisplayList()

	file, err := os.Create("image.png")
	defer file.Close()

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	for _, item := range displayList {
		render.PaintNode(file, img, item)
	}

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}
