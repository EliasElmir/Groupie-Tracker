package view

import (
	DRL "groupietracker/controller/infoDRL"
	"groupietracker/controller/infoPage"
	testmodel "groupietracker/controller/modelController"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	containerVertical1 = container.NewVBox()
	containerVertical2 = container.NewVBox()
)

func HomePage() *fyne.Container {
	retour := widget.NewButton("Accueil", func() {
		MainPage()
		DRL.DRLINFO.RemoveAll()
	})
	contain := container.NewVBox(retour)
	return contain
}

func SecondPage(id int) {
	r, _ := fyne.LoadResourceFromURLString(testmodel.GetArtistsID(id).Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal
	img.SetMinSize(fyne.NewSize(180, 180))
	GroupeInfo := infoPage.InfoGroupe(id)
	locationButton := DRL.LocationButton(id)
	//dateButton := DRL.DateButton(id)
	//relationButton := DRL.RelationButton(id)
	homeButton := HomePage()
	containerVertical1 = container.NewVBox(img, locationButton, homeButton)
	containerVertical2 = container.NewVBox(GroupeInfo, DRL.DRLINFO)
	containerH1 := container.NewHBox(containerVertical1, containerVertical2)
	containerH1.Refresh()
	myApp.Window.SetContent(containerH1)
}
