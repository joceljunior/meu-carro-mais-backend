package models

import "time"

type Anuncio struct {
	ID                uint       `gorm:"primaryKey"`
	Titulo            string     `gorm:"size:255"`
	Descricao         string     `gorm:"size:255"`
	Preco             float64    `gorm:"type:decimal(10,2)"`
	Imagem            string     `gorm:"size:255"`
	Destaque          bool       `gorm:"default:false"`
	DataCadastro      time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao   time.Time  `gorm:"autoUpdateTime"`
	DataExclusao      *time.Time `gorm:"index"`
	IDLoja            uint
	IDCategoria       uint
	IDProduto         *uint              `gorm:"null"`             // Pode ser null se for serviço ou veículo
	IDServico         *uint              `gorm:"null"`             // Pode ser null se for produto ou veículo
	IDVeiculo         *uint              `gorm:"null"`             // Pode ser null se for produto ou serviço
	TipoAnuncio       string             `gorm:"size:20;not null"` // "produto", "servico", "veiculo"
	Categoria         CategoriaAnuncio   `gorm:"foreignKey:IDCategoria"`
	Loja              Loja               `gorm:"foreignKey:IDLoja"`
	Produto           *Produto           `gorm:"foreignKey:IDProduto"`
	Servico           *Servico           `gorm:"foreignKey:IDServico"`
	Veiculo           *Veiculo           `gorm:"foreignKey:IDVeiculo"`
	HistoricosVeiculo []HistoricoVeiculo `gorm:"foreignKey:IDAnuncio"`
}
