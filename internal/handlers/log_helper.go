package handlers

import (
	"encoding/json"
	"fmt"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogAction registra uma ação do usuário de forma assíncrona
// Esta função não bloqueia a resposta ao usuário
func LogAction(c *gin.Context, tipoAcao, entidade string, idEntidade *uint, descricao string, dadosAntigos, dadosNovos interface{}) {
	// Executa em goroutine para não bloquear a resposta
	go func() {
		// Extrai informações do contexto
		idUsuario := extractUserID(c)
		ip := c.ClientIP()
		userAgent := c.GetHeader("User-Agent")
		metodoHTTP := c.Request.Method
		endpoint := c.FullPath()
		if endpoint == "" {
			endpoint = c.Request.URL.Path
		}
		statusHTTP := c.Writer.Status()
		if statusHTTP == 0 {
			statusHTTP = http.StatusOK // Default se ainda não foi definido
		}

		// Converte dados para JSONB
		var dadosAntigosJSONB, dadosNovosJSONB models.JSONB
		if dadosAntigos != nil {
			dadosAntigosJSONB = services.ToJSONB(dadosAntigos)
		}
		if dadosNovos != nil {
			dadosNovosJSONB = services.ToJSONB(dadosNovos)
		}

		// Registra o log
		_ = services.LogAction(services.LogData{
			IDUsuario:   idUsuario,
			TipoAcao:    tipoAcao,
			Entidade:    entidade,
			IDEntidade:  idEntidade,
			Descricao:   descricao,
			DadosAntigos: dadosAntigosJSONB,
			DadosNovos:   dadosNovosJSONB,
			IP:          ip,
			UserAgent:   userAgent,
			MetodoHTTP:  metodoHTTP,
			Endpoint:    endpoint,
			StatusHTTP:  statusHTTP,
		})
	}()
}

// LogActionSync registra uma ação de forma síncrona (bloqueia até completar)
// Use apenas quando necessário garantir que o log foi salvo antes de continuar
func LogActionSync(c *gin.Context, tipoAcao, entidade string, idEntidade *uint, descricao string, dadosAntigos, dadosNovos interface{}) error {
	idUsuario := extractUserID(c)
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	metodoHTTP := c.Request.Method
	endpoint := c.FullPath()
	if endpoint == "" {
		endpoint = c.Request.URL.Path
	}
	statusHTTP := c.Writer.Status()
	if statusHTTP == 0 {
		statusHTTP = http.StatusOK
	}

	var dadosAntigosJSONB, dadosNovosJSONB models.JSONB
	if dadosAntigos != nil {
		dadosAntigosJSONB = services.ToJSONB(dadosAntigos)
	}
	if dadosNovos != nil {
		dadosNovosJSONB = services.ToJSONB(dadosNovos)
	}

	return services.LogAction(services.LogData{
		IDUsuario:   idUsuario,
		TipoAcao:    tipoAcao,
		Entidade:    entidade,
		IDEntidade:  idEntidade,
		Descricao:   descricao,
		DadosAntigos: dadosAntigosJSONB,
		DadosNovos:   dadosNovosJSONB,
		IP:          ip,
		UserAgent:   userAgent,
		MetodoHTTP:  metodoHTTP,
		Endpoint:    endpoint,
		StatusHTTP:  statusHTTP,
	})
}

// extractUserID extrai o ID do usuário do contexto
// Tenta várias formas de obter o ID do usuário:
// 1. Do parâmetro da URL (id_usuario, id)
// 2. Do body da requisição (se disponível)
// 3. Do header Authorization (se implementado)
func extractUserID(c *gin.Context) *uint {
	// Tenta obter do parâmetro id_usuario
	if idStr := c.Param("id_usuario"); idStr != "" {
		if id, err := parseUint(idStr); err == nil {
			return &id
		}
	}

	// Tenta obter do parâmetro id
	if idStr := c.Param("id"); idStr != "" {
		if id, err := parseUint(idStr); err == nil {
			return &id
		}
	}

	// Tenta obter do query parameter user_id
	if idStr := c.Query("user_id"); idStr != "" {
		if id, err := parseUint(idStr); err == nil {
			return &id
		}
	}

	// Tenta obter do body JSON (para alguns endpoints específicos)
	// Usa GetRawData para não consumir o body
	if c.Request.Body != nil {
		bodyBytes, _ := c.GetRawData()
		if len(bodyBytes) > 0 {
			var body map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &body); err == nil {
				if idUsuario, ok := body["id_usuario"].(float64); ok {
					id := uint(idUsuario)
					return &id
				}
			}
		}
	}

	// Se não conseguir obter, retorna nil (log sem usuário)
	return nil
}

// parseUint converte string para uint
func parseUint(s string) (uint, error) {
	var result uint64
	_, err := fmt.Sscanf(s, "%d", &result)
	return uint(result), err
}

