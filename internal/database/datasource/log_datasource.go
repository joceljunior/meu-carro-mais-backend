package datasource

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
)

// CreateLog cria um novo registro de log
func CreateLog(log models.Log) error {
	return database.DB.Create(&log).Error
}

// GetLogsByUsuarioID retorna todos os logs de um usuário específico
func GetLogsByUsuarioID(idUsuario uint) ([]models.Log, error) {
	var logs []models.Log
	err := database.DB.
		Preload("Usuario").
		Where("id_usuario = ? AND data_exclusao IS NULL", idUsuario).
		Order("data_acao DESC").
		Find(&logs).Error
	return logs, err
}

// GetLogsByEntidade retorna todos os logs de uma entidade específica
func GetLogsByEntidade(entidade string, idEntidade uint) ([]models.Log, error) {
	var logs []models.Log
	err := database.DB.
		Preload("Usuario").
		Where("entidade = ? AND id_entidade = ? AND data_exclusao IS NULL", entidade, idEntidade).
		Order("data_acao DESC").
		Find(&logs).Error
	return logs, err
}

// GetLogsByTipoAcao retorna todos os logs de um tipo de ação específico
func GetLogsByTipoAcao(tipoAcao string) ([]models.Log, error) {
	var logs []models.Log
	err := database.DB.
		Preload("Usuario").
		Where("tipo_acao = ? AND data_exclusao IS NULL", tipoAcao).
		Order("data_acao DESC").
		Find(&logs).Error
	return logs, err
}

// GetAllLogs retorna todos os logs ativos
func GetAllLogs(limit, offset int) ([]models.Log, int64, error) {
	var logs []models.Log
	var total int64

	query := database.DB.Where("data_exclusao IS NULL")
	
	// Conta total
	if err := query.Model(&models.Log{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Busca logs com paginação
	err := query.
		Preload("Usuario").
		Order("data_acao DESC").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error

	return logs, total, err
}

// GetLogByID retorna um log específico por ID
func GetLogByID(id uint) (*models.Log, error) {
	var log models.Log
	err := database.DB.
		Preload("Usuario").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// SoftDeleteLog realiza soft delete do log
func SoftDeleteLog(id uint) error {
	return database.DB.Model(&models.Log{}).
		Where("id = ?", id).
		Update("data_exclusao", "NOW()").Error
}

