#!/bin/bash

# Script para gerenciar seeds da API
# Uso: ./scripts/seed.sh [comando]

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Função para imprimir mensagens coloridas
print_message() {
    echo -e "${GREEN}[SEED]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

# Verifica se o comando foi fornecido
if [ $# -eq 0 ]; then
    echo "Uso: $0 <comando>"
    echo ""
    echo "Comandos disponíveis:"
    echo "  run   - Executa todos os seeds"
    echo "  help  - Mostra esta ajuda"
    echo "  status - Mostra informações sobre os seeds"
    exit 1
fi

COMMAND=$1

case $COMMAND in
    "run")
        print_message "Executando seeds..."
        print_info "Isso irá popular as tabelas com dados de exemplo"
        read -p "Tem certeza que deseja continuar? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            go run cmd/seed/main.go run
            print_message "Seeds executados com sucesso!"
        else
            print_message "Seeds cancelados."
        fi
        ;;
    "help")
        print_message "Mostrando ajuda..."
        go run cmd/seed/main.go help
        ;;
    "status")
        print_info "Status dos Seeds"
        echo "=================="
        echo ""
        echo "Este comando irá inserir os seguintes dados:"
        echo ""
        echo "📋 Tipos de Plano:"
        echo "  - Gratuito"
        echo "  - Básico"
        echo "  - Premium"
        echo "  - Enterprise"
        echo ""
        echo "🏪 Categorias de Lojista:"
        echo "  - Concessionária"
        echo "  - Loja de Peças"
        echo "  - Oficina Mecânica"
        echo "  - Lava Jato"
        echo "  - Seguros"
        echo "  - Financiamento"
        echo ""
        echo "🚗 Categorias de Anúncio:"
        echo "  - Carros"
        echo "  - Motos"
        echo "  - Caminhões"
        echo "  - Peças"
        echo "  - Serviços"
        echo "  - Acessórios"
        echo ""
        echo "👥 Usuários de Exemplo:"
        echo "  - João Silva (joao@email.com)"
        echo "  - Maria Santos (maria@email.com)"
        echo "  - Pedro Costa (pedro@email.com)"
        echo ""
        echo "🏢 Lojas de Exemplo:"
        echo "  - Auto Center São Paulo"
        echo "  - Oficina do João"
        echo "  - Carros Premium"
        echo ""
        echo "🚙 Anúncios de Exemplo:"
        echo "  - Honda Civic 2020 (R$ 85.000)"
        echo "  - Toyota Corolla 2019 (R$ 75.000)"
        echo "  - Volkswagen Golf GTI (R$ 95.000)"
        echo ""
        echo "💳 Carteiras de Exemplo:"
        echo "  - Carteira para João Silva (R$ 1.000)"
        ;;
    *)
        print_error "Comando '$COMMAND' não reconhecido"
        echo "Use '$0 help' para ver os comandos disponíveis"
        exit 1
        ;;
esac 