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

// GetUserByID busca um usuário por ID
func GetUserByID(id uint) (*json.UserResponse, error) {
	user, err := datasource.GetUserByID(id)
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
		Mensagem:       "Usuário encontrado com sucesso",
	}

	return response, nil
}

// GetAllUsers retorna todos os usuários ativos
func GetAllUsers() ([]json.UserResponse, error) {
	users, err := datasource.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var responses []json.UserResponse
	for _, user := range users {
		response := json.UserResponse{
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
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateUser atualiza um usuário existente
func UpdateUser(id uint, req json.UserRequest) (*json.UserResponse, error) {
	user, err := datasource.UpdateUser(id, req)
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
		Mensagem:       "Usuário atualizado com sucesso",
	}

	return response, nil
}

// SoftDeleteUser realiza soft delete do usuário
func SoftDeleteUser(id uint) error {
	return datasource.SoftDeleteUser(id)
}

// RestoreUser restaura um usuário que foi soft deleted
func RestoreUser(id uint) error {
	return datasource.RestoreUser(id)
}
