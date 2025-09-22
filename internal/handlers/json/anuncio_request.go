package json

type AnuncioRequest struct {
	Titulo      string  `json:"titulo" binding:"required"`
	Descricao   string  `json:"descricao" binding:"required"`
	Preco       float64 `json:"preco" binding:"required,min=0"`
	Imagem      string  `json:"imagem,omitempty"`
	Destaque    bool    `json:"destaque,omitempty"`
	IDLoja      uint    `json:"id_loja" binding:"required"`
	IDCategoria uint    `json:"id_categoria" binding:"required"`
	IDProduto   *uint   `json:"id_produto,omitempty"`
	IDServico   *uint   `json:"id_servico,omitempty"`
	IDVeiculo   *uint   `json:"id_veiculo,omitempty"`
	TipoAnuncio string  `json:"tipo_anuncio" binding:"required,oneof=produto servico veiculo"`
}
