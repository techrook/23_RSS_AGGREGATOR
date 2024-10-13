package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/techrook/23_RSS_AGGREGATOR/internal/auth"
	"github.com/techrook/23_RSS_AGGREGATOR/internal/database"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}
	decoder :=  json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %S", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating user: %s", err))
		return
	}
	
	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request){
	apikey, err:= auth.GetApiKey(r.Header)

	if err != nil {
		respondWithError(w, 403,fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apikey)

	if err != nil {
		respondWithError(w, 403,fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}