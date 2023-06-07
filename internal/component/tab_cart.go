package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func setDefaultColumnsWidthCart(table *widget.Table) {
	colWidths := []float32{150, 150, 150, 150, 50, 150, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func GetTabCarts(top *widget.Button, cart *widget.Button) *container.TabItem {
	dataTblCart := [][]string{
		{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED_AT", "UPDATED_AT"},
		{"name_1", "Visa", "1234567890", "Artur", "123", "01-01-2023 10:30", "01-01-2023 08:30", "01-01-2023 10:30"},
		{"name_2", "Visa", "1234567890", "Artur", "123", "01-01-2023 10:30", "01-01-2023 08:30", "01-01-2023 10:30"},
		{"name_N", "Visa", "1234567890", "Artur", "123", "01-01-2023 10:30", "01-01-2023 08:30", "01-01-2023 10:30"}}

	tableDataCart := widget.NewTable(
		func() (int, int) {
			return len(dataTblCart), len(dataTblCart[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblCart[i.Row][i.Col])
		})

	setDefaultColumnsWidthCart(tableDataCart)

	containerTblCart := layout.NewBorderLayout(top, cart, nil, nil)
	boxCart := fyne.NewContainerWithLayout(containerTblCart, top, tableDataCart, cart)

	return container.NewTabItem("Банковские карты", boxCart)
}
