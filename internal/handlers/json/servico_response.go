package json

type ServicoResponse struct {
	ID          uint         `json:"id"`
	Titulo      string       `json:"titulo"`
	Descricao   string       `json:"descricao"`
	Preco       float64      `json:"preco"`
	Imagem      string       `json:"imagem"`
	Destaque    bool         `json:"destaque"`
	Distancia   float64      `json:"distancia,omitempty"`
	Categoria   string       `json:"categoria"`
	IDCategoria uint         `json:"id_categoria"`
	Loja        LojaResponse `json:"loja"`
}

type ServicosResponse struct {
	Servicos []ServicoResponse `json:"servicos"`
	Total    int               `json:"total"`
}

type CategoriaServicoResponse struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}
