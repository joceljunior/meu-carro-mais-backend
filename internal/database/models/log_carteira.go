package models

import "time"

type LogCarteira struct {
	ID          uint      `gorm:"primaryKey"`
	CarteiraID  uint
	Valor       float64
	Tipo        string    `gorm:"size:255"`
	Descricao   string    `gorm:"size:255"`
	DataCriacao time.Time `gorm:"autoCreateTime"`
} 