package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/seancampbell3161/RSSaggregator/internal/database"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONNECTION_STRING")

	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	config := apiConfig{dbQueries}

	r := chi.NewRouter()
	r.Use()
	apiRouter := chi.NewRouter()
	r.Mount("/v1", apiRouter)
	corsMux := middlewareCors(r)

	apiRouter.Get("/readiness", readinessHandler)
	apiRouter.Get("/err", errorHandler)

	apiRouter.Post("/users", config.createUserHandler)
	apiRouter.Get("/users", config.getUserHandler)

	server := &http.Server{
		Handler: corsMux,
		Addr:    port,
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
