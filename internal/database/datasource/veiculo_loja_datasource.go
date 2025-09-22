package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateVeiculoLoja cria um novo veículo de loja
func CreateVeiculoLoja(req json.VeiculoLojaRequest) (*models.VeiculoLoja, error) {
	veiculo := models.VeiculoLoja{
		Modelo: req.Modelo,
		Ano:    req.Ano,
		Cor:    req.Cor,
		Placa:  req.Placa,
		IDLoja: req.IDLoja,
		Ativo:  true,
	}

	err := database.DB.Create(&veiculo).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o veículo com os relacionamentos
	return GetVeiculoLojaByID(veiculo.ID)
}

// GetVeiculoLojaByID busca veículo de loja por ID (apenas veículos não excluídos)
func GetVeiculoLojaByID(id uint) (*models.VeiculoLoja, error) {
	var veiculo models.VeiculoLoja
	err := database.DB.
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("id = ? AND ativo = ? AND data_exclusao IS NULL", id, true).
		First(&veiculo).Error
	if err != nil {
		return nil, err
	}
	return &veiculo, nil
}

// GetAllVeiculosLoja retorna todos os veículos de loja ativos (não excluídos)
func GetAllVeiculosLoja() ([]models.VeiculoLoja, error) {
	var veiculos []models.VeiculoLoja
	err := database.DB.
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}

// GetVeiculosLojaByLojaID retorna todos os veículos de uma loja específica
func GetVeiculosLojaByLojaID(idLoja uint) ([]models.VeiculoLoja, error) {
	var veiculos []models.VeiculoLoja
	err := database.DB.
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("id_loja = ? AND ativo = ? AND data_exclusao IS NULL", idLoja, true).
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}

// UpdateVeiculoLoja atualiza um veículo de loja existente
func UpdateVeiculoLoja(id uint, req json.VeiculoLojaRequest) (*models.VeiculoLoja, error) {
	// Verifica se o veículo existe e não foi excluído
	veiculo, err := GetVeiculoLojaByID(id)
	if err != nil {
		return nil, errors.New("veículo não encontrado")
	}

	// Atualiza os campos
	veiculo.Modelo = req.Modelo
	veiculo.Ano = req.Ano
	veiculo.Cor = req.Cor
	veiculo.Placa = req.Placa
	veiculo.IDLoja = req.IDLoja

	err = database.DB.Save(&veiculo).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o veículo com os relacionamentos
	return GetVeiculoLojaByID(id)
}

// SoftDeleteVeiculoLoja realiza soft delete do veículo de loja (marca como excluído)
func SoftDeleteVeiculoLoja(id uint) error {
	// Verifica se o veículo existe e não foi excluído
	_, err := GetVeiculoLojaByID(id)
	if err != nil {
		return errors.New("veículo não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.VeiculoLoja{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreVeiculoLoja restaura um veículo de loja que foi soft deleted
func RestoreVeiculoLoja(id uint) error {
	var veiculo models.VeiculoLoja
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&veiculo).Error
	if err != nil {
		return errors.New("veículo não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.VeiculoLoja{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
