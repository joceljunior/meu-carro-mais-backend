package models

import "time"

type Carteira struct {
	ID              uint `gorm:"primaryKey"`
	UsuarioID       uint
	SaldoGeral      int       `gorm:"column:saldo_geral;type:integer"` // Moedas gerais (uso em qualquer loja)
	DataCriacao     time.Time `gorm:"autoCreateTime"`
	DataAtualizacao time.Time `gorm:"autoUpdateTime"`

	// Relacionamentos
	Usuario *Usuario `gorm:"foreignKey:UsuarioID"`
}
