package models

import "time"

type OfertaAutoMais struct {
	ID              uint       `gorm:"primaryKey"`
	IDLoja          uint       `gorm:"not null;index"`
	Nome            string     `gorm:"size:255;not null"`
	Descricao       string     `gorm:"size:500"`
	Moedas          int        `gorm:"not null"` // Quantidade de moedas Auto Mais necessárias (campo obrigatório)
	Porcentagem     float64    `gorm:"type:decimal(5,2);not null"` // Porcentagem de desconto ao usar moedas (0-100)
	Ativo           bool       `gorm:"default:true;index"`
	DataValidade    *time.Time `gorm:"index"` // Pode ser null se não tiver validade
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Loja Loja `gorm:"foreignKey:IDLoja"`
}

