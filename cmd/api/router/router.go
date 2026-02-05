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

	// Configuração de CORS - mais permissiva para funcionar com Swagger
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://meu-carro-mais-production.up.railway.app",
			"http://meu-carro-mais-production.up.railway.app",
			"https://meu-carro-mais-backend-production.up.railway.app",
			"http://meu-carro-mais-backend-production.up.railway.app",
			"http://localhost:3000",
			"http://localhost:8080",
			"http://127.0.0.1:3000",
			"http://127.0.0.1:8080",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Stripe-Signature", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowWildcard:    true,
		MaxAge:           12 * time.Hour,
	}))

	// Configura o host do Swagger baseado no ambiente
	configureSwaggerHost()

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
		&UploadRouter{},
		&PagamentoRouter{},
		&CarteiraRouter{},
		&RegistroInteresseRouter{},
		&LogRouter{},
		&DescontoRouter{},
		&OfertaAutoMaisRouter{},
	}

	for _, rt := range routers {
		rt.RegisterRoutes(api)
	}

	return r
}

// configureSwaggerHost configura o host do Swagger baseado no ambiente
func configureSwaggerHost() {
	// Detecta o ambiente baseado em variáveis de ambiente
	ginMode := os.Getenv("GIN_MODE")
	railwayEnvironment := os.Getenv("RAILWAY_ENVIRONMENT")
	port := os.Getenv("PORT")

	// Se estiver rodando no Railway (produção)
	if railwayEnvironment == "production" || ginMode == "release" {
		docs.SwaggerInfo.Host = "meu-carro-mais-backend-production.up.railway.app"
		docs.SwaggerInfo.Schemes = []string{"https", "http"}
	} else {
		// Ambiente de desenvolvimento/local
		// Se a variável SWAGGER_HOST estiver definida, usa ela
		if swaggerHost := os.Getenv("SWAGGER_HOST"); swaggerHost != "" {
			docs.SwaggerInfo.Host = swaggerHost
		} else {
			// Detecta automaticamente o host baseado na porta
			if port == "" {
				port = "8080"
			}
			docs.SwaggerInfo.Host = "localhost:" + port
		}
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	docs.SwaggerInfo.BasePath = "/"
}
