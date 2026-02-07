package models

import "time"

// StatusTransferencia representa o status de uma transferência
type StatusTransferencia string

const (
	StatusTransferenciaPendente   StatusTransferencia = "pendente"
	StatusTransferenciaConfirmada StatusTransferencia = "confirmada"
	StatusTransferenciaCancelada  StatusTransferencia = "cancelada"
)

// TipoTransferencia representa o tipo/origem da transferência
type TipoTransferencia string

const (
	TipoTransferenciaManual    TipoTransferencia = "manual"    // Transferência manual entre usuários
	TipoTransferenciaVendaLoja TipoTransferencia = "venda_loja" // Transferência automática por venda em loja
)

// TransferenciaVeiculo representa uma transferência de propriedade de um veículo
type TransferenciaVeiculo struct {
	ID                   uint                `gorm:"primaryKey"`
	IDVeiculo            uint                `gorm:"not null;index"`
	IDUsuarioOrigem      uint                `gorm:"not null;index"` // Dono anterior do veículo
	IDUsuarioDestino     uint                `gorm:"not null;index"` // Novo dono do veículo
	IDLojaVenda          *uint               `gorm:"index"`          // Loja que realizou a venda (se for venda em loja)
	IDHistoricoResgate   *uint               `gorm:"index"`          // Histórico de resgate que originou a transferência (se for venda)
	TipoTransferencia    TipoTransferencia   `gorm:"size:20;not null;default:'manual'"`
	Status               StatusTransferencia `gorm:"size:20;not null;default:'confirmada'"`
	Observacoes          string              `gorm:"type:text"`
	DataTransferencia    time.Time           `gorm:"autoCreateTime"`
	DataAtualizacao      time.Time           `gorm:"autoUpdateTime"`
	DataExclusao         *time.Time          `gorm:"index"`

	// Relacionamentos
	Veiculo         Veiculo  `gorm:"foreignKey:IDVeiculo"`
	UsuarioOrigem   Usuario  `gorm:"foreignKey:IDUsuarioOrigem"`
	UsuarioDestino  Usuario  `gorm:"foreignKey:IDUsuarioDestino"`
	LojaVenda       *Loja    `gorm:"foreignKey:IDLojaVenda"`
}
