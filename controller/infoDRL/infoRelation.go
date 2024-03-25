package infoDRL //DRL = Date Relation Location
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

func RelationButton(id int) *fyne.Container {
	relation := widget.NewButton("Relation", func() {
		Relation(id)
	})
	contain := container.NewVBox(relation)
	return contain
}

func Relation(id int) {
	if buttonstatus == 0 {
		var ld string
		var stld string
		var dte string
		var dat string
		var tdate []string
		stld = "Concert : ,"
		test := container.NewGridWithColumns(1)
		for location, dates := range testmodel.Relation(id).DatesLocations {
			location = cases.Title(language.Und).String(location)
			if 1 < len(dates) {
				for _, temp := range dates {
					datetemp := Date(temp)
					tdate = append(tdate, datetemp)
				}
				dte = strings.Join(tdate, " et le ")
				dat = dte
			} else {
				dte = strings.Join(dates, " ")
				dat = Date(dte)
			}
			locationl := strings.Split(location, "_")
			locationli := strings.Join(locationl, " ")
			ld = "- " + cases.Title(language.Und).String(locationli) + " : le " + dat + ","
			stld = stld + ld
		}
		var lod []string = strings.Split(stld, ",")
		for _, val := range lod {
			contain := canvas.NewText(val, color.White)
			contain.TextSize = 24
			contain.Alignment = fyne.TextAlignLeading
			test.Add(contain)
		}
		buttonstatus = 1
	}
}
