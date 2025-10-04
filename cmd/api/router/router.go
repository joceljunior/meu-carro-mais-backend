package router

import (
	"meu-carro-mais/docs"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Configura o host do Swagger baseado no ambiente
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		// Se não houver variável de ambiente, usa produção como padrão
		swaggerHost = "meu-carro-mais-production.up.railway.app"
	}

	docs.SwaggerInfo.Host = swaggerHost
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

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
