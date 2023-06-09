package service

import (
	"log"
	"strconv"
	"time"
	"unicode/utf8"

	"fyne.io/fyne/v2/widget"
	"github.com/vasiliyantufev/gui/internal/model"
)

func ValidateLogin(exists bool, user model.User, passwordLoginEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if !exists {
		labelAlertAuth.SetText("Данного пользователя не существует")
		log.Println(labelAlertAuth.Text)
		return false
	}
	if user.Password != passwordLoginEntry.Text {
		labelAlertAuth.SetText("Неверный пароль")
		log.Println(labelAlertAuth.Text)
		return false
	}
	return true
}

func ValidateRegistration(exists bool, usernameRegistrationEntry *widget.Entry, passwordRegistrationEntry *widget.Entry,
	passwordConfirmationRegistrationEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if exists {
		labelAlertAuth.SetText("Данный пользователь уже существует")
		log.Println(labelAlertAuth.Text)
		return false
	}
	if utf8.RuneCountInString(usernameRegistrationEntry.Text) < 6 {
		labelAlertAuth.SetText("Длинна логина должна быть не менее шести символов")
		log.Println(labelAlertAuth.Text)
		return false
	}
	if utf8.RuneCountInString(passwordRegistrationEntry.Text) < 6 {
		labelAlertAuth.SetText("Длинна пароля должна быть не менее шести символов")
		log.Println(labelAlertAuth.Text)
		return false
	}
	if passwordRegistrationEntry.Text != passwordConfirmationRegistrationEntry.Text {
		labelAlertAuth.SetText("Пароли не совпали")
		log.Println(labelAlertAuth.Text)
		return false
	}
	return true
}

func ValidateText(exists bool, textNameEntry *widget.Entry, textEntry *widget.Entry, textDescriptionEntry *widget.Entry, labelAlertText *widget.Label) bool {
	if exists {
		labelAlertText.SetText("Текст с таким name уже существует")
		log.Println(labelAlertText)
		return false
	}
	if utf8.RuneCountInString(textNameEntry.Text) < 6 {
		labelAlertText.SetText("Длинна name должна быть не менее шести символов")
		log.Println(labelAlertText.Text)
		return false
	}
	if utf8.RuneCountInString(textEntry.Text) < 6 {
		labelAlertText.SetText("Длинна text должна быть не менее шести символов")
		log.Println(labelAlertText.Text)
		return false
	}
	return true
}

func ValidateCart(exists bool, cartNameEntry *widget.Entry, paymentSystemEntry *widget.Entry, numberEntry *widget.Entry,
	holderEntry *widget.Entry, endDateEntry *widget.Entry, cvcEntry *widget.Entry, labelAlertCart *widget.Label) (bool, time.Time, int) {

	var validEndDate time.Time
	var validCvc int
	var err error

	if exists {
		labelAlertCart.SetText("Карта с таким name уже существует")
		log.Println(labelAlertCart)
		return false, time.Time{}, 0
	}
	if utf8.RuneCountInString(cartNameEntry.Text) < 6 {
		labelAlertCart.SetText("Длинна name должна быть не менее шести символов")
		log.Println(labelAlertCart.Text)
		return false, validEndDate, validCvc
	}
	if paymentSystemEntry.Text == "" {
		labelAlertCart.SetText("Payment System не заполнен")
		log.Println(labelAlertCart.Text)
		return false, validEndDate, validCvc
	}
	if numberEntry.Text == "" {
		labelAlertCart.SetText("Number не заполнен")
		log.Println(labelAlertCart.Text)
		return false, validEndDate, validCvc
	}
	if holderEntry.Text == "" {
		labelAlertCart.SetText("Holder не заполнен")
		log.Println(labelAlertCart.Text)
		return false, validEndDate, validCvc
	}
	if endDateEntry.Text == "" {
		labelAlertCart.SetText("End date не заполнен")
		log.Println(labelAlertCart.Text)
		return false, validEndDate, validCvc
	} else {
		layout := "01/02/2006"
		validEndDate, err = time.Parse(layout, endDateEntry.Text)
		if err != nil {
			labelAlertCart.SetText("End Date не корректный (пример: 01/02/2006)")
			log.Println(labelAlertCart.Text)
			return false, validEndDate, validCvc
		}
	}
	if cvcEntry.Text == "" {
		labelAlertCart.SetText("CVC не заполнен")
		log.Println(labelAlertCart.Text)
		return false, validEndDate, validCvc
	} else {
		validCvc, err = strconv.Atoi(cvcEntry.Text)
		if err != nil {
			labelAlertCart.SetText("CVC не корректный (пример: 123)")
			log.Println(labelAlertCart.Text)
			return false, validEndDate, validCvc
		}
	}

	return true, validEndDate, validCvc
}
