package json

type AvaliacaoRequest struct {
	IDUsuario  uint   `json:"id_usuario" binding:"required"`
	IDLoja     *uint  `json:"id_loja,omitempty"`
	IDServico  *uint  `json:"id_servico,omitempty"`
	IDProduto  *uint  `json:"id_produto,omitempty"`
	IDCupom    *uint  `json:"id_cupom,omitempty"`
	Nota       int    `json:"nota" binding:"required,min=1,max=5"`
	Comentario string `json:"comentario,omitempty" binding:"max=500"`
}
