package main

import (
	"log"
	"meu-carro-mais/internal/config"
	"meu-carro-mais/internal/database"
	"os"

	"meu-carro-mais/cmd/api/router"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()
	config.InitStripe()

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	r := router.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
