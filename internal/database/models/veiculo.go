package models

import "time"

type Veiculo struct {
	ID        uint      `gorm:"primaryKey"`
	Modelo    string    `gorm:"size:255;not null"`
	Ano       int       `gorm:"not null"`
	Cor       string    `gorm:"size:100;not null"`
	Placa     string    `gorm:"size:10;unique;not null"`
	IDUsuario uint      `gorm:"not null"`
	DataCadastro time.Time `gorm:"autoCreateTime"`
	Ativo     bool      `gorm:"default:true"`
	
	// Relacionamentos
	Usuario   Usuario   `gorm:"foreignKey:IDUsuario"`
	Historicos []HistoricoVeiculo `gorm:"foreignKey:IDVeiculo"`
}
