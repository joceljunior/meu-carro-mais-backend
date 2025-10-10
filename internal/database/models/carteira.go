package models

import "time"

type Carteira struct {
	ID              uint `gorm:"primaryKey"`
	UsuarioID       uint
	Saldo           int       `gorm:"type:integer"` // Moedas do app (valores inteiros)
	DataCriacao     time.Time `gorm:"autoCreateTime"`
	DataAtualizacao time.Time `gorm:"autoUpdateTime"`

	// Relacionamentos
	Usuario *Usuario `gorm:"foreignKey:UsuarioID"`
}
