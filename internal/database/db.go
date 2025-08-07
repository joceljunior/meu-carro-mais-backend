package database

import (
	"fmt"
	"log"

	"meu-carro-mais/internal/database/migrations"

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

	// Executa as migrations
	migrator := migrations.NewMigrator(db)
	if err := migrator.Run(); err != nil {
		log.Fatalf("Erro ao executar migrations: %v", err)
	}

	fmt.Println("Banco conectado e migrations executadas com sucesso!")
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Conexão com o banco de dados não inicializada! Chame InitDB() primeiro.")
	}
	return DB
}

// RunMigrations executa as migrations manualmente (útil para comandos CLI)
func RunMigrations() error {
	if DB == nil {
		return fmt.Errorf("conexão com o banco de dados não inicializada")
	}

	migrator := migrations.NewMigrator(DB)
	return migrator.Run()
}

// RollbackMigration executa o rollback da última migration
func RollbackMigration() error {
	if DB == nil {
		return fmt.Errorf("conexão com o banco de dados não inicializada")
	}

	migrator := migrations.NewMigrator(DB)
	return migrator.Rollback()
}

// MigrationStatus mostra o status das migrations
func MigrationStatus() error {
	if DB == nil {
		return fmt.Errorf("conexão com o banco de dados não inicializada")
	}

	migrator := migrations.NewMigrator(DB)
	return migrator.Status()
}
