package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateHistoricoPagamento cria um novo histórico de pagamento
func CreateHistoricoPagamento(req json.CheckoutRequest, sessionID string) (*models.HistoricoPagamento, error) {
	// Verifica se o usuário existe
	var usuario models.Usuario
	err := database.DB.Where("id = ? AND data_exclusao IS NULL", req.IDUsuario).First(&usuario).Error
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Define o valor baseado no tipo de plano
	var valor float64
	switch req.TipoPlano {
	case models.TipoPlanoMonthly:
		valor = 29.90
	case models.TipoPlanoYearly:
		valor = 299.00
	default:
		return nil, errors.New("tipo de plano inválido")
	}

	historico := models.HistoricoPagamento{
		IDUsuario:       req.IDUsuario,
		StripeSessionID: sessionID,
		Status:          models.StatusPagamentoPending,
		TipoPlano:       req.TipoPlano,
		Valor:           valor,
		Moeda:           "BRL",
	}

	err = database.DB.Create(&historico).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o histórico com os relacionamentos
	return GetHistoricoPagamentoByID(historico.ID)
}

// GetHistoricoPagamentoByID busca histórico de pagamento por ID
func GetHistoricoPagamentoByID(id uint) (*models.HistoricoPagamento, error) {
	var historico models.HistoricoPagamento
	err := database.DB.
		Preload("Usuario").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&historico).Error
	if err != nil {
		return nil, err
	}
	return &historico, nil
}

// GetHistoricoPagamentoBySessionID busca histórico de pagamento por Stripe Session ID
func GetHistoricoPagamentoBySessionID(sessionID string) (*models.HistoricoPagamento, error) {
	var historico models.HistoricoPagamento
	err := database.DB.
		Preload("Usuario").
		Where("stripe_session_id = ? AND data_exclusao IS NULL", sessionID).
		First(&historico).Error
	if err != nil {
		return nil, err
	}
	return &historico, nil
}

// GetHistoricosPagamentoByUsuarioID retorna todos os históricos de pagamento de um usuário
func GetHistoricosPagamentoByUsuarioID(idUsuario uint) ([]models.HistoricoPagamento, error) {
	var historicos []models.HistoricoPagamento
	err := database.DB.
		Preload("Usuario").
		Where("id_usuario = ? AND data_exclusao IS NULL", idUsuario).
		Order("data_criacao DESC").
		Find(&historicos).Error
	if err != nil {
		return nil, err
	}
	return historicos, nil
}

// GetAllHistoricosPagamento retorna todos os históricos de pagamento ativos
func GetAllHistoricosPagamento() ([]models.HistoricoPagamento, error) {
	var historicos []models.HistoricoPagamento
	err := database.DB.
		Preload("Usuario").
		Where("data_exclusao IS NULL").
		Order("data_criacao DESC").
		Find(&historicos).Error
	if err != nil {
		return nil, err
	}
	return historicos, nil
}

// UpdateHistoricoPagamento atualiza um histórico de pagamento
func UpdateHistoricoPagamento(id uint, updates map[string]interface{}) (*models.HistoricoPagamento, error) {
	// Verifica se o histórico existe
	_, err := GetHistoricoPagamentoByID(id)
	if err != nil {
		return nil, errors.New("histórico de pagamento não encontrado")
	}

	// Adiciona a data de atualização
	updates["data_atualizacao"] = time.Now()

	err = database.DB.Model(&models.HistoricoPagamento{}).
		Where("id = ?", id).
		Updates(updates).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o histórico atualizado
	return GetHistoricoPagamentoByID(id)
}

// UpdateHistoricoPagamentoBySessionID atualiza um histórico de pagamento por Stripe Session ID
func UpdateHistoricoPagamentoBySessionID(sessionID string, updates map[string]interface{}) (*models.HistoricoPagamento, error) {
	// Verifica se o histórico existe
	_, err := GetHistoricoPagamentoBySessionID(sessionID)
	if err != nil {
		return nil, errors.New("histórico de pagamento não encontrado")
	}

	// Adiciona a data de atualização
	updates["data_atualizacao"] = time.Now()

	err = database.DB.Model(&models.HistoricoPagamento{}).
		Where("stripe_session_id = ?", sessionID).
		Updates(updates).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o histórico atualizado
	return GetHistoricoPagamentoBySessionID(sessionID)
}

// UpdateStatusHistoricoPagamento atualiza apenas o status de um histórico de pagamento
func UpdateStatusHistoricoPagamento(id uint, status string) error {
	// Verifica se o histórico existe
	_, err := GetHistoricoPagamentoByID(id)
	if err != nil {
		return errors.New("histórico de pagamento não encontrado")
	}

	updates := map[string]interface{}{
		"status":           status,
		"data_atualizacao": time.Now(),
	}

	// Se o status for completed, define a data de pagamento
	if status == models.StatusPagamentoCompleted {
		updates["data_pagamento"] = time.Now()

		// Define a data de vencimento baseada no tipo de plano
		historico, err := GetHistoricoPagamentoByID(id)
		if err == nil && historico != nil {
			var vencimento time.Time
			if historico.TipoPlano == models.TipoPlanoMonthly {
				vencimento = time.Now().AddDate(0, 1, 0) // 1 mês
			} else if historico.TipoPlano == models.TipoPlanoYearly {
				vencimento = time.Now().AddDate(1, 0, 0) // 1 ano
			}
			updates["data_vencimento"] = vencimento
		}
	}

	err = database.DB.Model(&models.HistoricoPagamento{}).
		Where("id = ?", id).
		Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateStatusHistoricoPagamentoBySessionID atualiza apenas o status de um histórico de pagamento por Stripe Session ID
func UpdateStatusHistoricoPagamentoBySessionID(sessionID string, status string) error {
	// Verifica se o histórico existe
	historico, err := GetHistoricoPagamentoBySessionID(sessionID)
	if err != nil {
		return errors.New("histórico de pagamento não encontrado")
	}

	updates := map[string]interface{}{
		"status":           status,
		"data_atualizacao": time.Now(),
	}

	// Se o status for completed, define a data de pagamento
	if status == models.StatusPagamentoCompleted {
		updates["data_pagamento"] = time.Now()

		// Define a data de vencimento baseada no tipo de plano
		var vencimento time.Time
		if historico.TipoPlano == models.TipoPlanoMonthly {
			vencimento = time.Now().AddDate(0, 1, 0) // 1 mês
		} else if historico.TipoPlano == models.TipoPlanoYearly {
			vencimento = time.Now().AddDate(1, 0, 0) // 1 ano
		}
		updates["data_vencimento"] = vencimento
	}

	err = database.DB.Model(&models.HistoricoPagamento{}).
		Where("stripe_session_id = ?", sessionID).
		Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleteHistoricoPagamento realiza soft delete do histórico de pagamento
func SoftDeleteHistoricoPagamento(id uint) error {
	// Verifica se o histórico existe
	_, err := GetHistoricoPagamentoByID(id)
	if err != nil {
		return errors.New("histórico de pagamento não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.HistoricoPagamento{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreHistoricoPagamento restaura um histórico de pagamento que foi soft deleted
func RestoreHistoricoPagamento(id uint) error {
	var historico models.HistoricoPagamento
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&historico).Error
	if err != nil {
		return errors.New("histórico de pagamento não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.HistoricoPagamento{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
