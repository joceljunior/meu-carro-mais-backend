package datasource

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"

	"gorm.io/gorm"
)

// CreateVendaProdutoAvulso registra venda de produto não cadastrado.
func CreateVendaProdutoAvulso(idLoja, idUsuario uint, valor float64, descricaoProduto string) (*models.VendaProdutoAvulso, error) {
	v := models.VendaProdutoAvulso{
		IDLoja:           idLoja,
		IDUsuario:        idUsuario,
		Valor:            valor,
		DescricaoProduto: descricaoProduto,
	}
	if err := database.DB.Create(&v).Error; err != nil {
		return nil, err
	}
	return GetVendaProdutoAvulsoByID(v.ID)
}

// CreateVendaProdutoAvulsoTx cria venda dentro de uma transação.
func CreateVendaProdutoAvulsoTx(tx *gorm.DB, idLoja, idUsuario uint, valor float64, descricaoProduto string) (*models.VendaProdutoAvulso, error) {
	v := models.VendaProdutoAvulso{
		IDLoja:           idLoja,
		IDUsuario:        idUsuario,
		Valor:            valor,
		DescricaoProduto: descricaoProduto,
	}
	if err := tx.Create(&v).Error; err != nil {
		return nil, err
	}
	return GetVendaProdutoAvulsoByIDWithDB(tx, v.ID)
}

// GetVendaProdutoAvulsoByIDWithDB busca venda por ID usando tx/db.
func GetVendaProdutoAvulsoByIDWithDB(db *gorm.DB, id uint) (*models.VendaProdutoAvulso, error) {
	var v models.VendaProdutoAvulso
	err := db.
		Preload("Usuario").
		Preload("Loja").
		Where("id = ?", id).
		First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// GetVendaProdutoAvulsoByID busca venda por ID com relacionamentos.
func GetVendaProdutoAvulsoByID(id uint) (*models.VendaProdutoAvulso, error) {
	var v models.VendaProdutoAvulso
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Where("id = ?", id).
		First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// GetVendasProdutoAvulsoByUsuarioID lista vendas avulsas do cliente.
func GetVendasProdutoAvulsoByUsuarioID(idUsuario uint) ([]models.VendaProdutoAvulso, error) {
	var rows []models.VendaProdutoAvulso
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Where("id_usuario = ?", idUsuario).
		Order("data_venda DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// GetVendasProdutoAvulsoByLojaID lista vendas avulsas da loja.
func GetVendasProdutoAvulsoByLojaID(idLoja uint) ([]models.VendaProdutoAvulso, error) {
	var rows []models.VendaProdutoAvulso
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Where("id_loja = ?", idLoja).
		Order("data_venda DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// GetAllVendasProdutoAvulso lista todas as vendas avulsas.
func GetAllVendasProdutoAvulso() ([]models.VendaProdutoAvulso, error) {
	var rows []models.VendaProdutoAvulso
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Order("data_venda DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}
