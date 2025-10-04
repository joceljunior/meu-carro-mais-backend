package main

import (
	"log"
	"meu-carro-mais/internal/config"
	"meu-carro-mais/internal/database"
	"os"

	"meu-carro-mais/cmd/api/router"

	"github.com/gin-gonic/gin"
)

// @title           Meu Carro Mais API
// @version         1.0
// @description     API para gerenciamento de veículos, anúncios, lojas e serviços.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      meu-carro-mais-production.up.railway.app
// @BasePath  /
// @schemes   https http

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
