package models

type CategoriaAnuncio struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `gorm:"size:255"`
}
