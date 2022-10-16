package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// CustomMinSize allow you to override the minimum size of a single fyne.CanvasObject.
// but it will use the contents min width or height if it is bigger then the dimension specified / todo investigate
//
// setting the width or height to 0 will make it use the contents min dimension
type CustomMinSize struct {
	Width, Height float32
}

func (m CustomMinSize) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := m.Width, m.Height
	oMin := objects[0].MinSize()
	if w == 0 {
		w = oMin.Width
	}
	if h == 0 {
		h = oMin.Height
	}
	return fyne.NewSize(fyne.Max(w, oMin.Width), fyne.Max(h, oMin.Height))
}

func (m CustomMinSize) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	o1 := objects[0]
	o1.Resize(containerSize)
}

// NewSetMinSize creates a new fyne.Container using the CustomMinSize layout
func NewSetMinSize(obj fyne.CanvasObject, width, height float32) *fyne.Container {
	return container.New(&CustomMinSize{width, height}, obj)
}
