package main

import (
	"io/ioutil"
	"log"
	"pandora/engine/html"
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

	root, err := html.ParseHTML(string(_html))

	check(err)

	html.PrintTree(root, 0)
}
