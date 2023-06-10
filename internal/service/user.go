package service

import "github.com/vasiliyantufev/gui/internal/model"

var userId int64

func UserExist(username string) (int64, bool) {
	userId = 1
	return userId, false
}

func Authentication(username, password string) (model.User, bool) {
	userId = 1
	user := model.User{ID: userId, Name: username, Password: password}
	return user, true
}
