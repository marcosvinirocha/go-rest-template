# Go REST API

Este é um exemplo de API REST criada em Go com uma estrutura organizada em pastas.

## Estrutura de Pastas

```
go-rest/
├── cmd/
│   └── main.go
├── api/
│   ├── handlers/
│   ├── models/
│   ├── routes/
│   ├── middleware/
│   ├── database/
│   └── config/
├── tests/
└── go.mod
```

## Começando

1. Clone este repositório
2. Instale as dependências: `go mod tidy`
3. Configure as variáveis de ambiente (veja `.env.example`)
4. Execute a aplicação: `go run cmd/main.go`

## Endpoints

- `GET /users` - Retorna todos os usuários
- `GET /users/{id}` - Retorna um usuário específico
- `POST /users` - Cria um novo usuário
- `PUT /users/{id}` - Atualiza um usuário existente
- `DELETE /users/{id}` - Exclui um usuário
- `GET /health` - Verifica a saúde da aplicação

## Dependências

- `github.com/lib/pq` - Driver PostgreSQL (apenas como exemplo)
