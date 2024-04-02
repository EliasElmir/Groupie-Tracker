package fyne

import (
	"fmt"
	DataAPI "groupietracker/API"
	"image/color"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
	txt := canvas.NewText("========================================================================================================", color.Transparent)
	img.Resize((fyne.NewSize(400, 400)))

	GroupeInfo := InfoGroupe(id)
	locationButton := LocationButton(id)

	homeButton := HomePage()

	previousbutton := widget.NewButton("previous", func() {
		if id != 1 {
			SecondPage(id - 1)
		}
	})

	nextButton := widget.NewButton("next", func() {
		if id != 52 {
			SecondPage(id + 1)
		}
	})

	nxtAndpreButtons := container.NewHBox(previousbutton, layout.NewSpacer(), nextButton)
	image := container.NewHBox(txt, img)

	content := container.NewVBox(txt, txt, txt, txt, image, txt, txt, txt, txt, txt, txt, txt, txt, txt, GroupeInfo, DRLINFO)

	content2 := container.NewVBox(homeButton, locationButton)
	centeredContainer2 := container.NewCenter(content2)

	finalcontent := container.NewVBox(content, nxtAndpreButtons, centeredContainer2)
	finalcontent.Refresh()

	myApp.Window.SetContent(finalcontent)
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
		locationList := DataAPI.Location(id).Location
		for _, location := range locationList {
			place := strings.Title(strings.Replace(location, "_", " ", -1))
			contain := widget.NewButton(place, func() {
				fmt.Print(place)
				DateLocation(id, place)
			})
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

	locat := strings.ToLower(strings.Replace(locate, " ", "_", -1))
	dates, ok := DataAPI.Relation(id).DatesLocations[locat]
	if !ok {
		return
	}
	for _, temp := range dates {
		datetemp := Date(temp)
		datetemp = " -" + datetemp
		contain := canvas.NewText(datetemp, color.White)
		contain.TextSize = 24
		contain.Alignment = fyne.TextAlignLeading
		DRLINFO.Add(contain)
	}

	btn := widget.NewButton("Retour", func() {
		buttonstatus = 0
		Location(id)
	})
	DRLINFO.Add(btn)
}

func Date(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date
	}
	return t.Format(" - 2 janvier 2006")
}

func InfoGroupe(id int) *fyne.Container {
	infog := container.NewVBox()

	name := "Nom du groupe : " + DataAPI.GetArtistByID(id).Name
	nameText := canvas.NewText(name, color.White)
	nameText.TextSize = 24
	nameText.Alignment = fyne.TextAlignCenter
	infog.Add(nameText)

	members := DataAPI.GetArtistByID(id).Members
	var stringmembers string
	for i, member := range members {
		if i < len(members)/2 {
			stringmembers += member + ", "
		} else if i == len(members)-1 {
			stringmembers += member + ". "
		} else if i == len(members)-2 {
			stringmembers += member + " et "
		} else {
			stringmembers += member + ", "
		}
	}
	membersText := canvas.NewText("Membres du groupe : "+stringmembers, color.White)
	membersText.TextSize = 24
	membersText.Alignment = fyne.TextAlignCenter
	infog.Add(membersText)

	creationdate := "Date de création du groupe : " + strconv.Itoa(DataAPI.GetArtistByID(id).CreationDate)
	creationDate := canvas.NewText(creationdate, color.White)
	creationDate.TextSize = 24
	creationDate.Alignment = fyne.TextAlignCenter
	infog.Add(creationDate)

	firstAlbumDate := "Date de sortie du premier album : " + Date(DataAPI.GetArtistByID(id).FirstAlbum)
	firstAlbumDateText := canvas.NewText(firstAlbumDate, color.White)
	firstAlbumDateText.TextSize = 24
	firstAlbumDateText.Alignment = fyne.TextAlignCenter
	infog.Add(firstAlbumDateText)

	return infog
}
