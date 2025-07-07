package database

import (
	"fmt"
	"log"
	"os"

	"meu-carro-mais/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=meu_carro_mais port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	// Auto-migrate
	err = db.AutoMigrate(&models.TipoPlano{}, &models.CategoriaLojista{}, &models.Usuario{}, &models.Loja{}, &models.HistoricoPlanoUsuario{}, &models.Carteira{}, &models.LogCarteira{})
	if err != nil {
		log.Fatalf("Erro ao migrar tabelas: %v", err)
	}

	fmt.Println("Banco conectado e tabelas migradas!")
	return db
}
