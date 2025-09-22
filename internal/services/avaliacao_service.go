package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// CreateAvaliacao cria uma nova avaliação
func CreateAvaliacao(req json.AvaliacaoRequest) (*json.AvaliacaoResponse, error) {
	avaliacao, err := datasource.CreateAvaliacao(req)
	if err != nil {
		return nil, err
	}

	response := &json.AvaliacaoResponse{
		ID:              avaliacao.ID,
		IDUsuario:       avaliacao.IDUsuario,
		IDLoja:          avaliacao.IDLoja,
		Nota:            avaliacao.Nota,
		Comentario:      avaliacao.Comentario,
		DataAvaliacao:   avaliacao.DataAvaliacao,
		DataAtualizacao: avaliacao.DataAtualizacao,
		Usuario: json.UserResponse{
			ID:             avaliacao.Usuario.ID,
			Nome:           avaliacao.Usuario.Nome,
			Email:          avaliacao.Usuario.Email,
			CPF:            avaliacao.Usuario.CPF,
			Imagem:         avaliacao.Usuario.Imagem,
			Telefone:       avaliacao.Usuario.Telefone,
			Endereco:       avaliacao.Usuario.Endereco,
			DataNascimento: avaliacao.Usuario.DataNascimento,
			DataCadastro:   avaliacao.Usuario.DataCadastro,
			Ativo:          avaliacao.Usuario.Ativo,
			Latitude:       avaliacao.Usuario.Latitude,
			Longitude:      avaliacao.Usuario.Longitude,
			IDPlano:        avaliacao.Usuario.IDPlano,
			IDLoja:         avaliacao.Usuario.IDLoja,
		},
		Loja: json.LojaResponse{
			ID:          avaliacao.Loja.ID,
			Nome:        avaliacao.Loja.Nome,
			CNPJ:        avaliacao.Loja.CNPJ,
			Imagem:      avaliacao.Loja.Imagem,
			Latitude:    avaliacao.Loja.Latitude,
			Longitude:   avaliacao.Loja.Longitude,
			IDCategoria: avaliacao.Loja.IDCategoria,
			Categoria:   avaliacao.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// GetAvaliacaoByID busca uma avaliação por ID
func GetAvaliacaoByID(id uint) (*json.AvaliacaoResponse, error) {
	avaliacao, err := datasource.GetAvaliacaoByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.AvaliacaoResponse{
		ID:              avaliacao.ID,
		IDUsuario:       avaliacao.IDUsuario,
		IDLoja:          avaliacao.IDLoja,
		Nota:            avaliacao.Nota,
		Comentario:      avaliacao.Comentario,
		DataAvaliacao:   avaliacao.DataAvaliacao,
		DataAtualizacao: avaliacao.DataAtualizacao,
		Usuario: json.UserResponse{
			ID:             avaliacao.Usuario.ID,
			Nome:           avaliacao.Usuario.Nome,
			Email:          avaliacao.Usuario.Email,
			CPF:            avaliacao.Usuario.CPF,
			Imagem:         avaliacao.Usuario.Imagem,
			Telefone:       avaliacao.Usuario.Telefone,
			Endereco:       avaliacao.Usuario.Endereco,
			DataNascimento: avaliacao.Usuario.DataNascimento,
			DataCadastro:   avaliacao.Usuario.DataCadastro,
			Ativo:          avaliacao.Usuario.Ativo,
			Latitude:       avaliacao.Usuario.Latitude,
			Longitude:      avaliacao.Usuario.Longitude,
			IDPlano:        avaliacao.Usuario.IDPlano,
			IDLoja:         avaliacao.Usuario.IDLoja,
		},
		Loja: json.LojaResponse{
			ID:          avaliacao.Loja.ID,
			Nome:        avaliacao.Loja.Nome,
			CNPJ:        avaliacao.Loja.CNPJ,
			Imagem:      avaliacao.Loja.Imagem,
			Latitude:    avaliacao.Loja.Latitude,
			Longitude:   avaliacao.Loja.Longitude,
			IDCategoria: avaliacao.Loja.IDCategoria,
			Categoria:   avaliacao.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// GetAllAvaliacoes retorna todas as avaliações ativas
func GetAllAvaliacoes() ([]json.AvaliacaoResponse, error) {
	avaliacoes, err := datasource.GetAllAvaliacoes()
	if err != nil {
		return nil, err
	}

	var responses []json.AvaliacaoResponse
	for _, avaliacao := range avaliacoes {
		response := json.AvaliacaoResponse{
			ID:              avaliacao.ID,
			IDUsuario:       avaliacao.IDUsuario,
			IDLoja:          avaliacao.IDLoja,
			Nota:            avaliacao.Nota,
			Comentario:      avaliacao.Comentario,
			DataAvaliacao:   avaliacao.DataAvaliacao,
			DataAtualizacao: avaliacao.DataAtualizacao,
			Usuario: json.UserResponse{
				ID:             avaliacao.Usuario.ID,
				Nome:           avaliacao.Usuario.Nome,
				Email:          avaliacao.Usuario.Email,
				CPF:            avaliacao.Usuario.CPF,
				Imagem:         avaliacao.Usuario.Imagem,
				Telefone:       avaliacao.Usuario.Telefone,
				Endereco:       avaliacao.Usuario.Endereco,
				DataNascimento: avaliacao.Usuario.DataNascimento,
				DataCadastro:   avaliacao.Usuario.DataCadastro,
				Ativo:          avaliacao.Usuario.Ativo,
				Latitude:       avaliacao.Usuario.Latitude,
				Longitude:      avaliacao.Usuario.Longitude,
				IDPlano:        avaliacao.Usuario.IDPlano,
				IDLoja:         avaliacao.Usuario.IDLoja,
			},
			Loja: json.LojaResponse{
				ID:          avaliacao.Loja.ID,
				Nome:        avaliacao.Loja.Nome,
				CNPJ:        avaliacao.Loja.CNPJ,
				Imagem:      avaliacao.Loja.Imagem,
				Latitude:    avaliacao.Loja.Latitude,
				Longitude:   avaliacao.Loja.Longitude,
				IDCategoria: avaliacao.Loja.IDCategoria,
				Categoria:   avaliacao.Loja.Categoria.Nome,
			},
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// GetAvaliacoesByLojaID retorna todas as avaliações de uma loja específica
func GetAvaliacoesByLojaID(idLoja uint) (*json.AvaliacoesResponse, error) {
	avaliacoes, err := datasource.GetAvaliacoesByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var avaliacoesResponse []json.AvaliacaoResponse
	var somaNotas float64

	for _, avaliacao := range avaliacoes {
		avaliacaoResp := json.AvaliacaoResponse{
			ID:              avaliacao.ID,
			IDUsuario:       avaliacao.IDUsuario,
			IDLoja:          avaliacao.IDLoja,
			Nota:            avaliacao.Nota,
			Comentario:      avaliacao.Comentario,
			DataAvaliacao:   avaliacao.DataAvaliacao,
			DataAtualizacao: avaliacao.DataAtualizacao,
			Usuario: json.UserResponse{
				ID:             avaliacao.Usuario.ID,
				Nome:           avaliacao.Usuario.Nome,
				Email:          avaliacao.Usuario.Email,
				CPF:            avaliacao.Usuario.CPF,
				Imagem:         avaliacao.Usuario.Imagem,
				Telefone:       avaliacao.Usuario.Telefone,
				Endereco:       avaliacao.Usuario.Endereco,
				DataNascimento: avaliacao.Usuario.DataNascimento,
				DataCadastro:   avaliacao.Usuario.DataCadastro,
				Ativo:          avaliacao.Usuario.Ativo,
				Latitude:       avaliacao.Usuario.Latitude,
				Longitude:      avaliacao.Usuario.Longitude,
				IDPlano:        avaliacao.Usuario.IDPlano,
				IDLoja:         avaliacao.Usuario.IDLoja,
			},
			Loja: json.LojaResponse{
				ID:          avaliacao.Loja.ID,
				Nome:        avaliacao.Loja.Nome,
				CNPJ:        avaliacao.Loja.CNPJ,
				Imagem:      avaliacao.Loja.Imagem,
				Latitude:    avaliacao.Loja.Latitude,
				Longitude:   avaliacao.Loja.Longitude,
				IDCategoria: avaliacao.Loja.IDCategoria,
				Categoria:   avaliacao.Loja.Categoria.Nome,
			},
		}
		avaliacoesResponse = append(avaliacoesResponse, avaliacaoResp)
		somaNotas += float64(avaliacao.Nota)
	}

	var mediaNota float64
	if len(avaliacoes) > 0 {
		mediaNota = somaNotas / float64(len(avaliacoes))
	}

	response := &json.AvaliacoesResponse{
		Avaliacoes: avaliacoesResponse,
		Total:      len(avaliacoesResponse),
		MediaNota:  mediaNota,
	}

	return response, nil
}

// GetAvaliacoesByUsuarioID retorna todas as avaliações de um usuário específico
func GetAvaliacoesByUsuarioID(idUsuario uint) (*json.AvaliacoesResponse, error) {
	avaliacoes, err := datasource.GetAvaliacoesByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}

	var avaliacoesResponse []json.AvaliacaoResponse
	for _, avaliacao := range avaliacoes {
		avaliacaoResp := json.AvaliacaoResponse{
			ID:              avaliacao.ID,
			IDUsuario:       avaliacao.IDUsuario,
			IDLoja:          avaliacao.IDLoja,
			Nota:            avaliacao.Nota,
			Comentario:      avaliacao.Comentario,
			DataAvaliacao:   avaliacao.DataAvaliacao,
			DataAtualizacao: avaliacao.DataAtualizacao,
			Usuario: json.UserResponse{
				ID:             avaliacao.Usuario.ID,
				Nome:           avaliacao.Usuario.Nome,
				Email:          avaliacao.Usuario.Email,
				CPF:            avaliacao.Usuario.CPF,
				Imagem:         avaliacao.Usuario.Imagem,
				Telefone:       avaliacao.Usuario.Telefone,
				Endereco:       avaliacao.Usuario.Endereco,
				DataNascimento: avaliacao.Usuario.DataNascimento,
				DataCadastro:   avaliacao.Usuario.DataCadastro,
				Ativo:          avaliacao.Usuario.Ativo,
				Latitude:       avaliacao.Usuario.Latitude,
				Longitude:      avaliacao.Usuario.Longitude,
				IDPlano:        avaliacao.Usuario.IDPlano,
				IDLoja:         avaliacao.Usuario.IDLoja,
			},
			Loja: json.LojaResponse{
				ID:          avaliacao.Loja.ID,
				Nome:        avaliacao.Loja.Nome,
				CNPJ:        avaliacao.Loja.CNPJ,
				Imagem:      avaliacao.Loja.Imagem,
				Latitude:    avaliacao.Loja.Latitude,
				Longitude:   avaliacao.Loja.Longitude,
				IDCategoria: avaliacao.Loja.IDCategoria,
				Categoria:   avaliacao.Loja.Categoria.Nome,
			},
		}
		avaliacoesResponse = append(avaliacoesResponse, avaliacaoResp)
	}

	response := &json.AvaliacoesResponse{
		Avaliacoes: avaliacoesResponse,
		Total:      len(avaliacoesResponse),
	}

	return response, nil
}

// GetAvaliacaoEstatisticasByLojaID retorna estatísticas das avaliações de uma loja
func GetAvaliacaoEstatisticasByLojaID(idLoja uint) (*json.AvaliacaoEstatisticasResponse, error) {
	return datasource.GetAvaliacaoEstatisticasByLojaID(idLoja)
}

// UpdateAvaliacao atualiza uma avaliação existente
func UpdateAvaliacao(id uint, req json.AvaliacaoRequest) (*json.AvaliacaoResponse, error) {
	avaliacao, err := datasource.UpdateAvaliacao(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.AvaliacaoResponse{
		ID:              avaliacao.ID,
		IDUsuario:       avaliacao.IDUsuario,
		IDLoja:          avaliacao.IDLoja,
		Nota:            avaliacao.Nota,
		Comentario:      avaliacao.Comentario,
		DataAvaliacao:   avaliacao.DataAvaliacao,
		DataAtualizacao: avaliacao.DataAtualizacao,
		Usuario: json.UserResponse{
			ID:             avaliacao.Usuario.ID,
			Nome:           avaliacao.Usuario.Nome,
			Email:          avaliacao.Usuario.Email,
			CPF:            avaliacao.Usuario.CPF,
			Imagem:         avaliacao.Usuario.Imagem,
			Telefone:       avaliacao.Usuario.Telefone,
			Endereco:       avaliacao.Usuario.Endereco,
			DataNascimento: avaliacao.Usuario.DataNascimento,
			DataCadastro:   avaliacao.Usuario.DataCadastro,
			Ativo:          avaliacao.Usuario.Ativo,
			Latitude:       avaliacao.Usuario.Latitude,
			Longitude:      avaliacao.Usuario.Longitude,
			IDPlano:        avaliacao.Usuario.IDPlano,
			IDLoja:         avaliacao.Usuario.IDLoja,
		},
		Loja: json.LojaResponse{
			ID:          avaliacao.Loja.ID,
			Nome:        avaliacao.Loja.Nome,
			CNPJ:        avaliacao.Loja.CNPJ,
			Imagem:      avaliacao.Loja.Imagem,
			Latitude:    avaliacao.Loja.Latitude,
			Longitude:   avaliacao.Loja.Longitude,
			IDCategoria: avaliacao.Loja.IDCategoria,
			Categoria:   avaliacao.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// SoftDeleteAvaliacao realiza soft delete da avaliação
func SoftDeleteAvaliacao(id uint) error {
	return datasource.SoftDeleteAvaliacao(id)
}

// RestoreAvaliacao restaura uma avaliação que foi soft deleted
func RestoreAvaliacao(id uint) error {
	return datasource.RestoreAvaliacao(id)
}
