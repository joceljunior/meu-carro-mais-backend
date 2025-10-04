# Configurações de Ambiente

## Variáveis de Ambiente para Swagger

### Desenvolvimento/Debug
```bash
GIN_MODE=debug
PORT=8080
SWAGGER_HOST=localhost:8080
```

### Produção
```bash
GIN_MODE=release
RAILWAY_ENVIRONMENT=production
# SWAGGER_HOST será automaticamente definido como meu-carro-mais-production.up.railway.app
```

## Detecção Automática

O sistema detecta automaticamente o ambiente baseado nas seguintes variáveis:

1. **RAILWAY_ENVIRONMENT=production** → Usa URL de produção
2. **GIN_MODE=release** → Usa URL de produção  
3. **SWAGGER_HOST definida** → Usa o valor especificado
4. **Caso contrário** → Usa localhost:PORT

## URLs do Swagger

- **Desenvolvimento**: http://localhost:8080/swagger/index.html
- **Produção**: https://meu-carro-mais-production.up.railway.app/swagger/index.html
