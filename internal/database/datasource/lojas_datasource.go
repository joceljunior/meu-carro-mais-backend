package datasource

import (
	"math"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"sort"
)

func GetLojasByLocations(latitude string, longitude string) ([]models.Loja, error) {
	var lojas []models.Loja
	err := database.DB.Where("latitude = ? AND longitude = ?", latitude, longitude).Find(&lojas).Error
	if err != nil {
		return nil, err
	}
	return lojas, nil
}

// GetLojasByProximidade busca lojas ordenadas por proximidade do usuário
func GetLojasByProximidade(userLat, userLng float64) ([]models.Loja, error) {
	var lojas []models.Loja
	
	// Busca todas as lojas com suas categorias
	err := database.DB.Preload("Categoria").Find(&lojas).Error
	if err != nil {
		return nil, err
	}

	// Calcula a distância para cada loja e ordena
	type lojaComDistancia struct {
		Loja      models.Loja
		Distancia float64
	}

	var lojasComDistancia []lojaComDistancia

	for _, loja := range lojas {
		distancia := calcularDistancia(userLat, userLng, loja.Latitude, loja.Longitude)
		lojasComDistancia = append(lojasComDistancia, lojaComDistancia{
			Loja:      loja,
			Distancia: distancia,
		})
	}

	// Ordena por distância (mais próxima primeiro)
	sort.Slice(lojasComDistancia, func(i, j int) bool {
		return lojasComDistancia[i].Distancia < lojasComDistancia[j].Distancia
	})

	// Converte de volta para slice de lojas
	var lojasOrdenadas []models.Loja
	for _, lcd := range lojasComDistancia {
		lojasOrdenadas = append(lojasOrdenadas, lcd.Loja)
	}

	return lojasOrdenadas, nil
}

// GetCategoriasLojista retorna todas as categorias de lojista
func GetCategoriasLojista() ([]models.CategoriaLojista, error) {
	var categorias []models.CategoriaLojista
	
	err := database.DB.Find(&categorias).Error
	if err != nil {
		return nil, err
	}
	
	return categorias, nil
}

// calcularDistancia calcula a distância entre dois pontos usando a fórmula de Haversine
func calcularDistancia(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371 // Raio da Terra em km

	// Converte para radianos
	lat1Rad := lat1 * math.Pi / 180
	lng1Rad := lng1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lng2Rad := lng2 * math.Pi / 180

	// Diferenças
	dlat := lat2Rad - lat1Rad
	dlng := lng2Rad - lng1Rad

	// Fórmula de Haversine
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
