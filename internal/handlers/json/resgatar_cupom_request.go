package json

type ResgatarCupomRequest struct {
	IDUsuario        uint  `json:"id_usuario" binding:"required"`
	IDVeiculoUsuario *uint `json:"id_veiculo_usuario"`
}
