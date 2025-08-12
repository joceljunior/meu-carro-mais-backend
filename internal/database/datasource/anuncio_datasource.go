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

// GetAnuncios retorna todos os anúncios com relacionamentos
func GetAnuncios() ([]models.Anuncio, error) {
	var anuncios []models.Anuncio
	
	err := database.DB.
		Preload("Categoria").
		Preload("Loja").
		Preload("Loja.Categoria").
		Find(&anuncios).Error
	
	if err != nil {
		return nil, err
	}
	
	return anuncios, nil
}

// GetCategoriasAnuncio retorna todas as categorias de anúncio
func GetCategoriasAnuncio() ([]models.CategoriaAnuncio, error) {
	var categorias []models.CategoriaAnuncio
	
	err := database.DB.Find(&categorias).Error
	if err != nil {
		return nil, err
	}
	
	return categorias, nil
}
