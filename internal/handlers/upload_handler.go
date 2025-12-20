package handlers

import (
	"net/http"
	"strconv"

	"meu-carro-mais/internal/handlers/json"
	"meu-carro-mais/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateUploadHandler godoc
// @Summary      Cria um novo upload
// @Description  Cria um novo upload (imagem ou documento) no sistema
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        upload body json.UploadRequest true "Dados do upload"
// @Success      201  {object}  json.UploadResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /uploads [post]
func CreateUploadHandler(c *gin.Context) {
	var req json.UploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateUpload(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Registra log da criação
	idUpload := uint(resp.ID)
	LogAction(c, "criar", "upload", &idUpload,
		"Upload criado: "+req.Tipo+" - "+req.NomeArquivo, nil, resp)

	c.JSON(http.StatusCreated, resp)
}

// GetUploadHandler godoc
// @Summary      Busca upload por ID
// @Description  Retorna um upload específico pelo ID
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do upload"
// @Success      200  {object}  json.UploadResponse
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Upload não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /uploads/{id} [get]
func GetUploadHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	resp, err := services.GetUploadByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Upload não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllUploadsHandler godoc
// @Summary      Lista todos os uploads
// @Description  Retorna todos os uploads ativos do sistema. Pode filtrar por tipo usando query parameter 'tipo' (Imagem ou Documento)
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        tipo query string false "Filtrar por tipo: Imagem ou Documento"
// @Success      200  {object}  json.UploadsResponse
// @Failure      400  {object}  map[string]interface{} "Tipo inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /uploads [get]
func GetAllUploadsHandler(c *gin.Context) {
	tipo := c.Query("tipo")
	
	var resp []json.UploadResponse
	var err error

	if tipo != "" {
		// Valida se o tipo é válido
		if tipo != "Imagem" && tipo != "Documento" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Tipo inválido. Use 'Imagem' ou 'Documento'",
			})
			return
		}
		resp, err = services.GetAllUploadsByTipo(tipo)
	} else {
		resp, err = services.GetAllUploads()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := json.UploadsResponse{
		Uploads: resp,
		Total:   len(resp),
	}

	c.JSON(http.StatusOK, response)
}

// GetUploadsByUsuarioIDHandler godoc
// @Summary      Lista uploads de um usuário
// @Description  Retorna todos os uploads de um usuário específico
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id_usuario path int true "ID do usuário"
// @Success      200  {object}  json.UploadsResponse
// @Failure      400  {object}  map[string]interface{} "ID de usuário inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /usuarios/{id_usuario}/uploads [get]
func GetUploadsByUsuarioIDHandler(c *gin.Context) {
	idUsuarioStr := c.Param("id_usuario")
	idUsuario, err := strconv.ParseUint(idUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	resp, err := services.GetUploadsByUsuarioID(uint(idUsuario))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUploadsByVeiculoIDHandler godoc
// @Summary      Lista uploads de um veículo
// @Description  Retorna todos os uploads de um veículo específico
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo"
// @Success      200  {object}  json.UploadsResponse
// @Failure      400  {object}  map[string]interface{} "ID de veículo inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos/{id}/uploads [get]
func GetUploadsByVeiculoIDHandler(c *gin.Context) {
	idVeiculoStr := c.Param("id")
	idVeiculo, err := strconv.ParseUint(idVeiculoStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de veículo inválido",
		})
		return
	}

	resp, err := services.GetUploadsByVeiculoID(uint(idVeiculo))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUploadsByVeiculoLojaIDHandler godoc
// @Summary      Lista uploads de um veículo de loja
// @Description  Retorna todos os uploads de um veículo de loja específico
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do veículo de loja"
// @Success      200  {object}  json.UploadsResponse
// @Failure      400  {object}  map[string]interface{} "ID de veículo de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /veiculos-loja/{id}/uploads [get]
func GetUploadsByVeiculoLojaIDHandler(c *gin.Context) {
	idVeiculoLojaStr := c.Param("id")
	idVeiculoLoja, err := strconv.ParseUint(idVeiculoLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de veículo de loja inválido",
		})
		return
	}

	resp, err := services.GetUploadsByVeiculoLojaID(uint(idVeiculoLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUploadsByProdutoIDHandler godoc
// @Summary      Lista uploads de um produto
// @Description  Retorna todos os uploads de um produto específico
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do produto"
// @Success      200  {object}  json.UploadsResponse
// @Failure      400  {object}  map[string]interface{} "ID de produto inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /produtos/{id}/uploads [get]
func GetUploadsByProdutoIDHandler(c *gin.Context) {
	idProdutoStr := c.Param("id")
	idProduto, err := strconv.ParseUint(idProdutoStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de produto inválido",
		})
		return
	}

	resp, err := services.GetUploadsByProdutoID(uint(idProduto))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUploadsByServicoIDHandler godoc
// @Summary      Lista uploads de um serviço
// @Description  Retorna todos os uploads de um serviço específico
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do serviço"
// @Success      200  {object}  json.UploadsResponse
// @Failure      400  {object}  map[string]interface{} "ID de serviço inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /servicos/{id}/uploads [get]
func GetUploadsByServicoIDHandler(c *gin.Context) {
	idServicoStr := c.Param("id")
	idServico, err := strconv.ParseUint(idServicoStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de serviço inválido",
		})
		return
	}

	resp, err := services.GetUploadsByServicoID(uint(idServico))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUploadsByLojaIDHandler godoc
// @Summary      Lista uploads de uma loja
// @Description  Retorna todos os uploads de uma loja específica
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID da loja"
// @Success      200  {object}  json.UploadsResponse
// @Failure      400  {object}  map[string]interface{} "ID de loja inválido"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /lojas/{id}/uploads [get]
func GetUploadsByLojaIDHandler(c *gin.Context) {
	idLojaStr := c.Param("id")
	idLoja, err := strconv.ParseUint(idLojaStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de loja inválido",
		})
		return
	}

	resp, err := services.GetUploadsByLojaID(uint(idLoja))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUploadPrincipalByEntidadeHandler godoc
// @Summary      Busca upload principal de uma entidade
// @Description  Retorna o upload principal (imagem) de uma entidade específica
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        tipo path string true "Tipo da entidade (usuario, veiculo, veiculo_loja, produto, servico, loja)"
// @Param        id path int true "ID da entidade"
// @Success      200  {object}  json.UploadResponse
// @Failure      400  {object}  map[string]interface{} "Parâmetros inválidos"
// @Failure      404  {object}  map[string]interface{} "Upload não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /uploads/principal/{tipo}/{id} [get]
func GetUploadPrincipalByEntidadeHandler(c *gin.Context) {
	tipoEntidade := c.Param("tipo")
	idEntidadeStr := c.Param("id")
	idEntidade, err := strconv.ParseUint(idEntidadeStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de entidade inválido",
		})
		return
	}

	resp, err := services.GetUploadPrincipalByEntidade(tipoEntidade, uint(idEntidade))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Upload não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUploadHandler godoc
// @Summary      Atualiza upload
// @Description  Atualiza os dados de um upload existente
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do upload"
// @Param        upload body json.UploadRequest true "Dados atualizados do upload"
// @Success      200  {object}  json.UploadResponse
// @Failure      400  {object}  map[string]interface{} "Dados inválidos"
// @Failure      404  {object}  map[string]interface{} "Upload não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /uploads/{id} [put]
func UpdateUploadHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Busca dados antigos para o log
	uploadAntigo, _ := services.GetUploadByID(uint(id))

	var req json.UploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	resp, err := services.UpdateUpload(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Upload não encontrado",
		})
		return
	}

	// Registra log da atualização
	idUpload := uint(id)
	LogAction(c, "atualizar", "upload", &idUpload,
		"Upload atualizado: "+req.Tipo+" - "+req.NomeArquivo, uploadAntigo, resp)

	c.JSON(http.StatusOK, resp)
}

// SetUploadPrincipalHandler godoc
// @Summary      Define upload como principal
// @Description  Define um upload como principal da entidade (apenas para imagens)
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do upload"
// @Success      200  {object}  map[string]interface{} "Upload definido como principal com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido ou upload não é imagem"
// @Failure      404  {object}  map[string]interface{} "Upload não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /uploads/{id}/principal [put]
func SetUploadPrincipalHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.SetUploadPrincipal(uint(id))
	if err != nil {
		if err.Error() == "apenas imagens podem ser definidas como principais" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Upload não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload definido como principal com sucesso",
	})
}

// SoftDeleteUploadHandler godoc
// @Summary      Remove upload (soft delete)
// @Description  Realiza soft delete de um upload, marcando-o como excluído
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do upload"
// @Success      200  {object}  map[string]interface{} "Upload removido com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Upload não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /uploads/{id} [delete]
func SoftDeleteUploadHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	// Busca dados antes de deletar para o log
	uploadAntigo, _ := services.GetUploadByID(uint(id))

	err = services.SoftDeleteUpload(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Upload não encontrado",
		})
		return
	}

	// Registra log da exclusão
	idUpload := uint(id)
	LogAction(c, "deletar", "upload", &idUpload,
		"Upload excluído (soft delete)", uploadAntigo, nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload removido com sucesso",
	})
}

// RestoreUploadHandler godoc
// @Summary      Restaura upload
// @Description  Restaura um upload que foi soft deleted
// @Tags         Uploads
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do upload"
// @Success      200  {object}  map[string]interface{} "Upload restaurado com sucesso"
// @Failure      400  {object}  map[string]interface{} "ID inválido"
// @Failure      404  {object}  map[string]interface{} "Upload não encontrado"
// @Failure      500  {object}  map[string]interface{} "Erro interno do servidor"
// @Router       /uploads/{id}/restore [post]
func RestoreUploadHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = services.RestoreUpload(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Upload não encontrado ou não foi excluído",
		})
		return
	}

	// Busca dados restaurados para o log
	uploadRestaurado, _ := services.GetUploadByID(uint(id))

	// Registra log da restauração
	idUpload := uint(id)
	LogAction(c, "restaurar", "upload", &idUpload,
		"Upload restaurado", nil, uploadRestaurado)

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload restaurado com sucesso",
	})
}

