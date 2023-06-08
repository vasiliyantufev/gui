package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{
	[]string{"top left", "top right"},
	[]string{"bottom left", "bottom right"}}

func setDefaultColumnsWidthCart(table *widget.Table) {
	colWidths := []float32{150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")
	myWindow.Resize(fyne.NewSize(400, 200))

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	setDefaultColumnsWidthCart(list)
	var index float32
	index = 2

	var c *fyne.Container

	buttonTop := widget.NewButton("Обновить данные", func() {
		index++
		data = append(data, []string{"top left", "top right"})

		list.Resize(fyne.NewSize(float32(len(data)), float32(len(data[0]))))
		list.Refresh()
		myWindow.SetContent(c)
	})

	c = container.NewVBox(list, buttonTop)

	myWindow.SetContent(c)
	myWindow.ShowAndRun()
}
