package router

import (
	"meu-carro-mais/docs"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Configuração de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todas as origens (ajuste conforme necessário)
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Stripe-Signature"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configura o host do Swagger baseado no ambiente
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		// Detecta se está rodando localmente ou em produção
		port := os.Getenv("PORT")
		ginMode := os.Getenv("GIN_MODE")

		// Se não está em modo release ou não tem porta definida, assume localhost
		if ginMode != "release" || port == "" {
			if port == "" {
				port = "8080"
			}
			swaggerHost = "localhost:" + port
			docs.SwaggerInfo.Schemes = []string{"http"}
		} else {
			// Em produção, usa a URL do Railway
			swaggerHost = "meu-carro-mais-production.up.railway.app"
			docs.SwaggerInfo.Schemes = []string{"https", "http"}
		}
	} else {
		// Se SWAGGER_HOST está definido, detecta o scheme baseado no host
		if swaggerHost == "localhost:8080" || swaggerHost == "127.0.0.1:8080" {
			docs.SwaggerInfo.Schemes = []string{"http"}
		} else {
			docs.SwaggerInfo.Schemes = []string{"https", "http"}
		}
	}

	docs.SwaggerInfo.Host = swaggerHost
	docs.SwaggerInfo.BasePath = "/"

	// Swagger UI na raiz e em /swagger
	r.GET("/", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/")

	routers := []IRouter{
		&LoginRouter{},
		&ExemploRouter{},
		&LojaRouter{},
		&ServicoRouter{},
		&UserRouter{},
		&AnuncioRouter{},
		&VeiculoRouter{},
		&VeiculoLojaRouter{},
		&ProdutoRouter{},
		&HistoricoResgateRouter{},
		&AvaliacaoRouter{},
		&FotoRouter{},
		&PagamentoRouter{},
	}

	for _, rt := range routers {
		rt.RegisterRoutes(api)
	}

	return r
}
