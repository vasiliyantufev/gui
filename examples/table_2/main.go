package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Resize Example")

	// Creating the table widget
	table := widget.NewTable(
		func() (int, int) {
			return 5, 5 // create a 5x5 grid
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("") // create a label widget for each cell
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
		},
	)

	// Adding the table widget to a container
	container := fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithRows(1),
		table,
	)

	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
