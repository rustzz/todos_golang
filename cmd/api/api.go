package main

import (
	"os"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/rustzz/todos/config"
	"github.com/rustzz/todos/pkg/ratelimit"
	"github.com/rustzz/todos/internal/routing"
	"github.com/rustzz/todos/internal/database"
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
