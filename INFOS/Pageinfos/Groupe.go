package infoPage

import (
	testmodel "groupietracker/INFOS/API"
	DRL "groupietracker/INFOS/infosgroupe"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func InfoGroupe(id int) *fyne.Container {
	infog := container.NewVBox()
	var listmember string
	var listmember2 string
	i := 0
	v := 0
	name := "Nom du groupe : " + testmodel.GetArtistsID(id).Name
	nameText := canvas.NewText(name, color.White)
	nameText.TextSize = 24
	nameText.Alignment = fyne.TextAlignLeading
	infog.Add(nameText)
	members := testmodel.GetArtistsID(id).Members
	y := len(members)
	for _, member := range members {
		if i < (y / 2) {
			listmember = listmember + member + ", "
		} else if i == (y - 1) {
			listmember2 = listmember2 + member + ". "
			v = 1
		} else if i == (y - 2) {
			listmember2 = listmember2 + member + " et "
		} else {
			listmember2 = listmember2 + member + ", "
			v = 1
		}
		i = i + 1
	}
	stringmembers := "Membres du groupe : " + listmember
	members1 := canvas.NewText(stringmembers, color.White)
	members1.TextSize = 24
	members1.Alignment = fyne.TextAlignLeading
	infog.Add(members1)
	if v == 1 {
		stringmember2 := listmember2
		members2 := canvas.NewText(stringmember2, color.White)
		members2.TextSize = 24
		members2.Alignment = fyne.TextAlignLeading
		infog.Add(members2)
	}
	creationdate := "Date de creation du groupe : " + strconv.Itoa(testmodel.GetArtistsID(id).CreationDate)
	creationDate := canvas.NewText(creationdate, color.White)
	creationDate.TextSize = 24
	creationDate.Alignment = fyne.TextAlignLeading
	infog.Add(creationDate)
	FirstAlbum := "Date de sortie premier album : " + DRL.Date(testmodel.GetArtistsID(id).FirstAlbum)
	firstAlbum := canvas.NewText(FirstAlbum, color.White)
	firstAlbum.TextSize = 24
	firstAlbum.Alignment = fyne.TextAlignLeading
	infog.Add(firstAlbum)
	return infog
}
