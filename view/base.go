package view

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
