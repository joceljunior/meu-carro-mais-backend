#!/bin/bash

# Script para regenerar a documentação do Swagger
echo "🔄 Regenerando documentação do Swagger..."

# Instala o swag se não estiver instalado
if ! command -v swag &> /dev/null; then
    echo "📦 Instalando swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Regenera a documentação
echo "📝 Gerando documentação..."
swag init -g cmd/api/main.go -o docs

echo "✅ Documentação do Swagger regenerada com sucesso!"
echo "🌐 Acesse: http://localhost:8080/swagger/index.html"
