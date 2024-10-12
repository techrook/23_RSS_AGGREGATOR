package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/techrook/23_RSS_AGGREGATOR/internal/database"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	Name     string `json:"name"`
}

func databaseUserToUser (dbUser database.User) User{
	return User{
		ID: dbUser.ID,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
		Name: dbUser.Name,
	}
}