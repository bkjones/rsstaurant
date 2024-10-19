package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bkjones/rsstaurant/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// make the handler a method on the apiCfg struct so the handler has access to the db connection, since
// in golang these handler signatures cannot be changed, so we can't pass it in.
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
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
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
	}
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error getting posts for user: %v", err))
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
