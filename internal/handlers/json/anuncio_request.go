package json

type AnuncioRequest struct {
	Titulo             string  `json:"titulo" binding:"required"`
	Descricao          string  `json:"descricao" binding:"required"`
	Preco              float64 `json:"preco" binding:"required,min=0"`
	Imagem             string  `json:"imagem,omitempty"`
	Destaque           bool    `json:"destaque,omitempty"`
	Categoria          string  `json:"categoria" binding:"required"`
	IDLoja             uint    `json:"id_loja" binding:"required"`
	IDProduto          *uint   `json:"id_produto,omitempty"`
	IDServico          *uint   `json:"id_servico,omitempty"`
	IDVeiculo          *uint   `json:"id_veiculo,omitempty"`
	IDOfertaAutoMais   *uint   `json:"id_oferta_auto_mais,omitempty"` // Oferta Auto Mais para pagamento com moedas do app
	TipoAnuncio        string  `json:"tipo_anuncio" binding:"required,oneof=produto servico veiculo"`
	PorcentagemDesconto float64 `json:"porcentagem_desconto" binding:"min=0,max=100"` // Porcentagem de desconto (0-100)
	PrecoComDesconto   float64 `json:"preco_com_desconto" binding:"min=0"` // Preço com desconto aplicado
}
