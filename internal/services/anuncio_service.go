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
