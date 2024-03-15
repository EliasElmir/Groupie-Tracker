package Fyne

import (
	"groupietracker/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	myApp model.AppData
)

func Begin() {
	myApp.App = app.New()
	myApp.Window = myApp.App.NewWindow("Groupie Tracker.")
	MainPage()
	myApp.Window.Resize(fyne.NewSize(225, 875))
	myApp.Window.ShowAndRun()
}
