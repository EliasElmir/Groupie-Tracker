package fyne

import (
	"fmt"
	DataAPI "groupietracker/API" // Import the DataAPI package from the groupietracker/API directory
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
	DRLINFO      = container.NewGridWithColumns(2) // Declares a grid with 2 columns
	buttonstatus = 0                               // Initializes a variable to track button state
)

func HomePage() *fyne.Container {
	// Creates a "Home" button to return to the home page
	retour := widget.NewButton("Home", func() {
		MainPage()
		DRLINFO.RemoveAll()
	})
	contain := container.NewVBox(retour) // Creates a vertical container for the "Home" button
	return contain
}

func SecondPage(id int) {
	// Creates widgets and buttons to display group details and navigate between pages
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for an artist...")
	r, _ := fyne.LoadResourceFromURLString(DataAPI.GetArtistByID(id).Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal
	txt := canvas.NewText("========================================================================================================", color.Transparent)
	img.Resize((fyne.NewSize(400, 400)))

	GroupeInfo := InfoGroupe(id)
	locationButton := LocationButton(id)

	homeButton := HomePage()

	previousbutton := widget.NewButton("PREVIOUS", func() {
		if id != 1 {
			SecondPage(id - 1)
		}
	})

	nextButton := widget.NewButton("NEXT", func() {
		if id != 52 {
			SecondPage(id + 1)
		}

	})

	favoritesButton := widget.NewButton("Add to Favorites", func() {
		// DataAPI.AddFavorite(id)
	})

	nxtAndpreButtons := container.NewHBox(previousbutton, layout.NewSpacer(), nextButton)
	image := container.NewHBox(txt, img)

	content := container.NewVBox(txt, txt, txt, txt, image, txt, txt, txt, txt, txt, txt, txt, txt, txt, GroupeInfo, DRLINFO, favoritesButton)

	content2 := container.NewVBox(homeButton, locationButton, favoritesButton)
	centeredContainer2 := container.NewCenter(content2)

	finalcontent := container.NewVBox(content, nxtAndpreButtons, centeredContainer2)
	finalcontent.Refresh()

	myApp.Window.SetContent(finalcontent)
}

func LocationButton(id int) *fyne.Container {
	// Creates a button to display concert locations
	Concert := widget.NewButton("Concert", func() {
		fmt.Print("location")
		Location(id)
	})
	contain := container.NewVBox(Concert)
	return contain
}

func Location(id int) {
	// Displays concert locations for a given group
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
	// Displays concert dates for a given location
	DRLINFO.RemoveAll()
	text := "In concert at " + locate + " : "
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

	btn := widget.NewButton("Return", func() {
		buttonstatus = 0
		Location(id)
	})
	DRLINFO.Add(btn)
}

func Date(date string) string {
	// Converts date from "YYYY-MM-DD" format to "DD month YYYY"
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date
	}
	return t.Format(" - 2 January 2006")
}

func InfoGroupe(id int) *fyne.Container {
	// Creates a container with group information
	infog := container.NewVBox()

	name := "Group Name: " + DataAPI.GetArtistByID(id).Name
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
			stringmembers += member + " and "
		} else {
			stringmembers += member + ", "
		}
	}
	membersText := canvas.NewText("Group Members: "+stringmembers, color.White)
	membersText.TextSize = 24
	membersText.Alignment = fyne.TextAlignCenter
	infog.Add(membersText)

	creationdate := "Group Creation Date: " + strconv.Itoa(DataAPI.GetArtistByID(id).CreationDate)
	creationDate := canvas.NewText(creationdate, color.White)
	creationDate.TextSize = 24
	creationDate.Alignment = fyne.TextAlignCenter
	infog.Add(creationDate)

	firstAlbumDate := "Release Date of First Album: " + Date(DataAPI.GetArtistByID(id).FirstAlbum)
	firstAlbumDateText := canvas.NewText(firstAlbumDate, color.White)
	firstAlbumDateText.TextSize = 24
	firstAlbumDateText.Alignment = fyne.TextAlignCenter
	infog.Add(firstAlbumDateText)

	return infog
}
