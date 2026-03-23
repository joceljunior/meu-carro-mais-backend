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
	LinkInstagram   string     `gorm:"size:500"`
	LinkFacebook    string     `gorm:"size:500"`
	LinkSite        string     `gorm:"size:500"`
	HorarioFuncionamento string `gorm:"type:text"`
	// DescontoGeralPorcentagem aplica-se a produtos, serviços e veículos da loja (0–100).
	DescontoGeralPorcentagem float64 `gorm:"type:decimal(5,2);not null;default:0"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`
	IDUsuario       uint

	// Campos para vínculo de indicação (opcional)
	IDUsuarioIndicador *uint      `gorm:"index"`
	DataVinculoUsuario *time.Time `gorm:"type:datetime"`

	Cupons           []Cupom  `gorm:"foreignKey:IDLoja"`
	UsuarioIndicador *Usuario `gorm:"foreignKey:IDUsuarioIndicador"`
}
