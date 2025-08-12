package datasource

import (
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
)

// GetVeiculosByUsuario retorna todos os veículos de um usuário
func GetVeiculosByUsuario(idUsuario uint) ([]models.Veiculo, error) {
	var veiculos []models.Veiculo
	
	err := database.DB.
		Where("id_usuario = ? AND ativo = ?", idUsuario, true).
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
		Where("id = ? AND ativo = ?", id, true).
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
