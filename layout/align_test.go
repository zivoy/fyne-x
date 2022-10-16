package layout

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
	"image/color"
	"testing"
)

var alignments = map[string]int{
	"leading":  alignLeading,
	"center":   alignCenter,
	"trailing": alignTrailing,
}

func TestAlignments(t *testing.T) {
	el := canvas.NewRectangle(color.Gray{Y: 0x7a})
	el.SetMinSize(fyne.NewSize(20, 20))

	for h, aH := range alignments {
		for v, aV := range alignments {
			t.Run(fmt.Sprintf("%s-horisontal_%s-vertical", h, v), func(t *testing.T) {
				container := NewAlign(el, aV, aH)

				w:= test.NewWindow(container)
				w.Resize(fyne.NewSize(50, 50))
				test.AssertImageMatches(t, fmt.Sprintf("align/%sH-%sV.png", h, v), w.Canvas().Capture())
			})
		}
	}
}
