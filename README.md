# test_lambda_go

Simple AWS Lambda functions in Go over a static in-memory movie catalog
(title, director, genre, actors).

## Functions

| Function | Path | Input | Output |
|---|---|---|---|
| `list-movies` | `cmd/list-movies` | none | `[]movies.Movie` |
| `get-by-title` | `cmd/get-by-title` | `{"title": "Inception"}` | `movies.Movie` (error if not found) |
| `get-by-genre` | `cmd/get-by-genre` | `{"genre": "Sci-Fi"}` | `[]movies.Movie` |
| `get-by-director` | `cmd/get-by-director` | `{"director": "Christopher Nolan"}` | `[]movies.Movie` |

Title, genre and director matching is case-insensitive.

## Test

```bash
go test ./...
```

## Build and package for AWS Lambda

Each function targets the `provided.al2023` runtime with handler `bootstrap`.
From the repository root, for each function directory:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./cmd/list-movies
zip function.zip bootstrap
```

Upload `function.zip` to a Lambda function configured with the
`provided.al2023` runtime. Repeat for `get-by-title`, `get-by-genre` and
`get-by-director`.
