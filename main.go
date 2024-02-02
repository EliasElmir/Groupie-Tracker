package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello word")

	w.SetContent(widget.NewLabel("Hello"))

	w.ShowAndRun()
}
