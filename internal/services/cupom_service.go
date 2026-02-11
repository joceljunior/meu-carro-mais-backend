package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// modelToCupomResponse converte o model de cupom para response
func modelToCupomResponse(cupom *models.Cupom) json.CupomResponse {
	var precoOriginal float64
	if cupom.Produto != nil && cupom.IDProduto != nil {
		precoOriginal = cupom.Produto.Preco
	} else if cupom.Servico != nil && cupom.IDServico != nil {
		precoOriginal = cupom.Servico.Preco
	} else if cupom.Veiculo != nil && cupom.IDVeiculo != nil {
		precoOriginal = cupom.Preco
	} else {
		precoOriginal = cupom.Preco
	}

	precoComDesconto := cupom.PrecoComDesconto
	if precoComDesconto == 0 && cupom.PorcentagemDesconto > 0 {
		precoComDesconto = precoOriginal * (1 - cupom.PorcentagemDesconto/100)
	} else if precoComDesconto == 0 {
		precoComDesconto = precoOriginal
	}

	var avaliacao *float64
	if cupom.IDLoja != nil && *cupom.IDLoja > 0 {
		estatisticas, err := datasource.GetAvaliacaoEstatisticasByLojaID(*cupom.IDLoja)
		if err == nil && estatisticas != nil && estatisticas.TotalAvaliacoes > 0 {
			avaliacao = &estatisticas.MediaNota
		}
	}

	response := json.CupomResponse{
		ID:                  cupom.ID,
		Titulo:              cupom.Titulo,
		Descricao:           cupom.Descricao,
		Preco:               cupom.Preco,
		Imagem:              cupom.Imagem,
		Destaque:            cupom.Destaque,
		Categoria:           cupom.Categoria,
		IDLoja:              cupom.IDLoja,
		IDProduto:           cupom.IDProduto,
		IDServico:           cupom.IDServico,
		IDVeiculo:           cupom.IDVeiculo,
		IDOfertaAutoMais:    cupom.IDOfertaAutoMais,
		TipoCupom:           cupom.TipoCupom,
		PrecoOriginal:       precoOriginal,
		PrecoComDesconto:    precoComDesconto,
		PorcentagemDesconto: cupom.PorcentagemDesconto,
		Avaliacao:           avaliacao,
	}

	if cupom.Loja != nil {
		response.Loja = &json.LojaResponse{
			ID:             cupom.Loja.ID,
			Nome:           cupom.Loja.Nome,
			CNPJ:           cupom.Loja.CNPJ,
			Imagem:         cupom.Loja.Imagem,
			Endereco:       cupom.Loja.Endereco,
			Latitude:       cupom.Loja.Latitude,
			Longitude:      cupom.Loja.Longitude,
			Rating:         cupom.Loja.Rating,
			IsMeuCarroMais: cupom.Loja.IsMeuCarroMais,
			Categoria:      cupom.Loja.Categoria,
			IDUsuario:      cupom.Loja.IDUsuario,
		}
	}

	if cupom.OfertaAutoMais != nil {
		response.OfertaAutoMais = &json.OfertaAutoMaisResponse{
			ID:              cupom.OfertaAutoMais.ID,
			IDLoja:          cupom.OfertaAutoMais.IDLoja,
			Nome:            cupom.OfertaAutoMais.Nome,
			Descricao:       cupom.OfertaAutoMais.Descricao,
			Moedas:          cupom.OfertaAutoMais.Moedas,
			Porcentagem:     cupom.OfertaAutoMais.Porcentagem,
			Ativo:           cupom.OfertaAutoMais.Ativo,
			DataValidade:    cupom.OfertaAutoMais.DataValidade,
			DataCadastro:    cupom.OfertaAutoMais.DataCadastro,
			DataAtualizacao: cupom.OfertaAutoMais.DataAtualizacao,
		}
	}

	if cupom.Produto != nil {
		response.Produto = &json.ProdutoResponse{
			ID:           cupom.Produto.ID,
			Nome:         cupom.Produto.Nome,
			Descricao:    cupom.Produto.Descricao,
			Preco:        cupom.Produto.Preco,
			Imagem:       cupom.Produto.Imagem,
			Estoque:      cupom.Produto.Estoque,
			Ativo:        cupom.Produto.Ativo,
			Categoria:    cupom.Produto.Categoria,
			IDLoja:       cupom.Produto.IDLoja,
			DataCadastro: cupom.Produto.DataCadastro,
		}
	}

	if cupom.Servico != nil {
		servicoResp := &json.ServicoResponse{
			ID:        cupom.Servico.ID,
			Titulo:    cupom.Servico.Titulo,
			Descricao: cupom.Servico.Descricao,
			Preco:     cupom.Servico.Preco,
			Imagem:    cupom.Servico.Imagem,
			Destaque:  cupom.Servico.Destaque,
			Categoria: cupom.Servico.Categoria,
		}
		if cupom.Loja != nil {
			servicoResp.Rate = cupom.Loja.Rating
			servicoResp.Loja = json.LojaResponse{
				ID:             cupom.Loja.ID,
				Nome:           cupom.Loja.Nome,
				CNPJ:           cupom.Loja.CNPJ,
				Imagem:         cupom.Loja.Imagem,
				Endereco:       cupom.Loja.Endereco,
				Latitude:       cupom.Loja.Latitude,
				Longitude:      cupom.Loja.Longitude,
				Rating:         cupom.Loja.Rating,
				IsMeuCarroMais: cupom.Loja.IsMeuCarroMais,
				Categoria:      cupom.Loja.Categoria,
				IDUsuario:      cupom.Loja.IDUsuario,
			}
		}
		response.Servico = servicoResp
	}

	if cupom.Veiculo != nil {
		response.Veiculo = &json.VeiculoResponse{
			ID:                  cupom.Veiculo.ID,
			Marca:               cupom.Veiculo.Marca,
			Modelo:              cupom.Veiculo.Modelo,
			AnoFabricacao:       cupom.Veiculo.AnoFabricacao,
			AnoModelo:           cupom.Veiculo.AnoModelo,
			Cor:                 cupom.Veiculo.Cor,
			Placa:               cupom.Veiculo.Placa,
			Renavam:             cupom.Veiculo.Renavam,
			Chassi:              cupom.Veiculo.Chassi,
			TipoVeiculo:         cupom.Veiculo.TipoVeiculo,
			Combustivel:         cupom.Veiculo.Combustivel,
			Quilometragem:       cupom.Veiculo.Quilometragem,
			Preco:               cupom.Veiculo.Preco,
			Licenciamento:       cupom.Veiculo.Licenciamento,
			IPVAPago:            cupom.Veiculo.IPVAPago,
			PossuiFinanciamento: cupom.Veiculo.PossuiFinanciamento,
			PossuiMultas:        cupom.Veiculo.PossuiMultas,
			Observacoes:         cupom.Veiculo.Observacoes,
			IDUsuario:           cupom.Veiculo.IDUsuario,
			Ativo:               cupom.Veiculo.Ativo,
			DataCadastro:        cupom.Veiculo.DataCadastro,
		}
	}

	return response
}

// GetCupons retorna todos os cupons
func GetCupons() (*json.CuponsResponse, error) {
	cupons, err := datasource.GetCupons()
	if err != nil {
		return nil, err
	}

	var cuponsResponse []json.CupomResponse
	for _, cupom := range cupons {
		cuponsResponse = append(cuponsResponse, modelToCupomResponse(&cupom))
	}

	response := &json.CuponsResponse{
		Cupons: cuponsResponse,
		Total:  len(cuponsResponse),
	}

	return response, nil
}

// CreateCupom cria um novo cupom
func CreateCupom(req json.CupomRequest) (*json.CupomResponse, error) {
	cupom, err := datasource.CreateCupom(req)
	if err != nil {
		return nil, err
	}

	response := modelToCupomResponse(cupom)
	return &response, nil
}

// GetCupomByID busca um cupom por ID
func GetCupomByID(id uint) (*json.CupomResponse, error) {
	cupom, err := datasource.GetCupomByID(id)
	if err != nil {
		return nil, err
	}

	response := modelToCupomResponse(cupom)
	return &response, nil
}

// GetAllCupons retorna todos os cupons ativos
func GetAllCupons() ([]json.CupomResponse, error) {
	cupons, err := datasource.GetAllCupons()
	if err != nil {
		return nil, err
	}

	var responses []json.CupomResponse
	for _, cupom := range cupons {
		responses = append(responses, modelToCupomResponse(&cupom))
	}

	return responses, nil
}

// GetCuponsByLojaID retorna todos os cupons de uma loja específica
func GetCuponsByLojaID(lojaID uint) (*json.CuponsResponse, error) {
	cupons, err := datasource.GetCuponsByLojaID(lojaID)
	if err != nil {
		return nil, err
	}

	var cuponsResponse []json.CupomResponse
	for _, cupom := range cupons {
		cuponsResponse = append(cuponsResponse, modelToCupomResponse(&cupom))
	}

	response := &json.CuponsResponse{
		Cupons: cuponsResponse,
		Total:  len(cuponsResponse),
	}

	return response, nil
}

// UpdateCupom atualiza um cupom existente
func UpdateCupom(id uint, req json.CupomRequest) (*json.CupomResponse, error) {
	cupom, err := datasource.UpdateCupom(id, req)
	if err != nil {
		return nil, err
	}

	response := modelToCupomResponse(cupom)
	return &response, nil
}

// SoftDeleteCupom realiza soft delete do cupom
func SoftDeleteCupom(id uint) error {
	return datasource.SoftDeleteCupom(id)
}

// RestoreCupom restaura um cupom que foi soft deleted
func RestoreCupom(id uint) error {
	return datasource.RestoreCupom(id)
}

// GetCuponsProdutos retorna todos os cupons de produtos com informações de desconto
func GetCuponsProdutos(latitude, longitude *float64) (*json.CuponsProdutoResponse, error) {
	var cuponsResponse []json.CupomProdutoResponse

	if latitude != nil && longitude != nil {
		cuponsComDistancia, err := datasource.GetCuponsProdutosByProximidade(*latitude, *longitude)
		if err != nil {
			return nil, err
		}

		for _, cupomComDist := range cuponsComDistancia {
			cupom := cupomComDist.Cupom
			if cupom.Produto == nil || cupom.IDProduto == nil {
				continue
			}

			precoOriginal := cupom.Preco
			porcentagemDesconto := cupom.PorcentagemDesconto
			precoComDesconto := precoOriginal

			if porcentagemDesconto > 0 {
				precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
			} else if cupom.PrecoComDesconto > 0 && cupom.PrecoComDesconto < precoOriginal {
				precoComDesconto = cupom.PrecoComDesconto
				porcentagemDesconto = ((precoOriginal - precoComDesconto) / precoOriginal) * 100
			} else {
				if cupom.IDLoja != nil && *cupom.IDLoja > 0 {
					desconto, _ := datasource.GetDescontoAtivoByLojaID(*cupom.IDLoja)
					if desconto != nil {
						porcentagemDesconto = desconto.Porcentagem
						precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
					}
				}
			}

			imagem := cupom.Imagem
			if imagem == "" && cupom.Produto != nil {
				imagem = cupom.Produto.Imagem
			}

			var moedasUtiliza *int
			if cupom.OfertaAutoMais != nil && cupom.OfertaAutoMais.Ativo {
				moedas := cupom.OfertaAutoMais.Moedas
				moedasUtiliza = &moedas
			}

			distancia := cupomComDist.Distancia
			response := json.CupomProdutoResponse{
				ID:                  cupom.ID,
				NomeProduto:         cupom.Produto.Nome,
				NomeLoja:            cupom.Loja.Nome,
				EnderecoLoja:        cupom.Loja.Endereco,
				Imagem:              imagem,
				PrecoOriginal:       precoOriginal,
				PrecoComDesconto:    precoComDesconto,
				PorcentagemDesconto: porcentagemDesconto,
				IsMeuCarroMais:      cupom.Loja.IsMeuCarroMais,
				Categoria:           cupom.Categoria,
				Descricao:           cupom.Descricao,
				Rate:                cupom.Loja.Rating,
				MoedasUtiliza:       moedasUtiliza,
				Distancia:           &distancia,
			}

			cuponsResponse = append(cuponsResponse, response)
		}
	} else {
		cupons, err := datasource.GetCuponsProdutos()
		if err != nil {
			return nil, err
		}

		for _, cupom := range cupons {
			if cupom.Produto == nil || cupom.IDProduto == nil {
				continue
			}

			precoOriginal := cupom.Preco
			porcentagemDesconto := cupom.PorcentagemDesconto
			precoComDesconto := precoOriginal

			if porcentagemDesconto > 0 {
				precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
			} else if cupom.PrecoComDesconto > 0 && cupom.PrecoComDesconto < precoOriginal {
				precoComDesconto = cupom.PrecoComDesconto
				porcentagemDesconto = ((precoOriginal - precoComDesconto) / precoOriginal) * 100
			} else {
				if cupom.IDLoja != nil && *cupom.IDLoja > 0 {
					desconto, _ := datasource.GetDescontoAtivoByLojaID(*cupom.IDLoja)
					if desconto != nil {
						porcentagemDesconto = desconto.Porcentagem
						precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
					}
				}
			}

			imagem := cupom.Imagem
			if imagem == "" && cupom.Produto != nil {
				imagem = cupom.Produto.Imagem
			}

			var moedasUtiliza *int
			if cupom.OfertaAutoMais != nil && cupom.OfertaAutoMais.Ativo {
				moedas := cupom.OfertaAutoMais.Moedas
				moedasUtiliza = &moedas
			}

			response := json.CupomProdutoResponse{
				ID:                  cupom.ID,
				NomeProduto:         cupom.Produto.Nome,
				NomeLoja:            cupom.Loja.Nome,
				EnderecoLoja:        cupom.Loja.Endereco,
				Imagem:              imagem,
				PrecoOriginal:       precoOriginal,
				PrecoComDesconto:    precoComDesconto,
				PorcentagemDesconto: porcentagemDesconto,
				IsMeuCarroMais:      cupom.Loja.IsMeuCarroMais,
				Categoria:           cupom.Categoria,
				Descricao:           cupom.Descricao,
				Rate:                cupom.Loja.Rating,
				MoedasUtiliza:       moedasUtiliza,
			}

			cuponsResponse = append(cuponsResponse, response)
		}
	}

	return &json.CuponsProdutoResponse{
		Cupons: cuponsResponse,
		Total:  len(cuponsResponse),
	}, nil
}

// GetCuponsVeiculos retorna todos os cupons de veículos
func GetCuponsVeiculos(latitude, longitude *float64) (*json.CuponsVeiculoResponse, error) {
	var cuponsResponse []json.CupomVeiculoResponse

	buildVeiculoResponse := func(cupom models.Cupom, distancia *float64) json.CupomVeiculoResponse {
		veiculo := cupom.Veiculo
		imagem := cupom.Imagem
		nomeVeiculo := cupom.Titulo
		if nomeVeiculo == "" {
			nomeVeiculo = veiculo.Modelo
		}

		var moedasUtiliza *int
		if cupom.OfertaAutoMais != nil && cupom.OfertaAutoMais.Ativo {
			moedas := cupom.OfertaAutoMais.Moedas
			moedasUtiliza = &moedas
		}

		isMeuCarroMais := false
		if cupom.Loja != nil {
			isMeuCarroMais = cupom.Loja.IsMeuCarroMais
		}

		var fotos []string
		if veiculo.ID != 0 {
			uploads, err := datasource.GetUploadsByVeiculoID(veiculo.ID)
			if err == nil {
				for _, upload := range uploads {
					fotos = append(fotos, upload.URL)
				}
			}
		}

		var emailAnunciante, telefoneAnunciante, nomeAnunciante *string
		if cupom.Loja == nil && veiculo.Usuario.ID != 0 {
			emailAnunciante = &veiculo.Usuario.Email
			telefoneAnunciante = &veiculo.Usuario.Telefone
			nomeAnunciante = &veiculo.Usuario.Nome
		}

		enderecoLoja := ""
		if cupom.Loja != nil {
			enderecoLoja = cupom.Loja.Endereco
		}

		return json.CupomVeiculoResponse{
			ID:                  cupom.ID,
			NomeVeiculo:         nomeVeiculo,
			KM:                  veiculo.Quilometragem,
			AnoModelo:           veiculo.AnoModelo,
			AnoFabricacao:       &veiculo.AnoFabricacao,
			IsMeuCarroMais:      isMeuCarroMais,
			Preco:               cupom.Preco,
			Imagem:              imagem,
			Fotos:               fotos,
			Modelo:              veiculo.Modelo,
			Marca:               &veiculo.Marca,
			Placa:               veiculo.Placa,
			Renavam:             veiculo.Renavam,
			Chassi:              veiculo.Chassi,
			Cor:                 veiculo.Cor,
			TipoVeiculo:         veiculo.TipoVeiculo,
			Licenciamento:       veiculo.Licenciamento,
			IPVAPago:            veiculo.IPVAPago,
			PossuiFinanciamento: veiculo.PossuiFinanciamento,
			PossuiMultas:        veiculo.PossuiMultas,
			Observacoes:         veiculo.Observacoes,
			Combustivel:         veiculo.Combustivel,
			MoedasUtiliza:       moedasUtiliza,
			Distancia:           distancia,
			EnderecoLoja:        enderecoLoja,
			EmailAnunciante:     emailAnunciante,
			TelefoneAnunciante:  telefoneAnunciante,
			NomeAnunciante:      nomeAnunciante,
		}
	}

	if latitude != nil && longitude != nil {
		cuponsComDistancia, err := datasource.GetCuponsVeiculosByProximidade(*latitude, *longitude)
		if err != nil {
			return nil, err
		}

		for _, cupomComDist := range cuponsComDistancia {
			cupom := cupomComDist.Cupom
			if cupom.Veiculo == nil || cupom.IDVeiculo == nil {
				continue
			}
			distancia := cupomComDist.Distancia
			cuponsResponse = append(cuponsResponse, buildVeiculoResponse(cupom, &distancia))
		}
	} else {
		cupons, err := datasource.GetCuponsVeiculos()
		if err != nil {
			return nil, err
		}

		for _, cupom := range cupons {
			if cupom.Veiculo == nil || cupom.IDVeiculo == nil {
				continue
			}
			cuponsResponse = append(cuponsResponse, buildVeiculoResponse(cupom, nil))
		}
	}

	return &json.CuponsVeiculoResponse{
		Cupons: cuponsResponse,
		Total:  len(cuponsResponse),
	}, nil
}

// GetCuponsServicos retorna todos os cupons de serviços
func GetCuponsServicos(latitude, longitude *float64) (*json.CuponsServicoResponse, error) {
	var cuponsResponse []json.CupomServicoResponse

	buildServicoResponse := func(cupom models.Cupom, distancia *float64) json.CupomServicoResponse {
		servico := cupom.Servico
		precoOriginal := cupom.Preco
		porcentagemDesconto := cupom.PorcentagemDesconto
		precoComDesconto := precoOriginal

		if porcentagemDesconto > 0 {
			precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
		} else if cupom.PrecoComDesconto > 0 && cupom.PrecoComDesconto < precoOriginal {
			precoComDesconto = cupom.PrecoComDesconto
			porcentagemDesconto = ((precoOriginal - precoComDesconto) / precoOriginal) * 100
		} else {
			if cupom.IDLoja != nil && *cupom.IDLoja > 0 {
				desconto, _ := datasource.GetDescontoAtivoByLojaID(*cupom.IDLoja)
				if desconto != nil {
					porcentagemDesconto = desconto.Porcentagem
					precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
				}
			}
		}

		imagem := cupom.Imagem
		if imagem == "" && servico.Imagem != "" {
			imagem = servico.Imagem
		}

		nomeServico := cupom.Titulo
		if nomeServico == "" {
			nomeServico = servico.Titulo
		}

		var moedasUtiliza *int
		if cupom.OfertaAutoMais != nil && cupom.OfertaAutoMais.Ativo {
			moedas := cupom.OfertaAutoMais.Moedas
			moedasUtiliza = &moedas
		}

		return json.CupomServicoResponse{
			ID:                  cupom.ID,
			NomeServico:         nomeServico,
			NomeLoja:            cupom.Loja.Nome,
			EnderecoLoja:        cupom.Loja.Endereco,
			Imagem:              imagem,
			PrecoOriginal:       precoOriginal,
			PrecoComDesconto:    precoComDesconto,
			PorcentagemDesconto: porcentagemDesconto,
			IsMeuCarroMais:      cupom.Loja.IsMeuCarroMais,
			Categoria:           cupom.Categoria,
			Descricao:           cupom.Descricao,
			Rate:                cupom.Loja.Rating,
			MoedasUtiliza:       moedasUtiliza,
			Distancia:           distancia,
		}
	}

	if latitude != nil && longitude != nil {
		cuponsComDistancia, err := datasource.GetCuponsServicosByProximidade(*latitude, *longitude)
		if err != nil {
			return nil, err
		}

		for _, cupomComDist := range cuponsComDistancia {
			cupom := cupomComDist.Cupom
			if cupom.Servico == nil || cupom.IDServico == nil {
				continue
			}
			distancia := cupomComDist.Distancia
			cuponsResponse = append(cuponsResponse, buildServicoResponse(cupom, &distancia))
		}
	} else {
		cupons, err := datasource.GetCuponsServicos()
		if err != nil {
			return nil, err
		}

		for _, cupom := range cupons {
			if cupom.Servico == nil || cupom.IDServico == nil {
				continue
			}
			cuponsResponse = append(cuponsResponse, buildServicoResponse(cupom, nil))
		}
	}

	return &json.CuponsServicoResponse{
		Cupons: cuponsResponse,
		Total:  len(cuponsResponse),
	}, nil
}
