package models

import "time"

type Desconto struct {
	ID              uint       `gorm:"primaryKey"`
	IDLoja          uint       `gorm:"not null;index"`
	Porcentagem     float64    `gorm:"type:decimal(5,2);not null"` // Porcentagem mínima de desconto (0-100)
	Ativo           bool       `gorm:"default:true;index"`
	DataValidade    time.Time  `gorm:"not null"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Loja Loja `gorm:"foreignKey:IDLoja"`
}
