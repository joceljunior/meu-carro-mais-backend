package json

type ResgatarCupomRequest struct {
	IDUsuario        uint  `json:"id_usuario" binding:"required" example:"1"`
	IDVeiculoUsuario *uint `json:"id_veiculo_usuario,omitempty" example:"19" description:"Veículo do usuário (deve pertencer a id_usuario). Gravado no resgate; usado para histórico do veículo ao efetivar produto/serviço."`
	MoedasUtilizadas int   `json:"moedas_utilizadas" binding:"omitempty,min=0" example:"0"`
}
