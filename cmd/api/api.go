package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rustzz/todos/config"
	"github.com/rustzz/todos/internal/database"
	"github.com/rustzz/todos/internal/routing"
	"github.com/rustzz/todos/pkg/ratelimit"
)

func main() {
	config.Load()
	database.Migrate()

	var router = mux.NewRouter()
	routing.InitRoutes(router)

	var cors = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	var handler = cors.Handler(router)

	log.Print("Server starting...")
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf("%s", os.Getenv("API_HOST")),
			ratelimit.Check(handler),
		),
	)
	return
}
