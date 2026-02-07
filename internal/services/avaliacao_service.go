package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// convertAvaliacaoToResponse converte um modelo de avaliação para response
func convertAvaliacaoToResponse(avaliacao *models.Avaliacao) *json.AvaliacaoResponse {
	response := &json.AvaliacaoResponse{
		ID:              avaliacao.ID,
		IDUsuario:       avaliacao.IDUsuario,
		IDLoja:          avaliacao.IDLoja,
		IDServico:       avaliacao.IDServico,
		IDProduto:       avaliacao.IDProduto,
		IDAnuncio:       avaliacao.IDAnuncio,
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
	}

	// Adiciona dados da loja se existir
	if avaliacao.Loja != nil {
		response.Loja = &json.LojaResponse{
			ID:        avaliacao.Loja.ID,
			Nome:      avaliacao.Loja.Nome,
			CNPJ:      avaliacao.Loja.CNPJ,
			Imagem:    avaliacao.Loja.Imagem,
			Latitude:  avaliacao.Loja.Latitude,
			Longitude: avaliacao.Loja.Longitude,
			Categoria: avaliacao.Loja.Categoria,
			IDUsuario: avaliacao.Loja.IDUsuario,
		}
	}

	// Adiciona dados do serviço se existir
	if avaliacao.Servico != nil {
		response.Servico = &json.ServicoResponse{
			ID:        avaliacao.Servico.ID,
			Titulo:    avaliacao.Servico.Titulo,
			Descricao: avaliacao.Servico.Descricao,
			Preco:     avaliacao.Servico.Preco,
			Imagem:    avaliacao.Servico.Imagem,
			Destaque:  avaliacao.Servico.Destaque,
			Categoria: avaliacao.Servico.Categoria,
		}
	}

	// Adiciona dados do produto se existir
	if avaliacao.Produto != nil {
		response.Produto = &json.ProdutoResponse{
			ID:           avaliacao.Produto.ID,
			Nome:         avaliacao.Produto.Nome,
			Descricao:    avaliacao.Produto.Descricao,
			Preco:        avaliacao.Produto.Preco,
			Imagem:       avaliacao.Produto.Imagem,
			Estoque:      avaliacao.Produto.Estoque,
			Ativo:        avaliacao.Produto.Ativo,
			IDLoja:       avaliacao.Produto.IDLoja,
			DataCadastro: avaliacao.Produto.DataCadastro,
		}
	}

	return response
}

// CreateAvaliacao cria uma nova avaliação
func CreateAvaliacao(req json.AvaliacaoRequest) (*json.AvaliacaoResponse, error) {
	avaliacao, err := datasource.CreateAvaliacao(req)
	if err != nil {
		return nil, err
	}

	return convertAvaliacaoToResponse(avaliacao), nil
}

// GetAvaliacaoByID busca uma avaliação por ID
func GetAvaliacaoByID(id uint) (*json.AvaliacaoResponse, error) {
	avaliacao, err := datasource.GetAvaliacaoByID(id)
	if err != nil {
		return nil, err
	}

	return convertAvaliacaoToResponse(avaliacao), nil
}

// GetAllAvaliacoes retorna todas as avaliações ativas
func GetAllAvaliacoes() ([]json.AvaliacaoResponse, error) {
	avaliacoes, err := datasource.GetAllAvaliacoes()
	if err != nil {
		return nil, err
	}

	var responses []json.AvaliacaoResponse
	for _, avaliacao := range avaliacoes {
		responses = append(responses, *convertAvaliacaoToResponse(&avaliacao))
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
		avaliacoesResponse = append(avaliacoesResponse, *convertAvaliacaoToResponse(&avaliacao))
		somaNotas += float64(avaliacao.Nota)
	}

	var mediaNota float64
	if len(avaliacoes) > 0 {
		mediaNota = somaNotas / float64(len(avaliacoes))
	}

	return &json.AvaliacoesResponse{
		Avaliacoes: avaliacoesResponse,
		Total:      len(avaliacoesResponse),
		MediaNota:  mediaNota,
	}, nil
}

// GetAvaliacoesByUsuarioID retorna todas as avaliações de um usuário específico
func GetAvaliacoesByUsuarioID(idUsuario uint) (*json.AvaliacoesResponse, error) {
	avaliacoes, err := datasource.GetAvaliacoesByUsuarioID(idUsuario)
	if err != nil {
		return nil, err
	}

	var avaliacoesResponse []json.AvaliacaoResponse
	for _, avaliacao := range avaliacoes {
		avaliacoesResponse = append(avaliacoesResponse, *convertAvaliacaoToResponse(&avaliacao))
	}

	return &json.AvaliacoesResponse{
		Avaliacoes: avaliacoesResponse,
		Total:      len(avaliacoesResponse),
	}, nil
}

// GetAvaliacoesByServicoID retorna todas as avaliações de um serviço específico
func GetAvaliacoesByServicoID(idServico uint) (*json.AvaliacoesResponse, error) {
	avaliacoes, err := datasource.GetAvaliacoesByServicoID(idServico)
	if err != nil {
		return nil, err
	}

	var avaliacoesResponse []json.AvaliacaoResponse
	var somaNotas float64

	for _, avaliacao := range avaliacoes {
		avaliacoesResponse = append(avaliacoesResponse, *convertAvaliacaoToResponse(&avaliacao))
		somaNotas += float64(avaliacao.Nota)
	}

	var mediaNota float64
	if len(avaliacoes) > 0 {
		mediaNota = somaNotas / float64(len(avaliacoes))
	}

	return &json.AvaliacoesResponse{
		Avaliacoes: avaliacoesResponse,
		Total:      len(avaliacoesResponse),
		MediaNota:  mediaNota,
	}, nil
}

// GetAvaliacoesByProdutoID retorna todas as avaliações de um produto específico
func GetAvaliacoesByProdutoID(idProduto uint) (*json.AvaliacoesResponse, error) {
	avaliacoes, err := datasource.GetAvaliacoesByProdutoID(idProduto)
	if err != nil {
		return nil, err
	}

	var avaliacoesResponse []json.AvaliacaoResponse
	var somaNotas float64

	for _, avaliacao := range avaliacoes {
		avaliacoesResponse = append(avaliacoesResponse, *convertAvaliacaoToResponse(&avaliacao))
		somaNotas += float64(avaliacao.Nota)
	}

	var mediaNota float64
	if len(avaliacoes) > 0 {
		mediaNota = somaNotas / float64(len(avaliacoes))
	}

	return &json.AvaliacoesResponse{
		Avaliacoes: avaliacoesResponse,
		Total:      len(avaliacoesResponse),
		MediaNota:  mediaNota,
	}, nil
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

	return convertAvaliacaoToResponse(avaliacao), nil
}

// SoftDeleteAvaliacao realiza soft delete da avaliação
func SoftDeleteAvaliacao(id uint) error {
	return datasource.SoftDeleteAvaliacao(id)
}

// RestoreAvaliacao restaura uma avaliação que foi soft deleted
func RestoreAvaliacao(id uint) error {
	return datasource.RestoreAvaliacao(id)
}
