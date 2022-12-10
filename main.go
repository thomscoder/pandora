package main

import (
	"io/ioutil"
	"log"
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
	_css, err := ioutil.ReadFile("example.css")

	check(err)

	root, err := html.ParseHTML(string(_html))
	c, err := css.ParseCSS(string(_css))

	renderTree := render.NewRenderTree(root, c)

	displayList := render.NewLayoutTree(renderTree).BuildDisplayList()

	for _, d := range displayList {
		render.PaintNode(d)
	}

	check(err)

	//html.PrintTree(root, 0)
}
