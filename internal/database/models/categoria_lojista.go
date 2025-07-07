package models

type CategoriaLojista struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `gorm:"size:255;unique"`
} 