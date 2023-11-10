# GO API <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

#### API desenvolvida em GO para fins de aprendizagem usando GORM para conexão com banco de dados e GIN para criação da API.

## Entidades
### User
### Post
#### Relacionamento: one-to-many (um usuário pode ter varios posts)

## Installation

```bash
$ go mod tidy
```

## Running the app

Para conseguir rodar a aplicação sem erros, devera criar o arquivo .env com as seguintes variáveis

```
# Porta onde a API vai rodar
PORT=3000

# uma string que o JWT vai usar para gerar o token
SECRET="eyJhbGciOiJdfdfdssfreteryhfghjhgjhgjytuurytd"

# URL de conexão com o banco de dados
DB_URL="host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
```

```bash
# start
$ go run main.go
```
