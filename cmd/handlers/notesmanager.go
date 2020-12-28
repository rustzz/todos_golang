package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rustzz/todos/internal/checkers"
	"github.com/rustzz/todos/internal/database"
	"github.com/rustzz/todos/internal/models"
)

func GetNotes(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user_local = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token")}

	checkerr := checkers.NotebookCheckerCoollection(&user_local, "notebook")
	if checkerr != nil {
		json.NewEncoder(writer).Encode(checkerr)
		return
	}

	db := database.ConnectDatabase()
	var response = make(map[string]map[string]interface{})
	var id int
	var title, text string
	var checked bool

	result, _ := db.Table("notes").
		Where("owner = ?", user_local.Username).
		Select("id, title, text, checked").Rows()

	for result.Next() {
		result.Scan(&id, &title, &text, &checked)
		response[strconv.Itoa(id)] = make(map[string]interface{})
		response[strconv.Itoa(id)] = map[string]interface{}{
			"title": title, "text": text, "checked": checked}
	}

	json.NewEncoder(writer).Encode(map[string]interface{}{
		"ok": true, "notes": response})
	return
}

func AddNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user_local = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token")}

	checkerr := checkers.NotebookCheckerCoollection(&user_local, "notebook")
	if checkerr != nil {
		json.NewEncoder(writer).Encode(checkerr)
		return
	}

	db := database.ConnectDatabase()
	db.Table("notes").Create(&models.Data{
		Owner: user_local.Username})

	json.NewEncoder(writer).Encode(map[string]interface{}{"ok": true})
	return
}

// to fix: returns "ok: true" if nothing deleted
func DeleteNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user_local = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token")}

	checkerr := checkers.NotebookCheckerCoollection(&user_local, "notebook")
	if checkerr != nil {
		json.NewEncoder(writer).Encode(checkerr)
		return
	}

	db := database.ConnectDatabase()
	var data models.Data
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(writer).Encode(api_errors.NOTEBOOK_DATA_NOT_VALID)
		return
	}
	db.Table("notes").Where("owner = ?", user_local.Username).
		Delete(&models.Data{}, "id = ?", data.ID)

	json.NewEncoder(writer).Encode(map[string]interface{}{"ok": true})
	return
}

func UpdateNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user_local = models.User{
		Username: request.FormValue("username"),
		Token:    request.FormValue("token")}

	checkerr := checkers.NotebookCheckerCoollection(&user_local, "notebook")
	if checkerr != nil {
		json.NewEncoder(writer).Encode(checkerr)
		return
	}

	db := database.ConnectDatabase()
	var data models.Data
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		json.NewEncoder(writer).Encode(api_errors.NOTEBOOK_DATA_NOT_VALID)
		return
	}
	db.Table("notes").Where("owner = ?", user_local.Username).Where("id = ?", data.ID).
		Updates(&models.Data{Title: data.Title, Text: data.Text, Checked: data.Checked})

	json.NewEncoder(writer).Encode(map[string]interface{}{"ok": true})
	return
}
