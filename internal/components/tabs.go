package components

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

func setDefaultColumnsWidthCart(table *widget.Table) {
	colWidths := []float32{150, 150, 150, 150, 50, 150, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func GetTabs(window fyne.Window, top *widget.Button, text *widget.Button, cart *widget.Button) *container.AppTabs {

	//buttonTop := widget.NewButton("Обновить данные", func() {
	//})
	//buttonText := widget.NewButton("Добавить текстовые данные", func() {
	//	window.SetContent(containerFormText)
	//	window.Show()
	//})
	//buttonCart := widget.NewButton("Добавить банковскую карту", func() {
	//	window.SetContent(containerFormCart)
	//	window.Show()
	//})
	//---------------------------------------------------------------------- tbl - текстовые данные
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
	//---------------------------------------------------------------------- tbl - банковские карты
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
	//---------------------------------------------------------------------- tabs
	labelPlug := widget.NewLabel("Список файлов")
	tab1 := container.NewTabItem("Текстовые данные", boxText)
	tab2 := container.NewTabItem("Банковские карты", boxCart)
	tab3 := container.NewTabItem("Файлы", labelPlug)

	return container.NewAppTabs(tab1, tab2, tab3)
}
