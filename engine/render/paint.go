package render

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type DisplayListItem struct {
	Box    *Box
	Styles map[string]string
}

func (dli *DisplayListItem) String() string {
	// Convert the Box and Styles properties to strings
	boxString := dli.Box.String()
	stylesString := fmt.Sprintf("%+v", dli.Styles)

	// Return a string representation of the DisplayListItem
	return fmt.Sprintf("DisplayListItem{Box: %s, Styles: %s}", boxString, stylesString)
}

type DisplayList []*DisplayListItem

func (lt *LayoutTree) BuildDisplayList() DisplayList {
	displayList := make(DisplayList, 0)
	buildDisplayListForNode(lt.Root, &displayList)
	return displayList
}

func buildDisplayListForNode(node *LayoutNode, displayList *DisplayList) {
	// Add the current node's Box and Styles to the display list
	displayListItem := &DisplayListItem{
		Box:    node.Box,
		Styles: node.RenderNode.Styles,
	}
	*displayList = append(*displayList, displayListItem)

	// Recursively call buildDisplayListForNode for each child node
	for _, child := range node.Children {
		buildDisplayListForNode(child, displayList)
	}
}

// -------------------------------------------------------
// -------------------------------------------------------
// -------------------------------------------------------
// -------------------------------------------------------
// -------------------------------------------------------
// -------------------------------------------------------
// -------------------------------------------------------
var c = canvas.New(300, 300)

// Create a canvas context used to keep drawing state
var ctx = canvas.NewContext(c)

// Create a new RGBA image with the specified dimensions
func PaintNode(displayListItem *DisplayListItem) {
	// Create new canvas of dimension 100x100 mm

	// Create a triangle path from an SVG path and draw it to the canvas
	triangle := canvas.Rectangle(float64(displayListItem.Box.Width), float64(displayListItem.Box.Height))

	fmt.Println(float64(displayListItem.Box.Width), float64(displayListItem.Box.Height))
	ctx.SetFillColor(canvas.RGBA(parseColor(displayListItem.Styles["background-color"])))

	fmt.Println(canvas.Aliceblue, canvas.RGBA(parseColor(displayListItem.Styles["background-color"])))
	ctx.DrawPath(float64(displayListItem.Box.Position.X), float64(displayListItem.Box.Position.Y), triangle)

	// Rasterize the canvas and write to a PNG file with 3.2 dots-per-mm (320x320 px)
	renderers.Write("getting-started.png", c, canvas.DPMM(3.2))
}

func parseColor(colorString string) (uint8, uint8, uint8, float64) {
	// Check if the color string is in the RGB format (e.g. "rgb(255, 0, 0)")
	if strings.HasPrefix(colorString, "rgb(") && strings.HasSuffix(colorString, ");") {
		// Strip the "rgb(" and ")" prefix and suffix from the color string
		colorString = strings.TrimPrefix(colorString, "rgb(")
		colorString = strings.TrimSuffix(colorString, ");")

		// Split the color string into the red, green, and blue components
		parts := strings.Split(colorString, ",")

		// Parse the red, green, and blue components as integers
		r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		b, err := strconv.Atoi(strings.TrimSpace(parts[2]))

		// If any errors occurred during parsing, return the default color (black)
		if err != nil {
			return 0, 0, 0, 1
		}

		// Return the parsed color components as integers in the range [0, 255]
		return uint8(r), uint8(g), uint8(b), 1
	}

	// If the color string is not in the RGB format, return the default color (black)
	return 0, 0, 0, 0
}
