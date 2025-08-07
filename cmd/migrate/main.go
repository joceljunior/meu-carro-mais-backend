package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"meu-carro-mais/internal/database"
)

func main() {
	// Define os comandos disponíveis
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	rollbackCmd := flag.NewFlagSet("rollback", flag.ExitOnError)
	statusCmd := flag.NewFlagSet("status", flag.ExitOnError)

	// Verifica se foi passado algum comando
	if len(os.Args) < 2 {
		fmt.Println("Uso: migrate <comando>")
		fmt.Println("Comandos disponíveis:")
		fmt.Println("  run      - Executa todas as migrations pendentes")
		fmt.Println("  rollback - Executa rollback da última migration")
		fmt.Println("  status   - Mostra o status das migrations")
		os.Exit(1)
	}

	// Inicializa a conexão com o banco
	database.InitDB()

	// Processa o comando
	switch os.Args[1] {
	case "run":
		runCmd.Parse(os.Args[2:])
		if err := database.RunMigrations(); err != nil {
			log.Fatalf("Erro ao executar migrations: %v", err)
		}
		fmt.Println("Migrations executadas com sucesso!")

	case "rollback":
		rollbackCmd.Parse(os.Args[2:])
		if err := database.RollbackMigration(); err != nil {
			log.Fatalf("Erro ao executar rollback: %v", err)
		}
		fmt.Println("Rollback executado com sucesso!")

	case "status":
		statusCmd.Parse(os.Args[2:])
		if err := database.MigrationStatus(); err != nil {
			log.Fatalf("Erro ao verificar status: %v", err)
		}

	default:
		fmt.Printf("Comando '%s' não reconhecido\n", os.Args[1])
		fmt.Println("Comandos disponíveis: run, rollback, status")
		os.Exit(1)
	}
}
