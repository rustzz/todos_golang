package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"hash"
	"net/http"

	"github.com/google/uuid"
	"github.com/rustzz/todos/internal/checkers"
	"github.com/rustzz/todos/internal/database"
	"github.com/rustzz/todos/internal/models"
)

func SigninUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user_local = models.User{
		Username: request.FormValue("username"),
		Password: request.FormValue("password")}

	if !checkers.DataValid(&user_local, "am") {
		json.NewEncoder(writer).Encode(api_errors.DATA_EMPTY_ERROR)
		return
	}

	if !checkers.UserExists(&user_local) {
		json.NewEncoder(writer).Encode(api_errors.USER_NOT_EXISTS_ERROR)
		return
	}

	if !checkers.PasswordValid(&user_local) {
		json.NewEncoder(writer).Encode(api_errors.PASSWORD_NOT_VALID_ERROR)
		return
	}

	var hash hash.Hash = sha256.New()
	hash.Write([]byte(uuid.New().String()))
	var byte_token []byte = hash.Sum(nil)
	var token string = hex.EncodeToString(byte_token)

	db := database.ConnectDatabase()
	db.Table("users").Model(&models.User{}).Where("username = ?", user_local.Username).
		Update("token", token)

	json.NewEncoder(writer).Encode(map[string]interface{}{"ok": true, "token": token})
	return
}

func SignupUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user_local = models.User{
		Username: request.FormValue("username"),
		Password: request.FormValue("password")}

	if !checkers.DataValid(&user_local, "am") {
		json.NewEncoder(writer).Encode(api_errors.DATA_EMPTY_ERROR)
		return
	}

	if checkers.UserExists(&user_local) {
		json.NewEncoder(writer).Encode(api_errors.USER_EXISTS_ERROR)
		return
	}

	var hash hash.Hash = sha256.New()
	hash.Write([]byte(uuid.New().String()))
	var byte_token []byte = hash.Sum(nil)
	var token string = hex.EncodeToString(byte_token)

	db := database.ConnectDatabase()
	db.Table("users").Create(&models.User{
		Username: user_local.Username,
		Password: user_local.Password,
		Token:    token})

	json.NewEncoder(writer).Encode(map[string]interface{}{"ok": true, "token": token})
	return
}
