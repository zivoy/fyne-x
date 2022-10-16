package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type Stack struct {
	vertical bool
	size     fyne.Size
}

func (s *Stack) Layout(objs []fyne.CanvasObject, cont fyne.Size) {
	s.size = cont
	s.MinSize(objs) // MinSize is called when you resize a normal window, but not on test windows

	pos := fyne.NewPos(0, 0)
	for i, o := range objs {
		// get the padding adjustment needed
		adjust := theme.Padding() / 2
		if i == len(objs)-1 {
			adjust = 0
		}

		if s.vertical {
			oHeight := o.MinSize().Height
			o.Resize(fyne.NewSize(cont.Width, oHeight-adjust))
			o.Move(pos)
			pos = pos.Add(fyne.NewPos(0, oHeight))
		} else {
			oWidth := cont.Width / float32(len(objs))
			o.Resize(fyne.NewSize(oWidth-adjust, cont.Height))
			o.Move(pos)
			pos = pos.Add(fyne.NewPos(oWidth, 0))
		}
	}
}
func (s *Stack) MinSize(objs []fyne.CanvasObject) fyne.Size {
	var w, h, sumHeight float32
	for _, o := range objs {
		oSize := o.MinSize()
		sumHeight += oSize.Height
		// get the widest part
		w = fyne.Max(w, oSize.Width)
		// and the tallest part
		h = fyne.Max(h, oSize.Height)
	}

	// sum of padding between elements
	padding := theme.Padding() * float32(len(objs)-1)
	// move to vertical layout once the average width - padding breaches the widest minimum width
	s.vertical = w > (s.size.Width-padding)/float32(len(objs))
	if s.vertical {
		// if vertical, change the minimum height from the tallest object to the sum of all the heights
		h = sumHeight + padding
	}
	return fyne.NewSize(w, h)
}

func NewStack(obj ...fyne.CanvasObject) *fyne.Container {
	return container.New(&Stack{}, obj...)
}
