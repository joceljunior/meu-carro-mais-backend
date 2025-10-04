# Script PowerShell para regenerar a documentação do Swagger
Write-Host "🔄 Regenerando documentação do Swagger..." -ForegroundColor Cyan

# Verifica se o swag está instalado
try {
    swag --version | Out-Null
    Write-Host "✅ Swag já está instalado" -ForegroundColor Green
} catch {
    Write-Host "📦 Instalando swag..." -ForegroundColor Yellow
    go install github.com/swaggo/swag/cmd/swag@latest
}

# Regenera a documentação
Write-Host "📝 Gerando documentação..." -ForegroundColor Cyan
swag init -g cmd/api/main.go -o docs

Write-Host "✅ Documentação do Swagger regenerada com sucesso!" -ForegroundColor Green
Write-Host "🌐 Acesse: http://localhost:8080/swagger/index.html" -ForegroundColor Blue
