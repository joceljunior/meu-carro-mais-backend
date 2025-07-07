package models

import "time"

type HistoricoPlanoUsuario struct {
	ID         uint      `gorm:"primaryKey"`
	UsuarioID  uint
	PlanoID    uint
	DataCriacao time.Time `gorm:"autoCreateTime"`
} 