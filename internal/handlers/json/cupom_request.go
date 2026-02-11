package json

type CupomRequest struct {
	Titulo              string  `json:"titulo" binding:"required"`
	Descricao           string  `json:"descricao" binding:"required"`
	Preco               float64 `json:"preco" binding:"required,min=0"`
	Imagem              string  `json:"imagem,omitempty"`
	Destaque            bool    `json:"destaque,omitempty"`
	Categoria           string  `json:"categoria" binding:"required"`
	IDLoja              *uint   `json:"id_loja,omitempty"`
	IDProduto           *uint   `json:"id_produto,omitempty"`
	IDServico           *uint   `json:"id_servico,omitempty"`
	IDVeiculo           *uint   `json:"id_veiculo,omitempty"`
	IDOfertaAutoMais    *uint   `json:"id_oferta_auto_mais,omitempty"`
	IDUsuario           *uint   `json:"id_usuario,omitempty"`
	TipoCupom           string  `json:"tipo_cupom" binding:"required,oneof=produto servico veiculo"`
	PorcentagemDesconto float64 `json:"porcentagem_desconto" binding:"min=0,max=100"`
	PrecoComDesconto    float64 `json:"preco_com_desconto" binding:"min=0"`
}
