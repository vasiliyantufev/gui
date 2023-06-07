package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	w.Resize(fyne.NewSize(400, 400))

	label := widget.NewLabel("label")

	formLogin := widget.NewForm(
		widget.NewFormItem("Username", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)

	formLogin.OnCancel = func() {
		label.Text = "Canceled"
		label.Refresh()
	}

	formLogin.OnSubmit = func() {
		label.Text = "submitted"
		label.Refresh()
	}

	w.SetContent(
		container.NewVBox(
			formLogin,
			label,
		),
	)

	w.ShowAndRun()
}
