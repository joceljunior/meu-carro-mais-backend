package datasource

import (
	"gorm.io/gorm"
	"meu-carro-mais/internal/database/models"
)

func BuscarUsuarioPorEmail(db *gorm.DB, email string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := db.Where("email = ?", email).First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
} 