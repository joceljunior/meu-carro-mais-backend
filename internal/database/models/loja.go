package models

type Loja struct {
	ID          uint    `gorm:"primaryKey"`
	Nome        string  `gorm:"size:255"`
	CNPJ        string  `gorm:"size:255;unique"`
	Imagem      string  `gorm:"size:255"`
	Latitude    float64 `gorm:"type:decimal(10,8)"`
	Longitude   float64 `gorm:"type:decimal(11,8)"`
	IDCategoria uint
	Categoria   CategoriaLojista `gorm:"foreignKey:IDCategoria"`
}
