package services

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"

	"gorm.io/gorm"
)

func usuarioModelToUserResponse(u models.Usuario) json.UserResponse {
	return json.UserResponse{
		ID:             u.ID,
		Nome:           u.Nome,
		Email:          u.Email,
		CPF:            u.CPF,
		Imagem:         u.Imagem,
		Telefone:       u.Telefone,
		Endereco:       u.Endereco,
		DataNascimento: u.DataNascimento,
		DataCadastro:   u.DataCadastro,
		Ativo:          u.Ativo,
		Latitude:       u.Latitude,
		Longitude:      u.Longitude,
		IDPlano:        u.IDPlano,
		IDLoja:         u.IDLoja,
	}
}

// CreateHistoricoResgateFromCupom cria um histórico de resgate a partir de um cupom
func CreateHistoricoResgateFromCupom(cupomID uint, usuarioID uint, moedasUtilizadas int) (*json.HistoricoResgateResponse, error) {
	historico, err := datasource.CreateHistoricoResgateFromCupom(cupomID, usuarioID, moedasUtilizadas)
	if err != nil {
		return nil, err
	}

	return convertHistoricoToResponse(historico), nil
}

// convertHistoricoToResponse converte um modelo de histórico para response
func convertHistoricoToResponse(historico *models.HistoricoResgate) *json.HistoricoResgateResponse {
	response := &json.HistoricoResgateResponse{
		ID:               historico.ID,
		IDCupom:          historico.IDCupom,
		IDUsuario:        historico.IDUsuario,
		MoedasUtilizadas: historico.MoedasUtilizadas,
		DataResgate:      historico.DataResgate,
		DataAtualizacao:  historico.DataAtualizacao,
		Status:           historico.Status,
		Usuario:          usuarioModelToUserResponse(historico.Usuario),
	}

	if historico.Cupom != nil {
		cupomResp := modelToCupomResponse(historico.Cupom)
		response.Cupom = &cupomResp
	}

	return response
}

func buildHistoricosResgateResponse(historicos []models.HistoricoResgate, vendas []models.VendaProdutoAvulso) *json.HistoricosResgateResponse {
	var historicosResponse []json.HistoricoResgateResponse
	for i := range historicos {
		historicosResponse = append(historicosResponse, *convertHistoricoToResponse(&historicos[i]))
	}
	vendasResp := VendasProdutoAvulsoModelsToResponses(vendas)
	return &json.HistoricosResgateResponse{
		Historicos:               historicosResponse,
		VendasProdutoAvulso:      vendasResp,
		Total:                    len(historicosResponse),
		TotalVendasProdutoAvulso: len(vendasResp),
	}
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

// GetAllHistoricosResgate retorna todos os históricos e vendas avulsas
func GetAllHistoricosResgate() (*json.HistoricosResgateResponse, error) {
	historicos, err := datasource.GetAllHistoricosResgate()
	if err != nil {
		return nil, err
	}
	vendas, err := datasource.GetAllVendasProdutoAvulso()
	if err != nil {
		return nil, err
	}
	return buildHistoricosResgateResponse(historicos, vendas), nil
}

// GetHistoricosResgateByUsuarioID retorna históricos de cupom e vendas avulsas do usuário
func GetHistoricosResgateByUsuarioID(idUsuario uint) (*json.HistoricosResgateResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}
	vendas, err := datasource.GetVendasProdutoAvulsoByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}
	return buildHistoricosResgateResponse(historicos, vendas), nil
}

// GetHistoricosResgateByLojaID retorna históricos de cupom e vendas avulsas da loja
func GetHistoricosResgateByLojaID(idLoja uint) (*json.HistoricosResgateResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByLojaID(idLoja)
	if err != nil {
		return nil, err
	}
	vendas, err := datasource.GetVendasProdutoAvulsoByLojaID(idLoja)
	if err != nil {
		return nil, err
	}
	return buildHistoricosResgateResponse(historicos, vendas), nil
}

// UpdateHistoricoResgate atualiza um histórico existente
func UpdateHistoricoResgate(id uint, req json.HistoricoResgateRequest) (*json.HistoricoResgateResponse, error) {
	historico, err := datasource.UpdateHistoricoResgate(id, req)
	if err != nil {
		return nil, err
	}

	return convertHistoricoToResponse(historico), nil
}

// UpdateStatusHistoricoResgate atualiza o status e, se efetivado, credita moedas por loja na mesma transação (50% do % de desconto geral da loja sobre o valor do cupom).
func UpdateStatusHistoricoResgate(id uint, status string) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := datasource.UpdateStatusHistoricoResgateWithDB(tx, id, status); err != nil {
			return err
		}
		if status == "efetivado" {
			return AplicarMoedasCreditoResgateEfetivadoTx(tx, id)
		}
		return nil
	})
}

// SoftDeleteHistoricoResgate realiza soft delete do histórico
func SoftDeleteHistoricoResgate(id uint) error {
	return datasource.SoftDeleteHistoricoResgate(id)
}

// RestoreHistoricoResgate restaura um histórico que foi soft deleted
func RestoreHistoricoResgate(id uint) error {
	return datasource.RestoreHistoricoResgate(id)
}

// GetHistoricosResgateClienteByUsuarioID retorna histórico simplificado do cliente e vendas avulsas
func GetHistoricosResgateClienteByUsuarioID(usuarioID uint) (*json.HistoricosResgateClienteResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByUsuarioID(usuarioID)
	if err != nil {
		return nil, err
	}
	vendas, err := datasource.GetVendasProdutoAvulsoByUsuarioID(usuarioID)
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoResgateClienteResponse
	for _, historico := range historicos {
		resp := json.HistoricoResgateClienteResponse{
			ID:               historico.ID,
			IDCupom:          historico.IDCupom,
			IDUsuario:        historico.IDUsuario,
			MoedasUtilizadas: historico.MoedasUtilizadas,
			DataResgate:      historico.DataResgate,
			DataAtualizacao:  historico.DataAtualizacao,
			Status:           historico.Status,
		}

		if historico.Cupom != nil {
			cupomResp := modelToCupomResponse(historico.Cupom)
			resp.Cupom = &cupomResp
		}

		historicosResponse = append(historicosResponse, resp)
	}

	vendasResp := VendasProdutoAvulsoModelsToResponses(vendas)

	return &json.HistoricosResgateClienteResponse{
		Historicos:               historicosResponse,
		VendasProdutoAvulso:      vendasResp,
		Total:                    len(historicosResponse),
		TotalVendasProdutoAvulso: len(vendasResp),
	}, nil
}

// GetHistoricosResgateByCupomID retorna todos os históricos de resgate de um cupom específico (sem vendas avulsas)
func GetHistoricosResgateByCupomID(cupomID uint) (*json.HistoricosResgateResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByCupomID(cupomID)
	if err != nil {
		return nil, err
	}

	return buildHistoricosResgateResponse(historicos, nil), nil
}
