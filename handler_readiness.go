package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	type readyRes struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, 200, readyRes{Status: "ok"})
}
