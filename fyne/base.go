package fyne

import (
	DataAPI "groupietracker/API"
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
		grid.RemoveAll()
		for id := 1; id < 53; id++ {
			grid.Add(ButtonImg(id))
		}
	case 2, 3, 4, 5:
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
				return time1.Year() < time2.Year()
			})
		case 5:
			sort.Slice(artists, func(i, j int) bool {
				time1, _ := time.Parse("02-01-2006", artists[i].FirstAlbum)
				time2, _ := time.Parse("02-01-2006", artists[j].FirstAlbum)
				return time1.Year() > time2.Year()
			})
		}
		for _, artist := range artists {
			grid.Add(ButtonImg(int(artist.Id)))
		}
	}
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")
	grid2 := container.NewVScroll(grid)
	grid2.Refresh()
	gridContent := container.NewMax()
	gridContent.Add(grid2)

	full := container.NewBorder(toolBar, nil, nil, nil, gridContent)
	backgroundContainer := container.NewBorder(searchEntry, nil, nil, nil, background, full)
	myApp.Window.SetContent(backgroundContainer)

	searchEntry.OnChanged = func(text string) {
		filteredArtists, _ := DataAPI.Search(text)
		grid.RemoveAll()
		for _, artist := range filteredArtists {
			grid.Add(ButtonImg(artist.Id))
		}
		grid.Refresh()
	}
}

func ButtonImg(id int) *fyne.Container {
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
	fullScreen = !verifZoombool
	window.SetFullScreen(fullScreen)
}

func Filter() *widget.RadioGroup {
	radio := widget.NewRadioGroup([]string{"Trier par Date de creation du groupe. Du plus ancien au plus récent", "Trier par Date de creation du groupe. Du plus récent au plus ancien",
		"Trier par Date de sortie du premier album. Du plus ancien au plus récent", "Trier par Date de sortie du premier album. Du plus récent au plus ancien",
		"Trier par Default"}, func(value string) {
		switch value {
		case "Trier par Date de creation du groupe. Du plus ancien au plus récent":
			filtre = 2
		case "Trier par Date de creation du groupe. Du plus récent au plus ancien":
			filtre = 3
		case "Trier par Date de sortie du premier album. Du plus ancien au plus récent":
			filtre = 4
		case "Trier par Date de sortie du premier album. Du plus récent au plus ancien":
			filtre = 5
		case "Trier par Default":
			filtre = 1
		}
		MainPage()
	})
	return radio
}
