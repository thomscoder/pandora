package render

import (
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func PaintNode(file *os.File, img *image.RGBA, displayListItem *DisplayListItem) {

	posX := displayListItem.Box.Position.X
	posY := displayListItem.Box.Position.Y
	width := displayListItem.Box.Width
	height := displayListItem.Box.Height

	// Define the rectangles to draw
	rect := image.Rect(posX, posY, posX+width, posY+height)

	// Define the color to use for the rectangles
	r, g, b, a := parseColor(displayListItem.Styles["background-color"])

	col := color.RGBA{R: r, G: g, B: b, A: a}
	draw.Draw(img, rect, &image.Uniform{col}, image.Point{}, draw.Src)
	drawText(img, displayListItem)
}

func drawText(img *image.RGBA, displayListItem *DisplayListItem) {

	f := basicfont.Face7x13
	text := displayListItem.Text

	// Draw the text on the image
	point := fixed.Point26_6{
		X: fixed.I(displayListItem.Box.Position.X) + ((fixed.I(displayListItem.Box.Width) / 2) - (font.MeasureString(f, text) / 2)),
		Y: fixed.I(displayListItem.Box.Position.Y) + (fixed.I(displayListItem.Box.Height) / 2),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: f,
		Dot:  point,
	}
	d.DrawString(text)
}

// get the color from the CSS property
func parseColor(colorString string) (uint8, uint8, uint8, uint8) {
	if strings.HasPrefix(colorString, "rgb(") && strings.HasSuffix(colorString, ");") {

		colorString = strings.TrimPrefix(colorString, "rgb(")
		colorString = strings.TrimSuffix(colorString, ");")

		parts := strings.Split(colorString, ",")

		r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return 0, 0, 0, 255
		}

		g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return 0, 0, 0, 255
		}

		b, err := strconv.Atoi(strings.TrimSpace(parts[2]))

		if err != nil {
			return 0, 0, 0, 255
		}

		return uint8(r), uint8(g), uint8(b), 255
	}

	return 0, 0, 0, 0
}
