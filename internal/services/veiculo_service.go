package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// getImagemVeiculo busca a imagem principal de um veículo
func getImagemVeiculo(idVeiculo uint) string {
	upload, err := datasource.GetUploadPrincipalByEntidade("veiculo", idVeiculo)
	if err != nil {
		return ""
	}
	return upload.URL
}

// getFotosVeiculo busca todas as fotos de um veículo
func getFotosVeiculo(idVeiculo uint) []json.VeiculoFotoResponse {
	uploads, err := datasource.GetUploadsByVeiculoID(idVeiculo)
	if err != nil {
		return nil
	}

	var fotos []json.VeiculoFotoResponse
	for _, upload := range uploads {
		// Apenas retorna imagens (não documentos)
		if upload.Tipo == "Imagem" {
			foto := json.VeiculoFotoResponse{
				ID:          upload.ID,
				URL:         upload.URL,
				NomeArquivo: upload.NomeArquivo,
				Tamanho:     upload.Tamanho,
				TipoMime:    upload.TipoMime,
				Principal:   upload.Principal,
				Ordem:       upload.Ordem,
				DataUpload:  upload.DataUpload,
			}
			fotos = append(fotos, foto)
		}
	}

	return fotos
}

// GetVeiculosByUsuario retorna todos os veículos de um usuário
func GetVeiculosByUsuario(idUsuario uint) (*json.VeiculosResponse, error) {
	veiculos, err := datasource.GetVeiculosByUsuario(idUsuario)
	if err != nil {
		return nil, err
	}

	var veiculosResponse []json.VeiculoResponse
	for _, veiculo := range veiculos {
		imagem := getImagemVeiculo(veiculo.ID)
		fotos := getFotosVeiculo(veiculo.ID)
		veiculoResp := json.VeiculoResponse{
			ID:                  veiculo.ID,
			Marca:               veiculo.Marca,
			Modelo:              veiculo.Modelo,
			AnoFabricacao:       veiculo.AnoFabricacao,
			AnoModelo:           veiculo.AnoModelo,
			Cor:                 veiculo.Cor,
			Placa:               veiculo.Placa,
			Renavam:             veiculo.Renavam,
			Chassi:              veiculo.Chassi,
			TipoVeiculo:         veiculo.TipoVeiculo,
			Combustivel:         veiculo.Combustivel,
			Quilometragem:       veiculo.Quilometragem,
			Preco:               veiculo.Preco,
			Licenciamento:       veiculo.Licenciamento,
			IPVAPago:            veiculo.IPVAPago,
			PossuiFinanciamento: veiculo.PossuiFinanciamento,
			PossuiMultas:        veiculo.PossuiMultas,
			Observacoes:         veiculo.Observacoes,
			Imagem:              imagem,
			Fotos:               fotos,
			IDUsuario:           veiculo.IDUsuario,
			DataCadastro:        veiculo.DataCadastro,
			Ativo:               veiculo.Ativo,
		}
		veiculosResponse = append(veiculosResponse, veiculoResp)
	}

	response := &json.VeiculosResponse{
		Veiculos: veiculosResponse,
		Total:    len(veiculosResponse),
	}

	return response, nil
}

// GetHistoricosByVeiculo retorna o histórico de um veículo específico
func GetHistoricosByVeiculo(idVeiculo uint) (*json.HistoricosVeiculoResponse, error) {
	historicos, err := datasource.GetHistoricosByVeiculo(idVeiculo)
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoVeiculoResponse
	for _, historico := range historicos {
		historicoResp := json.HistoricoVeiculoResponse{
			ID:           historico.ID,
			IDVeiculo:    historico.IDVeiculo,
			IDAnuncio:    historico.IDAnuncio,
			Descricao:    historico.Descricao,
			Data:         historico.Data,
			DataCadastro: historico.DataCadastro,
		}
		historicosResponse = append(historicosResponse, historicoResp)
	}

	response := &json.HistoricosVeiculoResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}

	return response, nil
}

// GetHistoricosByUsuario retorna o histórico de todos os veículos de um usuário
func GetHistoricosByUsuario(idUsuario uint) (*json.HistoricosVeiculoResponse, error) {
	historicos, err := datasource.GetHistoricosByUsuario(idUsuario)
	if err != nil {
		return nil, err
	}

	var historicosResponse []json.HistoricoVeiculoResponse
	for _, historico := range historicos {
		historicoResp := json.HistoricoVeiculoResponse{
			ID:           historico.ID,
			IDVeiculo:    historico.IDVeiculo,
			IDAnuncio:    historico.IDAnuncio,
			Descricao:    historico.Descricao,
			Data:         historico.Data,
			DataCadastro: historico.DataCadastro,
		}
		historicosResponse = append(historicosResponse, historicoResp)
	}

	response := &json.HistoricosVeiculoResponse{
		Historicos: historicosResponse,
		Total:      len(historicosResponse),
	}

	return response, nil
}

// CreateVeiculo cria um novo veículo
func CreateVeiculo(req json.VeiculoRequest) (*json.VeiculoResponse, error) {
	veiculo, err := datasource.CreateVeiculo(req)
	if err != nil {
		return nil, err
	}

	imagem := getImagemVeiculo(veiculo.ID)
	fotos := getFotosVeiculo(veiculo.ID)
	response := &json.VeiculoResponse{
		ID:                  veiculo.ID,
		Marca:               veiculo.Marca,
		Modelo:              veiculo.Modelo,
		AnoFabricacao:       veiculo.AnoFabricacao,
		AnoModelo:           veiculo.AnoModelo,
		Cor:                 veiculo.Cor,
		Placa:               veiculo.Placa,
		Renavam:             veiculo.Renavam,
		Chassi:              veiculo.Chassi,
		TipoVeiculo:         veiculo.TipoVeiculo,
		Combustivel:         veiculo.Combustivel,
		Quilometragem:       veiculo.Quilometragem,
		Preco:               veiculo.Preco,
		Licenciamento:       veiculo.Licenciamento,
		IPVAPago:            veiculo.IPVAPago,
		PossuiFinanciamento: veiculo.PossuiFinanciamento,
		PossuiMultas:        veiculo.PossuiMultas,
		Observacoes:         veiculo.Observacoes,
		Imagem:              imagem,
		Fotos:               fotos,
		IDUsuario:           veiculo.IDUsuario,
		DataCadastro:        veiculo.DataCadastro,
		Ativo:               veiculo.Ativo,
	}

	return response, nil
}

// GetVeiculoByID busca um veículo por ID
func GetVeiculoByID(id uint) (*json.VeiculoResponse, error) {
	veiculo, err := datasource.GetVeiculoByID(id)
	if err != nil {
		return nil, err
	}

	imagem := getImagemVeiculo(veiculo.ID)
	fotos := getFotosVeiculo(veiculo.ID)
	response := &json.VeiculoResponse{
		ID:                  veiculo.ID,
		Marca:               veiculo.Marca,
		Modelo:              veiculo.Modelo,
		AnoFabricacao:       veiculo.AnoFabricacao,
		AnoModelo:           veiculo.AnoModelo,
		Cor:                 veiculo.Cor,
		Placa:               veiculo.Placa,
		Renavam:             veiculo.Renavam,
		Chassi:              veiculo.Chassi,
		TipoVeiculo:         veiculo.TipoVeiculo,
		Combustivel:         veiculo.Combustivel,
		Quilometragem:       veiculo.Quilometragem,
		Preco:               veiculo.Preco,
		Licenciamento:       veiculo.Licenciamento,
		IPVAPago:            veiculo.IPVAPago,
		PossuiFinanciamento: veiculo.PossuiFinanciamento,
		PossuiMultas:        veiculo.PossuiMultas,
		Observacoes:         veiculo.Observacoes,
		Imagem:              imagem,
		Fotos:               fotos,
		IDUsuario:           veiculo.IDUsuario,
		DataCadastro:        veiculo.DataCadastro,
		Ativo:               veiculo.Ativo,
	}

	return response, nil
}

// GetAllVeiculos retorna todos os veículos ativos
func GetAllVeiculos() ([]json.VeiculoResponse, error) {
	veiculos, err := datasource.GetAllVeiculos()
	if err != nil {
		return nil, err
	}

	var responses []json.VeiculoResponse
	for _, veiculo := range veiculos {
		imagem := getImagemVeiculo(veiculo.ID)
		fotos := getFotosVeiculo(veiculo.ID)
		response := json.VeiculoResponse{
			ID:                  veiculo.ID,
			Marca:               veiculo.Marca,
			Modelo:              veiculo.Modelo,
			AnoFabricacao:       veiculo.AnoFabricacao,
			AnoModelo:           veiculo.AnoModelo,
			Cor:                 veiculo.Cor,
			Placa:               veiculo.Placa,
			Renavam:             veiculo.Renavam,
			Chassi:              veiculo.Chassi,
			TipoVeiculo:         veiculo.TipoVeiculo,
			Combustivel:         veiculo.Combustivel,
			Quilometragem:       veiculo.Quilometragem,
			Preco:               veiculo.Preco,
			Licenciamento:       veiculo.Licenciamento,
			IPVAPago:            veiculo.IPVAPago,
			PossuiFinanciamento: veiculo.PossuiFinanciamento,
			PossuiMultas:        veiculo.PossuiMultas,
			Observacoes:         veiculo.Observacoes,
			Imagem:              imagem,
			Fotos:               fotos,
			IDUsuario:           veiculo.IDUsuario,
			DataCadastro:        veiculo.DataCadastro,
			Ativo:               veiculo.Ativo,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateVeiculo atualiza um veículo existente
func UpdateVeiculo(id uint, req json.VeiculoRequest) (*json.VeiculoResponse, error) {
	veiculo, err := datasource.UpdateVeiculo(id, req)
	if err != nil {
		return nil, err
	}

	imagem := getImagemVeiculo(veiculo.ID)
	fotos := getFotosVeiculo(veiculo.ID)
	response := &json.VeiculoResponse{
		ID:                  veiculo.ID,
		Marca:               veiculo.Marca,
		Modelo:              veiculo.Modelo,
		AnoFabricacao:       veiculo.AnoFabricacao,
		AnoModelo:           veiculo.AnoModelo,
		Cor:                 veiculo.Cor,
		Placa:               veiculo.Placa,
		Renavam:             veiculo.Renavam,
		Chassi:              veiculo.Chassi,
		TipoVeiculo:         veiculo.TipoVeiculo,
		Combustivel:         veiculo.Combustivel,
		Quilometragem:       veiculo.Quilometragem,
		Preco:               veiculo.Preco,
		Licenciamento:       veiculo.Licenciamento,
		IPVAPago:            veiculo.IPVAPago,
		PossuiFinanciamento: veiculo.PossuiFinanciamento,
		PossuiMultas:        veiculo.PossuiMultas,
		Observacoes:         veiculo.Observacoes,
		Imagem:              imagem,
		Fotos:               fotos,
		IDUsuario:           veiculo.IDUsuario,
		DataCadastro:        veiculo.DataCadastro,
		Ativo:               veiculo.Ativo,
	}

	return response, nil
}

// SoftDeleteVeiculo realiza soft delete do veículo
func SoftDeleteVeiculo(id uint) error {
	return datasource.SoftDeleteVeiculo(id)
}

// RestoreVeiculo restaura um veículo que foi soft deleted
func RestoreVeiculo(id uint) error {
	return datasource.RestoreVeiculo(id)
}
