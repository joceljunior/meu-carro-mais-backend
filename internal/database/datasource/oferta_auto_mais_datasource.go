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

// CreateOfertaAutoMais cria uma nova oferta Auto Mais para uma loja
func CreateOfertaAutoMais(req json.OfertaAutoMaisRequest) (*models.OfertaAutoMais, error) {
	oferta := models.OfertaAutoMais{
		IDLoja:       req.IDLoja,
		Nome:         req.Nome,
		Descricao:    req.Descricao,
		Moedas:       req.Moedas,
		Porcentagem:  req.Porcentagem,
		DataValidade: req.DataValidade,
		Ativo:        true,
	}

	err := database.DB.Create(&oferta).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a oferta com os relacionamentos
	return GetOfertaAutoMaisByID(oferta.ID)
}

// GetOfertaAutoMaisByID busca oferta por ID (apenas ofertas não excluídas)
func GetOfertaAutoMaisByID(id uint) (*models.OfertaAutoMais, error) {
	var oferta models.OfertaAutoMais
	err := database.DB.
		Preload("Loja").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&oferta).Error
	if err != nil {
		return nil, err
	}
	return &oferta, nil
}

// GetAllOfertasAutoMais retorna todas as ofertas Auto Mais (não excluídas)
func GetAllOfertasAutoMais() ([]models.OfertaAutoMais, error) {
	var ofertas []models.OfertaAutoMais
	err := database.DB.
		Preload("Loja").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&ofertas).Error
	if err != nil {
		return nil, err
	}
	return ofertas, nil
}

// GetAllOfertasAutoMaisAtivas retorna todas as ofertas Auto Mais ativas
func GetAllOfertasAutoMaisAtivas() ([]models.OfertaAutoMais, error) {
	var ofertas []models.OfertaAutoMais
	now := time.Now()
	err := database.DB.
		Preload("Loja").
		Where("ativo = ? AND data_exclusao IS NULL AND (data_validade IS NULL OR data_validade > ?)", true, now).
		Order("data_cadastro DESC").
		Find(&ofertas).Error
	if err != nil {
		return nil, err
	}
	return ofertas, nil
}

// OfertaAutoMaisComDistancia representa uma oferta Auto Mais com sua distância calculada
type OfertaAutoMaisComDistancia struct {
	Oferta    models.OfertaAutoMais
	Distancia float64
}

// GetOfertasAutoMaisAtivasByProximidade retorna ofertas Auto Mais ativas ordenadas por proximidade
func GetOfertasAutoMaisAtivasByProximidade(latitude, longitude float64) ([]OfertaAutoMaisComDistancia, error) {
	ofertas, err := GetAllOfertasAutoMaisAtivas()
	if err != nil {
		return nil, err
	}

	var ofertasComDistancia []OfertaAutoMaisComDistancia
	for _, oferta := range ofertas {
		// Calcula a distância usando a fórmula de Haversine
		distancia := calcularDistanciaOferta(latitude, longitude, oferta.Loja.Latitude, oferta.Loja.Longitude)
		ofertasComDistancia = append(ofertasComDistancia, OfertaAutoMaisComDistancia{
			Oferta:    oferta,
			Distancia: distancia,
		})
	}

	// Ordena por distância (menor primeiro)
	sort.Slice(ofertasComDistancia, func(i, j int) bool {
		return ofertasComDistancia[i].Distancia < ofertasComDistancia[j].Distancia
	})

	return ofertasComDistancia, nil
}

// calcularDistanciaOferta calcula a distância entre dois pontos usando a fórmula de Haversine
func calcularDistanciaOferta(lat1, lng1, lat2, lng2 float64) float64 {
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

// GetOfertasAutoMaisByLojaID retorna todas as ofertas Auto Mais de uma loja específica
func GetOfertasAutoMaisByLojaID(idLoja uint) ([]models.OfertaAutoMais, error) {
	var ofertas []models.OfertaAutoMais
	err := database.DB.
		Preload("Loja").
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Order("data_cadastro DESC").
		Find(&ofertas).Error
	if err != nil {
		return nil, err
	}
	return ofertas, nil
}

// GetOfertasAutoMaisAtivasByLojaID retorna apenas as ofertas Auto Mais ativas de uma loja
func GetOfertasAutoMaisAtivasByLojaID(idLoja uint) ([]models.OfertaAutoMais, error) {
	var ofertas []models.OfertaAutoMais
	now := time.Now()
	err := database.DB.
		Preload("Loja").
		Where("id_loja = ? AND ativo = ? AND data_exclusao IS NULL AND (data_validade IS NULL OR data_validade > ?)", idLoja, true, now).
		Order("data_cadastro DESC").
		Find(&ofertas).Error
	if err != nil {
		return nil, err
	}
	return ofertas, nil
}

// UpdateOfertaAutoMais atualiza uma oferta Auto Mais existente
func UpdateOfertaAutoMais(id uint, req json.OfertaAutoMaisUpdateRequest) (*models.OfertaAutoMais, error) {
	// Verifica se a oferta existe
	oferta, err := GetOfertaAutoMaisByID(id)
	if err != nil {
		return nil, errors.New("oferta não encontrada")
	}

	// Atualiza apenas os campos fornecidos
	updates := make(map[string]interface{})
	
	if req.Nome != "" {
		updates["nome"] = req.Nome
	}
	if req.Descricao != "" {
		updates["descricao"] = req.Descricao
	}
	if req.Moedas != nil {
		updates["moedas"] = *req.Moedas
	}
	if req.Porcentagem != nil {
		updates["porcentagem"] = *req.Porcentagem
	}
	if req.Ativo != nil {
		updates["ativo"] = *req.Ativo
	}
	if req.DataValidade != nil {
		updates["data_validade"] = req.DataValidade
	}

	if len(updates) > 0 {
		err = database.DB.Model(&models.OfertaAutoMais{}).
			Where("id = ?", id).
			Updates(updates).Error
		if err != nil {
			return nil, err
		}
	}

	// Retorna a oferta atualizada
	return GetOfertaAutoMaisByID(oferta.ID)
}

// DesativarOfertaAutoMais desativa uma oferta Auto Mais
func DesativarOfertaAutoMais(id uint) error {
	// Verifica se a oferta existe
	oferta, err := GetOfertaAutoMaisByID(id)
	if err != nil {
		return errors.New("oferta não encontrada")
	}

	if !oferta.Ativo {
		return errors.New("esta oferta já está inativa")
	}

	// Desativa a oferta
	err = database.DB.Model(&models.OfertaAutoMais{}).
		Where("id = ?", id).
		Update("ativo", false).Error
	if err != nil {
		return err
	}

	return nil
}

// AtivarOfertaAutoMais ativa uma oferta Auto Mais
func AtivarOfertaAutoMais(id uint) error {
	// Verifica se a oferta existe
	oferta, err := GetOfertaAutoMaisByID(id)
	if err != nil {
		return errors.New("oferta não encontrada")
	}

	if oferta.Ativo {
		return errors.New("esta oferta já está ativa")
	}

	// Ativa a oferta
	err = database.DB.Model(&models.OfertaAutoMais{}).
		Where("id = ?", id).
		Update("ativo", true).Error
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleteOfertaAutoMais realiza soft delete da oferta (marca como excluída)
func SoftDeleteOfertaAutoMais(id uint) error {
	// Verifica se a oferta existe e não foi excluída
	_, err := GetOfertaAutoMaisByID(id)
	if err != nil {
		return errors.New("oferta não encontrada")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.OfertaAutoMais{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreOfertaAutoMais restaura uma oferta que foi soft deleted
func RestoreOfertaAutoMais(id uint) error {
	var oferta models.OfertaAutoMais
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&oferta).Error
	if err != nil {
		return errors.New("oferta não encontrada ou não foi excluída")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.OfertaAutoMais{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}

