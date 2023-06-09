package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/vasiliyantufev/gui/internal/component"
	"github.com/vasiliyantufev/gui/internal/model"
	"github.com/vasiliyantufev/gui/internal/service"
)

func main() {

	dataTblText := [][]string{
		{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"}}

	dataTblCart := [][]string{
		{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED_AT", "UPDATED_AT"}}

	users := make(map[string]model.User)
	texts := make(map[string]model.Text)
	carts := make(map[string]model.Cart)

	application := app.New()
	application.Settings().SetTheme(theme.LightTheme())

	window := application.NewWindow("GophKeeper")
	window.Resize(fyne.NewSize(250, 80))

	var containerRadio *fyne.Container
	var containerFormLogin *fyne.Container
	var containerFormRegistration *fyne.Container
	var containerFormText *fyne.Container
	var containerFormCart *fyne.Container

	//----------------------------------------------------------------------
	separator := widget.NewSeparator()

	usernameLoginEntry := widget.NewEntry()
	passwordLoginEntry := widget.NewPasswordEntry()
	usernameRegistrationEntry := widget.NewEntry()
	passwordRegistrationEntry := widget.NewPasswordEntry()
	passwordConfirmationRegistrationEntry := widget.NewPasswordEntry()
	textNameEntry := widget.NewEntry()
	textEntry := widget.NewEntry()
	textDescriptionEntry := widget.NewEntry()
	cartNameEntry := widget.NewEntry()
	paymentSystemEntry := widget.NewEntry()
	numberEntry := widget.NewEntry()
	holderEntry := widget.NewEntry()
	endDateEntry := widget.NewEntry()
	cvcEntry := widget.NewEntry()

	labelAlertAuth := widget.NewLabel("")
	labelAlertText := widget.NewLabel("")
	labelAlertCart := widget.NewLabel("")
	labelAlertAuth.Hide()
	labelAlertText.Hide()
	labelAlertCart.Hide()

	formLogin := component.GetFormLogin(usernameLoginEntry, passwordLoginEntry)
	formRegistration := component.GetFormRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry)
	formText := component.GetFormText(textNameEntry, textEntry, textDescriptionEntry)
	formCart := component.GetFormCart(cartNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry)
	//----------------------------------------------------------------------
	options := []string{"Login", "Registration"}
	radio := widget.NewRadioGroup(options, func(value string) {
		log.Println("Radio set to ", value)
		if value == "Login" {
			window.SetContent(containerFormLogin)
			window.Resize(fyne.NewSize(500, 100))
			window.Show()
		}
		if value == "Registration" {
			window.SetContent(containerFormRegistration)
			window.Resize(fyne.NewSize(500, 100))
			window.Show()
		}
	})
	//----------------------------------------------------------------------
	i := 1
	tblText := component.GetTableText(dataTblText)
	tblCart := component.GetTableCart(dataTblCart)
	buttonTop := widget.NewButton("Обновить данные", func() {
		i++
		dataTblText[0][0] = strconv.Itoa(i)
		dataTblCart[0][0] = strconv.Itoa(i)
		tblText.Refresh()
		tblCart.Refresh()

		tblText.Resize(fyne.NewSize(5, 2))
		dataTblText = append(dataTblText, []string{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"})
		tblText.Refresh()

	})
	//----------------------------------------------------------------------
	buttonText := widget.NewButton("Добавить текстовые данные", func() {
		window.SetContent(containerFormText)
		window.Show()
	})
	buttonCart := widget.NewButton("Добавить банковскую карту", func() {
		window.SetContent(containerFormCart)
		window.Show()
	})

	tabText := component.GetTabTexts(tblText, buttonTop, buttonText)
	tabCart := component.GetTabCarts(tblCart, buttonTop, buttonCart)
	tabFile := component.GetTabFiles()

	containerTabs := container.NewAppTabs(tabText, tabCart, tabFile)
	//----------------------------------------------------------------------
	button := widget.NewButton("Submit", func() {
		labelAlertAuth.Show()
		var valid = false
		if radio.Selected == "Login" {
			user, exists := users[usernameLoginEntry.Text]
			valid = service.ValidateLogin(exists, user, passwordLoginEntry, labelAlertAuth)
			if valid {
				window.SetContent(containerTabs)
				window.Resize(fyne.NewSize(1250, 300))
				window.Show()
			}
		}
		if radio.Selected == "Registration" {
			_, exists := users[passwordLoginEntry.Text]
			valid = service.ValidateRegistration(exists, usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry, labelAlertAuth)
			if valid {
				users[usernameRegistrationEntry.Text] = model.User{Name: usernameRegistrationEntry.Text, Password: passwordRegistrationEntry.Text}
				window.SetContent(containerTabs)
				window.Resize(fyne.NewSize(1250, 300))
				window.Show()
			}
		}
	})

	buttonTextAdd := widget.NewButton("Добавить", func() {
		labelAlertText.Show()
		var valid = false
		_, exists := texts[textNameEntry.Text]
		valid = service.ValidateText(exists, textNameEntry, textEntry, textDescriptionEntry, labelAlertText)
		if valid {
			texts[textNameEntry.Text] = model.Text{
				Name:        textNameEntry.Text,
				Text:        textEntry.Text,
				Description: textDescriptionEntry.Text}

			service.ClearText(textNameEntry, textEntry, textDescriptionEntry)
			log.Print("Текст добавлен")

			labelAlertText.Hide()
			formText.Refresh()
			window.SetContent(containerTabs)
			window.Show()
		}
		log.Print(texts)
	})

	buttonCartAdd := widget.NewButton("Добавить", func() {
		labelAlertCart.Show()
		var valid = false
		var endDate time.Time
		var cvc int
		_, exists := carts[cartNameEntry.Text]
		valid, endDate, cvc = service.ValidateCart(exists, cartNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry, labelAlertCart)
		if valid {
			carts[cartNameEntry.Text] = model.Cart{
				Name:          cartNameEntry.Text,
				PaymentSystem: paymentSystemEntry.Text,
				Number:        numberEntry.Text,
				Holder:        holderEntry.Text,
				EndData:       endDate,
				CVC:           cvc}
			log.Println("Текст добавлен")

			service.ClearCart(cartNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry)
			formCart.Refresh()

			labelAlertCart.Hide()
			window.SetContent(containerTabs)
			window.Show()
		}
		fmt.Print(carts)
	})

	containerRadio = container.NewVBox(radio)
	containerFormLogin = container.NewVBox(formLogin, button, labelAlertAuth, separator, radio)
	containerFormRegistration = container.NewVBox(formRegistration, button, labelAlertAuth, separator, radio)
	containerFormText = container.NewVBox(formText, buttonTextAdd, labelAlertText)
	containerFormCart = container.NewVBox(formCart, buttonCartAdd, labelAlertCart)

	window.SetContent(containerRadio)
	window.ShowAndRun()
}
