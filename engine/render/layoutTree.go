package render

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

func (b *Box) String() string {
	// Use the fmt.Sprintf function to generate a string representation of the Box properties
	return fmt.Sprintf("Box{Width: %d, Height: %d, Padding: %d, Margin: %d, Position: Point{X: %d, Y: %d}}", b.Width, b.Height, b.Padding, b.Margin, b.Position.X, b.Position.Y)
}

type Point struct {
	X int
	Y int
}

func (lt *LayoutTree) String() string {
	return layoutTreeToString(lt.Root, 0)
}

func layoutTreeToString(node *LayoutNode, indent int) string {
	var buffer bytes.Buffer
	// Add indentation
	for i := 0; i < indent; i++ {
		buffer.WriteString("  ")
	}

	// Add the box's dimensions and position
	buffer.WriteString(fmt.Sprintf("(%d,%d) %dx%d", node.Box.Position.X, node.Box.Position.Y, node.Box.Width, node.Box.Height))
	buffer.WriteString("\n")

	// Recursively add the node's children
	for _, child := range node.Children {
		buffer.WriteString(layoutTreeToString(child, indent+1))
	}

	return buffer.String()
}

func NewLayoutTree(renderTree *RenderTree) *LayoutTree {
	root := buildLayoutTree(renderTree.Root, Point{X: 0, Y: 0})
	return &LayoutTree{
		Root: root,
	}
}

func buildLayoutTree(renderNode *RenderNode, position Point) *LayoutNode {
	box := &Box{
		Width:    parseStyle(renderNode.Styles["width"], 100),
		Height:   parseStyle(renderNode.Styles["height"], 100),
		Padding:  parseStyle(renderNode.Styles["padding"], 0),
		Margin:   parseStyle(renderNode.Styles["margin"], 0),
		Position: Point{X: parseStyle(renderNode.Styles["top"], 0), Y: parseStyle(renderNode.Styles["left"], 0)},
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
