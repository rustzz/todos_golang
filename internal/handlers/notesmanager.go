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

	var checkErr = checkers.CheckerCoollection(&userLocal, "notebook")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	var response = make(map[string]map[string]interface{})
	var id int
	var title, text string
	var checked bool

	var db = database.ConnectDatabase()
	var result, _ = db.Table("notes").
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

	var dbConn, _ = db.DB()
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

	var checkErr = checkers.CheckerCoollection(&userLocal, "notebook")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	var note = models.Note{Owner: userLocal.Username}
	var db = database.ConnectDatabase()
	db.Table("notes").
		Create(&note)

	var dbConn, _ = db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(
		map[string]interface{}{"ok": true, "id": note.ID})
	return
}

// DeleteNote : ...
func DeleteNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userLocal = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token"),
	}

	var checkErr = checkers.CheckerCoollection(&userLocal, "notebook")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	var data models.Note
	var err = json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(writer).Encode(apiErrors.NotebookDataNotValidError)
		return
	}

	var notesCount int64

	var db = database.ConnectDatabase()
	db.Table("notes").
		Where("owner = ?", userLocal.Username).
		Where("id = ?", data.ID).
		Count(&notesCount)

	if notesCount < 1 {
		json.NewEncoder(writer).Encode(apiErrors.NoteNotFoundError)
		return
	}

	db.Table("notes").
		Where("owner = ?", userLocal.Username).
		Delete(&models.Note{}, "id = ?", data.ID)

	var dbConn, _ = db.DB()
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

	var checkErr = checkers.CheckerCoollection(&userLocal, "notebook")
	if checkErr != nil {
		json.NewEncoder(writer).Encode(checkErr)
		return
	}

	var data models.Note
	var err = json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(writer).Encode(apiErrors.NotebookDataNotValidError)
		return
	}

	var db = database.ConnectDatabase()
	db.Table("notes").
		Where("owner = ?", userLocal.Username).
		Where("id = ?", data.ID).
		Updates(&models.Note{
			Title:   data.Title,
			Text:    data.Text,
			Checked: data.Checked})

	var dbConn, _ = db.DB()
	dbConn.Close()

	json.NewEncoder(writer).Encode(
		map[string]interface{}{"ok": true})
	return
}
