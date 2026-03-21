package models

// UsuarioMoedasLoja saldo de moedas restritas a uma loja específica.
type UsuarioMoedasLoja struct {
	ID        uint `gorm:"primaryKey"`
	UsuarioID uint `gorm:"not null;index"`
	LojaID    uint `gorm:"not null;index"`
	Saldo     int  `gorm:"not null;default:0"`

	Usuario Usuario `gorm:"foreignKey:UsuarioID"`
	Loja    Loja    `gorm:"foreignKey:LojaID"`
}

func (UsuarioMoedasLoja) TableName() string {
	return "usuario_moedas_loja"
}
