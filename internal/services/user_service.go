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

	var lojaResponse *json.LojaUsuarioResponse
	if user.Loja.ID != 0 {
		lojaResponse = &json.LojaUsuarioResponse{
			Id:   user.Loja.ID,
			Nome: user.Loja.Nome,
			Logo: user.Loja.Imagem,
		}
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
		Loja:           lojaResponse,
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
		var lojaResponse *json.LojaUsuarioResponse
		if user.Loja.ID != 0 {
			lojaResponse = &json.LojaUsuarioResponse{
				Id:   user.Loja.ID,
				Nome: user.Loja.Nome,
				Logo: user.Loja.Imagem,
			}
		}

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
			Loja:           lojaResponse,
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

	var lojaResponse *json.LojaUsuarioResponse
	if user.Loja.ID != 0 {
		lojaResponse = &json.LojaUsuarioResponse{
			Id:   user.Loja.ID,
			Nome: user.Loja.Nome,
			Logo: user.Loja.Imagem,
		}
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
		Loja:           lojaResponse,
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

// GetUserPlanStatus verifica o status do plano de um usuário
func GetUserPlanStatus(id uint) (*json.UserPlanStatusResponse, error) {
	usuario, historico, err := datasource.GetUserPlanStatus(id)
	if err != nil {
		return nil, err
	}

	// Determina se o usuário é premium baseado no ID do plano
	isPremium := usuario.IDPlano > 1 // Assumindo que ID 1 é gratuito e IDs > 1 são premium

	response := &json.UserPlanStatusResponse{
		IDUsuario:    usuario.ID,
		NomeUsuario:  usuario.Nome,
		EmailUsuario: usuario.Email,
		IDPlano:      usuario.IDPlano,
		NomePlano:    usuario.Plano.Nome,
		IsPremium:    isPremium,
		Mensagem:     "Status do plano verificado com sucesso",
	}

	// Se há histórico de pagamento, adiciona informações adicionais
	if historico != nil {
		response.DataVencimento = historico.DataVencimento
		response.StatusPagamento = historico.Status
	}

	return response, nil
}
