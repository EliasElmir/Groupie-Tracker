package infoDRL // DRL = Date Relation Location

import (
	testmodel "groupietracker/controller/modelController"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	buttonstatus = 0
	DRLINFO      = container.NewGridWithColumns(2)
)

func LocationButton(id int) *fyne.Container {
	Concert := widget.NewButton("Concert", func() {
		Location(id)
	})
	contain := container.NewVBox(Concert)
	return contain
}

func Location(id int) {
	DRLINFO.RemoveAll()
	if buttonstatus == 0 {
		var loca string
		var list string
		fr := 0
		for _, place := range testmodel.Location(id).Location {
			locationl := strings.Split(place, "_")
			locationli := strings.Join(locationl, " ")
			loca = cases.Title(language.Und).String(locationli)
			fr = fr + 1
			if fr == len(testmodel.Location(id).Location) {
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
	for location, dates := range testmodel.Relation(id).DatesLocations {
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
