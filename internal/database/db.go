package database

import (
	"fmt"
	"log"

	"meu-carro-mais/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "postgres://postgres:CcbEg6cgbCg4BfCc64CafdaFEDEfaC3E@mainline.proxy.rlwy.net:36295/railway"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco com DSN '%s': %v", dsn, err)
	}

	DB = db

	// Auto-migrate
	err = db.AutoMigrate(
		&models.TipoPlano{},
		&models.CategoriaLojista{},
		&models.Usuario{},
		&models.Loja{},
		&models.HistoricoPlanoUsuario{},
		&models.Carteira{},
		&models.LogCarteira{},
		&models.Anuncio{},
	)
	if err != nil {
		log.Fatalf("Erro ao migrar tabelas: %v", err)
	}

	fmt.Println("Banco conectado e tabelas migradas com sucesso!")
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Conexão com o banco de dados não inicializada! Chame InitDB() primeiro.")
	}
	return DB
}
