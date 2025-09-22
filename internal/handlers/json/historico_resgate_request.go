package json

type HistoricoResgateRequest struct {
	IDUsuario   uint    `json:"id_usuario" binding:"required"`
	IDProduto   *uint   `json:"id_produto,omitempty"`
	IDServico   *uint   `json:"id_servico,omitempty"`
	IDVeiculo   *uint   `json:"id_veiculo,omitempty"`
	IDLoja      uint    `json:"id_loja" binding:"required"`
	TipoResgate string  `json:"tipo_resgate" binding:"required,oneof=produto servico veiculo"`
	Valor       float64 `json:"valor" binding:"required,min=0"`
	Status      string  `json:"status,omitempty" binding:"omitempty,oneof=pendente confirmado cancelado"`
}
