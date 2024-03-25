package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func home() {
	myApp.Window = myApp.App.NewWindow("Groupie Tracker.")
	testButton := widget.NewButton("Enter in Groupie Tracker", func() {
		MainPage()
	})
	content := container.NewVBox(testButton)
	centeredContent := container.NewCenter(content)
	myApp.Window.SetContent(centeredContent)
	myApp.Window.Resize(fyne.NewSize(225, 875))
	myApp.Window.Show()
}
