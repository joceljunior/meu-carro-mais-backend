package json

import "time"

type LoginResponse struct {
	ID                         uint       `json:"id"`
	Nome                       string     `json:"nome"`
	Email                      string     `json:"email"`
	CPF                        string     `json:"cpf"`
	Imagem                     string     `json:"imagem"`
	Telefone                   string     `json:"telefone"`
	Endereco                   string     `json:"endereco"`
	DataNascimento             *time.Time `json:"data_nascimento"`
	DataCadastro               time.Time  `json:"data_cadastro"`
	Ativo                      bool       `json:"ativo"`
	Latitude                   *float64   `json:"latitude"`
	Longitude                  *float64   `json:"longitude"`
	IDPlano                    uint       `json:"id_plano"`
	IDLoja                     *uint      `json:"id_loja"`
	Tipo                       string     `json:"tipo"`
	Status                     string     `json:"status"`
	NomePlano                  string     `json:"nome_plano"`
	IDExecutivo                *uint      `json:"id_executivo"`
	SolicitacaoExecutivo       string     `json:"solicitacao_executivo"`
	DataSolicitacaoExecutivo   *time.Time `json:"data_solicitacao_executivo"`
	MotivoSolicitacaoExecutivo string     `json:"motivo_solicitacao_executivo"`
	MoedasGerais               int        `json:"moedas_gerais"`
	MoedasPorLoja              []MoedaLojaUsuarioItem `json:"moedas_por_loja"`
	LojaUsuarioResponse        `json:"loja,omitempty"`
}

type CupomDestaqueResponse struct {
	ID        uint    `json:"id"`
	Titulo    string  `json:"titulo"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Imagem    string  `json:"imagem"`
	TipoCupom string  `json:"tipo_cupom"`
}

type LojaUsuarioResponse struct {
	Id                     uint   `json:"id"`
	Nome                   string `json:"nome"`
	Logo                   string `json:"logo,omitempty"`
	CupomDestaqueResponse `json:"cupom_destaque,omitempty"`
}
