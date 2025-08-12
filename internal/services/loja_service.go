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