package json

// CarteiraRequest representa a estrutura de dados para criação/atualização de carteira
type CarteiraRequest struct {
	UsuarioID uint    `json:"usuario_id" binding:"required" example:"1"`
	Saldo     float64 `json:"saldo" binding:"required,min=0" example:"1000.00"`
}

// CarteiraSaldoRequest representa a estrutura para atualização apenas do saldo
type CarteiraSaldoRequest struct {
	Saldo float64 `json:"saldo" binding:"required,min=0" example:"1500.00"`
}

// CarteiraOperacaoRequest representa a estrutura para operações de adição/subtração
type CarteiraOperacaoRequest struct {
	Valor float64 `json:"valor" binding:"required,min=0.01" example:"100.00"`
}
