package infoDRL // DRL = Date Relation Location

import (
	"fmt"
	testmodel "groupietracker/controller/modelController"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	mois string
	v    = 0
)

func DateButton(id int) *fyne.Container {
	date := widget.NewButton("Date", func() {
		fmt.Print("Date")
		infoDate(id)
	})
	contain := container.NewVBox(date)
	return contain
}

func infoDate(id int) {
	fmt.Println(buttonstatus)
	if buttonstatus == 0 {
		var container = container.NewGridWithColumns(1)
		var listDate string
		i := 0
		for _, idate := range testmodel.Dates(id).Date {
			idate = Date(idate)
			i = i + 1
			if i == len(testmodel.Dates(id).Date) {
				listDate = listDate + idate
			} else if i == 5 {
				listDate = listDate + idate + " - " + "\n"
			} else if i == 10 {
				listDate = listDate + idate + " - " + "\n"
			} else if i == 15 {
				listDate = listDate + idate + " - " + "\n"
			} else {
				listDate = listDate + idate + " - "
			}
		}
		fmt.Println(listDate)
		var testDate []string = strings.Split(listDate, "*")
		v = 0
		for _, AllDate := range testDate {
			if v != 0 {
				DateAll := AllDate
				contain := widget.NewButton(DateAll,
					func() {
						fmt.Print(DateAll)
					},
				)
				container.Add(contain)
			} else {
				v = 1
			}
			DRLINFO.RemoveAll()
			DRLINFO.Add(container)
		}
		buttonstatus = 1
	}
}
