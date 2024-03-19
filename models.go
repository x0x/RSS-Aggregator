package main

import (
	"time"

	"RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	ApiKey     string    `json:"api_key"`
}

func dbUsertoUser(dbUser database.User) User {
	return User{
		Id:         dbUser.ID,
		Created_at: dbUser.CreatedAt,
		Updated_at: dbUser.UpdatedAt,
		Name:       dbUser.Name,
		ApiKey:     dbUser.ApiKey,
	}
}
