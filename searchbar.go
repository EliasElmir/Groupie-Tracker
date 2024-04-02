package groupietracker

import (
	"fmt"
	"groupietracker/DataAPI"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func searchBar() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Data")

	data := binding.BindStringList(
		&[]string{},
	)
	data2, err := DataAPI.GetArtistData(false)
	if err != nil {
		fmt.Println(err)
	}

	entry := widget.NewEntry()
	entry.OnChanged = func(s string) {
		data.Set(getDataSuggestion(s, data2))
	}

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	list.OnSelected = func(id widget.ListItemID) {
		fmt.Println(id)
	}

	myWindow.SetContent(container.NewBorder(entry, nil, nil, nil, list))
	myWindow.ShowAndRun()
}

// getDataSuggestion get the suggestions for the user
func getDataSuggestion(s string, data DataAPI.Artist) []string {
	var suggestions []string
	s = strings.ToLower(s)
	for i := 0; i < len(data); i++ {
		if len(s) >= 2 {
			if strings.Contains(strings.ToLower(data[i].Name), s) {
				suggestions = append(suggestions, data[i].Name+" - Artist/Band")
			}
			if strings.Contains(strconv.Itoa(data[i].CreationDate), s) {
				suggestions = append(suggestions, data[i].Name+" - "+strconv.Itoa(data[i].CreationDate)+" - Creation Date")
			}
			if strings.Contains(data[i].FirstAlbum, s) {
				suggestions = append(suggestions, data[i].Name+" - "+data[i].FirstAlbum+" - First Album")
			}
		}

		for _, member := range data[i].Members {
			if len(s) > 0 && strings.Contains(strings.ToLower(member), s) {
				suggestions = append(suggestions, member+" - Member")
			}
		}
	}

	return suggestions
}
