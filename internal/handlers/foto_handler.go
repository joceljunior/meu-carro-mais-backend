package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateFotoHandler godoc
// @Summary      Cria uma nova foto
// @Description  Cria uma nova foto no sistema
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        foto body json.FotoRequest true "Dados da foto"
// @Success      201  {object}  json.FotoResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /fotos [post]
func CreateFotoHandler(c *gin.Context) {
	var req json.FotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateFoto(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetFotoHandler godoc
// @Summary      Busca foto por ID
// @Description  Retorna uma foto específica pelo ID
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da foto"
// @Success      200  {object}  json.FotoResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Foto não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /fotos/{id} [get]
func GetFotoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetFotoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Foto não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllFotosHandler godoc
// @Summary      Lista todas as fotos
// @Description  Retorna todas as fotos ativas do sistema
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.FotosResponse
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /fotos [get]
func GetAllFotosHandler(c *gin.Context) {
	resp, err := services.GetAllFotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.FotosResponse{
		Fotos: resp,
		Total: len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetFotosByVeiculoIDHandler godoc
// @Summary      Lista fotos de um veículo
// @Description  Retorna todas as fotos de um veículo específico
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo"
// @Success      200  {object}  json.FotosResponse
// @Failure      400  {object}  map[string]interface{} "ID de veículo inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id}/fotos [get]
func GetFotosByVeiculoIDHandler(c *gin.Context) {
	idVeiculoStr := c.Param("id")
	idVeiculo, err := strconv.ParseUint(idVeiculoStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de veículo inválido",
		})
		return
	}

	resp, err := services.GetFotosByVeiculoID(uint(idVeiculo))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetFotosByVeiculoLojaIDHandler godoc
// @Summary      Lista fotos de um veículo de loja
// @Description  Retorna todas as fotos de um veículo de loja específico
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo de loja"
// @Success      200  {object}  json.FotosResponse
// @Failure      400  {object}  map[string]interface{} "ID de veículo de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos-loja/{id}/fotos [get]
func GetFotosByVeiculoLojaIDHandler(c *gin.Context) {
	idVeiculoLojaStr := c.Param("id")
	idVeiculoLoja, err := strconv.ParseUint(idVeiculoLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de veículo de loja inválido",
		})
		return
	}

	resp, err := services.GetFotosByVeiculoLojaID(uint(idVeiculoLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetFotosByProdutoIDHandler godoc
// @Summary      Lista fotos de um produto
// @Description  Retorna todas as fotos de um produto específico
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do produto"
// @Success      200  {object}  json.FotosResponse
// @Failure      400  {object}  map[string]interface{} "ID de produto inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /produtos/{id}/fotos [get]
func GetFotosByProdutoIDHandler(c *gin.Context) {
	idProdutoStr := c.Param("id")
	idProduto, err := strconv.ParseUint(idProdutoStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de produto inválido",
		})
		return
	}

	resp, err := services.GetFotosByProdutoID(uint(idProduto))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetFotosByServicoIDHandler godoc
// @Summary      Lista fotos de um serviço
// @Description  Retorna todas as fotos de um serviço específico
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do serviço"
// @Success      200  {object}  json.FotosResponse
// @Failure      400  {object}  map[string]interface{} "ID de serviço inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /servicos/{id}/fotos [get]
func GetFotosByServicoIDHandler(c *gin.Context) {
	idServicoStr := c.Param("id")
	idServico, err := strconv.ParseUint(idServicoStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de serviço inválido",
		})
		return
	}

	resp, err := services.GetFotosByServicoID(uint(idServico))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetFotosByLojaIDHandler godoc
// @Summary      Lista fotos de uma loja
// @Description  Retorna todas as fotos de uma loja específica
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.FotosResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/fotos [get]
func GetFotosByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetFotosByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetFotoPrincipalByEntidadeHandler godoc
// @Summary      Busca foto principal de uma entidade
// @Description  Retorna a foto principal de uma entidade específica
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        tipo path string true "Tipo da entidade (veiculo, veiculo_loja, produto, servico, loja)"
// @Param        id path int true "ID da entidade"
// @Success      200  {object}  json.FotoResponse
// @Failure      400  {object}  map[string]interface{} "Parâmetros inválidos"
// @Failure      404  {object}  map[string]interface{} "Foto não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /fotos/principal/{tipo}/{id} [get]
func GetFotoPrincipalByEntidadeHandler(c *gin.Context) {
	tipoEntidade := c.Param("tipo")
	idEntidadeStr := c.Param("id")
	idEntidade, err := strconv.ParseUint(idEntidadeStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de entidade inválido",
		})
		return
	}

	resp, err := services.GetFotoPrincipalByEntidade(tipoEntidade, uint(idEntidade))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Foto não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateFotoHandler godoc
// @Summary      Atualiza foto
// @Description  Atualiza os dados de uma foto existente
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da foto"
// @Param        foto body json.FotoRequest true "Dados atualizados da foto"
// @Success      200  {object}  json.FotoResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Foto não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /fotos/{id} [put]
func UpdateFotoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req json.FotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.UpdateFoto(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Foto não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SetFotoPrincipalHandler godoc
// @Summary      Define foto como principal
// @Description  Define uma foto como principal da entidade
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da foto"
// @Success      200  {object}  map[string]interface{} "Foto definida como principal com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Foto não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /fotos/{id}/principal [put]
func SetFotoPrincipalHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SetFotoPrincipal(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Foto não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Foto definida como principal com sucesso",
	})
}

// SoftDeleteFotoHandler godoc
// @Summary      Remove foto (soft delete)
// @Description  Realiza soft delete de uma foto, marcando-a como excluída
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da foto"
// @Success      200  {object}  map[string]interface{} "Foto removida com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Foto não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /fotos/{id} [delete]
func SoftDeleteFotoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SoftDeleteFoto(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Foto não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Foto removida com sucesso",
	})
}

// RestoreFotoHandler godoc
// @Summary      Restaura foto
// @Description  Restaura uma foto que foi soft deleted
// @Tags         Fotos
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da foto"
// @Success      200  {object}  map[string]interface{} "Foto restaurada com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Foto não encontrada"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /fotos/{id}/restore [post]
func RestoreFotoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreFoto(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Foto não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Foto restaurada com sucesso",
	})
}
