package models

type Loja struct {
	ID           uint            `gorm:"primaryKey"`
	Nome         string          `gorm:"size:255"`
	CNPJ         string          `gorm:"size:255;unique"`
	Imagem       string          `gorm:"size:255"`
	IDCategoria  uint            // chave estrangeira
	Categoria    CategoriaLojista `gorm:"foreignKey:IDCategoria"`
} 