package fyne

import (
	testmodel "groupietracker/INFOS/API"
	infoPage "groupietracker/INFOS/Pageinfos"
	DRL "groupietracker/INFOS/infosgroupe"
	"image/color"

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
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")
	r, _ := fyne.LoadResourceFromURLString(testmodel.GetArtistsID(id).Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal
	img.Resize((fyne.NewSize(200, 200)))
	GroupeInfo := infoPage.InfoGroupe(id)
	locationButton := DRL.LocationButton(id)
	txt := canvas.NewText("==================================", color.Transparent)
	//dateButton := DRL.DateButton(id)
	//relationButton := DRL.RelationButton(id)
	homeButton := HomePage()
	containerVertical1 = container.NewVBox(searchEntry, img, locationButton, homeButton)
	containerVertical2 = container.NewVBox(GroupeInfo, DRL.DRLINFO)
	containerH1 := container.NewHBox(containerVertical1, txt, containerVertical2)
	containerH1.Refresh()

	//centeredContainer := container.NewCenter(containerH1)
	myApp.Window.SetContent(containerH1)
}
