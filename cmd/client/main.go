package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"unicode/utf8"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/vasiliyantufev/gui/internal/component"
	"github.com/vasiliyantufev/gui/internal/model"
)

func main() {

	dataTblText := [][]string{
		{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"},
		{"name_1", "text", "text", "01-01-2023 08:30", "01-01-2023 10:30"},
		{"name_2", "text", "text", "01-01-2023 08:30", "01-01-2023 10:30"},
		{"name_N", "text", "text", "01-01-2023 08:30", "01-01-2023 10:30"}}

	dataTblCart := [][]string{
		{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED_AT", "UPDATED_AT"},
		{"name_1", "Visa", "1234567890", "Artur", "123", "01-01-2023 10:30", "01-01-2023 08:30", "01-01-2023 10:30"},
		{"name_2", "Visa", "1234567890", "Artur", "123", "01-01-2023 10:30", "01-01-2023 08:30", "01-01-2023 10:30"},
		{"name_N", "Visa", "1234567890", "Artur", "123", "01-01-2023 10:30", "01-01-2023 08:30", "01-01-2023 10:30"}}

	users := make(map[string]model.User)
	texts := make(map[string]model.Text)
	carts := make(map[string]model.Cart)

	fmt.Print(users)
	fmt.Print(texts)
	fmt.Print(carts)

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
		dataTblText[1][1] = strconv.Itoa(i)
		dataTblCart[1][1] = strconv.Itoa(i)
		tblText.Refresh()
		tblCart.Refresh()
	})

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
		if radio.Selected == "Login" {
			user, exists := users[usernameLoginEntry.Text]
			if exists {
				if user.Password != passwordLoginEntry.Text {
					labelAlertAuth.SetText("Неверный пароль")
					log.Println(labelAlertAuth.Text)
				} else {
					window.SetContent(containerTabs)
					window.Resize(fyne.NewSize(1250, 300))
					window.Show()
				}
			} else {
				labelAlertAuth.SetText("Данного пользователя не существует")
				log.Println(labelAlertAuth.Text)
			}
		}
		if radio.Selected == "Registration" {
			_, exists := users[passwordLoginEntry.Text]

			if !exists {
				if utf8.RuneCountInString(usernameRegistrationEntry.Text) < 6 {
					labelAlertAuth.SetText("Длинна логина должна быть не менее шести символов")
					log.Println(labelAlertAuth.Text)
				} else if utf8.RuneCountInString(passwordRegistrationEntry.Text) < 6 {
					labelAlertAuth.SetText("Длинна пароля должна быть не менее шести символов")
					log.Println(labelAlertAuth.Text)
				} else if passwordRegistrationEntry.Text != passwordConfirmationRegistrationEntry.Text {
					labelAlertAuth.SetText("Пароли не совпали")
					log.Println(labelAlertAuth.Text)
				} else {
					users[usernameRegistrationEntry.Text] = model.User{Name: usernameRegistrationEntry.Text, Password: passwordRegistrationEntry.Text}
					window.SetContent(containerTabs)
					window.Resize(fyne.NewSize(1250, 300))
					window.Show()
				}
			} else {
				labelAlertAuth.SetText("Данный пользователь уже существует")
				log.Println(labelAlertAuth.Text)
			}
		}
	})

	buttonTextAdd := widget.NewButton("Добавить", func() {
		_, exists := texts[textNameEntry.Text]
		if exists {
			labelAlertText.SetText("Текст с таким name уже существует")
			log.Println(labelAlertText)
		} else {
			if utf8.RuneCountInString(textNameEntry.Text) < 6 {
				labelAlertText.SetText("Длинна name должна быть не менее шести символов")
				log.Println(labelAlertText.Text)
			} else if utf8.RuneCountInString(textEntry.Text) < 6 {
				labelAlertText.SetText("Длинна text должна быть не менее шести символов")
				log.Println(labelAlertText.Text)
			} else {
				texts[textNameEntry.Text] = model.Text{Name: textNameEntry.Text, Text: textEntry.Text, Description: textDescriptionEntry.Text}
				log.Println("Текст добавлен")

				textNameEntry.Text = ""
				textEntry.Text = ""
				textDescriptionEntry.Text = ""
				formText.Refresh()

				window.SetContent(containerTabs)
				window.Show()
			}
		}
		fmt.Print(texts)
	})

	buttonCartAdd := widget.NewButton("Добавить", func() {
		_, exists := carts[cartNameEntry.Text]
		if exists {
			labelAlertCart.SetText("Карта с таким name уже существует")
			log.Println(labelAlertCart)
		} else {
			if utf8.RuneCountInString(cartNameEntry.Text) < 6 {
				labelAlertCart.SetText("Длинна name должна быть не менее шести символов")
				log.Println(labelAlertCart.Text)
			} else if paymentSystemEntry.Text == "" {
				labelAlertCart.SetText("Payment System не заполнен")
				log.Println(labelAlertCart.Text)
			} else if numberEntry.Text == "" {
				labelAlertCart.SetText("Number не заполнен")
				log.Println(labelAlertCart.Text)
			} else if holderEntry.Text == "" {
				labelAlertCart.SetText("Holder не заполнен")
				log.Println(labelAlertCart.Text)
			} else if endDateEntry.Text == "" {
				labelAlertCart.SetText("End date не заполнен")
				log.Println(labelAlertCart.Text)
			} else if cvcEntry.Text == "" {
				labelAlertCart.SetText("CVC не заполнен")
				log.Println(labelAlertCart.Text)
			} else {
				layout := "01/02/2006"
				endDate, errData := time.Parse(layout, endDateEntry.Text)
				cvc, errCvc := strconv.Atoi(cvcEntry.Text)
				if errData != nil {
					labelAlertCart.SetText("End Date не корректный (пример: 01/02/2006)")
					log.Println(labelAlertCart.Text)
				} else if errCvc != nil {
					labelAlertCart.SetText("CVC не корректный (пример: 123)")
					log.Println(labelAlertCart.Text)
				} else {
					carts[cartNameEntry.Text] = model.Cart{
						Name:          cartNameEntry.Text,
						PaymentSystem: paymentSystemEntry.Text,
						Number:        numberEntry.Text,
						Holder:        holderEntry.Text,
						EndData:       endDate,
						CVC:           cvc}
					log.Println("Текст добавлен")

					cartNameEntry.Text = ""
					paymentSystemEntry.Text = ""
					numberEntry.Text = ""
					holderEntry.Text = ""
					endDateEntry.Text = ""
					cvcEntry.Text = ""
					formCart.Refresh()

					window.SetContent(containerTabs)
					window.Show()
				}
			}
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
