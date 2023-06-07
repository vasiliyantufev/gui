package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/vasiliyantufev/gui/internal/models"
)

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

	users := make(map[string]models.User)
	texts := make(map[string]models.Text)
	carts := make(map[string]models.Cart)

	fmt.Print(users)
	fmt.Print(texts)
	fmt.Print(carts)

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

	containerTblText := layout.NewBorderLayout(buttonTop, buttonText, nil, nil)
	boxText := fyne.NewContainerWithLayout(containerTblText, buttonTop, tableData, buttonText)
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
