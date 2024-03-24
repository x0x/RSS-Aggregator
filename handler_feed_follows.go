package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"RSS-Aggregator/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing Json", err))
		return
	}

	dbFeedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedId,
		UserID:    dbUser.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Cannot create feedFollow ", err))
	}
	respondWithJson(w, 200, dbFeedFollowtoFeedFollow(dbFeedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowID")
	feedFollowIdUUID, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to parse feedFollowId", err))
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID: feedFollowIdUUID,
		UserID: dbUser.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to delete Feed Follow", err))
	}
	respondWithJson(w, http.StatusOK, struct{}{})
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	dbFeedFollows, err := apiCfg.DB.GetFeedFollow(r.Context(), dbUser.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to get feed follow for the user", err))
	}
	respondWithJson(w, 200, dbFeedFollowstoFeedFollows(dbFeedFollows))
}
