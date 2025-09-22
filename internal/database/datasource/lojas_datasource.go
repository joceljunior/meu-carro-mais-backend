package datasource

import (
	"errors"
	"math"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"sort"
	"time"
)

func GetLojasByLocations(latitude string, longitude string) ([]models.Loja, error) {
	var lojas []models.Loja
	err := database.DB.Where("latitude = ? AND longitude = ? AND data_exclusao IS NULL", latitude, longitude).Find(&lojas).Error
	if err != nil {
		return nil, err
	}
	return lojas, nil
}

// GetLojasByProximidade busca lojas ordenadas por proximidade do usuário
func GetLojasByProximidade(userLat, userLng float64) ([]models.Loja, error) {
	var lojas []models.Loja

	// Busca todas as lojas com suas categorias (apenas não excluídas)
	err := database.DB.Preload("Categoria").Where("data_exclusao IS NULL").Find(&lojas).Error
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

// CreateLoja cria uma nova loja
func CreateLoja(req json.LojaRequest) (*models.Loja, error) {
	loja := models.Loja{
		Nome:        req.Nome,
		CNPJ:        req.CNPJ,
		Imagem:      req.Imagem,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		IDCategoria: req.IDCategoria,
	}

	err := database.DB.Create(&loja).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a loja com os relacionamentos
	return GetLojaByID(loja.ID)
}

// GetLojaByID busca loja por ID (apenas lojas não excluídas)
func GetLojaByID(id uint) (*models.Loja, error) {
	var loja models.Loja
	err := database.DB.
		Preload("Categoria").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&loja).Error
	if err != nil {
		return nil, err
	}
	return &loja, nil
}

// GetAllLojas retorna todas as lojas ativas (não excluídas)
func GetAllLojas() ([]models.Loja, error) {
	var lojas []models.Loja
	err := database.DB.
		Preload("Categoria").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&lojas).Error
	if err != nil {
		return nil, err
	}
	return lojas, nil
}

// UpdateLoja atualiza uma loja existente
func UpdateLoja(id uint, req json.LojaRequest) (*models.Loja, error) {
	// Verifica se a loja existe e não foi excluída
	loja, err := GetLojaByID(id)
	if err != nil {
		return nil, errors.New("loja não encontrada")
	}

	// Atualiza os campos
	loja.Nome = req.Nome
	loja.CNPJ = req.CNPJ
	loja.Imagem = req.Imagem
	loja.Latitude = req.Latitude
	loja.Longitude = req.Longitude
	loja.IDCategoria = req.IDCategoria

	err = database.DB.Save(&loja).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a loja com os relacionamentos
	return GetLojaByID(id)
}

// SoftDeleteLoja realiza soft delete da loja (marca como excluída)
func SoftDeleteLoja(id uint) error {
	// Verifica se a loja existe e não foi excluída
	_, err := GetLojaByID(id)
	if err != nil {
		return errors.New("loja não encontrada")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Loja{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreLoja restaura uma loja que foi soft deleted
func RestoreLoja(id uint) error {
	var loja models.Loja
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&loja).Error
	if err != nil {
		return errors.New("loja não encontrada ou não foi excluída")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Loja{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
