package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const HEIGHT float32 = 30

var user = "user"
var password = "password"

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

func main() {
	application := app.New()

	application.Settings().SetTheme(theme.LightTheme())

	window := application.NewWindow("Auth")
	window.Resize(fyne.NewSize(250, 80))

	var contentA *fyne.Container
	var contentB *fyne.Container
	var contentC *fyne.Container
	var contentD *fyne.Container
	var contentE *fyne.Container

	separator := widget.NewSeparator()

	//----------------------------------------------------------------------
	UsernameRegistration := widget.NewEntry()
	PasswordRegistration := widget.NewPasswordEntry()
	NewPasswordEntryRegistration := widget.NewPasswordEntry()
	formRegistration := widget.NewForm(
		widget.NewFormItem("Username", UsernameRegistration),
		widget.NewFormItem("Password", PasswordRegistration),
		widget.NewFormItem("Confirm password", NewPasswordEntryRegistration),
	)
	//----------------------------------------------------------------------
	UsernameLogin := widget.NewEntry()
	PasswordLogin := widget.NewPasswordEntry()
	formLogin := widget.NewForm(
		widget.NewFormItem("Username", UsernameLogin),
		widget.NewFormItem("Password", PasswordLogin),
	)
	//----------------------------------------------------------------------
	TextName := widget.NewEntry()
	Text := widget.NewEntry()
	TextDescription := widget.NewEntry()
	formText := widget.NewForm(
		widget.NewFormItem("Name", TextName),
		widget.NewFormItem("Text", Text),
		widget.NewFormItem("Description", TextDescription),
	)
	//----------------------------------------------------------------------
	CartName := widget.NewEntry()
	PaymentSystem := widget.NewEntry()
	Number := widget.NewEntry()
	Holder := widget.NewEntry()
	EndDate := widget.NewEntry()
	CVC := widget.NewEntry()
	formCart := widget.NewForm(
		widget.NewFormItem("Name", CartName),
		widget.NewFormItem("Payment system", PaymentSystem),
		widget.NewFormItem("Number", Number),
		widget.NewFormItem("Holder", Holder),
		widget.NewFormItem("End date", EndDate),
		widget.NewFormItem("CVC", CVC),
	)
	//----------------------------------------------------------------------
	options := []string{"Login", "Registration"}
	radio := widget.NewRadioGroup(options, func(value string) {
		log.Println("Radio set to ", value)
		if value == "Login" {
			window.SetContent(contentB)
			window.Resize(fyne.NewSize(500, 100))
			window.Show()
		}
		if value == "Registration" {
			window.SetContent(contentC)
			window.Resize(fyne.NewSize(500, 100))
			window.Show()
		}
	})
	//----------------------------------------------------------------------
	buttonTop := widget.NewButton("Обновить данные", func() {
	})
	buttonText := widget.NewButton("Добавить текстовые данные", func() {
		window.SetContent(contentD)
		window.Show()
	})
	buttonCart := widget.NewButton("Добавить банковскую карту", func() {
		window.SetContent(contentE)
		window.Show()
	})
	//---------------------------------------------------------------------- tbl - текстовые данные
	//dataTblText := [][]fyne.CanvasObject{
	//	{widget.NewLabel("NAME"), widget.NewLabel("DATA"), widget.NewLabel("DESCRIPTION"), widget.NewLabel("CREATED_AT"), widget.NewLabel("UPDATED_AT")},
	//	{widget.NewLabel("name_1"), widget.NewLabel("text"), widget.NewLabel("text"), widget.NewLabel("01-01-2023 08:30"), widget.NewLabel("01-01-2023 10:30")},
	//	{widget.NewLabel("name_2"), widget.NewLabel("text"), widget.NewLabel("text"), widget.NewLabel("01-01-2023 08:30"), widget.NewLabel("01-01-2023 10:30")},
	//	{widget.NewLabel("name_N"), widget.NewLabel("text"), widget.NewLabel("text"), widget.NewLabel("01-01-2023 08:30"), widget.NewLabel("01-01-2023 10:30")},
	//}
	dataTblText := [][]string{
		[]string{"top left", "top right"},
		[]string{"bottom left", "bottom right"}}

	//tableData := widget.NewTable(
	//	func() (int, int) {
	//		return len(dataTblText), len(dataTblText[0])
	//	},
	//	func() fyne.CanvasObject {
	//		c := container.NewWithoutLayout()
	//		r := canvas.NewRectangle(color.White)
	//		r.SetMinSize(fyne.NewSize(0, HEIGHT))
	//		r.Resize(fyne.NewSize(0, HEIGHT))
	//		c.Add(r)
	//		return c
	//	},
	//	func(cell widget.TableCellID, o fyne.CanvasObject) {
	//		container := o.(*fyne.Container)
	//		var obj = dataTblText[cell.Row][cell.Col]
	//		container.Add(obj)
	//		container.Refresh()
	//	})
	//setDefaultColumnsWidthText(tableData)

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

	containerTblText := layout.NewBorderLayout(buttonTop, buttonText, nil, nil)
	boxText := fyne.NewContainerWithLayout(containerTblText, buttonTop, tableData, buttonText)
	//---------------------------------------------------------------------- tbl - банковские карты
	dataTblCart := [][]fyne.CanvasObject{
		{widget.NewLabel("NAME"), widget.NewLabel("PAYMENT SYSTEM"), widget.NewLabel("NUMBER"), widget.NewLabel("HOLDER"), widget.NewLabel("CVC"), widget.NewLabel("END DATE"), widget.NewLabel("CREATED_AT"), widget.NewLabel("UPDATED_AT")},
		{widget.NewLabel("name_1"), widget.NewLabel("Visa"), widget.NewLabel("1234567890"), widget.NewLabel("Artur"), widget.NewLabel("123"), widget.NewLabel("01-01-2023 10:30"), widget.NewLabel("01-01-2023 08:30"), widget.NewLabel("01-01-2023 10:30")},
		{widget.NewLabel("name_2"), widget.NewLabel("Master cart"), widget.NewLabel("1234567890"), widget.NewLabel("Ura"), widget.NewLabel("123"), widget.NewLabel("01-01-2023 10:30"), widget.NewLabel("01-01-2023 08:30"), widget.NewLabel("01-01-2023 10:30")},
		{widget.NewLabel("name_N"), widget.NewLabel("Visa"), widget.NewLabel("1234567890"), widget.NewLabel("Pety"), widget.NewLabel("123"), widget.NewLabel("01-01-2023 10:30"), widget.NewLabel("01-01-2023 08:30"), widget.NewLabel("01-01-2023 10:30")},
	}

	tableDataCart := widget.NewTable(
		func() (int, int) {
			return len(dataTblCart), len(dataTblCart[0])
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
			var obj = dataTblCart[cell.Row][cell.Col]
			container.Add(obj)
			container.Refresh()
		})
	setDefaultColumnsWidthCart(tableDataCart)

	containerTblCart := layout.NewBorderLayout(buttonTop, buttonCart, nil, nil)
	boxCart := fyne.NewContainerWithLayout(containerTblCart, buttonTop, tableDataCart, buttonCart)
	//---------------------------------------------------------------------- cabinet
	labelFile := widget.NewLabel("Список файлов")
	tab1 := container.NewTabItem("Текстовые данные", boxText)
	tab2 := container.NewTabItem("Банковские карты", boxCart)
	tab3 := container.NewTabItem("Файлы", labelFile)
	tabs := container.NewAppTabs(tab1, tab2, tab3)
	//----------------------------------------------------------------------
	button := widget.NewButton("Submit", func() {
		log.Println("Submit")
		if radio.Selected == "Login" {
			if UsernameLogin.Text == user && PasswordLogin.Text == password {
				window.SetContent(tabs)
				window.Resize(fyne.NewSize(1250, 300))
				window.Show()
			}
		}
		if radio.Selected == "Registration" {
			if UsernameRegistration.Text == user && PasswordRegistration.Text == password && NewPasswordEntryRegistration.Text == password {
				window.SetContent(tabs)
				window.Resize(fyne.NewSize(1250, 300))
				window.Show()
			}
		}
	})

	buttonTextAdd := widget.NewButton("Добавить", func() {
		window.SetContent(tabs)
		window.Show()
	})
	buttonCartAdd := widget.NewButton("Добавить", func() {
		window.SetContent(tabs)
		window.Show()
	})

	contentA = container.NewVBox(radio)
	contentB = container.NewVBox(formLogin, button, separator, radio)
	contentC = container.NewVBox(formRegistration, button, separator, radio)
	contentD = container.NewVBox(formText, buttonTextAdd)
	contentE = container.NewVBox(formCart, buttonCartAdd)

	window.SetContent(contentA)
	window.ShowAndRun()
}
