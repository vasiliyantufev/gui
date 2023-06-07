package component

import "fyne.io/fyne/v2/widget"

func GetFormCart(name *widget.Entry, paymentSystem *widget.Entry, number *widget.Entry, holder *widget.Entry, endDate *widget.Entry, cvc *widget.Entry) *widget.Form {
	formCart := widget.NewForm(
		widget.NewFormItem("Name", name),
		widget.NewFormItem("Payment system", paymentSystem),
		widget.NewFormItem("Number", number),
		widget.NewFormItem("Holder", holder),
		widget.NewFormItem("End date", endDate),
		widget.NewFormItem("CVC", cvc),
	)
	return formCart
}
