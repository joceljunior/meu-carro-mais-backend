package json

type AvaliacaoRequest struct {
	IDUsuario  uint   `json:"id_usuario" binding:"required"`
	IDLoja     uint   `json:"id_loja" binding:"required"`
	Nota       int    `json:"nota" binding:"required,min=1,max=5"`
	Comentario string `json:"comentario,omitempty" binding:"max=500"`
}
