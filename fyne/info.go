package fyne

import (
	"fmt"
	DataAPI "groupietracker/API"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	DRLINFO      = container.NewGridWithColumns(2)
	buttonstatus = 0
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
	txt := canvas.NewText("=======================", color.Transparent)
	img.Resize((fyne.NewSize(400, 400)))
	image := container.NewHBox(txt, img)

	GroupeInfo := InfoGroupe(id)
	locationButton := LocationButton(id)

	homeButton := HomePage()
	content := container.NewVBox(image, txt, txt, txt, txt, txt, txt, txt, txt, txt, GroupeInfo, DRLINFO, homeButton, locationButton)
	content.Refresh()

	centeredContainer := container.NewCenter(content)
	myApp.Window.SetContent(centeredContainer)

}

func LocationButton(id int) *fyne.Container {
	Concert := widget.NewButton("Concert", func() {
		fmt.Print("location")
		Location(id)
	})
	contain := container.NewVBox(Concert)
	return contain
}

func Location(id int) {
	DRLINFO.RemoveAll()
	fmt.Print(buttonstatus)
	if buttonstatus == 0 {
		var loca string
		var list string
		fr := 0
		for _, place := range DataAPI.Location(id).Location {
			locationl := strings.Split(place, "_")
			locationli := strings.Join(locationl, " ")
			loca = cases.Title(language.Und).String(locationli)
			fr = fr + 1
			if fr == len(DataAPI.Location(id).Location) {
				list = list + loca
			} else {
				list = list + loca + ","
			}
		}
		var testLoca []string = strings.Split(list, ",")
		for _, varloca := range testLoca {
			place := varloca
			contain := widget.NewButton(place,
				func() {
					fmt.Print(place)
					DateLocation(id, place)
				},
			)
			DRLINFO.Add(contain)
		}
		buttonstatus = 1
	} else {
		DRLINFO.RemoveAll()
		buttonstatus = 0
	}
}

func DateLocation(id int, locate string) {
	DRLINFO.RemoveAll()
	text := "En concert à " + locate + " : "
	contain := canvas.NewText(text, color.White)
	contain.TextSize = 24
	DRLINFO.Add(contain)
	contain = canvas.NewText(" ", color.White)
	contain.TextSize = 24
	DRLINFO.Add(contain)
	varloca := strings.Split(locate, " ")
	locat := strings.Join(varloca, "_")
	ocat := strings.ToLower(locat)
	for location, dates := range DataAPI.Relation(id).DatesLocations {
		if ocat == location {
			for _, temp := range dates {
				datetemp := Date(temp)
				datetemp = " -" + datetemp
				contain := canvas.NewText(datetemp, color.White)
				contain.TextSize = 24
				contain.Alignment = fyne.TextAlignLeading
				DRLINFO.Add(contain)
			}
		}
	}
	btn := widget.NewButton("Retour", func() {
		buttonstatus = 0
		Location(id)
	})
	DRLINFO.Add(btn)
}

func Date(date string) string {
	var mois string
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
