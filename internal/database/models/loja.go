package models

import "time"

type Loja struct {
	ID              uint       `gorm:"primaryKey"`
	Nome            string     `gorm:"size:255"`
	CNPJ            string     `gorm:"size:255;unique"`
	Imagem          string     `gorm:"size:255"`
	Latitude        float64    `gorm:"type:decimal(10,8)"`
	Longitude       float64    `gorm:"type:decimal(11,8)"`
	Rating          int        `gorm:"default:5"`
	IsMeuCarroMais  bool       `gorm:"default:false"`
	Categoria       string     `gorm:"size:255"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`
	IDUsuario uint
	Anuncios  []Anuncio `gorm:"foreignKey:IDLoja"`
}
