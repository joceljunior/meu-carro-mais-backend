package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateDesconto cria um novo desconto para uma loja
// Valida que a loja não tem outro desconto ativo
func CreateDesconto(req json.DescontoRequest) (*models.Desconto, error) {
	// Verifica se já existe um desconto ativo para esta loja
	descontoAtivo, _ := GetDescontoAtivoByLojaID(req.IDLoja)
	if descontoAtivo != nil {
		return nil, errors.New("esta loja já possui um desconto ativo. Cancele o desconto atual antes de criar um novo")
	}

	desconto := models.Desconto{
		IDLoja:       req.IDLoja,
		Porcentagem:  req.Porcentagem,
		DataValidade: req.DataValidade,
		Ativo:        true,
	}

	err := database.DB.Create(&desconto).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o desconto com os relacionamentos
	return GetDescontoByID(desconto.ID)
}

// GetDescontoByID busca desconto por ID (apenas descontos não excluídos)
func GetDescontoByID(id uint) (*models.Desconto, error) {
	var desconto models.Desconto
	err := database.DB.
		Preload("Loja").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&desconto).Error
	if err != nil {
		return nil, err
	}
	return &desconto, nil
}

// GetDescontoAtivoByLojaID busca o desconto ativo de uma loja
func GetDescontoAtivoByLojaID(idLoja uint) (*models.Desconto, error) {
	var desconto models.Desconto
	err := database.DB.
		Preload("Loja").
		Where("id_loja = ? AND ativo = ? AND data_exclusao IS NULL AND data_validade > ?", idLoja, true, time.Now()).
		First(&desconto).Error
	if err != nil {
		return nil, err
	}
	return &desconto, nil
}

// GetAllDescontos retorna todos os descontos (não excluídos)
func GetAllDescontos() ([]models.Desconto, error) {
	var descontos []models.Desconto
	err := database.DB.
		Preload("Loja").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&descontos).Error
	if err != nil {
		return nil, err
	}
	return descontos, nil
}

// GetAllDescontosAtivos retorna todos os descontos ativos
func GetAllDescontosAtivos() ([]models.Desconto, error) {
	var descontos []models.Desconto
	err := database.DB.
		Preload("Loja").
		Where("ativo = ? AND data_exclusao IS NULL AND data_validade > ?", true, time.Now()).
		Order("data_cadastro DESC").
		Find(&descontos).Error
	if err != nil {
		return nil, err
	}
	return descontos, nil
}

// GetDescontosByLojaID retorna todos os descontos de uma loja específica (histórico)
func GetDescontosByLojaID(idLoja uint) ([]models.Desconto, error) {
	var descontos []models.Desconto
	err := database.DB.
		Preload("Loja").
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Order("data_cadastro DESC").
		Find(&descontos).Error
	if err != nil {
		return nil, err
	}
	return descontos, nil
}

// CancelarDesconto cancela (desativa) um desconto ativo
func CancelarDesconto(id uint) error {
	// Verifica se o desconto existe
	desconto, err := GetDescontoByID(id)
	if err != nil {
		return errors.New("desconto não encontrado")
	}

	if !desconto.Ativo {
		return errors.New("este desconto já está inativo")
	}

	// Desativa o desconto
	err = database.DB.Model(&models.Desconto{}).
		Where("id = ?", id).
		Update("ativo", false).Error
	if err != nil {
		return err
	}

	return nil
}

// CancelarDescontoAtivoByLojaID cancela o desconto ativo de uma loja
func CancelarDescontoAtivoByLojaID(idLoja uint) error {
	// Verifica se existe um desconto ativo
	desconto, err := GetDescontoAtivoByLojaID(idLoja)
	if err != nil {
		return errors.New("nenhum desconto ativo encontrado para esta loja")
	}

	// Desativa o desconto
	err = database.DB.Model(&models.Desconto{}).
		Where("id = ?", desconto.ID).
		Update("ativo", false).Error
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleteDesconto realiza soft delete do desconto (marca como excluído)
func SoftDeleteDesconto(id uint) error {
	// Verifica se o desconto existe e não foi excluído
	_, err := GetDescontoByID(id)
	if err != nil {
		return errors.New("desconto não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Desconto{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreDesconto restaura um desconto que foi soft deleted
func RestoreDesconto(id uint) error {
	var desconto models.Desconto
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&desconto).Error
	if err != nil {
		return errors.New("desconto não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Desconto{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}

