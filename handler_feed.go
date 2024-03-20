package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"RSS-Aggregator/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type createFeedParams struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := createFeedParams{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing Json", err))
		return
	}

	dbFeed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    dbUser.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Cannot create feed ", err))
	}

	respondWithJson(w, 200, dbFeedtoFeed(dbFeed))
}

func (apiCfg *apiConfig) handlerListFeed(w http.ResponseWriter, r *http.Request) {
	dbFeeds, err := apiCfg.DB.ListFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Cannot create feed ", err))
	}

	respondWithJson(w, 200, dbFeedstoFeeds(dbFeeds))
}
