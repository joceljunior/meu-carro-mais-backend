package models

import "time"

type Servico struct {
	ID              uint       `gorm:"primaryKey"`
	Titulo          string     `gorm:"size:255"`
	Descricao       string     `gorm:"size:255"`
	Preco           float64    `gorm:"type:decimal(10,2)"`
	Imagem          string     `gorm:"size:255"`
	Destaque        bool       `gorm:"default:false"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`
	IDLoja          uint
	IDCategoria     uint
	Categoria       CategoriaServico `gorm:"foreignKey:IDCategoria"`
	Loja            Loja             `gorm:"foreignKey:IDLoja"`
}
