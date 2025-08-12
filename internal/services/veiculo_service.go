package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

// GetVeiculosByUsuario retorna todos os veículos de um usuário
func GetVeiculosByUsuario(idUsuario uint) (*json.VeiculosResponse, error) {
	veiculos, err := datasource.GetVeiculosByUsuario(idUsuario)
	if err != nil {
		return nil, err
	}

	var veiculosResponse []json.VeiculoResponse
	for _, veiculo := range veiculos {
		veiculoResp := json.VeiculoResponse{
			ID:           veiculo.ID,
			Modelo:       veiculo.Modelo,
			Ano:          veiculo.Ano,
			Cor:          veiculo.Cor,
			Placa:        veiculo.Placa,
			IDUsuario:    veiculo.IDUsuario,
			DataCadastro: veiculo.DataCadastro,
			Ativo:        veiculo.Ativo,
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
