package json

type AvaliacaoRequest struct {
	IDUsuario  uint   `json:"id_usuario" binding:"required"`
	IDLoja     *uint  `json:"id_loja,omitempty"`     // Opcional - avaliar loja
	IDServico  *uint  `json:"id_servico,omitempty"`  // Opcional - avaliar serviço
	IDProduto  *uint  `json:"id_produto,omitempty"`  // Opcional - avaliar produto
	IDAnuncio  *uint  `json:"id_anuncio,omitempty"`  // Opcional - referência ao anúncio
	Nota       int    `json:"nota" binding:"required,min=1,max=5"`
	Comentario string `json:"comentario,omitempty" binding:"max=500"`
}
