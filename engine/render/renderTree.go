package render

import (
	"bytes"
	"pandora/engine/css"
	"pandora/engine/html"
	"strings"
)

// A render tree is a data structure that represents the visual structure and content of a web page.
// It is used by a web browser to render the page on the user's screen.
// The render tree is constructed by combining the DOM tree (which represents the structural hierarchy of the page)
// with the CSSOM (which represents the styles applied to the page elements).
// The resulting render tree specifies the precise position and appearance of each element on the page,
// as well as any interactions or animations that are defined in the page's CSS or JavaScript code.

// - Parse the HTML and CSS code to generate a DOM tree and a CSSOM.
// - Traverse the DOM tree and match each element with the appropriate CSS rules from the CSSOM.
// - Create a new tree-like structure that represents the final visual representation of the page, including the layout and styles of each element.
// - Use the render tree to generate the final visual representation of the page on the user's screen.

type RenderTree struct {
	Root *RenderNode
}

func (rt *RenderTree) String() string {
	return renderTreeToString(rt.Root, 0)
}

func renderTreeToString(node *RenderNode, indent int) string {
	var buffer bytes.Buffer

	// Add indentation
	for i := 0; i < indent; i++ {
		buffer.WriteString("  ")
	}

	// Add the node's tag and attributes
	buffer.WriteString(node.Tag)
	for key, value := range node.Attributes {
		buffer.WriteString(" ")
		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(value)
	}

	// Add the node's text and styles
	if node.Text != "" {
		buffer.WriteString(": ")
		buffer.WriteString(strings.TrimSpace(node.Text))
	}
	if len(node.Styles) > 0 {
		buffer.WriteString(" {")
		for key, value := range node.Styles {
			buffer.WriteString(" ")
			buffer.WriteString(key)
			buffer.WriteString(": ")
			buffer.WriteString(value)
		}
		buffer.WriteString(" }")
	}
	buffer.WriteString("\n")

	// Recursively add the node's children
	for _, child := range node.Children {
		buffer.WriteString(renderTreeToString(child, indent+1))
	}

	return buffer.String()
}

type RenderNode struct {
	Tag        string
	Attributes map[string]string
	Text       string
	Styles     map[string]string
	Children   []*RenderNode
}

func NewRenderTree(dom *html.Node, stylesheet *css.Stylesheet) *RenderTree {
	root := buildRenderTree(dom, stylesheet)
	return &RenderTree{
		Root: root,
	}
}

func buildRenderTree(domNode *html.Node, stylesheet *css.Stylesheet) *RenderNode {
	renderNode := &RenderNode{
		Tag:        domNode.Name,
		Attributes: domNode.Attributes,
		Text:       domNode.Text,
		Children:   []*RenderNode{},
	}

	styles := lookupStyles(domNode, stylesheet)
	renderNode.Styles = styles

	for _, child := range domNode.Children {
		renderChild := buildRenderTree(child, stylesheet)
		renderNode.Children = append(renderNode.Children, renderChild)
	}

	return renderNode
}

func lookupStyles(domNode *html.Node, stylesheet *css.Stylesheet) map[string]string {
	// Look up the styles for the element in the stylesheet
	styles := map[string]string{}
	for _, rule := range stylesheet.Rules {
		if strings.TrimSpace(rule.Selector) == domNode.Name {
			styles = rule.Properties
			break
		}

		// Check if the selector matches the element's class attribute
		if strings.HasPrefix(rule.Selector, ".") {
			class := strings.TrimPrefix(rule.Selector, ".")
			if strings.ReplaceAll(domNode.Attributes["class"], "\"", "") == strings.TrimSpace(class) {
				styles = rule.Properties
				break
			}
		}

		// Check if the selector matches the element's ID attribute
		if strings.HasPrefix(rule.Selector, "#") {
			id := strings.TrimPrefix(rule.Selector, "#")
			if strings.ReplaceAll(domNode.Attributes["id"], "\"", "") == strings.TrimSpace(id) {
				styles = rule.Properties
				break
			}
		}
	}
	return styles
}
