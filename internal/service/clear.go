package service

import "fyne.io/fyne/v2/widget"

func ClearText(textNameEntry *widget.Entry, textEntry *widget.Entry, textDescriptionEntry *widget.Entry) {
	textNameEntry.Text = ""
	textEntry.Text = ""
	textDescriptionEntry.Text = ""
}

func ClearCart(cartNameEntry *widget.Entry, paymentSystemEntry *widget.Entry, numberEntry *widget.Entry,
	holderEntry *widget.Entry, endDateEntry *widget.Entry, cvcEntry *widget.Entry) {
	cartNameEntry.Text = ""
	paymentSystemEntry.Text = ""
	numberEntry.Text = ""
	holderEntry.Text = ""
	endDateEntry.Text = ""
	cvcEntry.Text = ""
}
