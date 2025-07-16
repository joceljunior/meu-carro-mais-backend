package datasource

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
)

func GetLojasByLocations(latitude string, longitude string) ([]models.Loja, error) {
	var lojas []models.Loja
	err := database.DB.Where("latitude = ? AND longitude = ?", latitude, longitude).Find(&lojas).Error
	if err != nil {
		return nil, err
	}
	return lojas, nil
}
