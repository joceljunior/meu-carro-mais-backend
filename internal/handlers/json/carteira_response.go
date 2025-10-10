package json

import "time"

// CarteiraResponse representa a estrutura de resposta da carteira
type CarteiraResponse struct {
	ID              uint      `json:"id" example:"1"`
	UsuarioID       uint      `json:"usuario_id" example:"1"`
	Saldo           int       `json:"saldo" example:"1000"` // Moedas do app (valores inteiros)
	DataCriacao     time.Time `json:"data_criacao" example:"2023-10-07T13:30:00Z"`
	DataAtualizacao time.Time `json:"data_atualizacao" example:"2023-10-07T13:30:00Z"`
	Mensagem        string    `json:"mensagem,omitempty" example:"Carteira criada com sucesso"`
}

// CarteiraComUsuarioResponse representa a carteira com dados do usuário
type CarteiraComUsuarioResponse struct {
	ID              uint      `json:"id" example:"1"`
	UsuarioID       uint      `json:"usuario_id" example:"1"`
	Saldo           int       `json:"saldo" example:"1000"` // Moedas do app (valores inteiros)
	DataCriacao     time.Time `json:"data_criacao" example:"2023-10-07T13:30:00Z"`
	DataAtualizacao time.Time `json:"data_atualizacao" example:"2023-10-07T13:30:00Z"`
	Usuario         struct {
		ID    uint   `json:"id" example:"1"`
		Nome  string `json:"nome" example:"João Silva"`
		Email string `json:"email" example:"joao@email.com"`
	} `json:"usuario"`
	Mensagem string `json:"mensagem,omitempty" example:"Carteira encontrada com sucesso"`
}

// CarteirasResponse representa uma lista de carteiras
type CarteirasResponse struct {
	Carteiras []CarteiraComUsuarioResponse `json:"carteiras"`
	Total     int                          `json:"total" example:"5"`
	Mensagem  string                       `json:"mensagem,omitempty" example:"Carteiras listadas com sucesso"`
}

// CarteiraOperacaoResponse representa a resposta de uma operação na carteira
type CarteiraOperacaoResponse struct {
	ID              uint      `json:"id" example:"1"`
	UsuarioID       uint      `json:"usuario_id" example:"1"`
	SaldoAnterior   int       `json:"saldo_anterior" example:"1000"` // Moedas do app (valores inteiros)
	SaldoAtual      int       `json:"saldo_atual" example:"1100"`    // Moedas do app (valores inteiros)
	ValorOperacao   int       `json:"valor_operacao" example:"100"`  // Moedas do app (valores inteiros)
	TipoOperacao    string    `json:"tipo_operacao" example:"adicao"`
	DataAtualizacao time.Time `json:"data_atualizacao" example:"2023-10-07T13:30:00Z"`
	Mensagem        string    `json:"mensagem" example:"Saldo adicionado com sucesso"`
}
