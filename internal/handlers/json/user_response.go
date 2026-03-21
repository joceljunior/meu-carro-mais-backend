package json

import "time"

type UserResponse struct {
	ID                         uint                 `json:"id"`
	Nome                       string               `json:"nome"`
	Email                      string               `json:"email"`
	CPF                        string               `json:"cpf"`
	Imagem                     string               `json:"imagem,omitempty"`
	Telefone                   string               `json:"telefone,omitempty"`
	Endereco                   string               `json:"endereco,omitempty"`
	DataNascimento             *time.Time           `json:"data_nascimento,omitempty"`
	DataCadastro               time.Time            `json:"data_cadastro"`
	Ativo                      bool                 `json:"ativo"`
	Latitude                   *float64             `json:"latitude,omitempty"`
	Longitude                  *float64             `json:"longitude,omitempty"`
	IDPlano                    uint                 `json:"id_plano"`
	IDLoja                     *uint                `json:"id_loja,omitempty"`
	Tipo                       string               `json:"tipo"`
	Status                     string               `json:"status"`
	SolicitacaoExecutivo       string               `json:"solicitacao_executivo,omitempty"`        // Status da solicitação: "", pendente, aprovada, rejeitada
	DataSolicitacaoExecutivo   *time.Time           `json:"data_solicitacao_executivo,omitempty"`   // Data da solicitação
	MotivoSolicitacaoExecutivo string               `json:"motivo_solicitacao_executivo,omitempty"` // Motivo/justificativa da solicitação
	IDLojaIndicadora           *uint                `json:"id_loja_indicadora,omitempty"`           // ID da loja que indicou este usuário (opcional)
	DataVinculoLoja            *time.Time           `json:"data_vinculo_loja,omitempty"`            // Data do vínculo com a loja indicadora
	LojaIndicadora             *LojaUsuarioResponse `json:"loja_indicadora,omitempty"`              // Dados da loja que indicou
	Loja                       *LojaUsuarioResponse `json:"loja,omitempty"`
	Mensagem                   string               `json:"mensagem,omitempty"`
	MoedasGerais               int                  `json:"moedas_gerais"`
	MoedasPorLoja              []MoedaLojaUsuarioItem `json:"moedas_por_loja"`
}
