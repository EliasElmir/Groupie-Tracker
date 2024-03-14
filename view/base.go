package view

import (
	"fmt"
	testmodel "groupietracker/controller/modelController"
	"groupietracker/model"
	"sort"
	"strconv"
	"time"

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
