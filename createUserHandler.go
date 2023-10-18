package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/seancampbell3161/RSSaggregator/internal/database"
	"net/http"
	"time"
)

type userCreateReq struct {
	Name string `json:"name"`
}

type userCreateRes struct {
	ID         string    `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	userReq := userCreateReq{}
	err := decoder.Decode(&userReq)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error decoding request"))
		return
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userReq.Name,
	}
	user, err := cfg.DB.CreateUser(r.Context(), userParams)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	response := userCreateRes{
		ID:         user.ID.String(),
		Created_at: user.CreatedAt,
		Updated_at: user.UpdatedAt,
		Name:       user.Name,
	}
	data, err := json.Marshal(response)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(201)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
	}
}
