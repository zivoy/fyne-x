package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

const (
	alignLeading = iota
	alignCenter
	alignTrailing
)

// Align aligns a single fyne.CanvasObject vertically and horizontally in its container
// while shrinking it to its minimum size
type Align struct {
	VAlignment, HAlignment int
}

func (c Align) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return objects[0].MinSize()
}

func (c Align) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	size := c.MinSize(objects)
	var xPos, yPos float32
	// horizontal alignment
	switch c.HAlignment {
	case alignLeading:
		xPos = 0
	case alignCenter:
		xPos = containerSize.Width/2 - size.Width/2
	case alignTrailing:
		xPos = containerSize.Width - size.Width
	}

	// vertical alignment
	switch c.VAlignment {
	case alignLeading:
		yPos = 0
	case alignCenter:
		yPos = containerSize.Height/2 - size.Height/2
	case alignTrailing:
		yPos = containerSize.Height - size.Height
	}

	pos := fyne.NewPos(xPos, yPos)
	o1 := objects[0]
	o1.Resize(size)
	o1.Move(pos)
}

// NewAlign creates a new fyne.Container using the Align layout
func NewAlign(obj fyne.CanvasObject, verticalAlignment, horizontalAlignment int) *fyne.Container {
	return container.New(&Align{verticalAlignment, horizontalAlignment}, obj)
}
