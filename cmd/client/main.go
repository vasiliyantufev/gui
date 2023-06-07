package main

import (
	"fmt"
	"log"
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
	usernameLogin := widget.NewEntry()
	passwordLogin := widget.NewPasswordEntry()
	usernameRegistration := widget.NewEntry()
	passwordRegistration := widget.NewPasswordEntry()
	passwordEntryRegistration := widget.NewPasswordEntry()
	labelAlertAuth := widget.NewLabel("Text")
	formLogin := component.GetFormLogin(usernameLogin, passwordLogin)
	formRegistration := component.GetFormRegistration(usernameRegistration, passwordRegistration, passwordEntryRegistration)
	formText := component.GetFormText()
	formCart := component.GetFormCart()
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
	buttonTop := widget.NewButton("Обновить данные", func() {})
	buttonText := widget.NewButton("Добавить текстовые данные", func() {
		window.SetContent(containerFormText)
		window.Show()
	})
	buttonCart := widget.NewButton("Добавить банковскую карту", func() {
		window.SetContent(containerFormCart)
		window.Show()
	})
	tabText := component.GetTabTexts(buttonTop, buttonText)
	tabCart := component.GetTabCarts(buttonTop, buttonCart)
	tabFile := component.GetTabFiles()
	containerTabs := container.NewAppTabs(tabText, tabCart, tabFile)
	//----------------------------------------------------------------------
	button := widget.NewButton("Submit", func() {
		if radio.Selected == "Login" {
			user, exists := users[usernameLogin.Text]
			if exists {
				if user.Password != passwordLogin.Text {
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
			_, exists := users[passwordLogin.Text]

			if !exists {
				if utf8.RuneCountInString(usernameRegistration.Text) < 6 {
					labelAlertAuth.SetText("Длинна логина должна быть не менее шести символов")
					log.Println(labelAlertAuth.Text)
				} else if utf8.RuneCountInString(passwordRegistration.Text) < 6 {
					labelAlertAuth.SetText("Длинна пароля должна быть не менее шести символов")
					log.Println(labelAlertAuth.Text)
				} else if passwordRegistration.Text != passwordEntryRegistration.Text {
					labelAlertAuth.SetText("Пароли не совпали")
					log.Println(labelAlertAuth.Text)
				} else {
					users[usernameRegistration.Text] = model.User{Name: usernameRegistration.Text, Password: passwordRegistration.Text}
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
		window.SetContent(containerTabs)
		window.Show()
	})
	buttonCartAdd := widget.NewButton("Добавить", func() {
		window.SetContent(containerTabs)
		window.Show()
	})

	containerRadio = container.NewVBox(radio)
	containerFormLogin = container.NewVBox(formLogin, button, labelAlertAuth, separator, radio)
	containerFormRegistration = container.NewVBox(formRegistration, button, labelAlertAuth, separator, radio)
	containerFormText = container.NewVBox(formText, buttonTextAdd)
	containerFormCart = container.NewVBox(formCart, buttonCartAdd)

	window.SetContent(containerRadio)
	window.ShowAndRun()
}
