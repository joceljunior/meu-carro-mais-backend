package json

// CarteiraRequest representa a estrutura de dados para criação/atualização de carteira
type CarteiraRequest struct {
	UsuarioID uint `json:"usuario_id" binding:"required" example:"1"`
	Saldo     int  `json:"saldo" binding:"required,min=0" example:"1000"` // Moedas do app (valores inteiros)
}

// CarteiraSaldoRequest representa a estrutura para atualização apenas do saldo
type CarteiraSaldoRequest struct {
	Saldo int `json:"saldo" binding:"required,min=0" example:"1500"` // Moedas do app (valores inteiros)
}

// CarteiraOperacaoRequest representa a estrutura para operações de adição/subtração
type CarteiraOperacaoRequest struct {
	Valor int `json:"valor" binding:"required,min=1" example:"100"` // Moedas do app (valores inteiros)
}
