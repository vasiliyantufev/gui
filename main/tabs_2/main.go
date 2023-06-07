package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const HEIGHT float32 = 30

func setDefaultColumnsWidth(table *widget.Table) {
	colWidths := []float32{130, 150, 160, 200, 400, 150, 250, 110, 80}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func main() {

	application := app.New()
	win := application.NewWindow("Test GUI")

	data := [][]fyne.CanvasObject{
		{widget.NewLabel("ffffffffffffff"), widget.NewHyperlink("TUTU", nil), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff")},
		{widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewHyperlink("ffffffffffffff", nil)},
		{widget.NewHyperlink("TUTU", nil), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff"), widget.NewLabel("ffffffffffffff")},
	}

	tableData := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			c := container.NewWithoutLayout()
			r := canvas.NewRectangle(color.White)
			r.SetMinSize(fyne.NewSize(0, HEIGHT))
			r.Resize(fyne.NewSize(0, HEIGHT))
			c.Add(r)
			return c
		},
		func(cell widget.TableCellID, o fyne.CanvasObject) {
			container := o.(*fyne.Container)
			var obj = data[cell.Row][cell.Col]
			container.Add(obj)
			container.Refresh()
		})

	setDefaultColumnsWidth(tableData)

	label := widget.NewLabel("My table :")
	button := widget.NewButton("Close", func() {
		application.Quit()
	})

	container := layout.NewBorderLayout(label, button, nil, nil)
	box := fyne.NewContainerWithLayout(container, label, tableData, button)
	//box := container.New(container, label, tableData, button)
	win.SetContent(box)

	win.Resize(fyne.NewSize(1500, 400))
	win.ShowAndRun()

}
