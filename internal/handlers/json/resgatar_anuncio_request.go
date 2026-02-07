package json

type ResgatarAnuncioRequest struct {
	IDUsuario        uint  `json:"id_usuario" binding:"required"`
	IDVeiculoUsuario *uint `json:"id_veiculo_usuario"` // Veículo do usuário para vincular o resgate (obrigatório para produto/serviço)
}

