package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use()
	apiRouter := chi.NewRouter()
	r.Mount("/v1", apiRouter)

	corsMux := middlewareCors(r)

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
