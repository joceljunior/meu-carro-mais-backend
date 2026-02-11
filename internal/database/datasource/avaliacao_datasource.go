package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateAvaliacao cria uma nova avaliação
func CreateAvaliacao(req json.AvaliacaoRequest) (*models.Avaliacao, error) {
	// Verifica duplicação baseado no tipo de avaliação
	var existingAvaliacao models.Avaliacao
	query := database.DB.Where("id_usuario = ? AND data_exclusao IS NULL", req.IDUsuario)

	// Verifica se já existe avaliação para o mesmo item
	if req.IDLoja != nil {
		if err := query.Where("id_loja = ?", *req.IDLoja).First(&existingAvaliacao).Error; err == nil {
			return nil, errors.New("usuário já avaliou esta loja")
		}
	}
	if req.IDServico != nil {
		if err := database.DB.Where("id_usuario = ? AND id_servico = ? AND data_exclusao IS NULL", req.IDUsuario, *req.IDServico).First(&existingAvaliacao).Error; err == nil {
			return nil, errors.New("usuário já avaliou este serviço")
		}
	}
	if req.IDProduto != nil {
		if err := database.DB.Where("id_usuario = ? AND id_produto = ? AND data_exclusao IS NULL", req.IDUsuario, *req.IDProduto).First(&existingAvaliacao).Error; err == nil {
			return nil, errors.New("usuário já avaliou este produto")
		}
	}

	avaliacao := models.Avaliacao{
		IDUsuario:  req.IDUsuario,
		IDLoja:     req.IDLoja,
		IDServico:  req.IDServico,
		IDProduto:  req.IDProduto,
		IDCupom:  req.IDCupom,
		Nota:       req.Nota,
		Comentario: req.Comentario,
	}

	err := database.DB.Create(&avaliacao).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a avaliação com os relacionamentos
	return GetAvaliacaoByID(avaliacao.ID)
}

// GetAvaliacaoByID busca avaliação por ID (apenas não excluídas)
func GetAvaliacaoByID(id uint) (*models.Avaliacao, error) {
	var avaliacao models.Avaliacao
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Preload("Servico").
		Preload("Produto").
		Preload("Cupom").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&avaliacao).Error
	if err != nil {
		return nil, err
	}
	return &avaliacao, nil
}

// GetAllAvaliacoes retorna todas as avaliações ativas (não excluídas)
func GetAllAvaliacoes() ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Preload("Servico").
		Preload("Produto").
		Preload("Cupom").
		Where("data_exclusao IS NULL").
		Order("data_avaliacao DESC").
		Find(&avaliacoes).Error
	if err != nil {
		return nil, err
	}
	return avaliacoes, nil
}

// GetAvaliacoesByLojaID retorna todas as avaliações de uma loja específica
func GetAvaliacoesByLojaID(idLoja uint) ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Preload("Servico").
		Preload("Produto").
		Preload("Cupom").
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Order("data_avaliacao DESC").
		Find(&avaliacoes).Error
	if err != nil {
		return nil, err
	}
	return avaliacoes, nil
}

// GetAvaliacoesByUsuarioID retorna todas as avaliações de um usuário específico
func GetAvaliacoesByUsuarioID(idUsuario uint) ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Preload("Servico").
		Preload("Produto").
		Preload("Cupom").
		Where("id_usuario = ? AND data_exclusao IS NULL", idUsuario).
		Order("data_avaliacao DESC").
		Find(&avaliacoes).Error
	if err != nil {
		return nil, err
	}
	return avaliacoes, nil
}

// GetAvaliacoesByServicoID retorna todas as avaliações de um serviço específico
func GetAvaliacoesByServicoID(idServico uint) ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Preload("Servico").
		Preload("Produto").
		Preload("Cupom").
		Where("id_servico = ? AND data_exclusao IS NULL", idServico).
		Order("data_avaliacao DESC").
		Find(&avaliacoes).Error
	if err != nil {
		return nil, err
	}
	return avaliacoes, nil
}

// GetAvaliacoesByProdutoID retorna todas as avaliações de um produto específico
func GetAvaliacoesByProdutoID(idProduto uint) ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Preload("Servico").
		Preload("Produto").
		Preload("Cupom").
		Where("id_produto = ? AND data_exclusao IS NULL", idProduto).
		Order("data_avaliacao DESC").
		Find(&avaliacoes).Error
	if err != nil {
		return nil, err
	}
	return avaliacoes, nil
}

// GetAvaliacaoByUsuarioELoja busca uma avaliação específica de um usuário para uma loja
func GetAvaliacaoByUsuarioELoja(idUsuario uint, idLoja uint) (*models.Avaliacao, error) {
	var avaliacao models.Avaliacao
	err := database.DB.
		Preload("Usuario").
		Preload("Loja").
		Preload("Servico").
		Preload("Produto").
		Preload("Cupom").
		Where("id_usuario = ? AND id_loja = ? AND data_exclusao IS NULL", idUsuario, idLoja).
		First(&avaliacao).Error
	if err != nil {
		return nil, err
	}
	return &avaliacao, nil
}

// GetAvaliacaoEstatisticasByLojaID retorna estatísticas das avaliações de uma loja
func GetAvaliacaoEstatisticasByLojaID(idLoja uint) (*json.AvaliacaoEstatisticasResponse, error) {
	var total int64
	var media float64
	var nota1, nota2, nota3, nota4, nota5 int64

	// Conta total de avaliações
	err := database.DB.Model(&models.Avaliacao{}).
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Count(&total).Error
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return &json.AvaliacaoEstatisticasResponse{
			TotalAvaliacoes: 0,
			MediaNota:       0,
			Nota1:           0,
			Nota2:           0,
			Nota3:           0,
			Nota4:           0,
			Nota5:           0,
		}, nil
	}

	// Calcula média
	err = database.DB.Model(&models.Avaliacao{}).
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Select("AVG(nota)").Scan(&media).Error
	if err != nil {
		return nil, err
	}

	// Conta por nota
	database.DB.Model(&models.Avaliacao{}).
		Where("id_loja = ? AND nota = 1 AND data_exclusao IS NULL", idLoja).
		Count(&nota1)
	database.DB.Model(&models.Avaliacao{}).
		Where("id_loja = ? AND nota = 2 AND data_exclusao IS NULL", idLoja).
		Count(&nota2)
	database.DB.Model(&models.Avaliacao{}).
		Where("id_loja = ? AND nota = 3 AND data_exclusao IS NULL", idLoja).
		Count(&nota3)
	database.DB.Model(&models.Avaliacao{}).
		Where("id_loja = ? AND nota = 4 AND data_exclusao IS NULL", idLoja).
		Count(&nota4)
	database.DB.Model(&models.Avaliacao{}).
		Where("id_loja = ? AND nota = 5 AND data_exclusao IS NULL", idLoja).
		Count(&nota5)

	return &json.AvaliacaoEstatisticasResponse{
		TotalAvaliacoes: int(total),
		MediaNota:       media,
		Nota1:           int(nota1),
		Nota2:           int(nota2),
		Nota3:           int(nota3),
		Nota4:           int(nota4),
		Nota5:           int(nota5),
	}, nil
}

// UpdateAvaliacao atualiza uma avaliação existente
func UpdateAvaliacao(id uint, req json.AvaliacaoRequest) (*models.Avaliacao, error) {
	// Verifica se a avaliação existe e não foi excluída
	avaliacao, err := GetAvaliacaoByID(id)
	if err != nil {
		return nil, errors.New("avaliação não encontrada")
	}

	// Atualiza os campos
	avaliacao.IDLoja = req.IDLoja
	avaliacao.IDServico = req.IDServico
	avaliacao.IDProduto = req.IDProduto
	avaliacao.IDCupom = req.IDCupom
	avaliacao.Nota = req.Nota
	avaliacao.Comentario = req.Comentario

	err = database.DB.Save(&avaliacao).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a avaliação com os relacionamentos
	return GetAvaliacaoByID(id)
}

// SoftDeleteAvaliacao realiza soft delete da avaliação (marca como excluída)
func SoftDeleteAvaliacao(id uint) error {
	// Verifica se a avaliação existe e não foi excluída
	_, err := GetAvaliacaoByID(id)
	if err != nil {
		return errors.New("avaliação não encontrada")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Avaliacao{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreAvaliacao restaura uma avaliação que foi soft deleted
func RestoreAvaliacao(id uint) error {
	var avaliacao models.Avaliacao
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&avaliacao).Error
	if err != nil {
		return errors.New("avaliação não encontrada ou não foi excluída")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Avaliacao{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
