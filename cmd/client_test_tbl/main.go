package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	application := app.New()
	window := application.NewWindow("GophKeeper")
	application.Settings().SetTheme(theme.LightTheme())
	window.Resize(fyne.NewSize(400, 200))

	var index = 1
	//---------------------------------------------------------------------- slices
	var dataTblText = [][]string{[]string{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"}}
	//---------------------------------------------------------------------- containers
	var containerTest *fyne.Container
	//---------------------------------------------------------------------- buttons
	var buttonTop *widget.Button
	//----------------------------------------------------------------------
	var tblText *widget.Table
	//---------------------------------------------------------------------- buttons event
	buttonTop = widget.NewButton("Обновить данные", func() {
		index++
		dataTblText = append(dataTblText, []string{"name_" + strconv.Itoa(index), "data", "description", "01/02/2006", "01/02/2006"})
		tblText.Resize(fyne.NewSize(float32(len(dataTblText)), float32(len(dataTblText[0]))))

		log.Print(dataTblText)

		tblText.Refresh()
		window.SetContent(containerTest)
	})
	//---------------------------------------------------------------------- tbl
	tblText = widget.NewTable(
		func() (int, int) {
			return len(dataTblText), len(dataTblText[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblText[i.Row][i.Col])
		})
	SetDefaultColumnsWidthText(tblText)

	containerTest = container.NewVBox(tblText, buttonTop)
	//----------------------------------------------------------------------
	window.SetContent(containerTest)
	window.ShowAndRun()
}

func SetDefaultColumnsWidthText(table *widget.Table) {
	colWidths := []float32{150, 500, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}
