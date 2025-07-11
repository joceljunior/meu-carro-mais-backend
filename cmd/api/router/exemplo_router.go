package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ExemploRouter struct{}

func (er *ExemploRouter) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/ping", func(c *gin.Context) {
		request := c.Request
		log.Printf("Recebida requisição: %s %s", request.Method, request.URL.Path)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
} 