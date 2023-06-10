package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var slice1 = [][]string{
	{"top left", "top right"},
	{"bottom left", "bottom right"},
	{"left", "right"}}

var slice2 = [][]string{
	{"bottom left", "bottom right"},
	{"top left", "top right"}}

var slice3 = [][]string{
	{"left", "right"},
	{"top", "bbb"},
	{"aaa", "kkk"},
	{"kkk", "aaa"},
	{"top", "bbb"},
	{"aaa", "kkk"},
	{"kkk", "aaa"}}

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

	slice1 = slice2

	list := widget.NewTable(
		func() (int, int) {
			return len(slice1), len(slice1[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(slice1[i.Row][i.Col])
		})

	setDefaultColumnsWidthCart(list)
	var index float32
	index = 2

	var c *fyne.Container

	buttonTop := widget.NewButton("Обновить данные", func() {
		index++

		slice1 = slice3

		list.Refresh()
		myWindow.SetContent(c)
	})

	c = container.NewVBox(list, buttonTop)

	myWindow.SetContent(c)
	myWindow.ShowAndRun()
}
