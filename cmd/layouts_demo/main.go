package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	w := a.NewWindow("X Layouts")
}

func makeAlignedLayout() *fyne.Container {
	theme.FyneLogo()
}
