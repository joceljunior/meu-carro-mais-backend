package models

import "time"

type Carteira struct {
	ID              uint      `gorm:"primaryKey"`
	UsuarioID       uint
	Saldo           float64
	DataCriacao     time.Time `gorm:"autoCreateTime"`
	DataAtualizacao time.Time `gorm:"autoUpdateTime"`
} 