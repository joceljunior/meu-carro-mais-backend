package json

// TransferenciaVeiculoRequest representa a requisição para transferir um veículo
type TransferenciaVeiculoRequest struct {
	IDVeiculo        uint   `json:"id_veiculo" binding:"required"`
	IDUsuarioDestino uint   `json:"id_usuario_destino" binding:"required"`
	Observacoes      string `json:"observacoes,omitempty"`
}

// BuscarUsuariosRequest representa a requisição para buscar usuários
type BuscarUsuariosRequest struct {
	Termo string `form:"termo"` // Termo para buscar por nome, email ou CPF
}
