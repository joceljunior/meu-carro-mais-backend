package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

func GetServicosByProximidade(latitude, longitude float64) (*json.ServicosResponse, error) {
	servicos, err := datasource.GetServicosByProximidade(latitude, longitude)
	if err != nil {
		return nil, err
	}

	response := &json.ServicosResponse{
		Servicos: servicos,
		Total:    len(servicos),
	}

	return response, nil
}

func GetCategoriasServico() (*json.CategoriasServicoResponse, error) {
	categorias, err := datasource.GetCategoriasServico()
	if err != nil {
		return nil, err
	}

	var categoriasResponse []json.CategoriaServicoResponse
	for _, categoria := range categorias {
		categoriaResp := json.CategoriaServicoResponse{
			ID:   categoria.ID,
			Nome: categoria.Nome,
		}
		categoriasResponse = append(categoriasResponse, categoriaResp)
	}

	response := &json.CategoriasServicoResponse{
		Categorias: categoriasResponse,
		Total:      len(categoriasResponse),
	}

	return response, nil
}
