package json

type CupomResponse struct {
	ID                  uint                    `json:"id"`
	Titulo              string                  `json:"titulo"`
	Descricao           string                  `json:"descricao"`
	Preco               float64                 `json:"preco"`
	Imagem              string                  `json:"imagem"`
	Destaque            bool                    `json:"destaque"`
	Categoria           string                  `json:"categoria"`
	IDLoja              *uint                   `json:"id_loja,omitempty"`
	IDProduto           *uint                   `json:"id_produto,omitempty"`
	IDServico           *uint                   `json:"id_servico,omitempty"`
	IDVeiculo           *uint                   `json:"id_veiculo,omitempty"`
	IDOfertaAutoMais    *uint                   `json:"id_oferta_auto_mais,omitempty"`
	IDUsuario           *uint                   `json:"id_usuario,omitempty"`
	TipoCupom           string                  `json:"tipo_cupom"`
	PrecoOriginal       float64                 `json:"preco_original"`
	PrecoComDesconto    float64                 `json:"preco_com_desconto"`
	PorcentagemDesconto float64                 `json:"porcentagem_desconto"`
	Avaliacao           *float64                `json:"avaliacao,omitempty"`
	EmailCriador        *string                 `json:"email_criador,omitempty"`
	TelefoneCriador     *string                 `json:"telefone_criador,omitempty"`
	Loja                *LojaResponse           `json:"loja,omitempty"`
	Produto             *ProdutoResponse        `json:"produto,omitempty"`
	Servico             *ServicoResponse        `json:"servico,omitempty"`
	Veiculo             *VeiculoResponse        `json:"veiculo,omitempty"`
	OfertaAutoMais      *OfertaAutoMaisResponse `json:"oferta_auto_mais,omitempty"`
}

type CuponsResponse struct {
	Cupons []CupomResponse `json:"cupons"`
	Total  int             `json:"total"`
}
