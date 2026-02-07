package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// modelToAnuncioResponse converte o model de anúncio para response
func modelToAnuncioResponse(anuncio *models.Anuncio) json.AnuncioResponse {
	// Calcula o preço original baseado no produto/serviço/veículo
	var precoOriginal float64
	if anuncio.Produto != nil && anuncio.IDProduto != nil {
		precoOriginal = anuncio.Produto.Preco
	} else if anuncio.Servico != nil && anuncio.IDServico != nil {
		precoOriginal = anuncio.Servico.Preco
	} else if anuncio.Veiculo != nil && anuncio.IDVeiculo != nil {
		// Para veículos, usa o preço do anúncio como original (veículos não têm preço próprio)
		precoOriginal = anuncio.Preco
	} else {
		// Fallback: usa o preço do anúncio
		precoOriginal = anuncio.Preco
	}

	// Usa o preço com desconto do anúncio, ou calcula se não estiver definido
	precoComDesconto := anuncio.PrecoComDesconto
	if precoComDesconto == 0 && anuncio.PorcentagemDesconto > 0 {
		precoComDesconto = precoOriginal * (1 - anuncio.PorcentagemDesconto/100)
	} else if precoComDesconto == 0 {
		precoComDesconto = precoOriginal
	}

	// Busca a avaliação média da loja (apenas se tiver loja)
	var avaliacao *float64
	if anuncio.IDLoja != nil && *anuncio.IDLoja > 0 {
		estatisticas, err := datasource.GetAvaliacaoEstatisticasByLojaID(*anuncio.IDLoja)
		if err == nil && estatisticas != nil && estatisticas.TotalAvaliacoes > 0 {
			avaliacao = &estatisticas.MediaNota
		}
	}

	response := json.AnuncioResponse{
		ID:                 anuncio.ID,
		Titulo:             anuncio.Titulo,
		Descricao:          anuncio.Descricao,
		Preco:              anuncio.Preco,
		Imagem:             anuncio.Imagem,
		Destaque:           anuncio.Destaque,
		Categoria:          anuncio.Categoria,
		IDLoja:             anuncio.IDLoja,
		IDProduto:          anuncio.IDProduto,
		IDServico:          anuncio.IDServico,
		IDVeiculo:          anuncio.IDVeiculo,
		IDOfertaAutoMais:   anuncio.IDOfertaAutoMais,
		TipoAnuncio:        anuncio.TipoAnuncio,
		PrecoOriginal:      precoOriginal,
		PrecoComDesconto:   precoComDesconto,
		PorcentagemDesconto: anuncio.PorcentagemDesconto,
		Avaliacao:          avaliacao,
	}

	// Adiciona a loja apenas se existir
	if anuncio.Loja != nil {
		response.Loja = &json.LojaResponse{
			ID:             anuncio.Loja.ID,
			Nome:           anuncio.Loja.Nome,
			CNPJ:           anuncio.Loja.CNPJ,
			Imagem:         anuncio.Loja.Imagem,
			Latitude:       anuncio.Loja.Latitude,
			Longitude:      anuncio.Loja.Longitude,
			Rating:         anuncio.Loja.Rating,
			IsMeuCarroMais: anuncio.Loja.IsMeuCarroMais,
			Categoria:      anuncio.Loja.Categoria,
			IDUsuario:      anuncio.Loja.IDUsuario,
		}
	}

	// Inclui a oferta Auto Mais se existir
	if anuncio.OfertaAutoMais != nil {
		response.OfertaAutoMais = &json.OfertaAutoMaisResponse{
			ID:              anuncio.OfertaAutoMais.ID,
			IDLoja:          anuncio.OfertaAutoMais.IDLoja,
			Nome:            anuncio.OfertaAutoMais.Nome,
			Descricao:       anuncio.OfertaAutoMais.Descricao,
			Moedas:          anuncio.OfertaAutoMais.Moedas,
			Porcentagem:     anuncio.OfertaAutoMais.Porcentagem,
			Ativo:           anuncio.OfertaAutoMais.Ativo,
			DataValidade:    anuncio.OfertaAutoMais.DataValidade,
			DataCadastro:    anuncio.OfertaAutoMais.DataCadastro,
			DataAtualizacao: anuncio.OfertaAutoMais.DataAtualizacao,
		}
	}

	// Inclui produto, serviço ou veículo se existir
	if anuncio.Produto != nil {
		response.Produto = &json.ProdutoResponse{
			ID:           anuncio.Produto.ID,
			Nome:         anuncio.Produto.Nome,
			Descricao:    anuncio.Produto.Descricao,
			Preco:        anuncio.Produto.Preco,
			Imagem:       anuncio.Produto.Imagem,
			Estoque:      anuncio.Produto.Estoque,
			Ativo:        anuncio.Produto.Ativo,
			Categoria:    anuncio.Produto.Categoria,
			IDLoja:       anuncio.Produto.IDLoja,
			DataCadastro: anuncio.Produto.DataCadastro,
		}
	}

	if anuncio.Servico != nil {
		servicoResp := &json.ServicoResponse{
			ID:        anuncio.Servico.ID,
			Titulo:    anuncio.Servico.Titulo,
			Descricao: anuncio.Servico.Descricao,
			Preco:     anuncio.Servico.Preco,
			Imagem:    anuncio.Servico.Imagem,
			Destaque:  anuncio.Servico.Destaque,
			Categoria: anuncio.Servico.Categoria,
		}
		// Adiciona dados da loja apenas se existir
		if anuncio.Loja != nil {
			servicoResp.Rate = anuncio.Loja.Rating
			servicoResp.Loja = json.LojaResponse{
				ID:             anuncio.Loja.ID,
				Nome:           anuncio.Loja.Nome,
				CNPJ:           anuncio.Loja.CNPJ,
				Imagem:         anuncio.Loja.Imagem,
				Latitude:       anuncio.Loja.Latitude,
				Longitude:      anuncio.Loja.Longitude,
				Rating:         anuncio.Loja.Rating,
				IsMeuCarroMais: anuncio.Loja.IsMeuCarroMais,
				Categoria:      anuncio.Loja.Categoria,
				IDUsuario:      anuncio.Loja.IDUsuario,
			}
		}
		response.Servico = servicoResp
	}

	if anuncio.Veiculo != nil {
		response.Veiculo = &json.VeiculoResponse{
			ID:                  anuncio.Veiculo.ID,
			Marca:               anuncio.Veiculo.Marca,
			Modelo:              anuncio.Veiculo.Modelo,
			AnoFabricacao:       anuncio.Veiculo.AnoFabricacao,
			AnoModelo:           anuncio.Veiculo.AnoModelo,
			Cor:                 anuncio.Veiculo.Cor,
			Placa:               anuncio.Veiculo.Placa,
			Renavam:             anuncio.Veiculo.Renavam,
			Chassi:              anuncio.Veiculo.Chassi,
			TipoVeiculo:         anuncio.Veiculo.TipoVeiculo,
			Combustivel:         anuncio.Veiculo.Combustivel,
			Quilometragem:       anuncio.Veiculo.Quilometragem,
			Preco:               anuncio.Veiculo.Preco,
			Licenciamento:       anuncio.Veiculo.Licenciamento,
			IPVAPago:            anuncio.Veiculo.IPVAPago,
			PossuiFinanciamento: anuncio.Veiculo.PossuiFinanciamento,
			PossuiMultas:        anuncio.Veiculo.PossuiMultas,
			Observacoes:         anuncio.Veiculo.Observacoes,
			IDUsuario:           anuncio.Veiculo.IDUsuario,
			Ativo:               anuncio.Veiculo.Ativo,
			DataCadastro:        anuncio.Veiculo.DataCadastro,
		}
	}

	return response
}

// GetAnuncios retorna todos os anúncios
func GetAnuncios() (*json.AnunciosResponse, error) {
	anuncios, err := datasource.GetAnuncios()
	if err != nil {
		return nil, err
	}

	var anunciosResponse []json.AnuncioResponse
	for _, anuncio := range anuncios {
		anunciosResponse = append(anunciosResponse, modelToAnuncioResponse(&anuncio))
	}

	response := &json.AnunciosResponse{
		Anuncios: anunciosResponse,
		Total:    len(anunciosResponse),
	}

	return response, nil
}

// CreateAnuncio cria um novo anúncio
func CreateAnuncio(req json.AnuncioRequest) (*json.AnuncioResponse, error) {
	anuncio, err := datasource.CreateAnuncio(req)
	if err != nil {
		return nil, err
	}

	response := modelToAnuncioResponse(anuncio)
	return &response, nil
}

// GetAnuncioByID busca um anúncio por ID
func GetAnuncioByID(id uint) (*json.AnuncioResponse, error) {
	anuncio, err := datasource.GetAnuncioByID(id)
	if err != nil {
		return nil, err
	}

	response := modelToAnuncioResponse(anuncio)
	return &response, nil
}

// GetAllAnuncios retorna todos os anúncios ativos
func GetAllAnuncios() ([]json.AnuncioResponse, error) {
	anuncios, err := datasource.GetAllAnuncios()
	if err != nil {
		return nil, err
	}

	var responses []json.AnuncioResponse
	for _, anuncio := range anuncios {
		responses = append(responses, modelToAnuncioResponse(&anuncio))
	}

	return responses, nil
}

// GetAnunciosByLojaID retorna todos os anúncios de uma loja específica
func GetAnunciosByLojaID(lojaID uint) (*json.AnunciosResponse, error) {
	anuncios, err := datasource.GetAnunciosByLojaID(lojaID)
	if err != nil {
		return nil, err
	}

	var anunciosResponse []json.AnuncioResponse
	for _, anuncio := range anuncios {
		anunciosResponse = append(anunciosResponse, modelToAnuncioResponse(&anuncio))
	}

	response := &json.AnunciosResponse{
		Anuncios: anunciosResponse,
		Total:    len(anunciosResponse),
	}

	return response, nil
}

// UpdateAnuncio atualiza um anúncio existente
func UpdateAnuncio(id uint, req json.AnuncioRequest) (*json.AnuncioResponse, error) {
	anuncio, err := datasource.UpdateAnuncio(id, req)
	if err != nil {
		return nil, err
	}

	response := modelToAnuncioResponse(anuncio)
	return &response, nil
}

// SoftDeleteAnuncio realiza soft delete do anúncio
func SoftDeleteAnuncio(id uint) error {
	return datasource.SoftDeleteAnuncio(id)
}

// RestoreAnuncio restaura um anúncio que foi soft deleted
func RestoreAnuncio(id uint) error {
	return datasource.RestoreAnuncio(id)
}

// GetAnunciosProdutos retorna todos os anúncios de produtos com informações de desconto
// Se latitude e longitude forem fornecidos, ordena por proximidade
func GetAnunciosProdutos(latitude, longitude *float64) (*json.AnunciosProdutoResponse, error) {
	var anunciosResponse []json.AnuncioProdutoResponse

	// Se latitude e longitude foram fornecidos, busca por proximidade
	if latitude != nil && longitude != nil {
		anunciosComDistancia, err := datasource.GetAnunciosProdutosByProximidade(*latitude, *longitude)
		if err != nil {
			return nil, err
		}

		for _, anuncioComDist := range anunciosComDistancia {
			anuncio := anuncioComDist.Anuncio
			// Verifica se o anúncio tem produto associado
			if anuncio.Produto == nil || anuncio.IDProduto == nil {
				continue
			}

			// Calcula preço com desconto
			// Prioridade: 1) porcentagem do anúncio, 2) preço com desconto do anúncio, 3) desconto da loja
			precoOriginal := anuncio.Preco
			porcentagemDesconto := anuncio.PorcentagemDesconto
			precoComDesconto := precoOriginal

			if porcentagemDesconto > 0 {
				// Usa a porcentagem do próprio anúncio
				precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
			} else if anuncio.PrecoComDesconto > 0 && anuncio.PrecoComDesconto < precoOriginal {
				// Usa o preço com desconto já definido no anúncio
				precoComDesconto = anuncio.PrecoComDesconto
				porcentagemDesconto = ((precoOriginal - precoComDesconto) / precoOriginal) * 100
			} else {
				// Busca desconto ativo da loja como fallback
				if anuncio.IDLoja != nil && *anuncio.IDLoja > 0 {
					desconto, _ := datasource.GetDescontoAtivoByLojaID(*anuncio.IDLoja)
					if desconto != nil {
						porcentagemDesconto = desconto.Porcentagem
						precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
					}
				}
			}

			// Determina a imagem (prioridade: anúncio > produto)
			imagem := anuncio.Imagem
			if imagem == "" && anuncio.Produto != nil {
				imagem = anuncio.Produto.Imagem
			}

			// Moedas da oferta Auto Mais
			var moedasUtiliza *int
			if anuncio.OfertaAutoMais != nil && anuncio.OfertaAutoMais.Ativo {
				moedas := anuncio.OfertaAutoMais.Moedas
				moedasUtiliza = &moedas
			}

			distancia := anuncioComDist.Distancia
			response := json.AnuncioProdutoResponse{
				ID:                  anuncio.ID,
				NomeProduto:         anuncio.Produto.Nome,
				NomeLoja:            anuncio.Loja.Nome,
				Imagem:              imagem,
				PrecoOriginal:       precoOriginal,
				PrecoComDesconto:    precoComDesconto,
				PorcentagemDesconto: porcentagemDesconto,
				IsMeuCarroMais:      anuncio.Loja.IsMeuCarroMais,
				Categoria:           anuncio.Categoria,
				Descricao:           anuncio.Descricao,
				Rate:                anuncio.Loja.Rating,
				MoedasUtiliza:       moedasUtiliza,
				Distancia:           &distancia,
			}

			anunciosResponse = append(anunciosResponse, response)
		}
	} else {
		// Busca sem ordenação por proximidade
		anuncios, err := datasource.GetAnunciosProdutos()
		if err != nil {
			return nil, err
		}

		for _, anuncio := range anuncios {
			// Verifica se o anúncio tem produto associado
			if anuncio.Produto == nil || anuncio.IDProduto == nil {
				continue
			}

			// Calcula preço com desconto
			// Prioridade: 1) porcentagem do anúncio, 2) preço com desconto do anúncio, 3) desconto da loja
			precoOriginal := anuncio.Preco
			porcentagemDesconto := anuncio.PorcentagemDesconto
			precoComDesconto := precoOriginal

			if porcentagemDesconto > 0 {
				// Usa a porcentagem do próprio anúncio
				precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
			} else if anuncio.PrecoComDesconto > 0 && anuncio.PrecoComDesconto < precoOriginal {
				// Usa o preço com desconto já definido no anúncio
				precoComDesconto = anuncio.PrecoComDesconto
				porcentagemDesconto = ((precoOriginal - precoComDesconto) / precoOriginal) * 100
			} else {
				// Busca desconto ativo da loja como fallback
				if anuncio.IDLoja != nil && *anuncio.IDLoja > 0 {
					desconto, _ := datasource.GetDescontoAtivoByLojaID(*anuncio.IDLoja)
					if desconto != nil {
						porcentagemDesconto = desconto.Porcentagem
						precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
					}
				}
			}

			// Determina a imagem (prioridade: anúncio > produto)
			imagem := anuncio.Imagem
			if imagem == "" && anuncio.Produto != nil {
				imagem = anuncio.Produto.Imagem
			}

			// Moedas da oferta Auto Mais
			var moedasUtiliza *int
			if anuncio.OfertaAutoMais != nil && anuncio.OfertaAutoMais.Ativo {
				moedas := anuncio.OfertaAutoMais.Moedas
				moedasUtiliza = &moedas
			}

			response := json.AnuncioProdutoResponse{
				ID:                  anuncio.ID,
				NomeProduto:         anuncio.Produto.Nome,
				NomeLoja:            anuncio.Loja.Nome,
				Imagem:              imagem,
				PrecoOriginal:       precoOriginal,
				PrecoComDesconto:    precoComDesconto,
				PorcentagemDesconto: porcentagemDesconto,
				IsMeuCarroMais:      anuncio.Loja.IsMeuCarroMais,
				Categoria:           anuncio.Categoria,
				Descricao:           anuncio.Descricao,
				Rate:                anuncio.Loja.Rating,
				MoedasUtiliza:       moedasUtiliza,
			}

			anunciosResponse = append(anunciosResponse, response)
		}
	}

	return &json.AnunciosProdutoResponse{
		Anuncios: anunciosResponse,
		Total:    len(anunciosResponse),
	}, nil
}

// GetAnunciosVeiculos retorna todos os anúncios de veículos com informações detalhadas
// Se latitude e longitude forem fornecidos, ordena por proximidade
func GetAnunciosVeiculos(latitude, longitude *float64) (*json.AnunciosVeiculoResponse, error) {
	var anunciosResponse []json.AnuncioVeiculoResponse

	// Se latitude e longitude foram fornecidos, busca por proximidade
	if latitude != nil && longitude != nil {
		anunciosComDistancia, err := datasource.GetAnunciosVeiculosByProximidade(*latitude, *longitude)
		if err != nil {
			return nil, err
		}

		for _, anuncioComDist := range anunciosComDistancia {
			anuncio := anuncioComDist.Anuncio
			// Verifica se o anúncio tem veículo associado
			if anuncio.Veiculo == nil || anuncio.IDVeiculo == nil {
				continue
			}

			veiculo := anuncio.Veiculo

			// Determina a imagem (prioridade: anúncio > veículo)
			imagem := anuncio.Imagem
			if imagem == "" {
				// Se não houver imagem no anúncio, pode buscar de outro lugar se necessário
				imagem = ""
			}

			// Nome do veículo (usa Titulo do anúncio ou Modelo do veículo)
			nomeVeiculo := anuncio.Titulo
			if nomeVeiculo == "" {
				nomeVeiculo = veiculo.Modelo
			}

			// Moedas da oferta Auto Mais
			var moedasUtiliza *int
			if anuncio.OfertaAutoMais != nil && anuncio.OfertaAutoMais.Ativo {
				moedas := anuncio.OfertaAutoMais.Moedas
				moedasUtiliza = &moedas
			}

			// IsMeuCarroMais só existe se tiver loja
			isMeuCarroMais := false
			if anuncio.Loja != nil {
				isMeuCarroMais = anuncio.Loja.IsMeuCarroMais
			}

			// Busca fotos do veículo
			var fotos []string
			if veiculo.ID != 0 {
				uploads, err := datasource.GetUploadsByVeiculoID(veiculo.ID)
				if err == nil {
					for _, upload := range uploads {
						fotos = append(fotos, upload.URL)
					}
				}
			}

			// Dados do anunciante (apenas quando não é loja)
			var emailAnunciante, telefoneAnunciante, nomeAnunciante *string
			if anuncio.Loja == nil && veiculo.Usuario.ID != 0 {
				emailAnunciante = &veiculo.Usuario.Email
				telefoneAnunciante = &veiculo.Usuario.Telefone
				nomeAnunciante = &veiculo.Usuario.Nome
			}

			distancia := anuncioComDist.Distancia
			response := json.AnuncioVeiculoResponse{
				ID:                  anuncio.ID,
				NomeVeiculo:         nomeVeiculo,
				KM:                  veiculo.Quilometragem,
				AnoModelo:           veiculo.AnoModelo,
				AnoFabricacao:       &veiculo.AnoFabricacao,
				IsMeuCarroMais:      isMeuCarroMais,
				Preco:               anuncio.Preco,
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
				Distancia:           &distancia,
				EmailAnunciante:     emailAnunciante,
				TelefoneAnunciante:  telefoneAnunciante,
				NomeAnunciante:      nomeAnunciante,
			}

			anunciosResponse = append(anunciosResponse, response)
		}
	} else {
		// Busca sem ordenação por proximidade
		anuncios, err := datasource.GetAnunciosVeiculos()
		if err != nil {
			return nil, err
		}

		for _, anuncio := range anuncios {
			// Verifica se o anúncio tem veículo associado
			if anuncio.Veiculo == nil || anuncio.IDVeiculo == nil {
				continue
			}

			veiculo := anuncio.Veiculo

			// Determina a imagem (prioridade: anúncio > veículo)
			imagem := anuncio.Imagem
			if imagem == "" {
				// Se não houver imagem no anúncio, pode buscar de outro lugar se necessário
				imagem = ""
			}

			// Nome do veículo (usa Titulo do anúncio ou Modelo do veículo)
			nomeVeiculo := anuncio.Titulo
			if nomeVeiculo == "" {
				nomeVeiculo = veiculo.Modelo
			}

			// Moedas da oferta Auto Mais
			var moedasUtiliza *int
			if anuncio.OfertaAutoMais != nil && anuncio.OfertaAutoMais.Ativo {
				moedas := anuncio.OfertaAutoMais.Moedas
				moedasUtiliza = &moedas
			}

			// IsMeuCarroMais só existe se tiver loja
			isMeuCarroMais := false
			if anuncio.Loja != nil {
				isMeuCarroMais = anuncio.Loja.IsMeuCarroMais
			}

			// Busca fotos do veículo
			var fotos []string
			if veiculo.ID != 0 {
				uploads, err := datasource.GetUploadsByVeiculoID(veiculo.ID)
				if err == nil {
					for _, upload := range uploads {
						fotos = append(fotos, upload.URL)
					}
				}
			}

			// Dados do anunciante (apenas quando não é loja)
			var emailAnunciante, telefoneAnunciante, nomeAnunciante *string
			if anuncio.Loja == nil && veiculo.Usuario.ID != 0 {
				emailAnunciante = &veiculo.Usuario.Email
				telefoneAnunciante = &veiculo.Usuario.Telefone
				nomeAnunciante = &veiculo.Usuario.Nome
			}

			response := json.AnuncioVeiculoResponse{
				ID:                  anuncio.ID,
				NomeVeiculo:         nomeVeiculo,
				KM:                  veiculo.Quilometragem,
				AnoModelo:           veiculo.AnoModelo,
				AnoFabricacao:       &veiculo.AnoFabricacao,
				IsMeuCarroMais:      isMeuCarroMais,
				Preco:               anuncio.Preco,
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
				EmailAnunciante:     emailAnunciante,
				TelefoneAnunciante:  telefoneAnunciante,
				NomeAnunciante:      nomeAnunciante,
			}

			anunciosResponse = append(anunciosResponse, response)
		}
	}

	return &json.AnunciosVeiculoResponse{
		Anuncios: anunciosResponse,
		Total:    len(anunciosResponse),
	}, nil
}

// GetAnunciosServicos retorna todos os anúncios de serviços com informações de desconto
// Se latitude e longitude forem fornecidos, ordena por proximidade
func GetAnunciosServicos(latitude, longitude *float64) (*json.AnunciosServicoResponse, error) {
	var anunciosResponse []json.AnuncioServicoResponse

	// Se latitude e longitude foram fornecidos, busca por proximidade
	if latitude != nil && longitude != nil {
		anunciosComDistancia, err := datasource.GetAnunciosServicosByProximidade(*latitude, *longitude)
		if err != nil {
			return nil, err
		}

		for _, anuncioComDist := range anunciosComDistancia {
			anuncio := anuncioComDist.Anuncio
			// Verifica se o anúncio tem serviço associado
			if anuncio.Servico == nil || anuncio.IDServico == nil {
				continue
			}

			servico := anuncio.Servico

			// Calcula preço com desconto
			// Prioridade: 1) porcentagem do anúncio, 2) preço com desconto do anúncio, 3) desconto da loja
			precoOriginal := anuncio.Preco
			porcentagemDesconto := anuncio.PorcentagemDesconto
			precoComDesconto := precoOriginal

			if porcentagemDesconto > 0 {
				// Usa a porcentagem do próprio anúncio
				precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
			} else if anuncio.PrecoComDesconto > 0 && anuncio.PrecoComDesconto < precoOriginal {
				// Usa o preço com desconto já definido no anúncio
				precoComDesconto = anuncio.PrecoComDesconto
				porcentagemDesconto = ((precoOriginal - precoComDesconto) / precoOriginal) * 100
			} else {
				// Busca desconto ativo da loja como fallback
				if anuncio.IDLoja != nil && *anuncio.IDLoja > 0 {
					desconto, _ := datasource.GetDescontoAtivoByLojaID(*anuncio.IDLoja)
					if desconto != nil {
						porcentagemDesconto = desconto.Porcentagem
						precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
					}
				}
			}

			// Determina a imagem (prioridade: anúncio > serviço)
			imagem := anuncio.Imagem
			if imagem == "" && servico.Imagem != "" {
				imagem = servico.Imagem
			}

			// Nome do serviço (usa Titulo do anúncio ou Titulo do serviço)
			nomeServico := anuncio.Titulo
			if nomeServico == "" {
				nomeServico = servico.Titulo
			}

			// Moedas da oferta Auto Mais
			var moedasUtiliza *int
			if anuncio.OfertaAutoMais != nil && anuncio.OfertaAutoMais.Ativo {
				moedas := anuncio.OfertaAutoMais.Moedas
				moedasUtiliza = &moedas
			}

			distancia := anuncioComDist.Distancia
			response := json.AnuncioServicoResponse{
				ID:                  anuncio.ID,
				NomeServico:         nomeServico,
				NomeLoja:            anuncio.Loja.Nome,
				Imagem:              imagem,
				PrecoOriginal:       precoOriginal,
				PrecoComDesconto:    precoComDesconto,
				PorcentagemDesconto: porcentagemDesconto,
				IsMeuCarroMais:      anuncio.Loja.IsMeuCarroMais,
				Categoria:           anuncio.Categoria,
				Descricao:           anuncio.Descricao,
				Rate:                anuncio.Loja.Rating,
				MoedasUtiliza:       moedasUtiliza,
				Distancia:           &distancia,
			}

			anunciosResponse = append(anunciosResponse, response)
		}
	} else {
		// Busca sem ordenação por proximidade
		anuncios, err := datasource.GetAnunciosServicos()
		if err != nil {
			return nil, err
		}

		for _, anuncio := range anuncios {
			// Verifica se o anúncio tem serviço associado
			if anuncio.Servico == nil || anuncio.IDServico == nil {
				continue
			}

			servico := anuncio.Servico

			// Calcula preço com desconto
			// Prioridade: 1) porcentagem do anúncio, 2) preço com desconto do anúncio, 3) desconto da loja
			precoOriginal := anuncio.Preco
			porcentagemDesconto := anuncio.PorcentagemDesconto
			precoComDesconto := precoOriginal

			if porcentagemDesconto > 0 {
				// Usa a porcentagem do próprio anúncio
				precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
			} else if anuncio.PrecoComDesconto > 0 && anuncio.PrecoComDesconto < precoOriginal {
				// Usa o preço com desconto já definido no anúncio
				precoComDesconto = anuncio.PrecoComDesconto
				porcentagemDesconto = ((precoOriginal - precoComDesconto) / precoOriginal) * 100
			} else {
				// Busca desconto ativo da loja como fallback
				if anuncio.IDLoja != nil && *anuncio.IDLoja > 0 {
					desconto, _ := datasource.GetDescontoAtivoByLojaID(*anuncio.IDLoja)
					if desconto != nil {
						porcentagemDesconto = desconto.Porcentagem
						precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
					}
				}
			}

			// Determina a imagem (prioridade: anúncio > serviço)
			imagem := anuncio.Imagem
			if imagem == "" && servico.Imagem != "" {
				imagem = servico.Imagem
			}

			// Nome do serviço (usa Titulo do anúncio ou Titulo do serviço)
			nomeServico := anuncio.Titulo
			if nomeServico == "" {
				nomeServico = servico.Titulo
			}

			// Moedas da oferta Auto Mais
			var moedasUtiliza *int
			if anuncio.OfertaAutoMais != nil && anuncio.OfertaAutoMais.Ativo {
				moedas := anuncio.OfertaAutoMais.Moedas
				moedasUtiliza = &moedas
			}

			response := json.AnuncioServicoResponse{
				ID:                  anuncio.ID,
				NomeServico:         nomeServico,
				NomeLoja:            anuncio.Loja.Nome,
				Imagem:              imagem,
				PrecoOriginal:       precoOriginal,
				PrecoComDesconto:    precoComDesconto,
				PorcentagemDesconto: porcentagemDesconto,
				IsMeuCarroMais:      anuncio.Loja.IsMeuCarroMais,
				Categoria:           anuncio.Categoria,
				Descricao:           anuncio.Descricao,
				Rate:                anuncio.Loja.Rating,
				MoedasUtiliza:       moedasUtiliza,
			}

			anunciosResponse = append(anunciosResponse, response)
		}
	}

	return &json.AnunciosServicoResponse{
		Anuncios: anunciosResponse,
		Total:    len(anunciosResponse),
	}, nil
}