package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

func CreateUser(req json.UserRequest) (*json.UserResponse, error) {
	user, err := datasource.CreateNewUser(req)
	if err != nil {
		return nil, err
	}

	response := &json.UserResponse{
		ID:             user.ID,
		Nome:           user.Nome,
		Email:          user.Email,
		CPF:            user.CPF,
		Imagem:         user.Imagem,
		Telefone:       user.Telefone,
		Endereco:       user.Endereco,
		DataNascimento: user.DataNascimento,
		DataCadastro:   user.DataCadastro,
		Ativo:          user.Ativo,
		Latitude:       user.Latitude,
		Longitude:      user.Longitude,
		IDPlano:        user.IDPlano,
		IDLoja:         user.IDLoja,
		Mensagem:       "Usuário criado com sucesso",
	}

	return response, nil
}
