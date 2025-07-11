package database

import (
	"fmt"
	"log"
	"os"

	"meu-carro-mais/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	fmt.Println("DSN (do ambiente ou fallback):", dsn)

	if dsn == "" {
		dsn = "postgresql://postgres:password@localhost:5432/meucarromais?sslmode=disable"
		fmt.Println("Usando DSN de fallback local.")
	} else {
	}

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
