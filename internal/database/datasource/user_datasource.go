package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

func GetUserByEmail(email string, senha string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("email = ? AND senha = ? AND data_exclusao IS NULL", email, senha).
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
		Where("email = ? AND data_exclusao IS NULL", email).
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

// GetUserByID busca usuário por ID (apenas usuários não excluídos)
func GetUserByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

// GetAllUsers retorna todos os usuários ativos (não excluídos)
func GetAllUsers() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := database.DB.
		Preload("Loja").
		Preload("Plano").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&usuarios).Error
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// UpdateUser atualiza um usuário existente
func UpdateUser(id uint, req json.UserRequest) (*models.Usuario, error) {
	// Verifica se o usuário existe e não foi excluído
	usuario, err := GetUserByID(id)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Atualiza os campos
	usuario.Nome = req.Nome
	usuario.Email = req.Email
	usuario.Senha = req.Senha
	usuario.CPF = req.CPF
	usuario.Imagem = req.Imagem
	usuario.Telefone = req.Telefone
	usuario.Endereco = req.Endereco
	usuario.DataNascimento = req.DataNascimento
	usuario.Latitude = req.Latitude
	usuario.Longitude = req.Longitude

	err = database.DB.Save(&usuario).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o usuário com os relacionamentos
	return GetUserByID(id)
}

// SoftDeleteUser realiza soft delete do usuário (marca como excluído)
func SoftDeleteUser(id uint) error {
	// Verifica se o usuário existe e não foi excluído
	_, err := GetUserByID(id)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreUser restaura um usuário que foi soft deleted
func RestoreUser(id uint) error {
	var usuario models.Usuario
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&usuario).Error
	if err != nil {
		return errors.New("usuário não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserPlano atualiza apenas o plano de um usuário
func UpdateUserPlano(id uint, planoID uint) error {
	// Verifica se o usuário existe
	_, err := GetUserByID(id)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	// Atualiza apenas o plano
	err = database.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("id_plano", planoID).Error
	if err != nil {
		return err
	}

	return nil
}
