package models

type Usuario struct {
	ID      uint   `gorm:"primaryKey"`
	Nome    string `gorm:"size:255"`
	Email   string `gorm:"size:255;unique"`
	Senha   string `gorm:"size:255"`
	CPF     string `gorm:"size:255;unique"`
	Imagem  string `gorm:"size:255"`
	IDPlano uint
	IDLoja  *uint
	Plano   TipoPlano `gorm:"foreignKey:IDPlano"`
	Loja    Loja      `gorm:"foreignKey:IDLoja"`
}
