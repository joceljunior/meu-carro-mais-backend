package datasource

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
)

func GetUserByEmail(email string, senha string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("email = ? AND senha = ?", email, senha).
		First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func CreateNewUser(email string, password string) (*models.Usuario, error) {
	user := models.Usuario{
		Email:   email,
		Senha:   password,
		IDLoja:  nil, // ou 1, conforme sua lógica
		IDPlano: 1,
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
