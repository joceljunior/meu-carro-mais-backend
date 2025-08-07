package datasource

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
)

// GetServicosByProximidade retorna lista de serviços ordenados por proximidade
func GetServicosByProximidade(latitude, longitude float64) ([]json.ServicoResponse, error) {
	var servicos []models.Servico

	err := database.DB.
		Preload("Categoria").
		Preload("Loja").
		Preload("Loja.Categoria").
		Find(&servicos).Error

	if err != nil {
		return nil, err
	}

	var servicosResponse []json.ServicoResponse

	for _, servico := range servicos {
		// Calcula a distância usando a fórmula de Haversine
		distancia := calcularDistancia(latitude, longitude, servico.Loja.Latitude, servico.Loja.Longitude)

		servicoResp := json.ServicoResponse{
			ID:          servico.ID,
			Titulo:      servico.Titulo,
			Descricao:   servico.Descricao,
			Preco:       servico.Preco,
			Imagem:      servico.Imagem,
			Destaque:    servico.Destaque,
			Distancia:   distancia,
			Categoria:   servico.Categoria.Nome,
			IDCategoria: servico.IDCategoria,
			Loja: json.LojaResponse{
				ID:          servico.Loja.ID,
				Nome:        servico.Loja.Nome,
				CNPJ:        servico.Loja.CNPJ,
				Imagem:      servico.Loja.Imagem,
				Latitude:    servico.Loja.Latitude,
				Longitude:   servico.Loja.Longitude,
				IDCategoria: servico.Loja.IDCategoria,
				Categoria:   servico.Loja.Categoria.Nome,
				Distancia:   distancia,
			},
		}

		servicosResponse = append(servicosResponse, servicoResp)
	}

	// Ordena por distância (menor primeiro)
	for i := 0; i < len(servicosResponse)-1; i++ {
		for j := i + 1; j < len(servicosResponse); j++ {
			if servicosResponse[i].Distancia > servicosResponse[j].Distancia {
				servicosResponse[i], servicosResponse[j] = servicosResponse[j], servicosResponse[i]
			}
		}
	}

	return servicosResponse, nil
}
