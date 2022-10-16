package layout

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"testing"
)

var stackItemSizes = [...]fyne.Size{
	{150, 100},
	{100, 200},
	{130, 130},
	{120, 100},
}

func makeStackItems(t *testing.T) *fyne.Container {
	t.Helper()
	cont := NewStack()
	for i, s := range stackItemSizes {
		rectangle := canvas.NewRectangle(color.Gray{Y: 0x7a})
		rectangle.SetMinSize(s)
		text := widget.NewLabel(fmt.Sprintf("item #%d", i+1))
		cont.Add(container.NewMax(rectangle, container.NewCenter(text)))
	}
	return cont
}

func TestStackLayoutVertical(t *testing.T) {
	w := test.NewWindow(makeStackItems(t))
	c := w.Canvas()

	container.NewAdaptiveGrid()

	// wider
	w.Resize(fyne.NewSize(400, 600))
	test.AssertImageMatches(t, "stack/verticalWide.png", c.Capture())

	// thinner
	w.Resize(fyne.NewSize(250, 600))
	test.AssertImageMatches(t, "stack/verticalThin.png", c.Capture())
}

func TestStackLayoutHorizontal(t *testing.T) {
	w := test.NewWindow(makeStackItems(t))
	c := w.Canvas()

	// wider
	w.Resize(fyne.NewSize(1500, 250))
	test.AssertImageMatches(t, "stack/horizontalWide.png", c.Capture())

	// thinner
	w.Resize(fyne.NewSize(800, 250))
	test.AssertImageMatches(t, "stack/horizontalThin.png", c.Capture())
}
