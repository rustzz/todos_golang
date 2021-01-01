package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rustzz/todos/internal/checkers"
	"github.com/rustzz/todos/internal/database"
	"github.com/rustzz/todos/internal/models"
)

// GetNotes : ...
func GetNotes(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userLocal = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token"),
	}

	checkErr := checkers.NotebookCheckerCoollection(&userLocal, "notebook")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	db := database.ConnectDatabase()
	var response = make(map[string]map[string]interface{})
	var id int
	var title, text string
	var checked bool

	result, _ := db.Table("notes").
		Where("owner = ?", userLocal.Username).
		Select("id, title, text, checked").
		Rows()

	for result.Next() {
		result.Scan(&id, &title, &text, &checked)
		response[strconv.Itoa(id)] = make(map[string]interface{})
		response[strconv.Itoa(id)] = map[string]interface{}{
			"title": title, "text": text, "checked": checked,
		}
	}
	dbConn, _ := db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(
		map[string]interface{}{"ok": true, "notes": response})
	return
}

// AddNote : ...
func AddNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userLocal = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token"),
	}

	checkErr := checkers.NotebookCheckerCoollection(&userLocal, "notebook")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	db := database.ConnectDatabase()
	db.Table("notes").
		Create(&models.Note{
			Owner: userLocal.Username})

	dbConn, _ := db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(
		map[string]interface{}{"ok": true})
	return
}

// DeleteNote : ...
// to fix: returns "ok: true" if nothing deleted
func DeleteNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var userLocal = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token"),
	}

	checkErr := checkers.NotebookCheckerCoollection(&userLocal, "notebook")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	db := database.ConnectDatabase()
	var data models.Note
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(writer).Encode(apiErrors.NotebookDataNotValidError)
		return
	}

	db.Table("notes").
		Where("owner = ?", userLocal.Username).
		Delete(&models.Note{}, "id = ?", data.ID)

	dbConn, _ := db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(
		map[string]interface{}{"ok": true})
	return
}

// UpdateNote : ...
func UpdateNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userLocal = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token"),
	}

	checkErr := checkers.NotebookCheckerCoollection(&userLocal, "notebook")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	db := database.ConnectDatabase()
	var data models.Note
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(writer).Encode(apiErrors.NotebookDataNotValidError)
		return
	}

	db.Table("notes").
		Where("owner = ?", userLocal.Username).
		Where("id = ?", data.ID).
		Updates(&models.Note{
			Title:   data.Title,
			Text:    data.Text,
			Checked: data.Checked})

	dbConn, _ := db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(
		map[string]interface{}{"ok": true})
	return
}
