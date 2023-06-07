package models

import "time"

type User struct {
	name     string
	password string
}

type Text struct {
	name        string
	description string
	text        string
}

type Cart struct {
	name          string
	paymentSystem string
	number        string
	holder        string
	cvc           int
	endData       time.Time
}
