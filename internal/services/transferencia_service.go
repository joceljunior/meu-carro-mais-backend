package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// getImagemVeiculoTransferencia busca a imagem principal de um veículo para transferência
func getImagemVeiculoTransferencia(idVeiculo uint) string {
	upload, err := datasource.GetUploadPrincipalByEntidade("veiculo", idVeiculo)
	if err != nil {
		return ""
	}
	return upload.URL
}

// buildVeiculoTransferenciaResponse constrói a resposta de veículo para transferência
func buildVeiculoTransferenciaResponse(veiculo *models.Veiculo) *json.VeiculoTransferenciaResponse {
	if veiculo == nil {
		return nil
	}

	imagem := getImagemVeiculoTransferencia(veiculo.ID)

	return &json.VeiculoTransferenciaResponse{
		ID:            veiculo.ID,
		Marca:         veiculo.Marca,
		Modelo:        veiculo.Modelo,
		AnoFabricacao: veiculo.AnoFabricacao,
		AnoModelo:     veiculo.AnoModelo,
		Cor:           veiculo.Cor,
		Placa:         veiculo.Placa,
		Imagem:        imagem,
	}
}

// buildUsuarioTransferenciaResponse constrói a resposta de usuário para transferência
func buildUsuarioTransferenciaResponse(usuario *models.Usuario) *json.UsuarioTransferenciaResponse {
	if usuario == nil {
		return nil
	}

	return &json.UsuarioTransferenciaResponse{
		ID:       usuario.ID,
		Nome:     usuario.Nome,
		Email:    usuario.Email,
		CPF:      usuario.CPF,
		Telefone: usuario.Telefone,
		Imagem:   usuario.Imagem,
	}
}

// buildLojaTransferenciaResponse constrói a resposta de loja para transferência
func buildLojaTransferenciaResponse(loja *models.Loja) *json.LojaUsuarioResponse {
	if loja == nil {
		return nil
	}

	return &json.LojaUsuarioResponse{
		Id:   loja.ID,
		Nome: loja.Nome,
		Logo: loja.Imagem,
	}
}

// buildTransferenciaResponse constrói a resposta de transferência
func buildTransferenciaResponse(transferencia *models.TransferenciaVeiculo, mensagem string) *json.TransferenciaVeiculoResponse {
	resp := &json.TransferenciaVeiculoResponse{
		ID:                 transferencia.ID,
		IDVeiculo:          transferencia.IDVeiculo,
		IDUsuarioOrigem:    transferencia.IDUsuarioOrigem,
		IDUsuarioDestino:   transferencia.IDUsuarioDestino,
		IDLojaVenda:        transferencia.IDLojaVenda,
		IDHistoricoResgate: transferencia.IDHistoricoResgate,
		TipoTransferencia:  string(transferencia.TipoTransferencia),
		Status:             string(transferencia.Status),
		Observacoes:        transferencia.Observacoes,
		DataTransferencia:  transferencia.DataTransferencia,
		Mensagem:           mensagem,
	}

	// Adiciona dados do veículo
	resp.Veiculo = buildVeiculoTransferenciaResponse(&transferencia.Veiculo)

	// Adiciona dados do usuário origem
	resp.UsuarioOrigem = buildUsuarioTransferenciaResponse(&transferencia.UsuarioOrigem)

	// Adiciona dados do usuário destino
	resp.UsuarioDestino = buildUsuarioTransferenciaResponse(&transferencia.UsuarioDestino)

	// Adiciona dados da loja (se houver)
	resp.LojaVenda = buildLojaTransferenciaResponse(transferencia.LojaVenda)

	return resp
}

// TransferirVeiculoManual realiza a transferência manual de um veículo
func TransferirVeiculoManual(req json.TransferenciaVeiculoRequest, idUsuarioOrigem uint) (*json.TransferenciaVeiculoResponse, error) {
	transferencia, err := datasource.TransferirVeiculo(
		req.IDVeiculo,
		idUsuarioOrigem,
		req.IDUsuarioDestino,
		models.TipoTransferenciaManual,
		req.Observacoes,
		nil, // Sem loja (transferência manual)
		nil, // Sem histórico de resgate
	)
	if err != nil {
		return nil, err
	}

	return buildTransferenciaResponse(transferencia, "Veículo transferido com sucesso"), nil
}

// TransferirVeiculoVendaLoja realiza a transferência automática por venda em loja
func TransferirVeiculoVendaLoja(idVeiculo, idUsuarioOrigem, idUsuarioDestino uint, idLojaVenda, idHistoricoResgate *uint) (*json.TransferenciaVeiculoResponse, error) {
	observacoes := "Transferência automática por venda em loja"

	transferencia, err := datasource.TransferirVeiculo(
		idVeiculo,
		idUsuarioOrigem,
		idUsuarioDestino,
		models.TipoTransferenciaVendaLoja,
		observacoes,
		idLojaVenda,
		idHistoricoResgate,
	)
	if err != nil {
		return nil, err
	}

	return buildTransferenciaResponse(transferencia, "Veículo transferido automaticamente por venda"), nil
}

// GetTransferenciaByID busca uma transferência por ID
func GetTransferenciaByID(id uint) (*json.TransferenciaVeiculoResponse, error) {
	transferencia, err := datasource.GetTransferenciaByID(id)
	if err != nil {
		return nil, err
	}

	return buildTransferenciaResponse(transferencia, ""), nil
}

// GetTransferenciasByVeiculo retorna todas as transferências de um veículo
func GetTransferenciasByVeiculo(idVeiculo uint) (*json.TransferenciasResponse, error) {
	transferencias, err := datasource.GetTransferenciasByVeiculo(idVeiculo)
	if err != nil {
		return nil, err
	}

	var responses []json.TransferenciaVeiculoResponse
	for _, t := range transferencias {
		responses = append(responses, *buildTransferenciaResponse(&t, ""))
	}

	return &json.TransferenciasResponse{
		Transferencias: responses,
		Total:          len(responses),
	}, nil
}

// GetTransferenciasByUsuario retorna todas as transferências de um usuário
func GetTransferenciasByUsuario(idUsuario uint) (*json.TransferenciasResponse, error) {
	transferencias, err := datasource.GetTransferenciasByUsuario(idUsuario)
	if err != nil {
		return nil, err
	}

	var responses []json.TransferenciaVeiculoResponse
	for _, t := range transferencias {
		responses = append(responses, *buildTransferenciaResponse(&t, ""))
	}

	return &json.TransferenciasResponse{
		Transferencias: responses,
		Total:          len(responses),
	}, nil
}

// GetAllTransferencias retorna todas as transferências
func GetAllTransferencias() (*json.TransferenciasResponse, error) {
	transferencias, err := datasource.GetAllTransferencias()
	if err != nil {
		return nil, err
	}

	var responses []json.TransferenciaVeiculoResponse
	for _, t := range transferencias {
		responses = append(responses, *buildTransferenciaResponse(&t, ""))
	}

	return &json.TransferenciasResponse{
		Transferencias: responses,
		Total:          len(responses),
	}, nil
}

// BuscarUsuariosParaTransferencia busca usuários para selecionar como novo dono
func BuscarUsuariosParaTransferencia(termo string) (*json.UsuariosBuscaResponse, error) {
	usuarios, err := datasource.BuscarUsuariosParaTransferencia(termo)
	if err != nil {
		return nil, err
	}

	var responses []json.UsuarioTransferenciaResponse
	for _, u := range usuarios {
		responses = append(responses, json.UsuarioTransferenciaResponse{
			ID:       u.ID,
			Nome:     u.Nome,
			Email:    u.Email,
			CPF:      u.CPF,
			Telefone: u.Telefone,
			Imagem:   u.Imagem,
		})
	}

	return &json.UsuariosBuscaResponse{
		Usuarios: responses,
		Total:    len(responses),
	}, nil
}
