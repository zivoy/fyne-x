package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type AdvancedWrap struct {
	size fyne.Size
}

func (a *AdvancedWrap) Layout(objs []fyne.CanvasObject, cont fyne.Size) {
	a.size = cont

	var x, y, h float32
	for _, o := range objs {
		if !o.Visible() {
			continue
		}
		s := o.MinSize()
		if cont.Width < x+s.Width {
			y += h
			x = 0
			h = 0
		}
		o.Resize(s)
		o.Move(fyne.NewPos(x, y))
		h = fyne.Max(s.Height, h)
		x += s.Width
	}
}

func (a AdvancedWrap) MinSize(objs []fyne.CanvasObject) fyne.Size {
	//            row, size
	items := make([][]fyne.Size, 0)

	var w float32
	row := make([]fyne.Size, 0)
	for _, o := range objs {
		if !o.Visible() {
			continue
		}
		s := o.MinSize()
		if a.size.Width < w+s.Width {
			w = 0
			items = append(items, row)
			row = make([]fyne.Size, 0)
		}
		row = append(row, s)
		w += s.Width
	}
	items = append(items, row)

	var h, maxH float32
	w = 0
	for _, row := range items {
		maxH = 0
		for _, s := range row {
			maxH = fyne.Max(s.Height, maxH)
			w = fyne.Max(w, s.Width)
		}
		h += maxH
	}
	return fyne.NewSize(w, h)
}

func NewAdvancedWrap(objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(&AdvancedWrap{}, objects...)
}

