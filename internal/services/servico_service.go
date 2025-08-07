package services

import (
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/handlers/json"
)

func GetServicosByProximidade(latitude, longitude float64) (*json.ServicosResponse, error) {
	servicos, err := datasource.GetServicosByProximidade(latitude, longitude)
	if err != nil {
		return nil, err
	}

	resp := &json.ServicosResponse{
		Servicos: servicos,
		Total:    len(servicos),
	}

	return resp, nil
}
