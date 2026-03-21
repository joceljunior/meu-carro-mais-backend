package json

type ResgatarCupomRequest struct {
	IDUsuario          uint  `json:"id_usuario" binding:"required"`
	IDVeiculoUsuario   *uint `json:"id_veiculo_usuario"`
	MoedasUtilizadas   int   `json:"moedas_utilizadas" binding:"omitempty,min=0"`
}
