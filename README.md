# Simple GO oapi-codegen, SQLc, and Sqlite3 example

This is a simple example of how to use `oapi-codegen` to generate a Go server from an OpenAPI 3.0 spec,
and use `sqlc` to generate Go code from SQL queries. The server uses `sqlite3` as the database.

## Project structure

- `tools` - Contains the `oapi-codegen` and `sqlc` import for generating Go code.
- `api` - Contains the OpenAPI 3.0 spec & generated code.
- `internal/database` - Contains the SQL queries and generated code.
- `internal/server` - Contains the server implementation code.
- `internal/domain` - Contains the domain services.
- (PLANNED) `internal/integration` - Contains the integration tests.

## Development

### Prerequisites
- [Go](https://golang.org/dl/)
- [Make](https://www.gnu.org/software/make/)

### Generate code (API & SQL)
```bash
make gen
```

### Build the server
```bash
make build
```

### Run the server
```bash
./api-server
```

## Todo
- [ ] Add tests
- [ ] Add more endpoints (DELETE & PATCH)
- [ ] Add more complex SQL queries
- [ ] Add OAuth2 support
