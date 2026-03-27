package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"
)

// GetCuponsHandler godoc
// @Summary      Lista todos os cupons
// @Description  Retorna todos os cupons disponíveis com informações da loja e categoria, incluindo preço original, preço com desconto, porcentagem de desconto e avaliação da loja
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.CuponsResponse "Lista de cupons"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /cupons [get]
func GetCuponsHandler(c *gin.Context) {
	resp, err := services.GetCupons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateCupomHandler godoc
// @Summary      Criação do cupom completo
// @Description  Cria um novo cupom com todos os dados fornecidos, incluindo porcentagem de desconto e preço com desconto
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        request body json.CupomRequest true "Dados completos do cupom"
// @Success      201  {object}  json.CupomResponse "Cupom criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /cupons [post]
func CreateCupomHandler(c *gin.Context) {
	var req json.CupomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	// Valida que IDLoja é obrigatório para cupons de produto ou serviço
	if req.TipoCupom == "produto" || req.TipoCupom == "servico" {
		if req.IDLoja == nil || *req.IDLoja == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id_loja é obrigatório para cupons de produto ou serviço",
			})
			return
		}
	}

	// Para cupons de veículo do usuário, se IDLoja for 0, converte para nil
	if req.TipoCupom == "veiculo" && req.IDLoja != nil && *req.IDLoja == 0 {
		req.IDLoja = nil
	}

	resp, err := services.CreateCupom(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Registra log da criação
	idCupom := uint(resp.ID)
	LogAction(c, "criar", "cupom", &idCupom,
		"Cupom criado: "+resp.Titulo, nil, resp)

	c.JSON(http.StatusCreated, resp)
}

// GetCupomHandler godoc
// @Summary      Busca cupom por ID
// @Description  Retorna os dados de um cupom específico pelo ID
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do cupom"
// @Success      200 {object} json.CupomResponse "Cupom encontrado"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Cupom não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/{id} [get]
func GetCupomHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetCupomByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cupom não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllCuponsHandler godoc
// @Summary      Lista todos os cupons
// @Description  Retorna uma lista com todos os cupons ativos
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Success      200 {array} json.CupomResponse "Lista de cupons"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /cupons [get]
func GetAllCuponsHandler(c *gin.Context) {
	resp, err := services.GetAllCupons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateCupomHandler godoc
// @Summary      Atualiza cupom
// @Description  Atualiza os dados de um cupom existente
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do cupom"
// @Param        request body json.CupomRequest true "Dados atualizados do cupom"
// @Success      200 {object} json.CupomResponse "Cupom atualizado com sucesso"
// @Failure      400 {object} map[string]interface{} "Dados inválidos"
// @Failure      404 {object} map[string]interface{} "Cupom não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/{id} [put]
func UpdateCupomHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.CupomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	// Para cupons de veículo do usuário, se IDLoja for 0, converte para nil
	if req.TipoCupom == "veiculo" && req.IDLoja != nil && *req.IDLoja == 0 {
		req.IDLoja = nil
	}

	// Busca dados antigos para o log
	cupomAntigo, _ := services.GetCupomByID(uint(id))

	resp, err := services.UpdateCupom(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cupom não encontrado",
		})
		return
	}

	// Registra log da atualização
	idCupom := uint(id)
	LogAction(c, "atualizar", "cupom", &idCupom,
		"Cupom atualizado: "+resp.Titulo, cupomAntigo, resp)

	c.JSON(http.StatusOK, resp)
}

// SoftDeleteCupomHandler godoc
// @Summary      Exclui cupom (soft delete)
// @Description  Realiza soft delete do cupom, marcando como excluído sem remover do banco
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do cupom"
// @Success      200 {object} map[string]interface{} "Cupom excluído com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Cupom não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/{id} [delete]
func SoftDeleteCupomHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Busca dados antes de deletar para o log
	cupomAntigo, _ := services.GetCupomByID(uint(id))

	err = services.SoftDeleteCupom(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cupom não encontrado",
		})
		return
	}

	// Registra log da exclusão
	idCupom := uint(id)
	LogAction(c, "deletar", "cupom", &idCupom,
		"Cupom excluído (soft delete)", cupomAntigo, nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "Cupom excluído com sucesso",
	})
}

// RestoreCupomHandler godoc
// @Summary      Restaura cupom excluído
// @Description  Restaura um cupom que foi soft deleted
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do cupom"
// @Success      200 {object} map[string]interface{} "Cupom restaurado com sucesso"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Cupom não encontrado"
// @Failure      500 {object} map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/{id}/restore [post]
func RestoreCupomHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreCupom(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Cupom não encontrado ou não foi excluído",
		})
		return
	}

	// Busca dados restaurados para o log
	cupomRestaurado, _ := services.GetCupomByID(uint(id))

	// Registra log da restauração
	idCupom := uint(id)
	LogAction(c, "restaurar", "cupom", &idCupom,
		"Cupom restaurado", nil, cupomRestaurado)

	c.JSON(http.StatusOK, gin.H{
		"message": "Cupom restaurado com sucesso",
	})
}

// GetCuponsProdutosHandler godoc
// @Summary      Lista cupons de produtos por proximidade
// @Description  Retorna cupons de produtos ordenados por proximidade da loja. Cada item inclui `id_loja` (loja do cupom).
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        latitude  query     number  false  "Latitude do usuário (opcional)"
// @Param        longitude query     number  false  "Longitude do usuário (opcional)"
// @Success      200       {object}  json.CuponsProdutoResponse
// @Failure      400       {object}  map[string]interface{} "Parâmetros inválidos"
// @Failure      500       {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/produtos [get]
func GetCuponsProdutosHandler(c *gin.Context) {
	latitudeStr := c.Query("latitude")
	longitudeStr := c.Query("longitude")

	var latitude, longitude *float64

	if latitudeStr != "" {
		var lat float64
		if _, err := fmt.Sscanf(latitudeStr, "%f", &lat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Latitude deve ser um número válido",
			})
			return
		}
		latitude = &lat
	}

	if longitudeStr != "" {
		var lng float64
		if _, err := fmt.Sscanf(longitudeStr, "%f", &lng); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Longitude deve ser um número válido",
			})
			return
		}
		longitude = &lng
	}

	if (latitude != nil && longitude == nil) || (latitude == nil && longitude != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Latitude e longitude devem ser fornecidos juntos",
		})
		return
	}

	resp, err := services.GetCuponsProdutos(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetCuponsVeiculosHandler godoc
// @Summary      Lista cupons de veículos por proximidade
// @Description  Retorna cupons de veículos ordenados por proximidade da loja
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        latitude  query     number  false  "Latitude do usuário (opcional)"
// @Param        longitude query     number  false  "Longitude do usuário (opcional)"
// @Success      200       {object}  json.CuponsVeiculoResponse
// @Failure      400       {object}  map[string]interface{} "Parâmetros inválidos"
// @Failure      500       {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/veiculos [get]
func GetCuponsVeiculosHandler(c *gin.Context) {
	latitudeStr := c.Query("latitude")
	longitudeStr := c.Query("longitude")

	var latitude, longitude *float64

	if latitudeStr != "" {
		var lat float64
		if _, err := fmt.Sscanf(latitudeStr, "%f", &lat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Latitude deve ser um número válido",
			})
			return
		}
		latitude = &lat
	}

	if longitudeStr != "" {
		var lng float64
		if _, err := fmt.Sscanf(longitudeStr, "%f", &lng); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Longitude deve ser um número válido",
			})
			return
		}
		longitude = &lng
	}

	if (latitude != nil && longitude == nil) || (latitude == nil && longitude != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Latitude e longitude devem ser fornecidos juntos",
		})
		return
	}

	resp, err := services.GetCuponsVeiculos(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetCuponsServicosHandler godoc
// @Summary      Lista cupons de serviços por proximidade
// @Description  Retorna cupons de serviços ordenados por proximidade da loja. Cada item inclui `id_loja` (loja do cupom).
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        latitude  query     number  false  "Latitude do usuário (opcional)"
// @Param        longitude query     number  false  "Longitude do usuário (opcional)"
// @Success      200       {object}  json.CuponsServicoResponse
// @Failure      400       {object}  map[string]interface{} "Parâmetros inválidos"
// @Failure      500       {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/servicos [get]
func GetCuponsServicosHandler(c *gin.Context) {
	latitudeStr := c.Query("latitude")
	longitudeStr := c.Query("longitude")

	var latitude, longitude *float64

	if latitudeStr != "" {
		var lat float64
		if _, err := fmt.Sscanf(latitudeStr, "%f", &lat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Latitude deve ser um número válido",
			})
			return
		}
		latitude = &lat
	}

	if longitudeStr != "" {
		var lng float64
		if _, err := fmt.Sscanf(longitudeStr, "%f", &lng); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Longitude deve ser um número válido",
			})
			return
		}
		longitude = &lng
	}

	if (latitude != nil && longitude == nil) || (latitude == nil && longitude != nil) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Latitude e longitude devem ser fornecidos juntos",
		})
		return
	}

	resp, err := services.GetCuponsServicos(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetCuponsByLojaIDHandler godoc
// @Summary      Lista cupons de uma loja
// @Description  Retorna todos os cupons ativos de uma loja específica
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        loja_id path int true "ID da loja"
// @Success      200  {object}  json.CuponsResponse "Lista de cupons"
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/loja/{loja_id} [get]
func GetCuponsByLojaIDHandler(c *gin.Context) {
	lojaIDStr := c.Param("loja_id")
	lojaID, err := strconv.ParseUint(lojaIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetCuponsByLojaID(uint(lojaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ResgatarCupomHandler godoc
// @Summary      Resgata um cupom
// @Description  Cria um histórico de resgate com status pendente. Envie `id_veiculo_usuario` (veículo do cliente) em cupons de produto/serviço para gravar em `id_veiculo` no resgate e permitir histórico no veículo ao efetivar.
// @Tags         Cupons
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do cupom"
// @Param        request body json.ResgatarCupomRequest true "Dados do resgate"
// @Success      201  {object}  json.HistoricoResgateResponse "Histórico de resgate criado com sucesso"
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Cupom não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /cupons/{id}/resgatar [post]
func ResgatarCupomHandler(c *gin.Context) {
	idStr := c.Param("id")
	cupomID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do cupom inválido",
		})
		return
	}

	var req json.ResgatarCupomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resp, err := services.CreateHistoricoResgateFromCupom(uint(cupomID), req.IDUsuario, req.MoedasUtilizadas, req.IDVeiculoUsuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Registra log do resgate
	idHistorico := uint(resp.ID)
	idCupomLog := uint(cupomID)
	LogAction(c, "resgatar", "cupom", &idCupomLog,
		"Cupom resgatado - histórico de resgate criado", nil, resp)
	LogAction(c, "criar", "historico_resgate", &idHistorico,
		"Histórico de resgate criado para cupom", nil, resp)

	c.JSON(http.StatusCreated, resp)
}
