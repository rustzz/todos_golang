package checkers

import (
	"strings"

	errs "github.com/rustzz/todos/cmd/errors"
	"github.com/rustzz/todos/internal/database"
	"github.com/rustzz/todos/internal/models"
)

var api_errors = errs.GetErrorsData()

func UserExists(user *models.User) bool {
	db := database.ConnectDatabase()
	var userDB models.User
	if err := db.Table("users").
		First(&userDB, "username = ?", user.Username).Error; err != nil {

		return false
	}
	return true
}

func DataValid(user *models.User, method string) bool {
	if method == "am" {
		var username string = strings.TrimSpace(user.Username)
		var password string = strings.TrimSpace(user.Password)
		if (len(username) < 3 || len(username) > 20) ||
			(len(password) < 6 || len(password) > 64) {

			return false
		}
	}
	if method == "notebook" {
		var username string = strings.TrimSpace(user.Username)
		var token string = strings.TrimSpace(user.Token)
		if (len(username) < 3 || len(username) > 20) ||
			(len(token) != 64) {

			return false
		}
	}
	return true
}

func PasswordValid(user *models.User) bool {
	db := database.ConnectDatabase()
	var userDB models.User
	if err := db.Table("users").Where("username = ?", user.Username).
		Find(&userDB).Error; err == nil {

		if userDB.Password == user.Password {
			return true
		}
		return false
	}
	return false
}

func TokenValid(user *models.User) bool {
	db := database.ConnectDatabase()
	var userDB models.User
	if err := db.Table("users").Where("username = ?", user.Username).
		Find(&userDB).Error; err == nil {

		if userDB.Token == user.Token {
			return true
		}
		return false
	}
	return false
}

func NotebookCheckerCoollection(user *models.User, method string) interface{} {
	if method == "notebook" {
		if !DataValid(user, "notebook") {
			return api_errors.NOTEBOOK_DATA_NOT_VALID
		}
		if !UserExists(user) {
			return api_errors.USER_NOT_EXISTS_ERROR
		}
		if !TokenValid(user) {
			return api_errors.TOKEN_NOT_VALID_ERROR
		}
		return nil
	}

	// auth
	if !DataValid(user, "am") {
		return api_errors.DATA_EMPTY_ERROR
	}
	if method == "am_signup" {
		if UserExists(user) {
			return api_errors.USER_EXISTS_ERROR
		}
		return nil
	}
	if method == "am_signin" {
		if !UserExists(user) {
			return api_errors.USER_NOT_EXISTS_ERROR
		}
		if !PasswordValid(user) {
			return api_errors.PASSWORD_NOT_VALID_ERROR
		}
		return nil
	}
	return nil
}
