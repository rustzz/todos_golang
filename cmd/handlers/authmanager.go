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

	checkerr := checkers.NotebookCheckerCoollection(&user_local, "am_signin")
	if checkerr != nil {
		json.NewEncoder(writer).Encode(checkerr)
		return
	}

	var hash hash.Hash = sha256.New()
	hash.Write([]byte(uuid.New().String()))
	var byte_token []byte = hash.Sum(nil)
	var token string = hex.EncodeToString(byte_token)

	db := database.ConnectDatabase()
	db.Table("users").Model(&models.User{}).Where("username = ?", user_local.Username).
		Update("token", token)
	dbConn, _ := db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(map[string]interface{}{"ok": true, "token": token})
	return
}

func SignupUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user_local = models.User{
		Username: request.FormValue("username"),
		Password: request.FormValue("password")}

	checkerr := checkers.NotebookCheckerCoollection(&user_local, "am_signup")
	if checkerr != nil {
		json.NewEncoder(writer).Encode(checkerr)
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
	dbConn, _ := db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(map[string]interface{}{"ok": true, "token": token})
	return
}
