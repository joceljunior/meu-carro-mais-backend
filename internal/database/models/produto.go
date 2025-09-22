package models

import "time"

type Produto struct {
	ID              uint       `gorm:"primaryKey"`
	Nome            string     `gorm:"size:255;not null"`
	Descricao       string     `gorm:"size:500"`
	Preco           float64    `gorm:"type:decimal(10,2);not null"`
	Imagem          string     `gorm:"size:255"`
	Estoque         int        `gorm:"default:0"`
	Ativo           bool       `gorm:"default:true"`
	IDLoja          uint       `gorm:"not null"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Loja Loja `gorm:"foreignKey:IDLoja"`
}
