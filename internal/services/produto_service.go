package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// CreateProduto cria um novo produto
func CreateProduto(req json.ProdutoRequest) (*json.ProdutoResponse, error) {
	produto, err := datasource.CreateProduto(req)
	if err != nil {
		return nil, err
	}

	response := &json.ProdutoResponse{
		ID:           produto.ID,
		Nome:         produto.Nome,
		Descricao:    produto.Descricao,
		Preco:        produto.Preco,
		Imagem:       produto.Imagem,
		Estoque:      produto.Estoque,
		Ativo:        produto.Ativo,
		Categoria:    produto.Categoria,
		IDLoja:       produto.IDLoja,
		DataCadastro: produto.DataCadastro,
		Loja: json.LojaResponse{
			ID:          produto.Loja.ID,
			Nome:        produto.Loja.Nome,
			CNPJ:        produto.Loja.CNPJ,
			Imagem:      produto.Loja.Imagem,
			Latitude:    produto.Loja.Latitude,
			Longitude:   produto.Loja.Longitude,
			IDCategoria: produto.Loja.IDCategoria,
			Categoria:   produto.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// GetProdutoByID busca um produto por ID
func GetProdutoByID(id uint) (*json.ProdutoResponse, error) {
	produto, err := datasource.GetProdutoByID(id)
	if err != nil {
		return nil, err
	}

	response := &json.ProdutoResponse{
		ID:           produto.ID,
		Nome:         produto.Nome,
		Descricao:    produto.Descricao,
		Preco:        produto.Preco,
		Imagem:       produto.Imagem,
		Estoque:      produto.Estoque,
		Ativo:        produto.Ativo,
		Categoria:    produto.Categoria,
		IDLoja:       produto.IDLoja,
		DataCadastro: produto.DataCadastro,
		Loja: json.LojaResponse{
			ID:          produto.Loja.ID,
			Nome:        produto.Loja.Nome,
			CNPJ:        produto.Loja.CNPJ,
			Imagem:      produto.Loja.Imagem,
			Latitude:    produto.Loja.Latitude,
			Longitude:   produto.Loja.Longitude,
			IDCategoria: produto.Loja.IDCategoria,
			Categoria:   produto.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// GetAllProdutos retorna todos os produtos ativos
func GetAllProdutos() ([]json.ProdutoResponse, error) {
	produtos, err := datasource.GetAllProdutos()
	if err != nil {
		return nil, err
	}

	var responses []json.ProdutoResponse
	for _, produto := range produtos {
		response := json.ProdutoResponse{
			ID:           produto.ID,
			Nome:         produto.Nome,
			Descricao:    produto.Descricao,
			Preco:        produto.Preco,
			Imagem:       produto.Imagem,
			Estoque:      produto.Estoque,
			Ativo:        produto.Ativo,
			Categoria:    produto.Categoria,
			IDLoja:       produto.IDLoja,
			DataCadastro: produto.DataCadastro,
			Loja: json.LojaResponse{
				ID:          produto.Loja.ID,
				Nome:        produto.Loja.Nome,
				CNPJ:        produto.Loja.CNPJ,
				Imagem:      produto.Loja.Imagem,
				Latitude:    produto.Loja.Latitude,
				Longitude:   produto.Loja.Longitude,
				IDCategoria: produto.Loja.IDCategoria,
				Categoria:   produto.Loja.Categoria.Nome,
			},
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// GetProdutosByLojaID retorna todos os produtos de uma loja específica
func GetProdutosByLojaID(idLoja uint) (*json.ProdutosResponse, error) {
	produtos, err := datasource.GetProdutosByLojaID(idLoja)
	if err != nil {
		return nil, err
	}

	var produtosResponse []json.ProdutoResponse
	for _, produto := range produtos {
		produtoResp := json.ProdutoResponse{
			ID:           produto.ID,
			Nome:         produto.Nome,
			Descricao:    produto.Descricao,
			Preco:        produto.Preco,
			Imagem:       produto.Imagem,
			Estoque:      produto.Estoque,
			Ativo:        produto.Ativo,
			Categoria:    produto.Categoria,
			IDLoja:       produto.IDLoja,
			DataCadastro: produto.DataCadastro,
			Loja: json.LojaResponse{
				ID:          produto.Loja.ID,
				Nome:        produto.Loja.Nome,
				CNPJ:        produto.Loja.CNPJ,
				Imagem:      produto.Loja.Imagem,
				Latitude:    produto.Loja.Latitude,
				Longitude:   produto.Loja.Longitude,
				IDCategoria: produto.Loja.IDCategoria,
				Categoria:   produto.Loja.Categoria.Nome,
			},
		}
		produtosResponse = append(produtosResponse, produtoResp)
	}

	response := &json.ProdutosResponse{
		Produtos: produtosResponse,
		Total:    len(produtosResponse),
	}

	return response, nil
}

// UpdateProduto atualiza um produto existente
func UpdateProduto(id uint, req json.ProdutoRequest) (*json.ProdutoResponse, error) {
	produto, err := datasource.UpdateProduto(id, req)
	if err != nil {
		return nil, err
	}

	response := &json.ProdutoResponse{
		ID:           produto.ID,
		Nome:         produto.Nome,
		Descricao:    produto.Descricao,
		Preco:        produto.Preco,
		Imagem:       produto.Imagem,
		Estoque:      produto.Estoque,
		Ativo:        produto.Ativo,
		Categoria:    produto.Categoria,
		IDLoja:       produto.IDLoja,
		DataCadastro: produto.DataCadastro,
		Loja: json.LojaResponse{
			ID:          produto.Loja.ID,
			Nome:        produto.Loja.Nome,
			CNPJ:        produto.Loja.CNPJ,
			Imagem:      produto.Loja.Imagem,
			Latitude:    produto.Loja.Latitude,
			Longitude:   produto.Loja.Longitude,
			IDCategoria: produto.Loja.IDCategoria,
			Categoria:   produto.Loja.Categoria.Nome,
		},
	}

	return response, nil
}

// SoftDeleteProduto realiza soft delete do produto
func SoftDeleteProduto(id uint) error {
	return datasource.SoftDeleteProduto(id)
}

// RestoreProduto restaura um produto que foi soft deleted
func RestoreProduto(id uint) error {
	return datasource.RestoreProduto(id)
}
