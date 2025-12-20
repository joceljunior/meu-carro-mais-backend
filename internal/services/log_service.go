package services

import (
	"encoding/json"
	"meu-carro-mais/internal/database/datasource"
	"meu-carro-mais/internal/database/models"
	jsonmodels "meu-carro-mais/internal/handlers/json"
	"time"
)

// LogAction registra uma ação do usuário no sistema
func LogAction(logData LogData) error {
	log := models.Log{
		IDUsuario:  logData.IDUsuario,
		TipoAcao:   logData.TipoAcao,
		Entidade:   logData.Entidade,
		IDEntidade: logData.IDEntidade,
		Descricao:  logData.Descricao,
		IP:         logData.IP,
		UserAgent:  logData.UserAgent,
		MetodoHTTP: logData.MetodoHTTP,
		Endpoint:   logData.Endpoint,
		StatusHTTP: logData.StatusHTTP,
		DataAcao:   time.Now(),
	}

	// Converte dados antigos para JSONB se fornecido
	if logData.DadosAntigos != nil {
		log.DadosAntigos = logData.DadosAntigos
	}

	// Converte dados novos para JSONB se fornecido
	if logData.DadosNovos != nil {
		log.DadosNovos = logData.DadosNovos
	}

	return datasource.CreateLog(log)
}

// LogData contém os dados para criar um log
type LogData struct {
	IDUsuario   *uint
	TipoAcao    string
	Entidade    string
	IDEntidade  *uint
	Descricao   string
	DadosAntigos models.JSONB
	DadosNovos   models.JSONB
	IP          string
	UserAgent   string
	MetodoHTTP  string
	Endpoint    string
	StatusHTTP  int
}

// GetLogsByUsuarioID retorna todos os logs de um usuário
func GetLogsByUsuarioID(idUsuario uint) ([]models.Log, error) {
	return datasource.GetLogsByUsuarioID(idUsuario)
}

// GetLogsByEntidade retorna todos os logs de uma entidade
func GetLogsByEntidade(entidade string, idEntidade uint) ([]models.Log, error) {
	return datasource.GetLogsByEntidade(entidade, idEntidade)
}

// GetLogsByTipoAcao retorna todos os logs de um tipo de ação
func GetLogsByTipoAcao(tipoAcao string) ([]models.Log, error) {
	return datasource.GetLogsByTipoAcao(tipoAcao)
}

// GetAllLogs retorna todos os logs com paginação
func GetAllLogs(limit, offset int) ([]models.Log, int64, error) {
	return datasource.GetAllLogs(limit, offset)
}

// GetLogByID retorna um log específico
func GetLogByID(id uint) (*models.Log, error) {
	return datasource.GetLogByID(id)
}

// ToJSONB converte uma interface para JSONB
func ToJSONB(data interface{}) models.JSONB {
	if data == nil {
		return nil
	}
	
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	
	var jsonb models.JSONB
	if err := json.Unmarshal(jsonBytes, &jsonb); err != nil {
		return nil
	}
	
	return jsonb
}

// ConvertLogsToResponse converte modelos de Log para LogResponse
func ConvertLogsToResponse(logs []models.Log) []jsonmodels.LogResponse {
	responses := make([]jsonmodels.LogResponse, len(logs))
	for i, log := range logs {
		responses[i] = ConvertLogToResponse(log)
	}
	return responses
}

// ConvertLogToResponse converte um modelo de Log para LogResponse
func ConvertLogToResponse(log models.Log) jsonmodels.LogResponse {
	response := jsonmodels.LogResponse{
		ID:              log.ID,
		TipoAcao:        log.TipoAcao,
		Entidade:        log.Entidade,
		Descricao:       log.Descricao,
		IP:              log.IP,
		UserAgent:       log.UserAgent,
		MetodoHTTP:      log.MetodoHTTP,
		Endpoint:        log.Endpoint,
		StatusHTTP:      log.StatusHTTP,
		DataAcao:        log.DataAcao,
		DataCadastro:    log.DataCadastro,
		DataAtualizacao: log.DataAtualizacao,
	}

	if log.IDUsuario != nil {
		response.IDUsuario = log.IDUsuario
	}
	if log.IDEntidade != nil {
		response.IDEntidade = log.IDEntidade
	}
	if log.DadosAntigos != nil {
		response.DadosAntigos = log.DadosAntigos
	}
	if log.DadosNovos != nil {
		response.DadosNovos = log.DadosNovos
	}
	if log.Usuario != nil {
		// Converte usuário para UserResponse se necessário
		response.Usuario = &jsonmodels.UserResponse{
			ID:    log.Usuario.ID,
			Nome:  log.Usuario.Nome,
			Email: log.Usuario.Email,
		}
	}

	return response
}
