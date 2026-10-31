![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

# FilmApi

API em Go que busca filmes pelo nome consultando a [OMDb API](https://www.omdbapi.com/).

## Requisitos

- Go 1.22 ou superior
- Chave de API do OMDb (gratuita em https://www.omdbapi.com/apikey.aspx)

## Configuração

A chave é lida da variável de ambiente `API_KEY`.

PowerShell:

```powershell
$env:API_KEY = "sua-chave"
```

Linux/macOS:

```bash
export API_KEY="sua-chave"
```

## Executando

```bash
go run ./cmd/api
```

O servidor sobe em `http://localhost:8080`.

## Endpoint

### GET /?s={nome}

Busca filmes pelo nome.

```bash
curl "http://localhost:8080/?s=matrix"
```

Resposta (200):

```json
{
  "data": [
    {
      "Title": "The Matrix",
      "Year": "1999",
      "ImdbID": "tt0133093",
      "Type": "movie",
      "Poster": "https://..."
    }
  ]
}
```

Erros:

| Status | Quando                                  |
|--------|-----------------------------------------|
| 400    | parâmetro `s` ausente ou vazio          |
| 404    | nenhum filme encontrado                 |
| 502    | falha ao consultar o OMDb               |
| 500    | erro inesperado                         |

Formato do erro:

```json
{ "error": "film name can't be empty" }
```

## Estrutura

```
cmd/api/                    ponto de entrada (main.go)
external/omdb/              cliente HTTP do OMDb (implementa FilmSearcher)
internal/domain/            entidade Film e validação do nome
internal/application/       caso de uso SearchFilms e interface FilmSearcher
internal/exception/         erros da aplicação
internal/interfaces/httpapi handler, DTOs e tradução de erros para HTTP
```

A `application` define a interface `FilmSearcher` e o `external/omdb` a implementa, então o caso de uso não depende do OMDb e pode ser testado com um fake.

## Testes

```bash
go test ./...
```

## CORS

Por padrão aceita requisições de `http://localhost:3000`. Para mudar, edite `AllowedOrigins` em `internal/interfaces/httpapi/handler.go`.