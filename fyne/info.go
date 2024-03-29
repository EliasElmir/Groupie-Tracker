package fyne

import (
	DataAPI "groupietracker/API"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	containerVertical1 = container.NewVBox()
	containerVertical2 = container.NewVBox()
	DRLINFO            = container.NewGridWithColumns(2)
	buttonstatus       = 0
)

func HomePage() *fyne.Container {
	retour := widget.NewButton("Accueil", func() {
		MainPage()
		DRLINFO.RemoveAll()
	})
	contain := container.NewVBox(retour)
	return contain
}

func SecondPage(id int) {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")
	r, _ := fyne.LoadResourceFromURLString(DataAPI.GetArtistByID(id).Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal
	img.Resize((fyne.NewSize(200, 200)))
	GroupeInfo := InfoGroupe(id)
	locationButton := LocationButton(id)
	txt := canvas.NewText("==================================", color.Transparent)
	//dateButton := DRL.DateButton(id)
	//relationButton := DRL.RelationButton(id)
	homeButton := HomePage()
	containerVertical1 = container.NewVBox(searchEntry, img, locationButton, homeButton)
	containerVertical2 = container.NewVBox(GroupeInfo, DRLINFO)
	containerH1 := container.NewHBox(containerVertical1, txt, containerVertical2)
	containerH1.Refresh()

	//centeredContainer := container.NewCenter(containerH1)
	myApp.Window.SetContent(containerH1)
}

func LocationButton(id int) *fyne.Container {
	Concert := widget.NewButton("Concert", func() {
		infoDate(id)
	})
	contain := container.NewVBox(Concert)
	return contain
}

func InfoGroupe(id int) *fyne.Container {
	infog := container.NewVBox()
	var listmember string
	var listmember2 string
	i := 0
	v := 0
	name := "Nom du groupe : " + DataAPI.GetArtistByID(id).Name
	nameText := canvas.NewText(name, color.White)
	nameText.TextSize = 24
	nameText.Alignment = fyne.TextAlignLeading
	infog.Add(nameText)
	members := DataAPI.GetArtistByID(id).Members
	y := len(members)
	for _, member := range members {
		if i < (y / 2) {
			listmember = listmember + member + ", "
		} else if i == (y - 1) {
			listmember2 = listmember2 + member + ". "
			v = 1
		} else if i == (y - 2) {
			listmember2 = listmember2 + member + " et "
		} else {
			listmember2 = listmember2 + member + ", "
			v = 1
		}
		i = i + 1
	}
	stringmembers := "Membres du groupe : " + listmember
	members1 := canvas.NewText(stringmembers, color.White)
	members1.TextSize = 24
	members1.Alignment = fyne.TextAlignLeading
	infog.Add(members1)
	if v == 1 {
		stringmember2 := listmember2
		members2 := canvas.NewText(stringmember2, color.White)
		members2.TextSize = 24
		members2.Alignment = fyne.TextAlignLeading
		infog.Add(members2)
	}
	creationdate := "Date de creation du groupe : " + strconv.Itoa(DataAPI.GetArtistByID(id).CreationDate)
	creationDate := canvas.NewText(creationdate, color.White)
	creationDate.TextSize = 24
	creationDate.Alignment = fyne.TextAlignLeading
	infog.Add(creationDate)
	FirstAlbum := "Date de sortie premier album : " + Date(DataAPI.GetArtistByID(id).FirstAlbum)
	firstAlbum := canvas.NewText(FirstAlbum, color.White)
	firstAlbum.TextSize = 24
	firstAlbum.Alignment = fyne.TextAlignLeading
	infog.Add(firstAlbum)
	return infog
}

func infoDate(id int) {
	var v int
	if buttonstatus == 0 {
		var container = container.NewGridWithColumns(1)
		var listDate string
		i := 0
		for _, idate := range DataAPI.GetDateByID(id).Dates {
			idate = Date(idate)
			i = i + 1
			if i == len(DataAPI.GetDateByID(id).Dates) {
				listDate = listDate + idate
			} else if i == 5 {
				listDate = listDate + idate + " - " + "\n"
			} else if i == 10 {
				listDate = listDate + idate + " - " + "\n"
			} else if i == 15 {
				listDate = listDate + idate + " - " + "\n"
			} else {
				listDate = listDate + idate + " - "
			}
		}
		var testDate []string = strings.Split(listDate, "*")
		v = 0
		for _, AllDate := range testDate {
			if v != 0 {
				DateAll := AllDate
				contain := widget.NewButton(DateAll,
					func() {
					},
				)
				container.Add(contain)
			} else {
				v = 1
			}
			DRLINFO.RemoveAll()
			DRLINFO.Add(container)
		}
		buttonstatus = 1
	}
}

func Date(date string) string {
	mois := ""
	dstring := strings.Split(date, "-")
	switch dstring[1] {
	case "01":
		mois = "janvier"
	case "02":
		mois = "février"
	case "03":
		mois = "mars"
	case "04":
		mois = "avril"
	case "05":
		mois = "mai"
	case "06":
		mois = "juin"
	case "07":
		mois = "juillet"
	case "08":
		mois = "août"
	case "09":
		mois = "septembre"
	case "10":
		mois = "octobre"
	case "11":
		mois = "novembre"
	case "12":
		mois = "décembre"
	}
	dstring[1] = mois
	ds := strings.Join(dstring, " ")
	return ds
}
