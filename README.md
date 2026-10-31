![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

# FilmApi

A Go API that searches films by title using the [OMDb API](https://www.omdbapi.com/).

## Requirements

- Go 1.22 or higher
- An OMDb API key (free at https://www.omdbapi.com/apikey.aspx)

## Configuration

The key is read from the `API_KEY` environment variable.

PowerShell:

```powershell
$env:API_KEY = "your-key"
```

Linux/macOS:

```bash
export API_KEY="your-key"
```

## Running

```bash
go run ./cmd/api
```

The server starts at `http://localhost:8080`.

## Endpoint

### GET /?s={title}

Searches films by title.

```bash
curl "http://localhost:8080/?s=matrix"
```

Response (200):

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

Errors:

| Status | When                                    |
|--------|-----------------------------------------|
| 400    | `s` parameter missing or empty          |
| 404    | no films found                          |
| 502    | failed to reach OMDb                    |
| 500    | unexpected error                        |

Error format:

```json
{ "error": "film name can't be empty" }
```

## Project structure

```
cmd/api/                    entry point (main.go)
external/omdb/              OMDb HTTP client (implements FilmSearcher)
internal/domain/            Film entity and name validation
internal/application/       SearchFilms use case and FilmSearcher interface
internal/exception/         application errors
internal/interfaces/httpapi handler, DTOs and error-to-HTTP mapping
```

The `application` layer defines the `FilmSearcher` interface and `external/omdb` implements it, so the use case does not depend on OMDb and can be tested with a fake.

## Tests

```bash
go test ./...
```

## CORS

By default, requests from `http://localhost:3000` are allowed. To change this, edit `AllowedOrigins` in `internal/interfaces/httpapi/handler.go`.