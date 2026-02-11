package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// CreateRegistroInteresse cria um novo registro de interesse
func CreateRegistroInteresse(req json.RegistroInteresseRequest) (*json.RegistroInteresseResponse, error) {
	registroInteresse, err := datasource.CreateRegistroInteresse(req)
	if err != nil {
		return nil, err
	}

	response := &json.RegistroInteresseResponse{
		ID:              registroInteresse.ID,
		IDCupom:         registroInteresse.IDCupom,
		Nome:            registroInteresse.Nome,
		Email:           registroInteresse.Email,
		Telefone:        registroInteresse.Telefone,
		Mensagem:        registroInteresse.Mensagem,
		DataCadastro:    registroInteresse.DataCadastro,
		DataAtualizacao: registroInteresse.DataAtualizacao,
	}

	// Se o cupom foi carregado, adiciona ao response
	if registroInteresse.Cupom.ID != 0 {
		cupomResp := modelToCupomResponse(&registroInteresse.Cupom)
		response.Cupom = &cupomResp
	}

	return response, nil
}

// GetRegistroInteresseByID busca um registro de interesse por ID
func GetRegistroInteresseByID(id uint) (*json.RegistroInteresseResponse, error) {
	registroInteresse, err := datasource.GetRegistroInteresseByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.RegistroInteresseResponse{
		ID:              registroInteresse.ID,
		IDCupom:         registroInteresse.IDCupom,
		Nome:            registroInteresse.Nome,
		Email:           registroInteresse.Email,
		Telefone:        registroInteresse.Telefone,
		Mensagem:        registroInteresse.Mensagem,
		DataCadastro:    registroInteresse.DataCadastro,
		DataAtualizacao: registroInteresse.DataAtualizacao,
	}

	// Se o cupom foi carregado, adiciona ao response
	if registroInteresse.Cupom.ID != 0 {
		cupomResp := modelToCupomResponse(&registroInteresse.Cupom)
		response.Cupom = &cupomResp
	}

	return response, nil
}

// GetAllRegistroInteresses retorna todos os registros de interesse ativos
func GetAllRegistroInteresses() ([]json.RegistroInteresseResponse, error) {
	registrosInteresse, err := datasource.GetAllRegistroInteresses()
	if err != nil {
		return nil, err
	}

	var responses []json.RegistroInteresseResponse
	for _, registroInteresse := range registrosInteresse {
		response := json.RegistroInteresseResponse{
			ID:              registroInteresse.ID,
			IDCupom:         registroInteresse.IDCupom,
			Nome:            registroInteresse.Nome,
			Email:           registroInteresse.Email,
			Telefone:        registroInteresse.Telefone,
			Mensagem:        registroInteresse.Mensagem,
			DataCadastro:    registroInteresse.DataCadastro,
			DataAtualizacao: registroInteresse.DataAtualizacao,
		}

		// Se o cupom foi carregado, adiciona ao response
		if registroInteresse.Cupom.ID != 0 {
			cupomResp := modelToCupomResponse(&registroInteresse.Cupom)
			response.Cupom = &cupomResp
		}

		responses = append(responses, response)
	}

	return responses, nil
}

// GetRegistroInteressesByCupomID retorna todos os registros de interesse de um cupom específico
func GetRegistroInteressesByCupomID(cupomID uint) ([]json.RegistroInteresseResponse, error) {
	registrosInteresse, err := datasource.GetRegistroInteressesByCupomID(cupomID)
	if err != nil {
		return nil, err
	}

	var responses []json.RegistroInteresseResponse
	for _, registroInteresse := range registrosInteresse {
		response := json.RegistroInteresseResponse{
			ID:              registroInteresse.ID,
			IDCupom:         registroInteresse.IDCupom,
			Nome:            registroInteresse.Nome,
			Email:           registroInteresse.Email,
			Telefone:        registroInteresse.Telefone,
			Mensagem:        registroInteresse.Mensagem,
			DataCadastro:    registroInteresse.DataCadastro,
			DataAtualizacao: registroInteresse.DataAtualizacao,
		}

		// Se o cupom foi carregado, adiciona ao response
		if registroInteresse.Cupom.ID != 0 {
			cupomResp := modelToCupomResponse(&registroInteresse.Cupom)
			response.Cupom = &cupomResp
		}

		responses = append(responses, response)
	}

	return responses, nil
}

// SoftDeleteRegistroInteresse realiza soft delete do registro de interesse
func SoftDeleteRegistroInteresse(id uint) error {
	return datasource.SoftDeleteRegistroInteresse(id)
}

// RestoreRegistroInteresse restaura um registro de interesse que foi soft deleted
func RestoreRegistroInteresse(id uint) error {
	return datasource.RestoreRegistroInteresse(id)
}
