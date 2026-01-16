package json

type AnuncioResponse struct {
	ID                 uint                    `json:"id"`
	Titulo             string                  `json:"titulo"`
	Descricao          string                  `json:"descricao"`
	Preco              float64                 `json:"preco"`
	Imagem             string                  `json:"imagem"`
	Destaque           bool                    `json:"destaque"`
	Categoria          string                  `json:"categoria"`
	IDLoja             uint                    `json:"id_loja"`
	IDProduto          *uint                   `json:"id_produto,omitempty"`
	IDServico          *uint                   `json:"id_servico,omitempty"`
	IDVeiculo          *uint                   `json:"id_veiculo,omitempty"`
	IDOfertaAutoMais   *uint                   `json:"id_oferta_auto_mais,omitempty"`
	TipoAnuncio        string                  `json:"tipo_anuncio"`
	PrecoOriginal      float64                 `json:"preco_original"` // Preço original do produto/serviço/veículo
	PrecoComDesconto   float64                 `json:"preco_com_desconto"` // Preço com desconto aplicado
	PorcentagemDesconto float64                `json:"porcentagem_desconto"` // Porcentagem de desconto
	Avaliacao          *float64                `json:"avaliacao,omitempty"` // Média de avaliações da loja
	Loja               LojaResponse            `json:"loja"`
	Produto            *ProdutoResponse        `json:"produto,omitempty"`
	Servico            *ServicoResponse        `json:"servico,omitempty"`
	Veiculo            *VeiculoResponse        `json:"veiculo,omitempty"`
	OfertaAutoMais     *OfertaAutoMaisResponse `json:"oferta_auto_mais,omitempty"` // Oferta Auto Mais para pagamento com moedas
}

type AnunciosResponse struct {
	Anuncios []AnuncioResponse `json:"anuncios"`
	Total    int               `json:"total"`
}
