package fyne

import (
	DataAPI "groupietracker/API" // Import the DataAPI package from the groupietracker/API directory
	"image/color"
	"sort"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	a          int  = 0
	grid            = container.NewGridWithColumns(5)
	fullScreen bool = false
	filtre     int
	filter     int
)

type AppData struct {
	App    fyne.App
	Window fyne.Window
}

func MainPage() {
	// Generates the main page of the application with a toolbar, an artists grid, and a search field.
	myApp.Window.SetFullScreen(fullScreen)

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
		loadArtistsIntoGrid(1, 53)
	case 2, 3, 4, 5:
		loadArtistsByFilter(filter)
	}

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for an artist...")

	grid2 := container.NewVScroll(grid)
	gridContent := container.NewMax(grid2)

	full := container.NewBorder(toolBar, nil, nil, nil, gridContent)
	backgroundContainer := container.NewBorder(searchEntry, nil, nil, nil, background, full)
	myApp.Window.SetContent(backgroundContainer)

	searchEntry.OnChanged = func(text string) {
		filterArtistsByText(text)
	}
}

func loadArtistsIntoGrid(start, end int) {
	// Loads artists into the grid.
	grid.RemoveAll()
	for id := start; id < end; id++ {
		grid.Add(ButtonImg(id))
	}
}

func loadArtistsByFilter(filter int) {
	// Loads artists into the grid based on a specific filter.
	grid.RemoveAll()
	artists := make([]DataAPI.DataArtist, 52)
	for i := 1; i <= 52; i++ {
		artists[i-1] = DataAPI.GetArtistByID(i)
	}
	switch filter {
	case 2:
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].CreationDate < artists[j].CreationDate
		})
	case 3:
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].CreationDate > artists[j].CreationDate
		})
	case 4:
		sort.Slice(artists, func(i, j int) bool {
			time1, _ := time.Parse("02-01-2006", artists[i].FirstAlbum)
			time2, _ := time.Parse("02-01-2006", artists[j].FirstAlbum)
			return time1.Before(time2)
		})
	case 5:
		sort.Slice(artists, func(i, j int) bool {
			time1, _ := time.Parse("02-01-2006", artists[i].FirstAlbum)
			time2, _ := time.Parse("02-01-2006", artists[j].FirstAlbum)
			return time1.After(time2)
		})
	}
	for _, artist := range artists {
		grid.Add(ButtonImg(int(artist.Id)))
	}
}

func filterArtistsByText(text string) {
	// Filters artists based on a search text.
	filteredArtists, _ := DataAPI.Search(text)
	grid.RemoveAll()
	for _, artist := range filteredArtists {
		grid.Add(ButtonImg(artist.Id))
	}
	grid.Refresh()
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
	radio := widget.NewRadioGroup([]string{"Sort by Group Creation Date. From oldest to newest", "Sort by Group Creation Date. From newest to oldest",
		"Sort by Release Date of First Album. From oldest to newest", "Sort by Release Date of First Album. From newest to oldest",
		"Sort by Default"}, func(value string) {
		switch value {
		case "Sort by Group Creation Date. From oldest to newest":
			filtre = 2
		case "Sort by Group Creation Date. From newest to oldest":
			filtre = 3
		case "Sort by Release Date of First Album. From oldest to newest":
			filtre = 4
		case "Sort by Release Date of First Album. From newest to oldest":
			filtre = 5
		case "Sort by Default":
			filtre = 1
		}
		MainPage()
	})
	return radio
}
