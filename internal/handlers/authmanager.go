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

// SigninUser : ...
func SigninUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userLocal = models.User{
		Username: request.FormValue("username"),
		Password: request.FormValue("password"),
	}

	var checkErr = checkers.NotebookCheckerCoollection(&userLocal, "am_signin")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	var hash hash.Hash = sha256.New()
	hash.Write([]byte(uuid.New().String()))
	var byteToken []byte = hash.Sum(nil)
	var token string = hex.EncodeToString(byteToken)

	var db = database.ConnectDatabase()
	db.Table("users").
		Model(&models.User{}).
		Where("username = ?", userLocal.Username).
		Update("token", token)

	var dbConn, _ = db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(
		map[string]interface{}{"ok": true, "token": token})
	return
}

// SignupUser : ...
func SignupUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userLocal = models.User{
		Username: request.FormValue("username"),
		Password: request.FormValue("password"),
	}

	var checkErr = checkers.NotebookCheckerCoollection(&userLocal, "am_signup")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	var hash hash.Hash = sha256.New()
	hash.Write([]byte(uuid.New().String()))
	var byteToken []byte = hash.Sum(nil)
	var token string = hex.EncodeToString(byteToken)

	var db = database.ConnectDatabase()
	db.Table("users").
		Create(&models.User{
			Username: userLocal.Username,
			Password: userLocal.Password,
			Token:    token})

	var dbConn, _ = db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(
		map[string]interface{}{"ok": true, "token": token})
	return
}
