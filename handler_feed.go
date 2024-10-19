package main

import (
	"encoding/json"
	"fmt"
	"github.com/bkjones/rsstaurant/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// make the handler a method on the apiCfg struct so the handler has access to the db connection, since
// in golang these handler signatures cannot be changed, so we can't pass it in.
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}
	// CreateUser originates from the users.sql file where we named our query CreateUser. sqlc generated
	// the golang CreateUser function in internal/database/users.sql.go
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting feeds: %v", err))
	}
	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
