package Fyne

import (
	"fmt"
	testmodel "groupietracker/controller/modelController"
	"groupietracker/model"
	"sort"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	b          int  = 0
	grid            = container.New(layout.NewGridLayoutWithColumns(5))
	fullScreen bool = false
	filtre     int
	filter     int
)

func MainPage() {
	myApp.Window.SetFullScreen(fullScreen)
	if b == 0 {
		filtre = 1
		b = 2
		fmt.Print("B:")
		fmt.Println(b)
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
	case 2:
		grid.RemoveAll()
		artists := make([]*model.Artist, 52)
		for i := 1; i <= 52; i++ {
			artists[i-1] = testmodel.GetArtistsID(i)
		}
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].CreationDate < artists[j].CreationDate
		})
		for i, artist := range artists {
			fmt.Println(strconv.Itoa(i+1)+".", artist.Name, "(", artist.CreationDate, ")")
			grid.Add(ButtonImg(int(artist.Id)))
		}
	case 3:
		grid.RemoveAll()
		artists := make([]*model.Artist, 52)
		for i := 1; i <= 52; i++ {
			artists[i-1] = testmodel.GetArtistsID(i)
		}
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].CreationDate > artists[j].CreationDate
		})
		for i, artist := range artists {
			fmt.Println(strconv.Itoa(i+1)+".", artist.Name, "(", artist.CreationDate, ")")
			grid.Add(ButtonImg(int(artist.Id)))
		}
	case 4:
		grid.RemoveAll()
		artists := make([]*model.Artist, 52)
		for i := 1; i <= 52; i++ {
			artists[i-1] = testmodel.GetArtistsID(i)
		}
		sort.Slice(artists, func(i, j int) bool {
			time1, _ := time.Parse("02-01-2006", artists[i].FirstAlbum)
			time2, _ := time.Parse("02-01-2006", artists[j].FirstAlbum)
			return time1.Year() < time2.Year()
		})
		for i, artist := range artists {
			fmt.Println(strconv.Itoa(i+1)+".", artist.Name, "(", artist.FirstAlbum, ")")
			grid.Add(ButtonImg(int(artist.Id)))
		}
	case 5:
		grid.RemoveAll()
		artists := make([]*model.Artist, 52)
		for i := 1; i <= 52; i++ {
			artists[i-1] = testmodel.GetArtistsID(i)
		}
		sort.Slice(artists, func(i, j int) bool {
			time1, _ := time.Parse("02-01-2006", artists[i].FirstAlbum)
			time2, _ := time.Parse("02-01-2006", artists[j].FirstAlbum)
			return time1.Year() > time2.Year()
		})
		for i, artist := range artists {
			fmt.Println(strconv.Itoa(i+1)+".", artist.Name, "(", artist.FirstAlbum, ")")
			grid.Add(ButtonImg(int(artist.Id)))
		}
	}
	grid2 := container.NewVScroll(grid)
	grid2.Refresh()
	full := container.NewBorder(toolBar, nil, nil, nil, grid2)
	myApp.Window.SetContent(full)
}

func ButtonImg(id int) *fyne.Container {
	r, _ := fyne.LoadResourceFromURLString(testmodel.GetArtistsID(id).Image)
	img := canvas.NewImageFromResource(r)
	img.SetMinSize(fyne.NewSize(300, 300))
	btn := widget.NewButton(" ", func() {
		SecondPage(id)
	})
	container1 := container.New(
		layout.NewMaxLayout(),
		btn,
		widget.NewCard(testmodel.GetArtistsID(id).Name, "", img),
	)
	return container1
}

func FullScreen(window fyne.Window, verifZoombool bool) {
	if verifZoombool {
		fullScreen = false
		window.SetFullScreen(fullScreen)
	} else {
		fullScreen = true
		window.SetFullScreen(fullScreen)
	}
}

func Filter() *widget.RadioGroup {
	var value string
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
	value = value
	return radio
}
