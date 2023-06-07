package models

import "time"

type Cart struct {
	name          string
	paymentSystem string
	number        string
	holder        string
	cvc           int
	endData       time.Time
}
