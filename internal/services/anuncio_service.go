package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// modelToAnuncioResponse converte o model de anúncio para response
func modelToAnuncioResponse(anuncio *models.Anuncio) json.AnuncioResponse {
	response := json.AnuncioResponse{
		ID:               anuncio.ID,
		Titulo:           anuncio.Titulo,
		Descricao:        anuncio.Descricao,
		Preco:            anuncio.Preco,
		Imagem:           anuncio.Imagem,
		Destaque:         anuncio.Destaque,
		Categoria:        anuncio.Categoria,
		IDLoja:           anuncio.IDLoja,
		IDProduto:        anuncio.IDProduto,
		IDServico:        anuncio.IDServico,
		IDVeiculo:        anuncio.IDVeiculo,
		IDOfertaAutoMais: anuncio.IDOfertaAutoMais,
		TipoAnuncio:      anuncio.TipoAnuncio,
		Loja: json.LojaResponse{
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
		},
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

			// Busca desconto ativo da loja
			desconto, _ := datasource.GetDescontoAtivoByLojaID(anuncio.IDLoja)
			
			// Calcula preço com desconto
			precoOriginal := anuncio.Preco
			porcentagemDesconto := 0.0
			precoComDesconto := precoOriginal

			if desconto != nil {
				porcentagemDesconto = desconto.Porcentagem
				precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
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

			// Busca desconto ativo da loja
			desconto, _ := datasource.GetDescontoAtivoByLojaID(anuncio.IDLoja)
			
			// Calcula preço com desconto
			precoOriginal := anuncio.Preco
			porcentagemDesconto := 0.0
			precoComDesconto := precoOriginal

			if desconto != nil {
				porcentagemDesconto = desconto.Porcentagem
				precoComDesconto = precoOriginal * (1 - porcentagemDesconto/100)
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

			distancia := anuncioComDist.Distancia
			response := json.AnuncioVeiculoResponse{
				ID:                  anuncio.ID,
				NomeVeiculo:         nomeVeiculo,
				KM:                  veiculo.Quilometragem,
				AnoModelo:            veiculo.Ano,
				AnoFabricacao:       &veiculo.Ano, // Usa o mesmo ano, pode ser ajustado se houver campo separado
				IsMeuCarroMais:      anuncio.Loja.IsMeuCarroMais,
				Preco:               anuncio.Preco,
				Imagem:              imagem,
				Modelo:              veiculo.Modelo,
				Marca:               nil, // Campo não existe no modelo atual
				Placa:               veiculo.Placa,
				Renavam:             nil, // Campo não existe no modelo atual
				Chassi:              nil, // Campo não existe no modelo atual
				Cor:                 veiculo.Cor,
				TipoVeiculo:         &anuncio.Categoria, // Usa categoria como tipo de veículo
				Licenciamento:       nil,                // Campo não existe no modelo atual
				IPVAPago:            nil,                // Campo não existe no modelo atual
				PossuiFinanciamento: nil,                // Campo não existe no modelo atual
				PossuiMultas:        nil,                // Campo não existe no modelo atual
				Observacoes:         veiculo.Observacoes,
				Combustivel:         nil, // Campo não existe no modelo atual
				MoedasUtiliza:        moedasUtiliza,
				Distancia:           &distancia,
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

			response := json.AnuncioVeiculoResponse{
				ID:                  anuncio.ID,
				NomeVeiculo:         nomeVeiculo,
				KM:                  veiculo.Quilometragem,
				AnoModelo:            veiculo.Ano,
				AnoFabricacao:       &veiculo.Ano, // Usa o mesmo ano, pode ser ajustado se houver campo separado
				IsMeuCarroMais:      anuncio.Loja.IsMeuCarroMais,
				Preco:               anuncio.Preco,
				Imagem:              imagem,
				Modelo:              veiculo.Modelo,
				Marca:               nil, // Campo não existe no modelo atual
				Placa:               veiculo.Placa,
				Renavam:             nil, // Campo não existe no modelo atual
				Chassi:              nil, // Campo não existe no modelo atual
				Cor:                 veiculo.Cor,
				TipoVeiculo:         &anuncio.Categoria, // Usa categoria como tipo de veículo
				Licenciamento:       nil,                // Campo não existe no modelo atual
				IPVAPago:            nil,                // Campo não existe no modelo atual
				PossuiFinanciamento: nil,                // Campo não existe no modelo atual
				PossuiMultas:        nil,                // Campo não existe no modelo atual
				Observacoes:         veiculo.Observacoes,
				Combustivel:         nil, // Campo não existe no modelo atual
				MoedasUtiliza:        moedasUtiliza,
			}

			anunciosResponse = append(anunciosResponse, response)
		}
	}

	return &json.AnunciosVeiculoResponse{
		Anuncios: anunciosResponse,
		Total:    len(anunciosResponse),
	}, nil
}