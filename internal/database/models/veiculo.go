package models

import "time"

type Veiculo struct {
	ID                  uint       `gorm:"primaryKey"`
	Marca               string     `gorm:"size:100;not null"`              // Marca do veículo
	Modelo              string     `gorm:"size:255;not null"`              // Modelo do veículo
	AnoFabricacao       int        `gorm:"not null"`                       // Ano de fabricação
	AnoModelo           int        `gorm:"not null"`                        // Ano modelo
	Cor                 string     `gorm:"size:100;not null"`               // Cor do veículo
	Placa               string     `gorm:"size:10;unique;not null"`         // Placa do veículo
	Renavam             *string    `gorm:"size:20"`                         // RENAVAM do veículo
	Chassi              *string    `gorm:"size:50"`                        // Chassi do veículo
	TipoVeiculo         *string    `gorm:"size:50"`                        // Tipo do veículo (carro, moto, etc)
	Combustivel         *string    `gorm:"size:50"`                        // Tipo de combustível
	Quilometragem       *int       `gorm:"default:null"`                  // Quilometragem (KM)
	Preco               *float64   `gorm:"type:decimal(10,2)"`             // Preço do veículo
	Licenciamento       *string    `gorm:"size:50"`                         // Status do licenciamento
	IPVAPago            *bool       `gorm:"default:false"`                  // Se o IPVA está pago
	PossuiFinanciamento *bool       `gorm:"default:false"`                   // Se possui financiamento
	PossuiMultas        *bool       `gorm:"default:false"`                  // Se possui multas
	Observacoes         *string    `gorm:"type:text"`                      // Observações gerais
	IDUsuario           uint       `gorm:"not null"`
	DataCadastro        time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao     time.Time  `gorm:"autoUpdateTime"`
	DataExclusao        *time.Time `gorm:"index"`
	Ativo               bool       `gorm:"default:true"`

	// Relacionamentos
	Usuario    Usuario            `gorm:"foreignKey:IDUsuario"`
	Historicos []HistoricoVeiculo `gorm:"foreignKey:IDVeiculo"`
}
