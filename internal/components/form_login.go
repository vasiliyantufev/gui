package components

import "fyne.io/fyne/v2/widget"

func GetFormLogin() (*widget.Form, *widget.Entry, *widget.Entry) {
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()
	formLogin := widget.NewForm(
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Password", password),
	)
	return formLogin, username, password
}
