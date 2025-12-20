package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// CreateRegistroInteresse cria um novo registro de interesse
func CreateRegistroInteresse(req json.RegistroInteresseRequest) (*json.RegistroInteresseResponse, error) {
	registroInteresse, err := datasource.CreateRegistroInteresse(req)
	if err != nil {
		return nil, err
	}

	response := &json.RegistroInteresseResponse{
		ID:              registroInteresse.ID,
		IDAnuncio:       registroInteresse.IDAnuncio,
		Nome:            registroInteresse.Nome,
		Email:           registroInteresse.Email,
		Telefone:        registroInteresse.Telefone,
		Mensagem:        registroInteresse.Mensagem,
		DataCadastro:    registroInteresse.DataCadastro,
		DataAtualizacao: registroInteresse.DataAtualizacao,
	}

	// Se o anúncio foi carregado, adiciona ao response
	if registroInteresse.Anuncio.ID != 0 {
		anuncioResp := &json.AnuncioResponse{
			ID:        registroInteresse.Anuncio.ID,
			Titulo:    registroInteresse.Anuncio.Titulo,
			Descricao: registroInteresse.Anuncio.Descricao,
			Preco:     registroInteresse.Anuncio.Preco,
			Imagem:    registroInteresse.Anuncio.Imagem,
			Destaque:  registroInteresse.Anuncio.Destaque,
			Categoria: registroInteresse.Anuncio.Categoria,
			IDLoja:    registroInteresse.Anuncio.IDLoja,
		}
		if registroInteresse.Anuncio.Loja.ID != 0 {
			anuncioResp.Loja = json.LojaResponse{
				ID:          registroInteresse.Anuncio.Loja.ID,
				Nome:        registroInteresse.Anuncio.Loja.Nome,
				CNPJ:        registroInteresse.Anuncio.Loja.CNPJ,
				Imagem:      registroInteresse.Anuncio.Loja.Imagem,
				Latitude:    registroInteresse.Anuncio.Loja.Latitude,
				Longitude:   registroInteresse.Anuncio.Loja.Longitude,
				IDCategoria: registroInteresse.Anuncio.Loja.IDCategoria,
				Categoria:   registroInteresse.Anuncio.Loja.Categoria.Nome,
			}
		}
		response.Anuncio = anuncioResp
	}

	return response, nil
}

// GetRegistroInteresseByID busca um registro de interesse por ID
func GetRegistroInteresseByID(id uint) (*json.RegistroInteresseResponse, error) {
	registroInteresse, err := datasource.GetRegistroInteresseByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.RegistroInteresseResponse{
		ID:              registroInteresse.ID,
		IDAnuncio:       registroInteresse.IDAnuncio,
		Nome:            registroInteresse.Nome,
		Email:           registroInteresse.Email,
		Telefone:        registroInteresse.Telefone,
		Mensagem:        registroInteresse.Mensagem,
		DataCadastro:    registroInteresse.DataCadastro,
		DataAtualizacao: registroInteresse.DataAtualizacao,
	}

	// Se o anúncio foi carregado, adiciona ao response
	if registroInteresse.Anuncio.ID != 0 {
		anuncioResp := &json.AnuncioResponse{
			ID:        registroInteresse.Anuncio.ID,
			Titulo:    registroInteresse.Anuncio.Titulo,
			Descricao: registroInteresse.Anuncio.Descricao,
			Preco:     registroInteresse.Anuncio.Preco,
			Imagem:    registroInteresse.Anuncio.Imagem,
			Destaque:  registroInteresse.Anuncio.Destaque,
			Categoria: registroInteresse.Anuncio.Categoria,
			IDLoja:    registroInteresse.Anuncio.IDLoja,
		}
		if registroInteresse.Anuncio.Loja.ID != 0 {
			anuncioResp.Loja = json.LojaResponse{
				ID:          registroInteresse.Anuncio.Loja.ID,
				Nome:        registroInteresse.Anuncio.Loja.Nome,
				CNPJ:        registroInteresse.Anuncio.Loja.CNPJ,
				Imagem:      registroInteresse.Anuncio.Loja.Imagem,
				Latitude:    registroInteresse.Anuncio.Loja.Latitude,
				Longitude:   registroInteresse.Anuncio.Loja.Longitude,
				IDCategoria: registroInteresse.Anuncio.Loja.IDCategoria,
				Categoria:   registroInteresse.Anuncio.Loja.Categoria.Nome,
			}
		}
		response.Anuncio = anuncioResp
	}

	return response, nil
}

// GetAllRegistroInteresses retorna todos os registros de interesse ativos
func GetAllRegistroInteresses() ([]json.RegistroInteresseResponse, error) {
	registrosInteresse, err := datasource.GetAllRegistroInteresses()
	if err != nil {
		return nil, err
	}

	var responses []json.RegistroInteresseResponse
	for _, registroInteresse := range registrosInteresse {
		response := json.RegistroInteresseResponse{
			ID:              registroInteresse.ID,
			IDAnuncio:       registroInteresse.IDAnuncio,
			Nome:            registroInteresse.Nome,
			Email:           registroInteresse.Email,
			Telefone:        registroInteresse.Telefone,
			Mensagem:        registroInteresse.Mensagem,
			DataCadastro:    registroInteresse.DataCadastro,
			DataAtualizacao: registroInteresse.DataAtualizacao,
		}

		// Se o anúncio foi carregado, adiciona ao response
		if registroInteresse.Anuncio.ID != 0 {
			anuncioResp := &json.AnuncioResponse{
				ID:        registroInteresse.Anuncio.ID,
				Titulo:    registroInteresse.Anuncio.Titulo,
				Descricao: registroInteresse.Anuncio.Descricao,
				Preco:     registroInteresse.Anuncio.Preco,
				Imagem:    registroInteresse.Anuncio.Imagem,
				Destaque:  registroInteresse.Anuncio.Destaque,
				Categoria: registroInteresse.Anuncio.Categoria,
				IDLoja:    registroInteresse.Anuncio.IDLoja,
			}
			if registroInteresse.Anuncio.Loja.ID != 0 {
				anuncioResp.Loja = json.LojaResponse{
					ID:          registroInteresse.Anuncio.Loja.ID,
					Nome:        registroInteresse.Anuncio.Loja.Nome,
					CNPJ:        registroInteresse.Anuncio.Loja.CNPJ,
					Imagem:      registroInteresse.Anuncio.Loja.Imagem,
					Latitude:    registroInteresse.Anuncio.Loja.Latitude,
					Longitude:   registroInteresse.Anuncio.Loja.Longitude,
					IDCategoria: registroInteresse.Anuncio.Loja.IDCategoria,
					Categoria:   registroInteresse.Anuncio.Loja.Categoria.Nome,
				}
			}
			response.Anuncio = anuncioResp
		}

		responses = append(responses, response)
	}

	return responses, nil
}

// GetRegistroInteressesByAnuncioID retorna todos os registros de interesse de um anúncio específico
func GetRegistroInteressesByAnuncioID(anuncioID uint) ([]json.RegistroInteresseResponse, error) {
	registrosInteresse, err := datasource.GetRegistroInteressesByAnuncioID(anuncioID)
	if err != nil {
		return nil, err
	}

	var responses []json.RegistroInteresseResponse
	for _, registroInteresse := range registrosInteresse {
		response := json.RegistroInteresseResponse{
			ID:              registroInteresse.ID,
			IDAnuncio:       registroInteresse.IDAnuncio,
			Nome:            registroInteresse.Nome,
			Email:           registroInteresse.Email,
			Telefone:        registroInteresse.Telefone,
			Mensagem:        registroInteresse.Mensagem,
			DataCadastro:    registroInteresse.DataCadastro,
			DataAtualizacao: registroInteresse.DataAtualizacao,
		}

		// Se o anúncio foi carregado, adiciona ao response
		if registroInteresse.Anuncio.ID != 0 {
			anuncioResp := &json.AnuncioResponse{
				ID:        registroInteresse.Anuncio.ID,
				Titulo:    registroInteresse.Anuncio.Titulo,
				Descricao: registroInteresse.Anuncio.Descricao,
				Preco:     registroInteresse.Anuncio.Preco,
				Imagem:    registroInteresse.Anuncio.Imagem,
				Destaque:  registroInteresse.Anuncio.Destaque,
				Categoria: registroInteresse.Anuncio.Categoria,
				IDLoja:    registroInteresse.Anuncio.IDLoja,
			}
			if registroInteresse.Anuncio.Loja.ID != 0 {
				anuncioResp.Loja = json.LojaResponse{
					ID:          registroInteresse.Anuncio.Loja.ID,
					Nome:        registroInteresse.Anuncio.Loja.Nome,
					CNPJ:        registroInteresse.Anuncio.Loja.CNPJ,
					Imagem:      registroInteresse.Anuncio.Loja.Imagem,
					Latitude:    registroInteresse.Anuncio.Loja.Latitude,
					Longitude:   registroInteresse.Anuncio.Loja.Longitude,
					IDCategoria: registroInteresse.Anuncio.Loja.IDCategoria,
					Categoria:   registroInteresse.Anuncio.Loja.Categoria.Nome,
				}
			}
			response.Anuncio = anuncioResp
		}

		responses = append(responses, response)
	}

	return responses, nil
}

// SoftDeleteRegistroInteresse realiza soft delete do registro de interesse
func SoftDeleteRegistroInteresse(id uint) error {
	return datasource.SoftDeleteRegistroInteresse(id)
}

// RestoreRegistroInteresse restaura um registro de interesse que foi soft deleted
func RestoreRegistroInteresse(id uint) error {
	return datasource.RestoreRegistroInteresse(id)
}
