package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateFoto cria uma nova foto
func CreateFoto(req json.FotoRequest) (*models.Foto, error) {
	// Validação: apenas um dos IDs deve ser preenchido
	count := 0
	if req.IDVeiculo != nil {
		count++
	}
	if req.IDVeiculoLoja != nil {
		count++
	}
	if req.IDProduto != nil {
		count++
	}
	if req.IDServico != nil {
		count++
	}
	if req.IDLoja != nil {
		count++
	}

	if count != 1 {
		return nil, errors.New("deve ser informado exatamente um ID: veiculo, veiculo_loja, produto, servico ou loja")
	}

	foto := models.Foto{
		IDVeiculo:     req.IDVeiculo,
		IDVeiculoLoja: req.IDVeiculoLoja,
		IDProduto:     req.IDProduto,
		IDServico:     req.IDServico,
		IDLoja:        req.IDLoja,
		TipoEntidade:  req.TipoEntidade,
		URL:           req.URL,
		NomeArquivo:   req.NomeArquivo,
		Tamanho:       req.Tamanho,
		TipoMime:      req.TipoMime,
		Principal:     req.Principal,
		Ordem:         req.Ordem,
	}

	// Se for a foto principal, remove a flag principal das outras fotos da mesma entidade
	if req.Principal {
		err := removePrincipalFlag(req)
		if err != nil {
			return nil, err
		}
	}

	err := database.DB.Create(&foto).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a foto com os relacionamentos
	return GetFotoByID(foto.ID)
}

// removePrincipalFlag remove a flag principal das outras fotos da mesma entidade
func removePrincipalFlag(req json.FotoRequest) error {
	var whereClause string
	var args []interface{}

	switch req.TipoEntidade {
	case "veiculo":
		whereClause = "id_veiculo = ? AND principal = true"
		args = []interface{}{*req.IDVeiculo}
	case "veiculo_loja":
		whereClause = "id_veiculo_loja = ? AND principal = true"
		args = []interface{}{*req.IDVeiculoLoja}
	case "produto":
		whereClause = "id_produto = ? AND principal = true"
		args = []interface{}{*req.IDProduto}
	case "servico":
		whereClause = "id_servico = ? AND principal = true"
		args = []interface{}{*req.IDServico}
	case "loja":
		whereClause = "id_loja = ? AND principal = true"
		args = []interface{}{*req.IDLoja}
	}

	return database.DB.Model(&models.Foto{}).
		Where(whereClause, args...).
		Update("principal", false).Error
}

// GetFotoByID busca foto por ID (apenas não excluídas)
func GetFotoByID(id uint) (*models.Foto, error) {
	var foto models.Foto
	err := database.DB.
		Preload("Veiculo").
		Preload("VeiculoLoja").
		Preload("Produto").
		Preload("Servico").
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&foto).Error
	if err != nil {
		return nil, err
	}
	return &foto, nil
}

// GetAllFotos retorna todas as fotos ativas (não excluídas)
func GetAllFotos() ([]models.Foto, error) {
	var fotos []models.Foto
	err := database.DB.
		Preload("Veiculo").
		Preload("VeiculoLoja").
		Preload("Produto").
		Preload("Servico").
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("data_exclusao IS NULL").
		Order("data_upload DESC").
		Find(&fotos).Error
	if err != nil {
		return nil, err
	}
	return fotos, nil
}

// GetFotosByVeiculoID retorna todas as fotos de um veículo específico
func GetFotosByVeiculoID(idVeiculo uint) ([]models.Foto, error) {
	var fotos []models.Foto
	err := database.DB.
		Preload("Veiculo").
		Where("id_veiculo = ? AND data_exclusao IS NULL", idVeiculo).
		Order("principal DESC, ordem ASC, data_upload ASC").
		Find(&fotos).Error
	if err != nil {
		return nil, err
	}
	return fotos, nil
}

// GetFotosByVeiculoLojaID retorna todas as fotos de um veículo de loja específico
func GetFotosByVeiculoLojaID(idVeiculoLoja uint) ([]models.Foto, error) {
	var fotos []models.Foto
	err := database.DB.
		Preload("VeiculoLoja").
		Where("id_veiculo_loja = ? AND data_exclusao IS NULL", idVeiculoLoja).
		Order("principal DESC, ordem ASC, data_upload ASC").
		Find(&fotos).Error
	if err != nil {
		return nil, err
	}
	return fotos, nil
}

// GetFotosByProdutoID retorna todas as fotos de um produto específico
func GetFotosByProdutoID(idProduto uint) ([]models.Foto, error) {
	var fotos []models.Foto
	err := database.DB.
		Preload("Produto").
		Where("id_produto = ? AND data_exclusao IS NULL", idProduto).
		Order("principal DESC, ordem ASC, data_upload ASC").
		Find(&fotos).Error
	if err != nil {
		return nil, err
	}
	return fotos, nil
}

// GetFotosByServicoID retorna todas as fotos de um serviço específico
func GetFotosByServicoID(idServico uint) ([]models.Foto, error) {
	var fotos []models.Foto
	err := database.DB.
		Preload("Servico").
		Where("id_servico = ? AND data_exclusao IS NULL", idServico).
		Order("principal DESC, ordem ASC, data_upload ASC").
		Find(&fotos).Error
	if err != nil {
		return nil, err
	}
	return fotos, nil
}

// GetFotosByLojaID retorna todas as fotos de uma loja específica
func GetFotosByLojaID(idLoja uint) ([]models.Foto, error) {
	var fotos []models.Foto
	err := database.DB.
		Preload("Loja").
		Preload("Loja.Categoria").
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Order("principal DESC, ordem ASC, data_upload ASC").
		Find(&fotos).Error
	if err != nil {
		return nil, err
	}
	return fotos, nil
}

// GetFotoPrincipalByEntidade retorna a foto principal de uma entidade
func GetFotoPrincipalByEntidade(tipoEntidade string, idEntidade uint) (*models.Foto, error) {
	var foto models.Foto
	var whereClause string
	var args []interface{}

	switch tipoEntidade {
	case "veiculo":
		whereClause = "id_veiculo = ? AND principal = true AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "veiculo_loja":
		whereClause = "id_veiculo_loja = ? AND principal = true AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "produto":
		whereClause = "id_produto = ? AND principal = true AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "servico":
		whereClause = "id_servico = ? AND principal = true AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "loja":
		whereClause = "id_loja = ? AND principal = true AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	default:
		return nil, errors.New("tipo de entidade inválido")
	}

	err := database.DB.
		Preload("Veiculo").
		Preload("VeiculoLoja").
		Preload("Produto").
		Preload("Servico").
		Preload("Loja").
		Preload("Loja.Categoria").
		Where(whereClause, args...).
		First(&foto).Error
	if err != nil {
		return nil, err
	}
	return &foto, nil
}

// UpdateFoto atualiza uma foto existente
func UpdateFoto(id uint, req json.FotoRequest) (*models.Foto, error) {
	// Verifica se a foto existe e não foi excluída
	foto, err := GetFotoByID(id)
	if err != nil {
		return nil, errors.New("foto não encontrada")
	}

	// Atualiza os campos
	foto.URL = req.URL
	foto.NomeArquivo = req.NomeArquivo
	foto.Tamanho = req.Tamanho
	foto.TipoMime = req.TipoMime
	foto.Principal = req.Principal
	foto.Ordem = req.Ordem

	// Se for a foto principal, remove a flag principal das outras fotos da mesma entidade
	if req.Principal {
		err := removePrincipalFlag(req)
		if err != nil {
			return nil, err
		}
	}

	err = database.DB.Save(&foto).Error
	if err != nil {
		return nil, err
	}

	// Recarrega a foto com os relacionamentos
	return GetFotoByID(id)
}

// SetFotoPrincipal define uma foto como principal
func SetFotoPrincipal(id uint) error {
	// Verifica se a foto existe e não foi excluída
	foto, err := GetFotoByID(id)
	if err != nil {
		return errors.New("foto não encontrada")
	}

	// Remove a flag principal das outras fotos da mesma entidade
	req := json.FotoRequest{
		TipoEntidade: foto.TipoEntidade,
	}

	switch foto.TipoEntidade {
	case "veiculo":
		req.IDVeiculo = foto.IDVeiculo
	case "veiculo_loja":
		req.IDVeiculoLoja = foto.IDVeiculoLoja
	case "produto":
		req.IDProduto = foto.IDProduto
	case "servico":
		req.IDServico = foto.IDServico
	case "loja":
		req.IDLoja = foto.IDLoja
	}

	err = removePrincipalFlag(req)
	if err != nil {
		return err
	}

	// Define esta foto como principal
	err = database.DB.Model(&models.Foto{}).
		Where("id = ?", id).
		Update("principal", true).Error
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleteFoto realiza soft delete da foto (marca como excluída)
func SoftDeleteFoto(id uint) error {
	// Verifica se a foto existe e não foi excluída
	_, err := GetFotoByID(id)
	if err != nil {
		return errors.New("foto não encontrada")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Foto{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreFoto restaura uma foto que foi soft deleted
func RestoreFoto(id uint) error {
	var foto models.Foto
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&foto).Error
	if err != nil {
		return errors.New("foto não encontrada ou não foi excluída")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Foto{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}
