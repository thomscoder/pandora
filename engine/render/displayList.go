package render

import "fmt"

// A display list, is a flat list of visual elements that are ready to be rendered.
// This makes it faster to render the elements on the screen,
// because the display list can be easily processed and rendered in a single pass and we don't have to traverse
// the layout tree all the time

type DisplayListItem struct {
	Box    *Box
	Styles map[string]string
	Text   string
}

func (dli *DisplayListItem) String() string {
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

	displayListItem := &DisplayListItem{
		Box:    node.Box,
		Styles: node.RenderNode.Styles,
		Text:   node.RenderNode.Text,
	}
	*displayList = append(*displayList, displayListItem)

	for _, child := range node.Children {
		buildDisplayListForNode(child, displayList)
	}
}
