package datasource

import (
	"errors"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"time"
)

// CreateUpload cria um novo upload
func CreateUpload(req json.UploadRequest) (*models.Upload, error) {
	// Validação: apenas um dos IDs deve ser preenchido
	count := 0
	if req.IDUsuario != nil {
		count++
	}
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
		return nil, errors.New("deve ser informado exatamente um ID: usuario, veiculo, veiculo_loja, produto, servico ou loja")
	}

	upload := models.Upload{
		IDUsuario:     req.IDUsuario,
		IDVeiculo:     req.IDVeiculo,
		IDVeiculoLoja: req.IDVeiculoLoja,
		IDProduto:     req.IDProduto,
		IDServico:     req.IDServico,
		IDLoja:        req.IDLoja,
		TipoEntidade:  req.TipoEntidade,
		Tipo:          req.Tipo,
		URL:           req.URL,
		NomeArquivo:   req.NomeArquivo,
		Tamanho:       req.Tamanho,
		TipoMime:      req.TipoMime,
		Principal:     req.Principal,
		Ordem:         req.Ordem,
	}

	// Se for imagem principal, remove a flag principal das outras imagens da mesma entidade
	if req.Principal && req.Tipo == "Imagem" {
		err := removeUploadPrincipalFlag(req)
		if err != nil {
			return nil, err
		}
	}

	err := database.DB.Create(&upload).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o upload com os relacionamentos
	return GetUploadByID(upload.ID)
}

// removeUploadPrincipalFlag remove a flag principal das outras imagens da mesma entidade
func removeUploadPrincipalFlag(req json.UploadRequest) error {
	var whereClause string
	var args []interface{}

	switch req.TipoEntidade {
	case "usuario":
		whereClause = "id_usuario = ? AND principal = true AND tipo = 'Imagem'"
		args = []interface{}{*req.IDUsuario}
	case "veiculo":
		whereClause = "id_veiculo = ? AND principal = true AND tipo = 'Imagem'"
		args = []interface{}{*req.IDVeiculo}
	case "veiculo_loja":
		whereClause = "id_veiculo_loja = ? AND principal = true AND tipo = 'Imagem'"
		args = []interface{}{*req.IDVeiculoLoja}
	case "produto":
		whereClause = "id_produto = ? AND principal = true AND tipo = 'Imagem'"
		args = []interface{}{*req.IDProduto}
	case "servico":
		whereClause = "id_servico = ? AND principal = true AND tipo = 'Imagem'"
		args = []interface{}{*req.IDServico}
	case "loja":
		whereClause = "id_loja = ? AND principal = true AND tipo = 'Imagem'"
		args = []interface{}{*req.IDLoja}
	}

	return database.DB.Model(&models.Upload{}).
		Where(whereClause, args...).
		Update("principal", false).Error
}

// GetUploadByID busca upload por ID (apenas não excluídos)
func GetUploadByID(id uint) (*models.Upload, error) {
	var upload models.Upload
	err := database.DB.
		Preload("Usuario").
		Preload("Veiculo").
		Preload("VeiculoLoja").
		Preload("Produto").
		Preload("Servico").
		Preload("Loja").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&upload).Error
	if err != nil {
		return nil, err
	}
	return &upload, nil
}

// GetAllUploads retorna todos os uploads ativos (não excluídos)
func GetAllUploads() ([]models.Upload, error) {
	var uploads []models.Upload
	err := database.DB.
		Preload("Usuario").
		Preload("Veiculo").
		Preload("VeiculoLoja").
		Preload("Produto").
		Preload("Servico").
		Preload("Loja").
		Where("data_exclusao IS NULL").
		Order("data_upload DESC").
		Find(&uploads).Error
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetAllUploadsByTipo retorna todos os uploads ativos filtrados por tipo (Imagem ou Documento)
func GetAllUploadsByTipo(tipo string) ([]models.Upload, error) {
	var uploads []models.Upload
	err := database.DB.
		Preload("Usuario").
		Preload("Veiculo").
		Preload("VeiculoLoja").
		Preload("Produto").
		Preload("Servico").
		Preload("Loja").
		Where("tipo = ? AND data_exclusao IS NULL", tipo).
		Order("data_upload DESC").
		Find(&uploads).Error
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetUploadsByUsuarioID retorna todos os uploads de um usuário específico
func GetUploadsByUsuarioID(idUsuario uint) ([]models.Upload, error) {
	var uploads []models.Upload
	err := database.DB.
		Preload("Usuario").
		Where("id_usuario = ? AND data_exclusao IS NULL", idUsuario).
		Order("tipo ASC, principal DESC, ordem ASC, data_upload ASC").
		Find(&uploads).Error
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetUploadsByVeiculoID retorna todos os uploads de um veículo específico
func GetUploadsByVeiculoID(idVeiculo uint) ([]models.Upload, error) {
	var uploads []models.Upload
	err := database.DB.
		Preload("Veiculo").
		Where("id_veiculo = ? AND data_exclusao IS NULL", idVeiculo).
		Order("tipo ASC, principal DESC, ordem ASC, data_upload ASC").
		Find(&uploads).Error
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetUploadsByVeiculoLojaID retorna todos os uploads de um veículo de loja específico
func GetUploadsByVeiculoLojaID(idVeiculoLoja uint) ([]models.Upload, error) {
	var uploads []models.Upload
	err := database.DB.
		Preload("VeiculoLoja").
		Where("id_veiculo_loja = ? AND data_exclusao IS NULL", idVeiculoLoja).
		Order("tipo ASC, principal DESC, ordem ASC, data_upload ASC").
		Find(&uploads).Error
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetUploadsByProdutoID retorna todos os uploads de um produto específico
func GetUploadsByProdutoID(idProduto uint) ([]models.Upload, error) {
	var uploads []models.Upload
	err := database.DB.
		Preload("Produto").
		Where("id_produto = ? AND data_exclusao IS NULL", idProduto).
		Order("tipo ASC, principal DESC, ordem ASC, data_upload ASC").
		Find(&uploads).Error
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetUploadsByServicoID retorna todos os uploads de um serviço específico
func GetUploadsByServicoID(idServico uint) ([]models.Upload, error) {
	var uploads []models.Upload
	err := database.DB.
		Preload("Servico").
		Where("id_servico = ? AND data_exclusao IS NULL", idServico).
		Order("tipo ASC, principal DESC, ordem ASC, data_upload ASC").
		Find(&uploads).Error
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetUploadsByLojaID retorna todos os uploads de uma loja específica
func GetUploadsByLojaID(idLoja uint) ([]models.Upload, error) {
	var uploads []models.Upload
	err := database.DB.
		Preload("Loja").
		Where("id_loja = ? AND data_exclusao IS NULL", idLoja).
		Order("tipo ASC, principal DESC, ordem ASC, data_upload ASC").
		Find(&uploads).Error
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// GetUploadPrincipalByEntidade retorna o upload principal (imagem) de uma entidade
func GetUploadPrincipalByEntidade(tipoEntidade string, idEntidade uint) (*models.Upload, error) {
	var upload models.Upload
	var whereClause string
	var args []interface{}

	switch tipoEntidade {
	case "usuario":
		whereClause = "id_usuario = ? AND principal = true AND tipo = 'Imagem' AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "veiculo":
		whereClause = "id_veiculo = ? AND principal = true AND tipo = 'Imagem' AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "veiculo_loja":
		whereClause = "id_veiculo_loja = ? AND principal = true AND tipo = 'Imagem' AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "produto":
		whereClause = "id_produto = ? AND principal = true AND tipo = 'Imagem' AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "servico":
		whereClause = "id_servico = ? AND principal = true AND tipo = 'Imagem' AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	case "loja":
		whereClause = "id_loja = ? AND principal = true AND tipo = 'Imagem' AND data_exclusao IS NULL"
		args = []interface{}{idEntidade}
	default:
		return nil, errors.New("tipo de entidade inválido")
	}

	err := database.DB.
		Preload("Usuario").
		Preload("Veiculo").
		Preload("VeiculoLoja").
		Preload("Produto").
		Preload("Servico").
		Preload("Loja").
		Where(whereClause, args...).
		First(&upload).Error
	if err != nil {
		return nil, err
	}
	return &upload, nil
}

// UpdateUpload atualiza um upload existente
func UpdateUpload(id uint, req json.UploadRequest) (*models.Upload, error) {
	// Verifica se o upload existe e não foi excluído
	upload, err := GetUploadByID(id)
	if err != nil {
		return nil, errors.New("upload não encontrado")
	}

	// Atualiza os campos
	upload.URL = req.URL
	upload.NomeArquivo = req.NomeArquivo
	upload.Tamanho = req.Tamanho
	upload.TipoMime = req.TipoMime
	upload.Tipo = req.Tipo
	upload.Principal = req.Principal
	upload.Ordem = req.Ordem

	// Se for imagem principal, remove a flag principal das outras imagens da mesma entidade
	if req.Principal && req.Tipo == "Imagem" {
		err := removeUploadPrincipalFlag(req)
		if err != nil {
			return nil, err
		}
	}

	err = database.DB.Save(&upload).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o upload com os relacionamentos
	return GetUploadByID(id)
}

// SetUploadPrincipal define um upload como principal (apenas para imagens)
func SetUploadPrincipal(id uint) error {
	// Verifica se o upload existe e não foi excluído
	upload, err := GetUploadByID(id)
	if err != nil {
		return errors.New("upload não encontrado")
	}

	// Apenas imagens podem ser principais
	if upload.Tipo != "Imagem" {
		return errors.New("apenas imagens podem ser definidas como principais")
	}

	// Remove a flag principal das outras imagens da mesma entidade
	req := json.UploadRequest{
		TipoEntidade: upload.TipoEntidade,
		Tipo:         upload.Tipo,
	}

	switch upload.TipoEntidade {
	case "usuario":
		req.IDUsuario = upload.IDUsuario
	case "veiculo":
		req.IDVeiculo = upload.IDVeiculo
	case "veiculo_loja":
		req.IDVeiculoLoja = upload.IDVeiculoLoja
	case "produto":
		req.IDProduto = upload.IDProduto
	case "servico":
		req.IDServico = upload.IDServico
	case "loja":
		req.IDLoja = upload.IDLoja
	}

	err = removeUploadPrincipalFlag(req)
	if err != nil {
		return err
	}

	// Define este upload como principal
	err = database.DB.Model(&models.Upload{}).
		Where("id = ?", id).
		Update("principal", true).Error
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleteUpload realiza soft delete do upload (marca como excluído)
func SoftDeleteUpload(id uint) error {
	// Verifica se o upload existe e não foi excluído
	_, err := GetUploadByID(id)
	if err != nil {
		return errors.New("upload não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Upload{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreUpload restaura um upload que foi soft deleted
func RestoreUpload(id uint) error {
	var upload models.Upload
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&upload).Error
	if err != nil {
		return errors.New("upload não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Upload{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}

