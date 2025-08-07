#!/bin/bash

# Script para gerenciar migrations da API
# Uso: ./scripts/migrate.sh [comando]

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Função para imprimir mensagens coloridas
print_message() {
    echo -e "${GREEN}[MIGRATE]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Verifica se o comando foi fornecido
if [ $# -eq 0 ]; then
    echo "Uso: $0 <comando>"
    echo ""
    echo "Comandos disponíveis:"
    echo "  run      - Executa todas as migrations pendentes"
    echo "  status   - Mostra o status das migrations"
    echo "  rollback - Executa rollback da última migration"
    echo "  help     - Mostra esta ajuda"
    exit 1
fi

COMMAND=$1

case $COMMAND in
    "run")
        print_message "Executando migrations..."
        go run cmd/migrate/main.go run
        print_message "Migrations executadas com sucesso!"
        ;;
    "status")
        print_message "Verificando status das migrations..."
        go run cmd/migrate/main.go status
        ;;
    "rollback")
        print_warning "ATENÇÃO: Isso irá reverter a última migration!"
        read -p "Tem certeza que deseja continuar? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            print_message "Executando rollback..."
            go run cmd/migrate/main.go rollback
            print_message "Rollback executado com sucesso!"
        else
            print_message "Rollback cancelado."
        fi
        ;;
    "help")
        echo "Sistema de Migrations - Meu Carro Mais"
        echo "======================================"
        echo ""
        echo "Este script facilita o gerenciamento de migrations do banco de dados."
        echo ""
        echo "Comandos:"
        echo "  run      - Executa todas as migrations pendentes"
        echo "  status   - Mostra o status das migrations"
        echo "  rollback - Executa rollback da última migration"
        echo "  help     - Mostra esta ajuda"
        echo ""
        echo "Exemplos:"
        echo "  $0 run      # Executa migrations"
        echo "  $0 status   # Verifica status"
        echo "  $0 rollback # Faz rollback"
        ;;
    *)
        print_error "Comando '$COMMAND' não reconhecido"
        echo "Use '$0 help' para ver os comandos disponíveis"
        exit 1
        ;;
esac 