package models

import "time"

type Loja struct {
	ID              uint       `gorm:"primaryKey"`
	Nome            string     `gorm:"size:255"`
	CNPJ            string     `gorm:"size:255;unique"`
	Imagem          string     `gorm:"size:255"`
	Endereco        string     `gorm:"size:500"`
	Latitude        float64    `gorm:"type:decimal(10,8)"`
	Longitude       float64    `gorm:"type:decimal(11,8)"`
	Rating          int        `gorm:"default:5"`
	IsMeuCarroMais  bool       `gorm:"default:false"`
	Categoria       string     `gorm:"size:255"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`
	IDUsuario uint
	
	// Campos para vínculo de indicação (opcional)
	// Uma loja pode ter sido indicada por um usuário (executivo ou outro usuário)
	IDUsuarioIndicador  *uint      `gorm:"index"`        // ID do usuário que indicou esta loja (opcional)
	DataVinculoUsuario  *time.Time `gorm:"type:datetime"` // Data do vínculo com o usuário indicador
	
	Anuncios         []Anuncio `gorm:"foreignKey:IDLoja"`
	UsuarioIndicador *Usuario  `gorm:"foreignKey:IDUsuarioIndicador"` // Referência ao usuário que indicou esta loja
}
