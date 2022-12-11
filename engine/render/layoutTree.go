package render

// A layout tree is a data structure used by a web browser to represent the hierarchical layout of the elements on the page.
// It is typically used by the browser to calculate the positions and sizes of the elements on the page,
// as well as their relationships with each other.

// When rendering a web page, the browser uses the layout tree to determine the positions and sizes of the elements on the page.
// This allows the browser to correctly position and size the elements, and to lay out the elements in a hierarchical manner.

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

type LayoutTree struct {
	Root *LayoutNode
}

type LayoutNode struct {
	Box        *Box
	Children   []*LayoutNode
	RenderNode *RenderNode
}

type Box struct {
	Width    int
	Height   int
	Padding  int
	Margin   int
	Position Point
}

type Point struct {
	X int
	Y int
}

func (b *Box) String() string {
	return fmt.Sprintf("Box{Width: %d, Height: %d, Padding: %d, Margin: %d, Position: Point{X: %d, Y: %d}}", b.Width, b.Height, b.Padding, b.Margin, b.Position.X, b.Position.Y)
}

func (lt *LayoutTree) String() string {
	return layoutTreeToString(lt.Root, 0)
}

func layoutTreeToString(node *LayoutNode, indent int) string {
	var buffer bytes.Buffer
	for i := 0; i < indent; i++ {
		buffer.WriteString("  ")
	}

	buffer.WriteString(fmt.Sprintf("(%d,%d) %dx%d", node.Box.Position.X, node.Box.Position.Y, node.Box.Width, node.Box.Height))
	buffer.WriteString("\n")

	for _, child := range node.Children {
		buffer.WriteString(layoutTreeToString(child, indent+1))
	}

	return buffer.String()
}

// create the layout tree
func NewLayoutTree(renderTree *RenderTree) *LayoutTree {
	root := buildLayoutTree(renderTree.Root, Point{X: 0, Y: 0})
	return &LayoutTree{
		Root: root,
	}
}

func buildLayoutTree(renderNode *RenderNode, position Point) *LayoutNode {
	pos := Point{
		Y: position.Y + parseStyle(renderNode.Styles["top"], 0) + parseStyle(renderNode.Styles["margin"], 0) + parseStyle(renderNode.Styles["margin-top"], 0) - parseStyle(renderNode.Styles["margin-bottom"], 0),
		X: position.X + parseStyle(renderNode.Styles["left"], 0) + parseStyle(renderNode.Styles["margin"], 0) - parseStyle(renderNode.Styles["margin-right"], 0),
	}

	box := &Box{
		Width:    parseStyle(renderNode.Styles["width"], 0),
		Height:   parseStyle(renderNode.Styles["height"], 0),
		Margin:   parseStyle(renderNode.Styles["margin"], 0),
		Position: pos,
	}

	layoutNode := &LayoutNode{
		Box:        box,
		Children:   []*LayoutNode{},
		RenderNode: renderNode,
	}

	childX := box.Position.X + box.Margin + box.Padding
	childY := box.Position.Y + box.Margin + box.Padding

	for _, child := range renderNode.Children {
		display := child.Styles["display"]
		if display == "block" {
			childY = layoutNode.Box.Position.Y + layoutNode.Box.Height + layoutNode.Box.Margin + layoutNode.Box.Padding
		}
		childLayout := buildLayoutTree(child, Point{X: childX, Y: childY})
		layoutNode.Children = append(layoutNode.Children, childLayout)
	}

	return layoutNode
}

func parseStyle(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}

	regex := regexp.MustCompile(`(\d+)`)

	// Use the FindString function to extract the numeric part of the string.
	numericPart := regex.FindString(value)

	val, err := strconv.Atoi(numericPart)

	if err != nil {
		return defaultValue
	}
	return val
}
