package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rustzz/todos/cmd/routing"
	"github.com/rustzz/todos/config"
	"github.com/rustzz/todos/internal/database"
	"github.com/rustzz/todos/pkg/ratelimit"
)

func main() {
	config.Load()
	database.Migrate()

	router := mux.NewRouter()
	routing.InitRoutes(router)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := cors.Handler(router)

	log.Print("Server starting...")
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf("%s", os.Getenv("API_HOST")),
			ratelimit.Check(handler),
		),
	)
	return
}
