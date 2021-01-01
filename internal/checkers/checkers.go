package checkers

import (
	"strings"

	errs "github.com/rustzz/todos/cmd/errors"
	"github.com/rustzz/todos/internal/database"
	"github.com/rustzz/todos/internal/models"
)

var apiErrors = errs.GetErrorsData()

// UserExists : ...
func UserExists(user *models.User) bool {
	db := database.ConnectDatabase()
	var userDB models.User
	if err := db.Table("users").
		First(&userDB, "username = ?", user.Username).Error; err != nil {

		dbConn, _ := db.DB()
		dbConn.Close()
		return false
	}
	dbConn, _ := db.DB()
	dbConn.Close()
	return true
}

// DataValid : ...
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

// PasswordValid : ...
func PasswordValid(user *models.User) bool {
	db := database.ConnectDatabase()
	var userDB models.User
	if err := db.Table("users").
		Where("username = ?", user.Username).
		Find(&userDB).Error; err == nil {

		if userDB.Password == user.Password {
			dbConn, _ := db.DB()
			dbConn.Close()
			return true
		}
		dbConn, _ := db.DB()
		dbConn.Close()
		return false
	}
	dbConn, _ := db.DB()
	dbConn.Close()
	return false
}

// TokenValid : ...
func TokenValid(user *models.User) bool {
	db := database.ConnectDatabase()
	var userDB models.User
	if err := db.Table("users").
		Where("username = ?", user.Username).
		Find(&userDB).Error; err == nil {

		if userDB.Token == user.Token {
			dbConn, _ := db.DB()
			dbConn.Close()
			return true
		}
		dbConn, _ := db.DB()
		dbConn.Close()
		return false
	}
	dbConn, _ := db.DB()
	dbConn.Close()
	return false
}

// NotebookCheckerCoollection : ...
func NotebookCheckerCoollection(user *models.User, method string) interface{} {
	if method == "notebook" {
		if !DataValid(user, "notebook") {
			return apiErrors.NotebookDataNotValidError
		}
		if !UserExists(user) {
			return apiErrors.UserNotExistsError
		}
		if !TokenValid(user) {
			return apiErrors.TokenNotValidError
		}
		return nil
	}

	// Auth
	if !DataValid(user, "am") {
		return apiErrors.DataEmptyError
	}
	if method == "am_signup" {
		if UserExists(user) {
			return apiErrors.UserExistsError
		}
		return nil
	}
	if method == "am_signin" {
		if !UserExists(user) {
			return apiErrors.UserNotExistsError
		}
		if !PasswordValid(user) {
			return apiErrors.PasswordNotValidError
		}
		return nil
	}
	return nil
}
