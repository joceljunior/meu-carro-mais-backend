package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// GetVeiculosByUsuario retorna todos os veículos de um usuário
func GetVeiculosByUsuario(idUsuario uint) ([]models.Veiculo, error) {
	var veiculos []models.Veiculo

	err := database.DB.
		Where("id_usuario = ? AND ativo = ? AND data_exclusao IS NULL", idUsuario, true).
		Find(&veiculos).Error

	if err != nil {
		return nil, err
	}

	return veiculos, nil
}

// GetVeiculoByID retorna um veículo específico por ID
func GetVeiculoByID(id uint) (*models.Veiculo, error) {
	var veiculo models.Veiculo

	err := database.DB.
		Preload("Usuario").
		Where("id = ? AND ativo = ? AND data_exclusao IS NULL", id, true).
		First(&veiculo).Error

	if err != nil {
		return nil, err
	}

	return &veiculo, nil
}

// GetHistoricosByVeiculo retorna o histórico de um veículo específico
func GetHistoricosByVeiculo(idVeiculo uint) ([]models.HistoricoVeiculo, error) {
	var historicos []models.HistoricoVeiculo

	err := database.DB.
		Preload("Anuncio").
		Preload("Anuncio.Loja").
		Where("id_veiculo = ?", idVeiculo).
		Order("data DESC").
		Find(&historicos).Error

	if err != nil {
		return nil, err
	}

	return historicos, nil
}

// GetHistoricosByUsuario retorna o histórico de todos os veículos de um usuário
func GetHistoricosByUsuario(idUsuario uint) ([]models.HistoricoVeiculo, error) {
	var historicos []models.HistoricoVeiculo

	err := database.DB.
		Joins("JOIN veiculos ON historico_veiculos.id_veiculo = veiculos.id").
		Preload("Veiculo").
		Preload("Anuncio").
		Preload("Anuncio.Loja").
		Where("veiculos.id_usuario = ?", idUsuario).
		Order("historico_veiculos.data DESC").
		Find(&historicos).Error

	if err != nil {
		return nil, err
	}

	return historicos, nil
}

// CreateVeiculo cria um novo veículo
func CreateVeiculo(req json.VeiculoRequest) (*models.Veiculo, error) {
	veiculo := models.Veiculo{
		Modelo:    req.Modelo,
		Ano:       req.Ano,
		Cor:       req.Cor,
		Placa:     req.Placa,
		IDUsuario: req.IDUsuario,
		Ativo:     true,
	}

	err := database.DB.Create(&veiculo).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o veículo com os relacionamentos
	return GetVeiculoByID(veiculo.ID)
}

// GetAllVeiculos retorna todos os veículos ativos (não excluídos)
func GetAllVeiculos() ([]models.Veiculo, error) {
	var veiculos []models.Veiculo
	err := database.DB.
		Preload("Usuario").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}

// UpdateVeiculo atualiza um veículo existente
func UpdateVeiculo(id uint, req json.VeiculoRequest) (*models.Veiculo, error) {
	// Verifica se o veículo existe e não foi excluído
	veiculo, err := GetVeiculoByID(id)
	if err != nil {
		return nil, errors.New("veículo não encontrado")
	}

	// Atualiza os campos
	veiculo.Modelo = req.Modelo
	veiculo.Ano = req.Ano
	veiculo.Cor = req.Cor
	veiculo.Placa = req.Placa
	veiculo.IDUsuario = req.IDUsuario

	err = database.DB.Save(&veiculo).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o veículo com os relacionamentos
	return GetVeiculoByID(id)
}

// SoftDeleteVeiculo realiza soft delete do veículo (marca como excluído)
func SoftDeleteVeiculo(id uint) error {
	// Verifica se o veículo existe e não foi excluído
	_, err := GetVeiculoByID(id)
	if err != nil {
		return errors.New("veículo não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Veiculo{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreVeiculo restaura um veículo que foi soft deleted
func RestoreVeiculo(id uint) error {
	var veiculo models.Veiculo
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&veiculo).Error
	if err != nil {
		return errors.New("veículo não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Veiculo{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
