package router

import "github.com/gin-gonic/gin"

type IRouter interface {
	RegisterRoutes(rg *gin.RouterGroup)
} 