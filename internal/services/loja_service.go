package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// GetLojasByProximidade busca lojas ordenadas por proximidade do usuário
func GetLojasByProximidade(latitude, longitude float64) (*json.LojasResponse, error) {
	lojas, err := datasource.GetLojasByProximidade(latitude, longitude)
	if err != nil {
		return nil, err
	}

	var lojasResponse []json.LojaResponse
	for _, loja := range lojas {
		lojaResp := json.LojaResponse{
			ID:          loja.ID,
			Nome:        loja.Nome,
			CNPJ:        loja.CNPJ,
			Imagem:      loja.Imagem,
			Latitude:    loja.Latitude,
			Longitude:   loja.Longitude,
			IDCategoria: loja.IDCategoria,
			Categoria:   loja.Categoria.Nome,
		}
		lojasResponse = append(lojasResponse, lojaResp)
	}

	return &json.LojasResponse{
		Lojas: lojasResponse,
		Total: len(lojasResponse),
	}, nil
}

// GetCategoriasLojista retorna todas as categorias de lojista
func GetCategoriasLojista() (*json.CategoriasLojistaResponse, error) {
	categorias, err := datasource.GetCategoriasLojista()
	if err != nil {
		return nil, err
	}

	var categoriasResponse []json.CategoriaLojistaResponse
	for _, categoria := range categorias {
		categoriaResp := json.CategoriaLojistaResponse{
			ID:   categoria.ID,
			Nome: categoria.Nome,
		}
		categoriasResponse = append(categoriasResponse, categoriaResp)
	}

	response := &json.CategoriasLojistaResponse{
		Categorias: categoriasResponse,
		Total:      len(categoriasResponse),
	}

	return response, nil
}

// CreateLoja cria uma nova loja
func CreateLoja(req json.LojaRequest) (*json.LojaResponse, error) {
	loja, err := datasource.CreateLoja(req)
	if err != nil {
		return nil, err
	}

	response := &json.LojaResponse{
		ID:          loja.ID,
		Nome:        loja.Nome,
		CNPJ:        loja.CNPJ,
		Imagem:      loja.Imagem,
		Latitude:    loja.Latitude,
		Longitude:   loja.Longitude,
		IDCategoria: loja.IDCategoria,
		Categoria:   loja.Categoria.Nome,
	}

	return response, nil
}

// GetLojaByID busca uma loja por ID
func GetLojaByID(id uint) (*json.LojaResponse, error) {
	loja, err := datasource.GetLojaByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.LojaResponse{
		ID:          loja.ID,
		Nome:        loja.Nome,
		CNPJ:        loja.CNPJ,
		Imagem:      loja.Imagem,
		Latitude:    loja.Latitude,
		Longitude:   loja.Longitude,
		IDCategoria: loja.IDCategoria,
		Categoria:   loja.Categoria.Nome,
	}

	return response, nil
}

// GetAllLojas retorna todas as lojas ativas
func GetAllLojas() ([]json.LojaResponse, error) {
	lojas, err := datasource.GetAllLojas()
	if err != nil {
		return nil, err
	}

	var responses []json.LojaResponse
	for _, loja := range lojas {
		response := json.LojaResponse{
			ID:          loja.ID,
			Nome:        loja.Nome,
			CNPJ:        loja.CNPJ,
			Imagem:      loja.Imagem,
			Latitude:    loja.Latitude,
			Longitude:   loja.Longitude,
			IDCategoria: loja.IDCategoria,
			Categoria:   loja.Categoria.Nome,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateLoja atualiza uma loja existente
func UpdateLoja(id uint, req json.LojaRequest) (*json.LojaResponse, error) {
	loja, err := datasource.UpdateLoja(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.LojaResponse{
		ID:          loja.ID,
		Nome:        loja.Nome,
		CNPJ:        loja.CNPJ,
		Imagem:      loja.Imagem,
		Latitude:    loja.Latitude,
		Longitude:   loja.Longitude,
		IDCategoria: loja.IDCategoria,
		Categoria:   loja.Categoria.Nome,
	}

	return response, nil
}

// SoftDeleteLoja realiza soft delete da loja
func SoftDeleteLoja(id uint) error {
	return datasource.SoftDeleteLoja(id)
}

// RestoreLoja restaura uma loja que foi soft deleted
func RestoreLoja(id uint) error {
	return datasource.RestoreLoja(id)
}
