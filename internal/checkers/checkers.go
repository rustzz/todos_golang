package checkers

import (
	"encoding/json"
	"net/http"
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

func NotebookCheckerCoollection(user *models.User,
	writer http.ResponseWriter) {

	if !DataValid(user, "notebook") {
		json.NewEncoder(writer).Encode(api_errors.NOTEBOOK_DATA_NOT_VALID)
		return
	}

	if !UserExists(user) {
		json.NewEncoder(writer).Encode(api_errors.USER_NOT_EXISTS_ERROR)
		return
	}

	if !TokenValid(user) {
		json.NewEncoder(writer).Encode(api_errors.TOKEN_NOT_VALID_ERROR)
		return
	}
}
