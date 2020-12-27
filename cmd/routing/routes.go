package routing

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rustzz/todos/cmd/handlers"
)

func InitRoutes(handler *mux.Router) {
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(""))
	}).Methods("GET")

	handler.HandleFunc("/am/signin", handlers.SigninUser).Methods("POST")
	handler.HandleFunc("/am/signup", handlers.SignupUser).Methods("POST")
	handler.HandleFunc("/notebook/get", handlers.GetNotes).Methods("POST")
	handler.HandleFunc("/notebook/add", handlers.AddNote).Methods("POST")
	handler.HandleFunc("/notebook/delete", handlers.DeleteNote).Methods("POST")
	handler.HandleFunc("/notebook/update", handlers.UpdateNote).Methods("POST")
	return
}
