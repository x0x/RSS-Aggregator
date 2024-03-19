package main

import (
	"RSS-Aggregator/internal/auth"
	"RSS-Aggregator/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing Json", err))
		return
	}

	dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create user", err))
	}

	respondWithJson(w, 200, dbUsertoUser(dbUser))
}

func (apiCfg *apiConfig) handlerGetUser( w http.ResponseWriter, r * http.Request) {
  
  apiKey,err :=  auth.GetApiKey(r.Header)

  if err != nil {
    respondWithError(w, 500 , fmt.Sprintf("Auth error: %v",err))
  }
  
  dbUser, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
  
  if err != nil {
    respondWithError(w, 400, fmt.Sprintf("Couldn't get user %v",err))
  }

  respondWithJson(w,200,dbUsertoUser(dbUser))

}

