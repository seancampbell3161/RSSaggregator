package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("content-type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshalling json: %s", err)
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(data)
	w.WriteHeader(status)
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Responding with status 500: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, status, errorResponse{Error: msg})
}
