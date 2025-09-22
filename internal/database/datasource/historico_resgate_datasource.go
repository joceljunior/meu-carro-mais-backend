package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateHistoricoResgate cria um novo histórico de resgate
func CreateHistoricoResgate(req json.HistoricoResgateRequest) (*models.HistoricoResgate, error) {
	// Validação: apenas um dos IDs deve ser preenchido
	count := 0
	if req.IDProduto != nil {
		count++
	}
	if req.IDServico != nil {
		count++
	}
	if req.IDVeiculo != nil {
		count++
	}

	if count != 1 {
		return nil, errors.New("deve ser informado exatamente um ID: produto, serviço ou veículo")
	}

	historico := models.HistoricoResgate{
		IDUsuario:   req.IDUsuario,
		IDProduto:   req.IDProduto,
		IDServico:   req.IDServico,
		IDVeiculo:   req.IDVeiculo,
		IDLoja:      req.IDLoja,
		TipoResgate: req.TipoResgate,
		Valor:       req.Valor,
		Status:      "pendente", // Status padrão
	}

	// Se status foi informado, usa o informado
	if req.Status != "" {
		historico.Status = req.Status
	}

	err := database.DB.Create(&historico).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o histórico com os relacionamentos
	return GetHistoricoResgateByID(historico.ID)
}

// GetHistoricoResgateByID busca histórico por ID (apenas não excluídos)
func GetHistoricoResgateByID(id uint) (*models.HistoricoResgate, error) {
	var historico models.HistoricoResgate
	err := database.DB.
		Preload("Usuario").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&historico).Error
	if err != nil {
		return nil, err
	}
	return &historico, nil
}

// GetAllHistoricosResgate retorna todos os históricos ativos (não excluídos)
func GetAllHistoricosResgate() ([]models.HistoricoResgate, error) {
	var historicos []models.HistoricoResgate
	err := database.DB.
		Preload("Usuario").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("data_exclusao IS NULL").
		Order("data_resgate DESC").
		Find(&historicos).Error
	if err != nil {
		return nil, err
	}
	return historicos, nil
}

// GetHistoricosResgateByUsuarioID retorna todos os históricos de um usuário específico
func GetHistoricosResgateByUsuarioID(idUsuario uint) ([]models.HistoricoResgate, error) {
	var historicos []models.HistoricoResgate
	err := database.DB.
		Preload("Usuario").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("id_usuario = ? AND data_exclusao IS NULL", idUsuario).
		Order("data_resgate DESC").
		Find(&historicos).Error
	if err != nil {
		return nil, err
	}
	return historicos, nil
}

// GetHistoricosResgateByLojaID retorna todos os históricos de uma loja específica
func GetHistoricosResgateByLojaID(idLoja uint) ([]models.HistoricoResgate, error) {
	var historicos []models.HistoricoResgate
	err := database.DB.
		Preload("Usuario").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Order("data_resgate DESC").
		Find(&historicos).Error
	if err != nil {
		return nil, err
	}
	return historicos, nil
}

// UpdateHistoricoResgate atualiza um histórico existente
func UpdateHistoricoResgate(id uint, req json.HistoricoResgateRequest) (*models.HistoricoResgate, error) {
	// Verifica se o histórico existe e não foi excluído
	historico, err := GetHistoricoResgateByID(id)
	if err != nil {
		return nil, errors.New("histórico não encontrado")
	}

	// Atualiza os campos
	historico.IDUsuario = req.IDUsuario
	historico.IDProduto = req.IDProduto
	historico.IDServico = req.IDServico
	historico.IDVeiculo = req.IDVeiculo
	historico.IDLoja = req.IDLoja
	historico.TipoResgate = req.TipoResgate
	historico.Valor = req.Valor
	if req.Status != "" {
		historico.Status = req.Status
	}

	err = database.DB.Save(&historico).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o histórico com os relacionamentos
	return GetHistoricoResgateByID(id)
}

// UpdateStatusHistoricoResgate atualiza apenas o status de um histórico
func UpdateStatusHistoricoResgate(id uint, status string) error {
	// Verifica se o histórico existe e não foi excluído
	_, err := GetHistoricoResgateByID(id)
	if err != nil {
		return errors.New("histórico não encontrado")
	}

	// Atualiza apenas o status
	err = database.DB.Model(&models.HistoricoResgate{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleteHistoricoResgate realiza soft delete do histórico (marca como excluído)
func SoftDeleteHistoricoResgate(id uint) error {
	// Verifica se o histórico existe e não foi excluído
	_, err := GetHistoricoResgateByID(id)
	if err != nil {
		return errors.New("histórico não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.HistoricoResgate{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreHistoricoResgate restaura um histórico que foi soft deleted
func RestoreHistoricoResgate(id uint) error {
	var historico models.HistoricoResgate
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&historico).Error
	if err != nil {
		return errors.New("histórico não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.HistoricoResgate{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
