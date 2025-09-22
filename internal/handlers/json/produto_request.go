package json

type ProdutoRequest struct {
	Nome      string  `json:"nome" binding:"required"`
	Descricao string  `json:"descricao,omitempty"`
	Preco     float64 `json:"preco" binding:"required,min=0"`
	Imagem    string  `json:"imagem,omitempty"`
	Estoque   int     `json:"estoque" binding:"min=0"`
	IDLoja    uint    `json:"id_loja" binding:"required"`
}
