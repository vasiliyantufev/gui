package components

import "fyne.io/fyne/v2/widget"

func GetFormRegistration() (*widget.Form, *widget.Entry, *widget.Entry, *widget.Entry) {
	UsernameRegistration := widget.NewEntry()
	PasswordRegistration := widget.NewPasswordEntry()
	NewPasswordEntryRegistration := widget.NewPasswordEntry()
	formRegistration := widget.NewForm(
		widget.NewFormItem("Username", UsernameRegistration),
		widget.NewFormItem("Password", PasswordRegistration),
		widget.NewFormItem("Confirm password", NewPasswordEntryRegistration),
	)

	return formRegistration, UsernameRegistration, PasswordRegistration, NewPasswordEntryRegistration
}