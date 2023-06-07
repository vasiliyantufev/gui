package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func setDefaultColumnsWidthText(table *widget.Table) {
	colWidths := []float32{150, 500, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func GetTabTexts(top *widget.Button, text *widget.Button) *container.TabItem {
	dataTblText := [][]string{
		{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"},
		{"name_1", "text", "text", "01-01-2023 08:30", "01-01-2023 10:30"},
		{"name_2", "text", "text", "01-01-2023 08:30", "01-01-2023 10:30"},
		{"name_N", "text", "text", "01-01-2023 08:30", "01-01-2023 10:30"}}

	tableData := widget.NewTable(
		func() (int, int) {
			return len(dataTblText), len(dataTblText[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblText[i.Row][i.Col])
		})
	setDefaultColumnsWidthText(tableData)

	containerTblText := layout.NewBorderLayout(top, text, nil, nil)
	boxText := fyne.NewContainerWithLayout(containerTblText, top, tableData, text)
	return container.NewTabItem("Текстовые данные", boxText)
}
