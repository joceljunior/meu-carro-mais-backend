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

// GetUserByEmailOnly busca usuário apenas por email (sem validação de senha)
func GetUserByEmailOnly(email string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("email = ?", email).
		First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func CreateNewUser(json json.UserRequest) (*models.Usuario, error) {
	user := models.Usuario{
		Email:          json.Email,
		Senha:          json.Senha,
		IDLoja:         nil,
		Nome:           json.Nome,
		CPF:            json.CPF,
		Imagem:         json.Imagem,
		Telefone:       json.Telefone,
		Endereco:       json.Endereco,
		DataNascimento: json.DataNascimento,
		Latitude:       json.Latitude,
		Longitude:      json.Longitude,
		IDPlano:        1, // Plano padrão (Gratuito)
		Ativo:          true,
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o usuário com os relacionamentos
	return GetUserByEmailOnly(json.Email)
}

// CreateUserFromLogin cria um novo usuário a partir dos dados do login
func CreateUserFromLogin(loginReq json.LoginRequest) (*models.Usuario, error) {
	// Gera um CPF temporário baseado no email para evitar problemas com constraint unique
	tempCPF := "TEMP_" + loginReq.Email

	user := models.Usuario{
		Email:   loginReq.Email,
		Senha:   loginReq.Senha, // Mantém a senha para compatibilidade, mas não será usada para validação
		IDLoja:  nil,
		Nome:    "",      // Será preenchido pelo frontend se necessário
		CPF:     tempCPF, // CPF temporário que pode ser atualizado depois
		Imagem:  "",
		IDPlano: 1, // Plano padrão (Gratuito)
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o usuário com os relacionamentos
	return GetUserByEmailOnly(loginReq.Email)
}
