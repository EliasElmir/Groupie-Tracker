package infosgroupe

import (
	testmodel "groupietracker/INFOS/API"
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
		infoDate(id)
	})
	contain := container.NewVBox(date)
	return contain
}

func infoDate(id int) {
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
		var testDate []string = strings.Split(listDate, "*")
		v = 0
		for _, AllDate := range testDate {
			if v != 0 {
				DateAll := AllDate
				contain := widget.NewButton(DateAll,
					func() {
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

func Date(date string) string {
	dstring := strings.Split(date, "-")
	switch dstring[1] {
	case "01":
		mois = "janvier"
	case "02":
		mois = "février"
	case "03":
		mois = "mars"
	case "04":
		mois = "avril"
	case "05":
		mois = "mai"
	case "06":
		mois = "juin"
	case "07":
		mois = "juillet"
	case "08":
		mois = "août"
	case "09":
		mois = "septembre"
	case "10":
		mois = "octobre"
	case "11":
		mois = "novembre"
	case "12":
		mois = "décembre"
	}
	dstring[1] = mois
	ds := strings.Join(dstring, " ")
	return ds
}
