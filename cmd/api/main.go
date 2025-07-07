package main

import (
	"log"
	"meu-carro-mais/internal/database"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o banco de dados e faz auto-migrate
	database.InitDB()

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)
	router := gin.Default()

	// Rota de exemplo
	router.GET("/ping", func(c *gin.Context) {

		request := c.Request
		log.Printf("Recebida requisição: %s %s", request.Method, request.URL.Path)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
