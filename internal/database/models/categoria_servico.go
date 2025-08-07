package models

type CategoriaServico struct {
	ID   uint   `gorm:"primaryKey"`
	Nome string `gorm:"size:255"`
}
