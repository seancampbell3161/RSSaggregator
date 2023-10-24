package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (cfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	if len(apiKey) > 0 {
		apiKey = strings.Split(apiKey, " ")[1]
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		data, err := json.Marshal(user)
		_, err = w.Write(data)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		w.Header().Set("content-type", "application/json")
		return
	}
	fmt.Println("apiKey does not exist")
	w.WriteHeader(404)
}
