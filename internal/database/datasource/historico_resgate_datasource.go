package datasource

import (
	"errors"

	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"

	"gorm.io/gorm"
)

// CreateHistoricoResgateFromCupom cria um histórico de resgate a partir de um cupom.
// idVeiculoUsuario é o veículo do cliente (ex.: produto/serviço); nil ou 0 ignora.
func CreateHistoricoResgateFromCupom(cupomID uint, usuarioID uint, moedasUtilizadas int, idVeiculoUsuario *uint) (*models.HistoricoResgate, error) {
	// Busca o cupom
	cupom, err := GetCupomByID(cupomID)
	if err != nil {
		return nil, errors.New("cupom não encontrado")
	}

	// Verifica se o cupom não foi excluído
	if cupom.DataExclusao != nil {
		return nil, errors.New("cupom não está disponível")
	}

	historico := models.HistoricoResgate{
		IDCupom:          &cupomID,
		IDUsuario:        usuarioID,
		Status:           "pendente",
		MoedasUtilizadas: moedasUtilizadas,
	}
	if idVeiculoUsuario != nil && *idVeiculoUsuario != 0 {
		v := *idVeiculoUsuario
		historico.IDVeiculo = &v
	}

	err = database.DB.Create(&historico).Error
	if err != nil {
		return nil, err
	}

	return GetHistoricoResgateByID(historico.ID)
}

// CreateHistoricoResgate cria um novo histórico de resgate
func CreateHistoricoResgate(req json.HistoricoResgateRequest) (*models.HistoricoResgate, error) {
	historico := models.HistoricoResgate{
		IDCupom:          req.IDCupom,
		IDUsuario:        req.IDUsuario,
		Status:           "pendente",
		MoedasUtilizadas: req.MoedasUtilizadas,
	}

	if req.Status != "" {
		historico.Status = req.Status
	}

	err := database.DB.Create(&historico).Error
	if err != nil {
		return nil, err
	}

	return GetHistoricoResgateByID(historico.ID)
}

// GetHistoricoResgateByID busca histórico por ID
func GetHistoricoResgateByID(id uint) (*models.HistoricoResgate, error) {
	var historico models.HistoricoResgate
	err := database.DB.
		Preload("Usuario").
		Preload("Veiculo").
		Preload("Cupom").
		Preload("Cupom.Loja").
		Preload("Cupom.Produto").
		Preload("Cupom.Servico").
		Preload("Cupom.Veiculo").
		Preload("Cupom.OfertaAutoMais").
		Where("id = ?", id).
		First(&historico).Error
	if err != nil {
		return nil, err
	}
	return &historico, nil
}

// GetHistoricosResgateEfetivadosByVeiculoID retorna resgates efetivados vinculados ao veículo do usuário.
func GetHistoricosResgateEfetivadosByVeiculoID(idVeiculo uint) ([]models.HistoricoResgate, error) {
	var historicos []models.HistoricoResgate
	err := database.DB.
		Preload("Cupom").
		Preload("Cupom.Produto").
		Preload("Cupom.Servico").
		Where("id_veiculo = ? AND status = ? AND id_cupom IS NOT NULL", idVeiculo, "efetivado").
		Order("data_resgate DESC").
		Find(&historicos).Error
	if err != nil {
		return nil, err
	}
	return historicos, nil
}

// GetAllHistoricosResgate retorna todos os históricos
func GetAllHistoricosResgate() ([]models.HistoricoResgate, error) {
	var historicos []models.HistoricoResgate
	err := database.DB.
		Preload("Usuario").
		Preload("Veiculo").
		Preload("Cupom").
		Preload("Cupom.Loja").
		Preload("Cupom.Produto").
		Preload("Cupom.Servico").
		Preload("Cupom.Veiculo").
		Preload("Cupom.OfertaAutoMais").
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
		Preload("Veiculo").
		Preload("Cupom").
		Preload("Cupom.Loja").
		Preload("Cupom.Produto").
		Preload("Cupom.Servico").
		Preload("Cupom.Veiculo").
		Preload("Cupom.OfertaAutoMais").
		Where("id_usuario = ?", idUsuario).
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
		Preload("Veiculo").
		Preload("Cupom").
		Preload("Cupom.Loja").
		Preload("Cupom.Produto").
		Preload("Cupom.Servico").
		Preload("Cupom.Veiculo").
		Preload("Cupom.OfertaAutoMais").
		Joins("JOIN cupons ON cupons.id = historico_resgates.id_cupom").
		Where("cupons.id_loja = ?", idLoja).
		Order("historico_resgates.data_resgate DESC").
		Find(&historicos).Error
	if err != nil {
		return nil, err
	}
	return historicos, nil
}

// UpdateHistoricoResgate atualiza um histórico existente
func UpdateHistoricoResgate(id uint, req json.HistoricoResgateRequest) (*models.HistoricoResgate, error) {
	historico, err := GetHistoricoResgateByID(id)
	if err != nil {
		return nil, errors.New("histórico não encontrado")
	}

	historico.IDCupom = req.IDCupom
	historico.IDUsuario = req.IDUsuario
	historico.MoedasUtilizadas = req.MoedasUtilizadas
	if req.Status != "" {
		historico.Status = req.Status
	}

	err = database.DB.Save(&historico).Error
	if err != nil {
		return nil, err
	}

	return GetHistoricoResgateByID(id)
}

// UpdateStatusHistoricoResgate atualiza apenas o status de um histórico
func UpdateStatusHistoricoResgate(id uint, status string) error {
	_, err := GetHistoricoResgateByID(id)
	if err != nil {
		return errors.New("histórico não encontrado")
	}

	err = database.DB.Model(&models.HistoricoResgate{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateStatusHistoricoResgateWithDB atualiza status usando uma sessão DB/tx.
func UpdateStatusHistoricoResgateWithDB(db *gorm.DB, id uint, status string) error {
	return db.Model(&models.HistoricoResgate{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// GetHistoricoResgateByIDWithDB busca histórico com os mesmos preloads que GetHistoricoResgateByID.
func GetHistoricoResgateByIDWithDB(db *gorm.DB, id uint) (*models.HistoricoResgate, error) {
	var historico models.HistoricoResgate
	err := db.
		Preload("Usuario").
		Preload("Veiculo").
		Preload("Cupom").
		Preload("Cupom.Loja").
		Preload("Cupom.Produto").
		Preload("Cupom.Servico").
		Preload("Cupom.Veiculo").
		Preload("Cupom.OfertaAutoMais").
		Where("id = ?", id).
		First(&historico).Error
	if err != nil {
		return nil, err
	}
	return &historico, nil
}

// SoftDeleteHistoricoResgate realiza soft delete do histórico
func SoftDeleteHistoricoResgate(id uint) error {
	_, err := GetHistoricoResgateByID(id)
	if err != nil {
		return errors.New("histórico não encontrado")
	}

	err = database.DB.Delete(&models.HistoricoResgate{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreHistoricoResgate restaura um histórico que foi deletado
func RestoreHistoricoResgate(id uint) error {
	// Since HistoricoResgate no longer has DataExclusao, we use hard delete/restore not applicable
	return errors.New("operação não suportada na nova estrutura")
}

// GetHistoricosResgateByLojaIDViaCupom retorna históricos filtrados pela loja do cupom
func GetHistoricosResgateByLojaIDViaCupom(idLoja uint) ([]models.HistoricoResgate, error) {
	return GetHistoricosResgateByLojaID(idLoja)
}

// GetHistoricosResgateByCupomID retorna históricos de resgate de um cupom específico
func GetHistoricosResgateByCupomID(cupomID uint) ([]models.HistoricoResgate, error) {
	var historicos []models.HistoricoResgate
	err := database.DB.
		Preload("Usuario").
		Preload("Veiculo").
		Preload("Cupom").
		Preload("Cupom.Loja").
		Preload("Cupom.Produto").
		Preload("Cupom.Servico").
		Preload("Cupom.Veiculo").
		Preload("Cupom.OfertaAutoMais").
		Where("id_cupom = ?", cupomID).
		Order("data_resgate DESC").
		Find(&historicos).Error
	if err != nil {
		return nil, err
	}
	return historicos, nil
}
