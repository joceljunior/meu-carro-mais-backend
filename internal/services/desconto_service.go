package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// CreateDesconto cria um novo desconto para uma loja
func CreateDesconto(req json.DescontoRequest) (*json.DescontoResponse, error) {
	desconto, err := datasource.CreateDesconto(req)
	if err != nil {
		return nil, err
	}

	response := &json.DescontoResponse{
		ID:              desconto.ID,
		IDLoja:          desconto.IDLoja,
		Porcentagem:     desconto.Porcentagem,
		Ativo:           desconto.Ativo,
		DataValidade:    desconto.DataValidade,
		DataCadastro:    desconto.DataCadastro,
		DataAtualizacao: desconto.DataAtualizacao,
		Loja: json.LojaResponse{
			ID:             desconto.Loja.ID,
			Nome:           desconto.Loja.Nome,
			CNPJ:           desconto.Loja.CNPJ,
			Imagem:         desconto.Loja.Imagem,
			Latitude:       desconto.Loja.Latitude,
			Longitude:      desconto.Loja.Longitude,
			Rating:         desconto.Loja.Rating,
			IsMeuCarroMais: desconto.Loja.IsMeuCarroMais,
			Categoria:      desconto.Loja.Categoria,
			IDUsuario:      desconto.Loja.IDUsuario,
		},
	}

	return response, nil
}

// GetDescontoByID busca um desconto por ID
func GetDescontoByID(id uint) (*json.DescontoResponse, error) {
	desconto, err := datasource.GetDescontoByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.DescontoResponse{
		ID:              desconto.ID,
		IDLoja:          desconto.IDLoja,
		Porcentagem:     desconto.Porcentagem,
		Ativo:           desconto.Ativo,
		DataValidade:    desconto.DataValidade,
		DataCadastro:    desconto.DataCadastro,
		DataAtualizacao: desconto.DataAtualizacao,
		Loja: json.LojaResponse{
			ID:             desconto.Loja.ID,
			Nome:           desconto.Loja.Nome,
			CNPJ:           desconto.Loja.CNPJ,
			Imagem:         desconto.Loja.Imagem,
			Latitude:       desconto.Loja.Latitude,
			Longitude:      desconto.Loja.Longitude,
			Rating:         desconto.Loja.Rating,
			IsMeuCarroMais: desconto.Loja.IsMeuCarroMais,
			Categoria:      desconto.Loja.Categoria,
			IDUsuario:      desconto.Loja.IDUsuario,
		},
	}

	return response, nil
}

// GetDescontoAtivoByLojaID busca o desconto ativo de uma loja
func GetDescontoAtivoByLojaID(idLoja uint) (*json.DescontoResponse, error) {
	desconto, err := datasource.GetDescontoAtivoByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	response := &json.DescontoResponse{
		ID:              desconto.ID,
		IDLoja:          desconto.IDLoja,
		Porcentagem:     desconto.Porcentagem,
		Ativo:           desconto.Ativo,
		DataValidade:    desconto.DataValidade,
		DataCadastro:    desconto.DataCadastro,
		DataAtualizacao: desconto.DataAtualizacao,
		Loja: json.LojaResponse{
			ID:             desconto.Loja.ID,
			Nome:           desconto.Loja.Nome,
			CNPJ:           desconto.Loja.CNPJ,
			Imagem:         desconto.Loja.Imagem,
			Latitude:       desconto.Loja.Latitude,
			Longitude:      desconto.Loja.Longitude,
			Rating:         desconto.Loja.Rating,
			IsMeuCarroMais: desconto.Loja.IsMeuCarroMais,
			Categoria:      desconto.Loja.Categoria,
			IDUsuario:      desconto.Loja.IDUsuario,
		},
	}

	return response, nil
}

// GetAllDescontos retorna todos os descontos
func GetAllDescontos() ([]json.DescontoResponse, error) {
	descontos, err := datasource.GetAllDescontos()
	if err != nil {
		return nil, err
	}

	var responses []json.DescontoResponse
	for _, desconto := range descontos {
		response := json.DescontoResponse{
			ID:              desconto.ID,
			IDLoja:          desconto.IDLoja,
			Porcentagem:     desconto.Porcentagem,
			Ativo:           desconto.Ativo,
			DataValidade:    desconto.DataValidade,
			DataCadastro:    desconto.DataCadastro,
			DataAtualizacao: desconto.DataAtualizacao,
			Loja: json.LojaResponse{
				ID:             desconto.Loja.ID,
				Nome:           desconto.Loja.Nome,
				CNPJ:           desconto.Loja.CNPJ,
				Imagem:         desconto.Loja.Imagem,
				Latitude:       desconto.Loja.Latitude,
				Longitude:      desconto.Loja.Longitude,
				Rating:         desconto.Loja.Rating,
				IsMeuCarroMais: desconto.Loja.IsMeuCarroMais,
				Categoria:      desconto.Loja.Categoria,
				IDUsuario:      desconto.Loja.IDUsuario,
			},
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// GetAllDescontosAtivos retorna todos os descontos ativos
func GetAllDescontosAtivos() ([]json.DescontoResponse, error) {
	descontos, err := datasource.GetAllDescontosAtivos()
	if err != nil {
		return nil, err
	}

	var responses []json.DescontoResponse
	for _, desconto := range descontos {
		response := json.DescontoResponse{
			ID:              desconto.ID,
			IDLoja:          desconto.IDLoja,
			Porcentagem:     desconto.Porcentagem,
			Ativo:           desconto.Ativo,
			DataValidade:    desconto.DataValidade,
			DataCadastro:    desconto.DataCadastro,
			DataAtualizacao: desconto.DataAtualizacao,
			Loja: json.LojaResponse{
				ID:             desconto.Loja.ID,
				Nome:           desconto.Loja.Nome,
				CNPJ:           desconto.Loja.CNPJ,
				Imagem:         desconto.Loja.Imagem,
				Latitude:       desconto.Loja.Latitude,
				Longitude:      desconto.Loja.Longitude,
				Rating:         desconto.Loja.Rating,
				IsMeuCarroMais: desconto.Loja.IsMeuCarroMais,
				Categoria:      desconto.Loja.Categoria,
				IDUsuario:      desconto.Loja.IDUsuario,
			},
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// GetDescontosByLojaID retorna todos os descontos de uma loja (histórico)
func GetDescontosByLojaID(idLoja uint) (*json.DescontosResponse, error) {
	descontos, err := datasource.GetDescontosByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var descontosResponse []json.DescontoResponse
	for _, desconto := range descontos {
		descontoResp := json.DescontoResponse{
			ID:              desconto.ID,
			IDLoja:          desconto.IDLoja,
			Porcentagem:     desconto.Porcentagem,
			Ativo:           desconto.Ativo,
			DataValidade:    desconto.DataValidade,
			DataCadastro:    desconto.DataCadastro,
			DataAtualizacao: desconto.DataAtualizacao,
			Loja: json.LojaResponse{
				ID:             desconto.Loja.ID,
				Nome:           desconto.Loja.Nome,
				CNPJ:           desconto.Loja.CNPJ,
				Imagem:         desconto.Loja.Imagem,
				Latitude:       desconto.Loja.Latitude,
				Longitude:      desconto.Loja.Longitude,
				Rating:         desconto.Loja.Rating,
				IsMeuCarroMais: desconto.Loja.IsMeuCarroMais,
				Categoria:      desconto.Loja.Categoria,
				IDUsuario:      desconto.Loja.IDUsuario,
			},
		}
		descontosResponse = append(descontosResponse, descontoResp)
	}

	response := &json.DescontosResponse{
		Descontos: descontosResponse,
		Total:     len(descontosResponse),
	}

	return response, nil
}

// CancelarDesconto cancela um desconto por ID
func CancelarDesconto(id uint) error {
	return datasource.CancelarDesconto(id)
}

// CancelarDescontoAtivoByLojaID cancela o desconto ativo de uma loja
func CancelarDescontoAtivoByLojaID(idLoja uint) error {
	return datasource.CancelarDescontoAtivoByLojaID(idLoja)
}

// SoftDeleteDesconto realiza soft delete do desconto
func SoftDeleteDesconto(id uint) error {
	return datasource.SoftDeleteDesconto(id)
}

// RestoreDesconto restaura um desconto que foi soft deleted
func RestoreDesconto(id uint) error {
	return datasource.RestoreDesconto(id)
}

