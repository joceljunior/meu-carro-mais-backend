package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateProduto cria um novo produto
func CreateProduto(req json.ProdutoRequest) (*models.Produto, error) {
	produto := models.Produto{
		Nome:      req.Nome,
		Descricao: req.Descricao,
		Preco:     req.Preco,
		Imagem:    req.Imagem,
		Estoque:   req.Estoque,
		Categoria: req.Categoria,
		Ativo:     true,
		IDLoja:    req.IDLoja,
	}

	err := database.DB.Create(&produto).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o produto com os relacionamentos
	return GetProdutoByID(produto.ID)
}

// GetProdutoByID busca produto por ID (apenas produtos não excluídos)
func GetProdutoByID(id uint) (*models.Produto, error) {
	var produto models.Produto
	err := database.DB.
		Preload("Loja").
		Where("id = ? AND ativo = ? AND data_exclusao IS NULL", id, true).
		First(&produto).Error
	if err != nil {
		return nil, err
	}
	return &produto, nil
}

// GetAllProdutos retorna todos os produtos ativos (não excluídos)
func GetAllProdutos() ([]models.Produto, error) {
	var produtos []models.Produto
	err := database.DB.
		Preload("Loja").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&produtos).Error
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

// GetProdutosByLojaID retorna todos os produtos de uma loja específica
func GetProdutosByLojaID(idLoja uint) ([]models.Produto, error) {
	var produtos []models.Produto
	err := database.DB.
		Preload("Loja").
		Where("id_loja = ? AND ativo = ? AND data_exclusao IS NULL", idLoja, true).
		Find(&produtos).Error
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

// UpdateProduto atualiza um produto existente
func UpdateProduto(id uint, req json.ProdutoRequest) (*models.Produto, error) {
	// Verifica se o produto existe e não foi excluído
	produto, err := GetProdutoByID(id)
	if err != nil {
		return nil, errors.New("produto não encontrado")
	}

	// Atualiza os campos
	produto.Nome = req.Nome
	produto.Descricao = req.Descricao
	produto.Preco = req.Preco
	produto.Imagem = req.Imagem
	produto.Estoque = req.Estoque
	produto.Categoria = req.Categoria
	produto.IDLoja = req.IDLoja

	err = database.DB.Save(&produto).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o produto com os relacionamentos
	return GetProdutoByID(id)
}

// SoftDeleteProduto realiza soft delete do produto (marca como excluído)
func SoftDeleteProduto(id uint) error {
	// Verifica se o produto existe e não foi excluído
	_, err := GetProdutoByID(id)
	if err != nil {
		return errors.New("produto não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Produto{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreProduto restaura um produto que foi soft deleted
func RestoreProduto(id uint) error {
	var produto models.Produto
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&produto).Error
	if err != nil {
		return errors.New("produto não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Produto{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
