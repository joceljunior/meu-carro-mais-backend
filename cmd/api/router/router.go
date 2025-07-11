package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "meu-carro-mais/docs"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Swagger UI na raiz e em /swagger
	r.GET("/", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/")

	routers := []IRouter{
		&LoginRouter{},
		&ExemploRouter{},
	}

	for _, rt := range routers {
		rt.RegisterRoutes(api)
	}

	return r
} 