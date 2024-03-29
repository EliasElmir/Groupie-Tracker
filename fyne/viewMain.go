package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	myApp AppData
)

func Begin() {
	myApp.App = app.New()
	myApp.Window = myApp.App.NewWindow("Groupie Tracker.")
	// MainPage()
	home()
	myApp.Window.Resize(fyne.NewSize(225, 875))
	myApp.Window.ShowAndRun()
}
