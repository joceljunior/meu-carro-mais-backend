package models

import "time"

type Usuario struct {
	ID              uint       `gorm:"primaryKey"`
	Nome            string     `gorm:"size:255"`
	Email           string     `gorm:"size:255;unique"`
	Senha           string     `gorm:"size:255"`
	CPF             string     `gorm:"size:255;unique"`
	Imagem          string     `gorm:"size:255"`
	Telefone        string     `gorm:"size:20"`
	Endereco        string     `gorm:"size:500"`
	DataNascimento  *time.Time `gorm:"type:date"`
	DataCadastro    time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`
	Ativo           bool       `gorm:"default:true"`
	Latitude        *float64   `gorm:"type:decimal(10,8)"`
	Longitude       *float64   `gorm:"type:decimal(11,8)"`
	IDPlano         uint
	IDLoja          *uint
	Plano           TipoPlano `gorm:"foreignKey:IDPlano"`
	Loja            Loja      `gorm:"foreignKey:IDLoja"`
	Veiculos        []Veiculo `gorm:"foreignKey:IDUsuario"`
}
