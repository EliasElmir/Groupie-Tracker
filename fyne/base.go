package fyne

import (
	DataAPI "groupietracker/API" // Import the DataAPI package from the groupietracker/API directory
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	a             int = 0
	grid              = container.NewGridWithColumns(5)
	cachedArtists DataAPI.Artist
	fullScreen    bool = false
	filtre        int
	filter        int
)

type AppData struct {
	App    fyne.App
	Window fyne.Window
}

func MainPage() {
	// Generates the main page of the application with a toolbar, an artists grid, and a search field.
	myApp.Window.SetFullScreen(fullScreen)
	var backgroundContainer *fyne.Container
	background := canvas.NewRectangle(color.White)

	if a == 0 {
		filtre = 1
		a = 2
	}
	filter = filtre

	toolBar := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.MenuExpandIcon(), func() {
				myApp.Window.SetContent(Filter())
			},
		),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(
			theme.ViewFullScreenIcon(), func() {
				FullScreen(myApp.Window, fullScreen)
			},
		),
		widget.NewToolbarAction(
			theme.CancelIcon(), func() {
				myApp.Window.Close()
			},
		),
	)

	switch filter {
	case 1:
		artist, _ := DataAPI.GetArtistData(true)
		grid = loadArtistsIntoGrid(artist)
	case 2, 3, 4, 5:
		loadArtistsByFilter(filter)
	}

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for an artist...")

	grid2 := container.NewVScroll(grid)

	searchEntry.OnChanged = func(text string) {
		if len(cachedArtists) == 0 {
			cachedArtists, _ = DataAPI.GetArtistData(true)
		}

		filteredArtists := filterArtistsByText(text)

		grid = loadArtistsIntoGrid(filteredArtists)
		backgroundContainer.Objects[1] = container.NewVScroll(grid)
		myApp.Window.SetContent(backgroundContainer)
		myApp.Window.Canvas().Refresh(backgroundContainer)
	}

	backgroundContainer = container.NewBorder(container.NewVBox(searchEntry, toolBar), nil, nil, nil, background, grid2)
	backgroundContainer.Refresh()
	myApp.Window.SetContent(backgroundContainer)
}

func loadArtistsIntoGrid(artists DataAPI.Artist) *fyne.Container {
	var artistCards []fyne.CanvasObject

	for _, artist := range artists {
		content := ButtonImg(artist.Id)

		card := container.NewVBox(content)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(5, artistCards...)

	return grid
}

func loadArtistsByFilter(filter int) {
	// Loads artists into the grid based on a specific filter.
	grid.RemoveAll()
	artists, _ := DataAPI.GetArtistData(true)

	switch filter {
	case 2:
		grid = loadArtistsIntoGrid(DataAPI.SortByCreationDateAscending(artists))
	case 3:
		grid = loadArtistsIntoGrid(DataAPI.SortByCreationDateAndFirstAlbum(artists))
	case 4:
		grid = loadArtistsIntoGrid(DataAPI.SortByNumberOfMembersDescending(artists))

	case 1:
		for _, artist := range artists {
			grid.Add(ButtonImg(int(artist.Id)))
		}
	}
}

func filterArtistsByText(text string) []DataAPI.DataArtist {
	// Vérifie d'abord si la recherche est vide
	if text == "" {
		return cachedArtists
	}

	// Effectue la recherche dans les données cache
	var filteredArtists []DataAPI.DataArtist
	for _, artist := range cachedArtists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(text)) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

func ButtonImg(id int) *fyne.Container {
	// Creates a container for a button representing an artist with their name and an image.
	r, _ := fyne.LoadResourceFromURLString(DataAPI.GetArtistByID(id).Image)
	img := canvas.NewImageFromResource(r)
	img.SetMinSize(fyne.NewSize(300, 300))
	btn := widget.NewButton(" ", func() {
		SecondPage(id)
	})
	container1 := container.New(
		layout.NewMaxLayout(),
		btn,
		widget.NewCard(DataAPI.GetArtistByID(id).Name, "", img),
	)
	return container1
}

func FullScreen(window fyne.Window, verifZoombool bool) {
	// Enables or disables fullscreen mode of the window.
	fullScreen = !verifZoombool
	window.SetFullScreen(fullScreen)
}

func Filter() *widget.RadioGroup {
	// Creates a radio button group to select different types of filters.
	radio := widget.NewRadioGroup([]string{"Sort by Group Creation Date.",
		"Sort by Release Date of First Album.", "Sort by number of members.",
		"Sort by Default"}, func(value string) {
		switch value {
		case "Sort by Group Creation Date.":
			filtre = 2
		case "Sort by Release Date of First Album.":
			filtre = 3
		case "Sort by number of members.":
			filtre = 4
		case "Sort by Default":
			filtre = 1
		}
		MainPage()
	})
	return radio

}
