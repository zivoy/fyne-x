package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
	"testing"
)

var minSizeListTestItems = [...]string{"element", "one", "two", "e", "something longer", "2"}

func makeList(t *testing.T) *widget.List {
	t.Helper()
	// a list is probably one of the better items to test the CustomMinSize layout on
	return widget.NewList(func() int {
		return len(minSizeListTestItems)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("") // placeholder is short to demonstrate the MinSize
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		o := object.(*widget.Label)
		o.SetText(minSizeListTestItems[id])
	})
}

func TestAdjustDefaultMinSize(t *testing.T) {
	list := makeList(t)
	var width, height float32
	width = 200
	height = 300
	item := NewSetMinSize(list, width, height)

	size := item.Size()
	assert.Equal(t, width, size.Width, "Width errored")
	assert.Equal(t, height, size.Height, "Height errored")
}

func TestMinSizeAutoHeight(t *testing.T) {
	list := makeList(t)
	_ = test.NewWindow(list) // populate the list
	height := list.Size().Height
	var width float32 = 200
	item := NewSetMinSize(list, width, 0)

	size := item.Size()
	assert.Equal(t, width, size.Width, "Width errored")
	assert.Equal(t, height, size.Height, "Height errored")
}

func TestMinSizeBiggerOfMinSize(t *testing.T) {
	list := makeList(t)
	_ = test.NewWindow(list) // populate the list
	ListSize := list.Size()
	var height float32 = 20 // smaller than one line
	item := NewSetMinSize(list, 0, height)

	size := item.Size()
	assert.Equal(t, ListSize.Width, size.Width, "Width errored")
	assert.Equal(t, ListSize.Height, size.Height, "Height is invalid")
}
