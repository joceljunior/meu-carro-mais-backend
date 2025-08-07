package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/seeds"
)

func main() {
	// Define os comandos disponíveis
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	// Verifica se foi passado algum comando
	if len(os.Args) < 2 {
		fmt.Println("Uso: seed <comando>")
		fmt.Println("Comandos disponíveis:")
		fmt.Println("  run   - Executa todos os seeds")
		fmt.Println("  help  - Mostra esta ajuda")
		os.Exit(1)
	}

	// Inicializa a conexão com o banco
	database.InitDB()

	// Processa o comando
	switch os.Args[1] {
	case "run":
		runCmd.Parse(os.Args[2:])
		fmt.Println("🌱 Iniciando execução dos seeds...")
		
		seeder := seeds.NewSeeder()
		if err := seeder.Run(); err != nil {
			log.Fatalf("Erro ao executar seeds: %v", err)
		}
		
		fmt.Println("✅ Seeds executados com sucesso!")

	case "help":
		helpCmd.Parse(os.Args[2:])
		fmt.Println("Sistema de Seeds - Meu Carro Mais")
		fmt.Println("=================================")
		fmt.Println("")
		fmt.Println("Este comando facilita a população das tabelas com dados de exemplo.")
		fmt.Println("")
		fmt.Println("Comandos:")
		fmt.Println("  run   - Executa todos os seeds")
		fmt.Println("  help  - Mostra esta ajuda")
		fmt.Println("")
		fmt.Println("Exemplos:")
		fmt.Println("  go run cmd/seed/main.go run    # Executa todos os seeds")
		fmt.Println("  go run cmd/seed/main.go help   # Mostra esta ajuda")
		fmt.Println("")
		fmt.Println("Dados que serão inseridos:")
		fmt.Println("  - Tipos de Plano (Gratuito, Básico, Premium, Enterprise)")
		fmt.Println("  - Categorias de Lojista (Concessionária, Oficina, etc.)")
		fmt.Println("  - Categorias de Anúncio (Carros, Motos, etc.)")
		fmt.Println("  - Usuários de exemplo (João, Maria, Pedro)")
		fmt.Println("  - Lojas de exemplo (Auto Center, Oficina do João, etc.)")
		fmt.Println("  - Anúncios de exemplo (Honda Civic, Toyota Corolla, etc.)")
		fmt.Println("  - Carteiras de exemplo")

	default:
		fmt.Printf("Comando '%s' não reconhecido\n", os.Args[1])
		fmt.Println("Comandos disponíveis: run, help")
		os.Exit(1)
	}
} 