package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func home() {
	// Creates a new window with the title "FestivaSync"
	// Creates a button "Enter in FestivaSync" which, when clicked, calls the MainPage() function.
	// Loads an image from a URL and places it in a canvas.Image.
	// Creates a vertical container for the image and the button.
	// Centers the vertical content in a container.
	// Sets the window content to be the centered content.
	myApp.Window = myApp.App.NewWindow("FestivaSync")
	testButton := widget.NewButton("Enter in FestivaSync", func() {
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
}
