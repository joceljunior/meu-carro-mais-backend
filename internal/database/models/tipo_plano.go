package models

type TipoPlano struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `gorm:"size:255;unique"`
} 