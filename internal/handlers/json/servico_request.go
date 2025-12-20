package json

type ServicoRequest struct {
	Titulo    string  `json:"titulo" binding:"required"`
	Descricao string  `json:"descricao" binding:"required"`
	Preco     float64 `json:"preco" binding:"required,min=0"`
	Imagem    string  `json:"imagem,omitempty"`
	Destaque  bool    `json:"destaque,omitempty"`
	Categoria string  `json:"categoria" binding:"required"`
	IDLoja    uint    `json:"id_loja" binding:"required"`
}
