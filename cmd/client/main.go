package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/vasiliyantufev/gui/internal/component"
	"github.com/vasiliyantufev/gui/internal/model"
)

var user = "user"
var password = "password"

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

	separator := widget.NewSeparator()

	//----------------------------------------------------------------------
	formLogin, UsernameLogin, PasswordLogin := component.GetFormLogin()
	formRegistration, UsernameRegistration, PasswordRegistration, NewPasswordEntryRegistration := component.GetFormRegistration()
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
	buttonTop := widget.NewButton("Обновить данные", func() {
	})
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
		log.Println("Submit")
		if radio.Selected == "Login" {
			if UsernameLogin.Text == user && PasswordLogin.Text == password {
				window.SetContent(containerTabs)
				window.Resize(fyne.NewSize(1250, 300))
				window.Show()
			}
		}
		if radio.Selected == "Registration" {
			if UsernameRegistration.Text == user && PasswordRegistration.Text == password && NewPasswordEntryRegistration.Text == password {
				window.SetContent(containerTabs)
				window.Resize(fyne.NewSize(1250, 300))
				window.Show()
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
	containerFormLogin = container.NewVBox(formLogin, button, separator, radio)
	containerFormRegistration = container.NewVBox(formRegistration, button, separator, radio)
	containerFormText = container.NewVBox(formText, buttonTextAdd)
	containerFormCart = container.NewVBox(formCart, buttonCartAdd)

	window.SetContent(containerRadio)
	window.ShowAndRun()
}
