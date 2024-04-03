package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	myApp AppData
)

func Begin() {
	// Initializes the Fyne application, creates a new window, starts the home page, sets the window size, and displays the window.
	myApp.App = app.New()
	myApp.Window = myApp.App.NewWindow("FestivaSync")
	// MainPage()
	home()
	myApp.Window.Resize(fyne.NewSize(225, 875))
	myApp.Window.ShowAndRun()
}
