package json

import "time"

// TransferenciaVeiculoResponse representa a resposta de uma transferência
type TransferenciaVeiculoResponse struct {
	ID                 uint                         `json:"id"`
	IDVeiculo          uint                         `json:"id_veiculo"`
	IDUsuarioOrigem    uint                         `json:"id_usuario_origem"`
	IDUsuarioDestino   uint                         `json:"id_usuario_destino"`
	IDLojaVenda        *uint                        `json:"id_loja_venda,omitempty"`
	IDHistoricoResgate *uint                        `json:"id_historico_resgate,omitempty"`
	TipoTransferencia  string                       `json:"tipo_transferencia"`
	Status             string                       `json:"status"`
	Observacoes        string                       `json:"observacoes,omitempty"`
	DataTransferencia  time.Time                    `json:"data_transferencia"`
	Veiculo            *VeiculoTransferenciaResponse `json:"veiculo,omitempty"`
	UsuarioOrigem      *UsuarioTransferenciaResponse `json:"usuario_origem,omitempty"`
	UsuarioDestino     *UsuarioTransferenciaResponse `json:"usuario_destino,omitempty"`
	LojaVenda          *LojaUsuarioResponse          `json:"loja_venda,omitempty"`
	Mensagem           string                        `json:"mensagem,omitempty"`
}

// TransferenciasResponse representa a resposta com lista de transferências
type TransferenciasResponse struct {
	Transferencias []TransferenciaVeiculoResponse `json:"transferencias"`
	Total          int                            `json:"total"`
}

// VeiculoTransferenciaResponse representa dados simplificados do veículo na transferência
type VeiculoTransferenciaResponse struct {
	ID            uint    `json:"id"`
	Marca         string  `json:"marca"`
	Modelo        string  `json:"modelo"`
	AnoFabricacao int     `json:"ano_fabricacao"`
	AnoModelo     int     `json:"ano_modelo"`
	Cor           string  `json:"cor"`
	Placa         string  `json:"placa"`
	Imagem        string  `json:"imagem,omitempty"`
}

// UsuarioTransferenciaResponse representa dados simplificados do usuário na transferência
type UsuarioTransferenciaResponse struct {
	ID       uint   `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	CPF      string `json:"cpf"`
	Telefone string `json:"telefone,omitempty"`
	Imagem   string `json:"imagem,omitempty"`
}

// UsuariosBuscaResponse representa a resposta da busca de usuários para transferência
type UsuariosBuscaResponse struct {
	Usuarios []UsuarioTransferenciaResponse `json:"usuarios"`
	Total    int                            `json:"total"`
}
