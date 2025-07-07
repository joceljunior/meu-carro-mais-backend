package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Define o modo do Gin (release, debug, test)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Cria uma instância do router Gin
	router := gin.Default()

	// Rota de exemplo
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Porta padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Inicia o servidor
	log.Printf("Servidor rodando na porta %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
