package main

import (
	"fmt"
	"net/http"

	"RSS-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	dbPosts, err := apiCfg.DB.GetPostsByuser(r.Context(), database.GetPostsByuserParams{
		UserID: dbUser.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts by user", err))
		return
	}

	respondWithJson(w, 200, dbPostsToPosts(dbPosts))
}
