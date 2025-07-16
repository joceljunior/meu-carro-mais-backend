package models

type Anuncio struct {
	ID          uint    `gorm:"primaryKey"`
	Titulo      string  `gorm:"size:255"`
	Descricao   string  `gorm:"size:255"`
	Preco       float64 `gorm:"type:decimal(10,2)"`
	Imagem      string  `gorm:"size:255"`
	Destaque    bool    `gorm:"default:false"`
	IDLoja      uint
	IDCategoria uint
	Categoria   CategoriaAnuncio `gorm:"foreignKey:IDCategoria"`
	Loja        Loja             `gorm:"foreignKey:IDLoja"`
}
