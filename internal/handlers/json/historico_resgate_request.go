package json

type HistoricoResgateRequest struct {
	IDCupom            *uint  `json:"id_cupom,omitempty"`
	IDUsuario          uint   `json:"id_usuario" binding:"required"`
	Status             string `json:"status,omitempty" binding:"omitempty,oneof=pendente efetivado cancelado"`
	MoedasUtilizadas   int    `json:"moedas_utilizadas" binding:"omitempty,min=0"`
}
