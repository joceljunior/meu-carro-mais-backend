package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// CreateHistoricoResgateFromAnuncio cria um histórico de resgate a partir de um anúncio
func CreateHistoricoResgateFromAnuncio(anuncioID uint, usuarioID uint) (*json.HistoricoResgateResponse, error) {
	historico, err := datasource.CreateHistoricoResgateFromAnuncio(anuncioID, usuarioID)
	if err != nil {
		return nil, err
	}

	// Usa a mesma lógica de conversão do CreateHistoricoResgate
	return convertHistoricoToResponse(historico), nil
}

// convertHistoricoToResponse converte um modelo de histórico para response
func convertHistoricoToResponse(historico *models.HistoricoResgate) *json.HistoricoResgateResponse {
	response := &json.HistoricoResgateResponse{
		ID:                  historico.ID,
		IDUsuario:           historico.IDUsuario,
		IDProduto:           historico.IDProduto,
		IDServico:           historico.IDServico,
		IDVeiculo:           historico.IDVeiculo,
		IDLoja:              historico.IDLoja,
		TipoResgate:         historico.TipoResgate,
		Quantidade:          historico.Quantidade,
		ValorUnitario:       historico.ValorUnitario,
		ValorOriginal:       historico.ValorOriginal,
		DescontoAplicado:    historico.DescontoAplicado,
		PorcentagemDesconto: historico.PorcentagemDesconto,
		Valor:               historico.Valor,
		Status:              historico.Status,
		DataResgate:         historico.DataResgate,
		DataAtualizacao:     historico.DataAtualizacao,
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
		Loja: json.LojaResponse{
			ID:        historico.Loja.ID,
			Nome:      historico.Loja.Nome,
			CNPJ:      historico.Loja.CNPJ,
			Imagem:    historico.Loja.Imagem,
			Latitude:  historico.Loja.Latitude,
			Longitude: historico.Loja.Longitude,
			Categoria: historico.Loja.Categoria,
			IDUsuario: historico.Loja.IDUsuario,
		},
	}

	// Adiciona dados do produto se existir
	if historico.Produto != nil {
		produtoResp := &json.ProdutoResponse{
			ID:           historico.Produto.ID,
			Nome:         historico.Produto.Nome,
			Descricao:    historico.Produto.Descricao,
			Preco:        historico.Produto.Preco,
			Imagem:       historico.Produto.Imagem,
			Estoque:      historico.Produto.Estoque,
			Ativo:        historico.Produto.Ativo,
			IDLoja:       historico.Produto.IDLoja,
			DataCadastro: historico.Produto.DataCadastro,
		}
		response.Produto = produtoResp
	}

	// Adiciona dados do serviço se existir
	if historico.Servico != nil {
		servicoResp := &json.ServicoResponse{
			ID:        historico.Servico.ID,
			Titulo:    historico.Servico.Titulo,
			Descricao: historico.Servico.Descricao,
			Preco:     historico.Servico.Preco,
			Imagem:    historico.Servico.Imagem,
			Destaque:  historico.Servico.Destaque,
			Categoria: historico.Servico.Categoria,
		}
		response.Servico = servicoResp
	}

	// Adiciona dados do veículo se existir
	if historico.Veiculo != nil {
		veiculoResp := &json.VeiculoResponse{
			ID:                  historico.Veiculo.ID,
			Marca:               historico.Veiculo.Marca,
			Modelo:              historico.Veiculo.Modelo,
			AnoFabricacao:       historico.Veiculo.AnoFabricacao,
			AnoModelo:           historico.Veiculo.AnoModelo,
			Cor:                 historico.Veiculo.Cor,
			Placa:               historico.Veiculo.Placa,
			Renavam:             historico.Veiculo.Renavam,
			Chassi:              historico.Veiculo.Chassi,
			TipoVeiculo:         historico.Veiculo.TipoVeiculo,
			Combustivel:         historico.Veiculo.Combustivel,
			Quilometragem:       historico.Veiculo.Quilometragem,
			Preco:               historico.Veiculo.Preco,
			Licenciamento:       historico.Veiculo.Licenciamento,
			IPVAPago:            historico.Veiculo.IPVAPago,
			PossuiFinanciamento: historico.Veiculo.PossuiFinanciamento,
			PossuiMultas:        historico.Veiculo.PossuiMultas,
			Observacoes:         historico.Veiculo.Observacoes,
			IDUsuario:           historico.Veiculo.IDUsuario,
			DataCadastro:        historico.Veiculo.DataCadastro,
			Ativo:               historico.Veiculo.Ativo,
		}
		response.Veiculo = veiculoResp
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
		// Monta os itens comprados
		var itens []json.ItemCompraResponse

		if historico.Produto != nil {
			itens = append(itens, json.ItemCompraResponse{
				ID:            historico.Produto.ID,
				Nome:          historico.Produto.Nome,
				Descricao:     historico.Produto.Descricao,
				Imagem:        historico.Produto.Imagem,
				TipoItem:      "produto",
				Quantidade:    historico.Quantidade,
				ValorUnitario: historico.ValorUnitario,
			})
		}

		if historico.Servico != nil {
			itens = append(itens, json.ItemCompraResponse{
				ID:            historico.Servico.ID,
				Nome:          historico.Servico.Titulo,
				Descricao:     historico.Servico.Descricao,
				Imagem:        historico.Servico.Imagem,
				TipoItem:      "servico",
				Quantidade:    historico.Quantidade,
				ValorUnitario: historico.ValorUnitario,
			})
		}

		if historico.Veiculo != nil {
			nomeVeiculo := historico.Veiculo.Marca + " " + historico.Veiculo.Modelo
			itens = append(itens, json.ItemCompraResponse{
				ID:            historico.Veiculo.ID,
				Nome:          nomeVeiculo,
				Descricao:     historico.Veiculo.Cor,
				Imagem:        "", // Veículo não tem campo imagem direto, buscar do upload
				TipoItem:      "veiculo",
				Quantidade:    historico.Quantidade,
				ValorUnitario: historico.ValorUnitario,
			})
		}

		// Busca a avaliação do cliente para esta loja
		var avaliacaoResp *json.AvaliacaoClienteResponse
		avaliacao, err := datasource.GetAvaliacaoByUsuarioELoja(usuarioID, historico.IDLoja)
		if err == nil && avaliacao != nil {
			avaliacaoResp = &json.AvaliacaoClienteResponse{
				ID:            avaliacao.ID,
				Nota:          avaliacao.Nota,
				Comentario:    avaliacao.Comentario,
				DataAvaliacao: avaliacao.DataAvaliacao,
			}
		}

		historicosResponse = append(historicosResponse, json.HistoricoResgateClienteResponse{
			ID:                  historico.ID,
			IDLoja:              historico.IDLoja,
			NomeLoja:            historico.Loja.Nome,
			ImagemLoja:          historico.Loja.Imagem,
			DataResgate:         historico.DataResgate,
			Status:              historico.Status,
			Itens:               itens,
			Quantidade:          historico.Quantidade,
			ValorUnitario:       historico.ValorUnitario,
			ValorOriginal:       historico.ValorOriginal,
			DescontoAplicado:    historico.DescontoAplicado,
			PorcentagemDesconto: historico.PorcentagemDesconto,
			ValorTotal:          historico.Valor,
			Avaliacao:           avaliacaoResp,
		})
	}

	return &json.HistoricosResgateClienteResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}, nil
}

// GetHistoricosResgateByAnuncioID retorna todos os históricos de resgate de um anúncio específico
func GetHistoricosResgateByAnuncioID(anuncioID uint) (*json.HistoricosResgateResponse, error) {
	historicos, err := datasource.GetHistoricosResgateByAnuncioID(anuncioID)
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