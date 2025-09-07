package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUser(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty parameters")
	}

	// does a user with this username already exist?
	userExist, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}
	if userExist {
		return fmt.Errorf("user already exists")
	}

	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}
	return nil
}
