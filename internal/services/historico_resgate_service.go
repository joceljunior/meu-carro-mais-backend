package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// CreateHistoricoResgateFromCupom cria um histórico de resgate a partir de um cupom
func CreateHistoricoResgateFromCupom(cupomID uint, usuarioID uint) (*json.HistoricoResgateResponse, error) {
	historico, err := datasource.CreateHistoricoResgateFromCupom(cupomID, usuarioID)
	if err != nil {
		return nil, err
	}

	return convertHistoricoToResponse(historico), nil
}

// convertHistoricoToResponse converte um modelo de histórico para response
func convertHistoricoToResponse(historico *models.HistoricoResgate) *json.HistoricoResgateResponse {
	response := &json.HistoricoResgateResponse{
		ID:              historico.ID,
		IDCupom:         historico.IDCupom,
		IDUsuario:       historico.IDUsuario,
		DataResgate:     historico.DataResgate,
		DataAtualizacao: historico.DataAtualizacao,
		Status:          historico.Status,
		Usuario: json.UserResponse{
			ID:             historico.Usuario.ID,
			Nome:           historico.Usuario.Nome,
			Email:          historico.Usuario.Email,
			CPF:            historico.Usuario.CPF,
			Imagem:         historico.Usuario.Imagem,
			Telefone:       historico.Usuario.Telefone,
			Endereco:       historico.Usuario.Endereco,
			DataNascimento: historico.Usuario.DataNascimento,
			DataCadastro:   historico.Usuario.DataCadastro,
			Ativo:          historico.Usuario.Ativo,
			Latitude:       historico.Usuario.Latitude,
			Longitude:      historico.Usuario.Longitude,
			IDPlano:        historico.Usuario.IDPlano,
			IDLoja:         historico.Usuario.IDLoja,
		},
	}

	// Adiciona dados do cupom se existir
	if historico.Cupom != nil {
		cupomResp := modelToCupomResponse(historico.Cupom)
		response.Cupom = &cupomResp
	}

	return response
}

// CreateHistoricoResgate cria um novo histórico de resgate
func CreateHistoricoResgate(req json.HistoricoResgateRequest) (*json.HistoricoResgateResponse, error) {
	historico, err := datasource.CreateHistoricoResgate(req)
	if err != nil {
		return nil, err
	}

	return convertHistoricoToResponse(historico), nil
}

// GetHistoricoResgateByID busca um histórico por ID
func GetHistoricoResgateByID(id uint) (*json.HistoricoResgateResponse, error) {
	historico, err := datasource.GetHistoricoResgateByID(id)
	if err != nil {
		return nil, err
	}

	return convertHistoricoToResponse(historico), nil
}

// GetAllHistoricosResgate retorna todos os históricos ativos
func GetAllHistoricosResgate() ([]json.HistoricoResgateResponse, error) {
	historicos, err := datasource.GetAllHistoricosResgate()
	if err != nil {
		return nil, err
	}

	var responses []json.HistoricoResgateResponse
	for _, historico := range historicos {
		responses = append(responses, *convertHistoricoToResponse(&historico))
	}

	return responses, nil
}

// GetHistoricosResgateByUsuarioID retorna todos os históricos de um usuário específico
func GetHistoricosResgateByUsuarioID(idUsuario uint) (*json.HistoricosResgateResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoResgateResponse
	for _, historico := range historicos {
		historicosResponse = append(historicosResponse, *convertHistoricoToResponse(&historico))
	}

	return &json.HistoricosResgateResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}, nil
}

// GetHistoricosResgateByLojaID retorna todos os históricos de uma loja específica
func GetHistoricosResgateByLojaID(idLoja uint) (*json.HistoricosResgateResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoResgateResponse
	for _, historico := range historicos {
		historicosResponse = append(historicosResponse, *convertHistoricoToResponse(&historico))
	}

	return &json.HistoricosResgateResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}, nil
}

// UpdateHistoricoResgate atualiza um histórico existente
func UpdateHistoricoResgate(id uint, req json.HistoricoResgateRequest) (*json.HistoricoResgateResponse, error) {
	historico, err := datasource.UpdateHistoricoResgate(id, req)
	if err != nil {
		return nil, err
	}

	return convertHistoricoToResponse(historico), nil
}

// UpdateStatusHistoricoResgate atualiza apenas o status de um histórico
func UpdateStatusHistoricoResgate(id uint, status string) error {
	return datasource.UpdateStatusHistoricoResgate(id, status)
}

// SoftDeleteHistoricoResgate realiza soft delete do histórico
func SoftDeleteHistoricoResgate(id uint) error {
	return datasource.SoftDeleteHistoricoResgate(id)
}

// RestoreHistoricoResgate restaura um histórico que foi soft deleted
func RestoreHistoricoResgate(id uint) error {
	return datasource.RestoreHistoricoResgate(id)
}

// GetHistoricosResgateClienteByUsuarioID retorna histórico simplificado do cliente
func GetHistoricosResgateClienteByUsuarioID(usuarioID uint) (*json.HistoricosResgateClienteResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByUsuarioID(usuarioID)
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoResgateClienteResponse
	for _, historico := range historicos {
		resp := json.HistoricoResgateClienteResponse{
			ID:              historico.ID,
			IDCupom:         historico.IDCupom,
			IDUsuario:       historico.IDUsuario,
			DataResgate:     historico.DataResgate,
			DataAtualizacao: historico.DataAtualizacao,
			Status:          historico.Status,
		}

		if historico.Cupom != nil {
			cupomResp := modelToCupomResponse(historico.Cupom)
			resp.Cupom = &cupomResp
		}

		historicosResponse = append(historicosResponse, resp)
	}

	return &json.HistoricosResgateClienteResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}, nil
}

// GetHistoricosResgateByCupomID retorna todos os históricos de resgate de um cupom específico
func GetHistoricosResgateByCupomID(cupomID uint) (*json.HistoricosResgateResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByCupomID(cupomID)
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoResgateResponse
	for _, historico := range historicos {
		historicosResponse = append(historicosResponse, *convertHistoricoToResponse(&historico))
	}

	return &json.HistoricosResgateResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}, nil
}
