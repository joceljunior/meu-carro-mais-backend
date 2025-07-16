package datasource

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
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

func CreateNewUser(json json.UserRequest) (*string, error) {
	user := models.Usuario{
		Email:   json.Email,
		Senha:   json.Password,
		IDLoja:  nil,
		Nome:    json.Nome,
		CPF:     json.CPF,
		Imagem:  json.Imagem,
		IDPlano: 1,
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	resp := "Usuário criado com sucesso"
	return &resp, nil
}
