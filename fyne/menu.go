package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func home() {
	myApp.Window = myApp.App.NewWindow("Groupie Tracker.")
	testButton := widget.NewButton("Enter in Groupie Tracker", func() {
		MainPage()
	})
	imageURL := "https://img.freepik.com/vecteurs-libre/illustration-vectorielle-micro-haut-parleurs-noirs-logo-promotionnel-vintage-pour-concert-festival-musique_74855-10591.jpg"
	imageContent, _ := fyne.LoadResourceFromURLString(imageURL)
	image := canvas.NewImageFromResource(imageContent)
	image.FillMode = canvas.ImageFillOriginal
	image.SetMinSize(fyne.NewSize(482, 482))
	content := container.NewVBox(image, testButton)
	centeredContent := container.NewCenter(content)
	myApp.Window.SetContent(centeredContent)
	myApp.Window.Resize(fyne.NewSize(100, 100))
	myApp.Window.CenterOnScreen()
	myApp.Window.Show()
}
