package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// GetServicosByProximidade retorna lista de serviços ordenados por proximidade
func GetServicosByProximidade(latitude, longitude float64) ([]json.ServicoResponse, error) {
	var servicos []models.Servico

	err := database.DB.
		Preload("Loja").
		Where("data_exclusao IS NULL").
		Find(&servicos).Error

	if err != nil {
		return nil, err
	}

	var servicosResponse []json.ServicoResponse

	for _, servico := range servicos {
		// Calcula a distância usando a fórmula de Haversine
		distancia := calcularDistancia(latitude, longitude, servico.Loja.Latitude, servico.Loja.Longitude)

		servicoResp := json.ServicoResponse{
			ID:        servico.ID,
			Titulo:    servico.Titulo,
			Descricao: servico.Descricao,
			Preco:     servico.Preco,
			Imagem:    servico.Imagem,
			Destaque:  servico.Destaque,
			Distancia: distancia,
			Categoria: servico.Categoria,
			Rate:      servico.Loja.Rating,
		Loja: json.LojaFromModel(servico.Loja),
		}

		servicosResponse = append(servicosResponse, servicoResp)
	}

	// Ordena por distância (menor primeiro)
	for i := 0; i < len(servicosResponse)-1; i++ {
		for j := i + 1; j < len(servicosResponse); j++ {
			if servicosResponse[i].Distancia > servicosResponse[j].Distancia {
				servicosResponse[i], servicosResponse[j] = servicosResponse[j], servicosResponse[i]
			}
		}
	}

	return servicosResponse, nil
}

// CreateServico cria um novo serviço
func CreateServico(req json.ServicoRequest) (*models.Servico, error) {
	servico := models.Servico{
		Titulo:    req.Titulo,
		Descricao: req.Descricao,
		Preco:     req.Preco,
		Imagem:    req.Imagem,
		Destaque:  req.Destaque,
		Categoria: req.Categoria,
		IDLoja:    req.IDLoja,
	}

	err := database.DB.Create(&servico).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o serviço com os relacionamentos
	return GetServicoByID(servico.ID)
}

// GetServicoByID busca serviço por ID (apenas serviços não excluídos)
func GetServicoByID(id uint) (*models.Servico, error) {
	var servico models.Servico
	err := database.DB.
		Preload("Loja").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&servico).Error
	if err != nil {
		return nil, err
	}
	return &servico, nil
}

// GetAllServicos retorna todos os serviços ativos (não excluídos)
func GetAllServicos() ([]models.Servico, error) {
	var servicos []models.Servico
	err := database.DB.
		Preload("Loja").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&servicos).Error
	if err != nil {
		return nil, err
	}
	return servicos, nil
}

// GetServicosByLojaID retorna todos os serviços de uma loja específica
func GetServicosByLojaID(idLoja uint) ([]models.Servico, error) {
	var servicos []models.Servico
	err := database.DB.
		Preload("Loja").
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Order("data_cadastro DESC").
		Find(&servicos).Error
	if err != nil {
		return nil, err
	}
	return servicos, nil
}

// UpdateServico atualiza um serviço existente
func UpdateServico(id uint, req json.ServicoRequest) (*models.Servico, error) {
	// Verifica se o serviço existe e não foi excluído
	servico, err := GetServicoByID(id)
	if err != nil {
		return nil, errors.New("serviço não encontrado")
	}

	// Atualiza os campos
	servico.Titulo = req.Titulo
	servico.Descricao = req.Descricao
	servico.Preco = req.Preco
	servico.Imagem = req.Imagem
	servico.Destaque = req.Destaque
	servico.Categoria = req.Categoria
	servico.IDLoja = req.IDLoja

	err = database.DB.Save(&servico).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o serviço com os relacionamentos
	return GetServicoByID(id)
}

// SoftDeleteServico realiza soft delete do serviço (marca como excluído)
func SoftDeleteServico(id uint) error {
	// Verifica se o serviço existe e não foi excluído
	_, err := GetServicoByID(id)
	if err != nil {
		return errors.New("serviço não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Servico{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreServico restaura um serviço que foi soft deleted
func RestoreServico(id uint) error {
	var servico models.Servico
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&servico).Error
	if err != nil {
		return errors.New("serviço não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Servico{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
