package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

func CreateUser(req json.UserRequest) (*string, error) {
	resp, err := datasource.CreateNewUser(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
