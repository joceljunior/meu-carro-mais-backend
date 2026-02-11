package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateRegistroInteresse cria um novo registro de interesse
func CreateRegistroInteresse(req json.RegistroInteresseRequest) (*models.RegistroInteresse, error) {
	// Verifica se o cupom existe e não foi excluído
	var cupom models.Cupom
	err := database.DB.Where("id = ? AND data_exclusao IS NULL", req.IDCupom).First(&cupom).Error
	if err != nil {
		return nil, errors.New("cupom não encontrado")
	}

	registroInteresse := models.RegistroInteresse{
		IDCupom:  req.IDCupom,
		Nome:     req.Nome,
		Email:    req.Email,
		Telefone: req.Telefone,
		Mensagem: req.Mensagem,
	}

	err = database.DB.Create(&registroInteresse).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o registro com os relacionamentos
	return GetRegistroInteresseByID(registroInteresse.ID)
}

// GetRegistroInteresseByID busca registro de interesse por ID (apenas não excluídos)
func GetRegistroInteresseByID(id uint) (*models.RegistroInteresse, error) {
	var registroInteresse models.RegistroInteresse
	err := database.DB.
		Preload("Cupom").
		Preload("Cupom.Loja").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&registroInteresse).Error
	if err != nil {
		return nil, err
	}
	return &registroInteresse, nil
}

// GetAllRegistroInteresses retorna todos os registros de interesse ativos (não excluídos)
func GetAllRegistroInteresses() ([]models.RegistroInteresse, error) {
	var registrosInteresse []models.RegistroInteresse
	err := database.DB.
		Preload("Cupom").
		Preload("Cupom.Loja").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&registrosInteresse).Error
	if err != nil {
		return nil, err
	}
	return registrosInteresse, nil
}

// GetRegistroInteressesByCupomID retorna todos os registros de interesse de um cupom específico
func GetRegistroInteressesByCupomID(cupomID uint) ([]models.RegistroInteresse, error) {
	var registrosInteresse []models.RegistroInteresse
	err := database.DB.
		Preload("Cupom").
		Preload("Cupom.Loja").
		Where("id_cupom = ? AND data_exclusao IS NULL", cupomID).
		Order("data_cadastro DESC").
		Find(&registrosInteresse).Error
	if err != nil {
		return nil, err
	}
	return registrosInteresse, nil
}

// SoftDeleteRegistroInteresse realiza soft delete do registro de interesse (marca como excluído)
func SoftDeleteRegistroInteresse(id uint) error {
	// Verifica se o registro existe e não foi excluído
	_, err := GetRegistroInteresseByID(id)
	if err != nil {
		return errors.New("registro de interesse não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.RegistroInteresse{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreRegistroInteresse restaura um registro de interesse que foi soft deleted
func RestoreRegistroInteresse(id uint) error {
	var registroInteresse models.RegistroInteresse
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&registroInteresse).Error
	if err != nil {
		return errors.New("registro de interesse não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.RegistroInteresse{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
