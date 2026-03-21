package models

import "time"

// VendaProdutoAvulso registra venda pela loja de produto não cadastrado no sistema.
type VendaProdutoAvulso struct {
	ID                 uint      `gorm:"primaryKey"`
	IDUsuario          uint      `gorm:"not null;index"`
	IDLoja             uint      `gorm:"not null;index"`
	Valor              float64   `gorm:"type:decimal(10,2);not null"`
	DescricaoProduto   string    `gorm:"size:500;not null"`
	DataVenda          time.Time `gorm:"autoCreateTime"`

	Usuario Usuario `gorm:"foreignKey:IDUsuario"`
	Loja    Loja    `gorm:"foreignKey:IDLoja"`
}

func (VendaProdutoAvulso) TableName() string {
	return "vendas_produto_avulso"
}
