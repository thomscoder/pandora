package html

import (
	"fmt"
	"strings"
)

// Parse the HTML code.
// This involves breaking the code up into individual tokens, such as tags, attributes, and text content,
// and organizing them into a tree-like structure that represents the hierarchical structure of the page

// Convert the tokens into a DOM tree.
// The DOM tree is a representation of the page's structure and content that can be understood by the browser engine.
// It includes nodes for each element on the page, such as headings, paragraphs, images, and links,
// as well as their attributes and text content.

// Traverse the DOM tree and generate output.
// Once the DOM tree has been constructed, you can traverse it and generate output that represents the tree's structure and content.
// This could be done using a simple recursive algorithm that visits each node in the tree
// and outputs its information in a suitable format.

// Optionally, perform additional processing on the DOM tree.
// Depending on the requirements of your browser engine, you may want to perform additional processing on the DOM tree,
// such as applying CSS styles, calculating layout, or executing JavaScript code.
// These steps are typically performed by the browser engine itself,
// but you could also add them to your HTML parser if desired.

type Node struct {
	Name string

	Attributes map[string]string

	Children []*Node

	Text string
}

// parses the html string and returns the root node
func ParseHTML(html string) (*Node, error) {
	stack := []*Node{{}}

	pos := 0

	for pos < len(html) {

		// we look for the next '<' character
		// if there is no more '<' means we reached the end since </html>
		// closes the document

		next := strings.IndexByte(html[pos:], '<')

		if next == -1 {
			// we passed the </html>
			break
		}

		next += pos

		// is it end of tag or start of tag?
		if html[next+1] == '/' {
			if len(stack) == 0 {
				return nil, fmt.Errorf("no matching closing tag at position %d", next)
			}
			// end of the tag we pop from the stack
			stack = stack[:len(stack)-1]
			pos = next + 1

		} else {
			// start of the tag
			node, end := parseTag(html[next+1:])

			if node == nil {
				// if the tag is not well formed skip
				pos = end + next + 1
				continue
			}

			// add the new node as a child of the current node
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, node)

			stack = append(stack, node)
			pos = end + next + 1
		}
	}

	// return the root node
	return stack[0], nil
}

// parses the tag and returns the node
func parseTag(html string) (*Node, int) {
	// find the index of the next opening < character
	textEnd := strings.IndexByte(html, '<')

	// find the index of the closing > character
	end := strings.IndexByte(html, '>')

	if end == -1 {
		return nil, 0
	}

	var text string

	if textEnd != -1 {
		text = html[:textEnd]
	}

	text = strings.TrimSpace(text)

	// remove the opening tag from the text

	// we parse the tag and the attributes
	partsOfTag := strings.Split(html[:end], " ")
	name := partsOfTag[0]
	text = strings.TrimPrefix(text, html[:end]+">")

	attributes := make(map[string]string)
	for _, part := range partsOfTag[1:] {
		// Attributes are written like class="something"
		// so we can use SplitN to get the attr name and the value
		attributeAtXRay := strings.SplitN(part, "=", 2)
		if len(attributeAtXRay) == 2 {
			attributes[attributeAtXRay[0]] = attributeAtXRay[1]
		}
	}

	node := &Node{
		Name:       name,
		Attributes: attributes,
		Children:   []*Node{},
		Text:       text,
	}

	return node, end
}

// build the DOM tree
func PrintTree(node *Node, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%s", node.Name)
	for k, v := range node.Attributes {
		fmt.Printf(" %s=%s", k, v)
	}
	fmt.Printf(" %s", node.Text)
	fmt.Println()

	// Print the node's children.
	for _, child := range node.Children {
		PrintTree(child, indent+1)
	}
}
