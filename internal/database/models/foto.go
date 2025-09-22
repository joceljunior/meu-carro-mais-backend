package models

import "time"

type Foto struct {
	ID              uint       `gorm:"primaryKey"`
	IDVeiculo       *uint      `gorm:"null"`             // Pode ser null se for de outro tipo
	IDVeiculoLoja   *uint      `gorm:"null"`             // Pode ser null se for de outro tipo
	IDProduto       *uint      `gorm:"null"`             // Pode ser null se for de outro tipo
	IDServico       *uint      `gorm:"null"`             // Pode ser null se for de outro tipo
	IDLoja          *uint      `gorm:"null"`             // Pode ser null se for de outro tipo
	TipoEntidade    string     `gorm:"size:20;not null"` // "veiculo", "veiculo_loja", "produto", "servico", "loja"
	URL             string     `gorm:"size:500;not null"`
	NomeArquivo     string     `gorm:"size:255;not null"`
	Tamanho         int64      `gorm:"not null"`          // Tamanho em bytes
	TipoMime        string     `gorm:"size:100;not null"` // image/jpeg, image/png, etc.
	Principal       bool       `gorm:"default:false"`     // Se é a foto principal
	Ordem           int        `gorm:"default:0"`         // Ordem de exibição
	DataUpload      time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao time.Time  `gorm:"autoUpdateTime"`
	DataExclusao    *time.Time `gorm:"index"`

	// Relacionamentos
	Veiculo     *Veiculo     `gorm:"foreignKey:IDVeiculo"`
	VeiculoLoja *VeiculoLoja `gorm:"foreignKey:IDVeiculoLoja"`
	Produto     *Produto     `gorm:"foreignKey:IDProduto"`
	Servico     *Servico     `gorm:"foreignKey:IDServico"`
	Loja        *Loja        `gorm:"foreignKey:IDLoja"`
}
