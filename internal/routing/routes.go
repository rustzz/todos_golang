package routing

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rustzz/todos/internal/handlers"
)

// InitRoutes : ...
func InitRoutes(handler *mux.Router) {
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(""))
	}).Methods(http.MethodGet)

	handler.HandleFunc("/am/signin", handlers.SigninUser).Methods(http.MethodPost)
	handler.HandleFunc("/am/signup", handlers.SignupUser).Methods(http.MethodPost)
	handler.HandleFunc("/notebook/get", handlers.GetNotes).Methods(http.MethodPost)
	handler.HandleFunc("/notebook/add", handlers.AddNote).Methods(http.MethodPost)
	handler.HandleFunc("/notebook/delete", handlers.DeleteNote).Methods(http.MethodPost)
	handler.HandleFunc("/notebook/update", handlers.UpdateNote).Methods(http.MethodPost)
	return
}
