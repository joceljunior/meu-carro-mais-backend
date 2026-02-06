package models

import "time"

type Anuncio struct {
	ID                uint       `gorm:"primaryKey"`
	Titulo            string     `gorm:"size:255"`
	Descricao         string     `gorm:"size:255"`
	Preco             float64    `gorm:"type:decimal(10,2)"`
	Imagem            string     `gorm:"size:255"`
	Destaque          bool       `gorm:"default:false"`
	Categoria         string     `gorm:"size:100"`
	DataCadastro      time.Time  `gorm:"autoCreateTime"`
	DataAtualizacao   time.Time  `gorm:"autoUpdateTime"`
	DataExclusao      *time.Time `gorm:"index"`
	IDLoja            *uint              `gorm:"null"`             // Pode ser null para anúncios de veículo do usuário
	IDProduto         *uint              `gorm:"null"`             // Pode ser null se for serviço ou veículo
	IDServico         *uint              `gorm:"null"`             // Pode ser null se for produto ou veículo
	IDVeiculo         *uint              `gorm:"null"`             // Pode ser null se for produto ou serviço
	IDOfertaAutoMais  *uint              `gorm:"null;index"`       // Referência à oferta Auto Mais (pagamento com moedas do app)
	TipoAnuncio       string             `gorm:"size:20;not null"` // "produto", "servico", "veiculo"
	PorcentagemDesconto float64          `gorm:"type:decimal(5,2);default:0"` // Porcentagem de desconto do anúncio
	PrecoComDesconto  float64            `gorm:"type:decimal(10,2)"` // Preço com desconto aplicado
	Loja              *Loja              `gorm:"foreignKey:IDLoja"`
	Produto           *Produto           `gorm:"foreignKey:IDProduto"`
	Servico           *Servico           `gorm:"foreignKey:IDServico"`
	Veiculo           *Veiculo           `gorm:"foreignKey:IDVeiculo"`
	OfertaAutoMais    *OfertaAutoMais    `gorm:"foreignKey:IDOfertaAutoMais"`
	HistoricosVeiculo []HistoricoVeiculo `gorm:"foreignKey:IDAnuncio"`
}
