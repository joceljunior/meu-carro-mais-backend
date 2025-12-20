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

// CreateServico cria um novo serviço
func CreateServico(req json.ServicoRequest) (*json.ServicoResponse, error) {
	servico, err := datasource.CreateServico(req)
	if err != nil {
		return nil, err
	}

	response := &json.ServicoResponse{
		ID:        servico.ID,
		Titulo:    servico.Titulo,
		Descricao: servico.Descricao,
		Preco:     servico.Preco,
		Imagem:    servico.Imagem,
		Destaque:  servico.Destaque,
		Categoria: servico.Categoria,
		Loja: json.LojaResponse{
			ID:          servico.Loja.ID,
			Nome:        servico.Loja.Nome,
			CNPJ:        servico.Loja.CNPJ,
			Imagem:      servico.Loja.Imagem,
			Latitude:    servico.Loja.Latitude,
			Longitude:   servico.Loja.Longitude,
			IDCategoria: servico.Loja.IDCategoria,
			Categoria:   servico.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// GetServicoByID busca um serviço por ID
func GetServicoByID(id uint) (*json.ServicoResponse, error) {
	servico, err := datasource.GetServicoByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.ServicoResponse{
		ID:        servico.ID,
		Titulo:    servico.Titulo,
		Descricao: servico.Descricao,
		Preco:     servico.Preco,
		Imagem:    servico.Imagem,
		Destaque:  servico.Destaque,
		Categoria: servico.Categoria,
		Loja: json.LojaResponse{
			ID:          servico.Loja.ID,
			Nome:        servico.Loja.Nome,
			CNPJ:        servico.Loja.CNPJ,
			Imagem:      servico.Loja.Imagem,
			Latitude:    servico.Loja.Latitude,
			Longitude:   servico.Loja.Longitude,
			IDCategoria: servico.Loja.IDCategoria,
			Categoria:   servico.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// GetAllServicos retorna todos os serviços ativos
func GetAllServicos() ([]json.ServicoResponse, error) {
	servicos, err := datasource.GetAllServicos()
	if err != nil {
		return nil, err
	}

	var responses []json.ServicoResponse
	for _, servico := range servicos {
		response := json.ServicoResponse{
			ID:        servico.ID,
			Titulo:    servico.Titulo,
			Descricao: servico.Descricao,
			Preco:     servico.Preco,
			Imagem:    servico.Imagem,
			Destaque:  servico.Destaque,
			Categoria: servico.Categoria,
			Loja: json.LojaResponse{
				ID:          servico.Loja.ID,
				Nome:        servico.Loja.Nome,
				CNPJ:        servico.Loja.CNPJ,
				Imagem:      servico.Loja.Imagem,
				Latitude:    servico.Loja.Latitude,
				Longitude:   servico.Loja.Longitude,
				IDCategoria: servico.Loja.IDCategoria,
				Categoria:   servico.Loja.Categoria.Nome,
			},
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateServico atualiza um serviço existente
func UpdateServico(id uint, req json.ServicoRequest) (*json.ServicoResponse, error) {
	servico, err := datasource.UpdateServico(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.ServicoResponse{
		ID:        servico.ID,
		Titulo:    servico.Titulo,
		Descricao: servico.Descricao,
		Preco:     servico.Preco,
		Imagem:    servico.Imagem,
		Destaque:  servico.Destaque,
		Categoria: servico.Categoria,
		Loja: json.LojaResponse{
			ID:          servico.Loja.ID,
			Nome:        servico.Loja.Nome,
			CNPJ:        servico.Loja.CNPJ,
			Imagem:      servico.Loja.Imagem,
			Latitude:    servico.Loja.Latitude,
			Longitude:   servico.Loja.Longitude,
			IDCategoria: servico.Loja.IDCategoria,
			Categoria:   servico.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// SoftDeleteServico realiza soft delete do serviço
func SoftDeleteServico(id uint) error {
	return datasource.SoftDeleteServico(id)
}

// RestoreServico restaura um serviço que foi soft deleted
func RestoreServico(id uint) error {
	return datasource.RestoreServico(id)
}
