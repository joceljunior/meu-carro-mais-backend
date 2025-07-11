package services

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
	"errors"
)

func Login(req json.LoginRequest) (*json.LoginResponse, error) {
	db := database.InitDB() // ou GetDB, dependendo da sua implementação
	usuario, err := datasource.BuscarUsuarioPorEmail(db, req.Email)
	if err != nil {
		return nil, errors.New("usuário ou senha inválidos")
	}
	if usuario.Senha != req.Senha {
		return nil, errors.New("usuário ou senha inválidos")
	}
	resp := &json.LoginResponse{
		ID:    usuario.ID,
		Nome:  usuario.Nome,
		Email: usuario.Email,
	}
	return resp, nil
} 