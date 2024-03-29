package main

import (
	"fmt"
	"net/http"

	"RSS-Aggregator/internal/auth"
	"RSS-Aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler, args ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Auth error: %v", err))
			return
		}

		dbUser, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user %v", err))
      return
		}
    
    role := "PUBLIC"
    if len(args) > 0 {
      role = args[0] 
    }

    if dbUser.Role != role {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Auth error RESTRICTED ACCESS"))
      return
    }

		handler(w, r, dbUser)
	}
}
