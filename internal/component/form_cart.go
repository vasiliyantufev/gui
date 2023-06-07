package component

import "fyne.io/fyne/v2/widget"

func GetFormCart() *widget.Form {
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
	return formCart
}
