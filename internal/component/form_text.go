package component

import "fyne.io/fyne/v2/widget"

func GetFormText() *widget.Form {
	TextName := widget.NewEntry()
	Text := widget.NewEntry()
	TextDescription := widget.NewEntry()
	formText := widget.NewForm(
		widget.NewFormItem("Name", TextName),
		widget.NewFormItem("Text", Text),
		widget.NewFormItem("Description", TextDescription),
	)
	return formText
}
