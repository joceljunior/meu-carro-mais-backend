package datasource

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
)

func GetAnuncioDestaqueByLojaID(lojaID uint) (*models.Anuncio, error) {
	var anuncio models.Anuncio
	err := database.DB.Where("id_loja = ? AND destaque = ?", lojaID, true).First(&anuncio).Error
	if err != nil {
		return nil, err
	}
	return &anuncio, nil
}
