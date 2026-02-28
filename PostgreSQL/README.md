# PostgreSQL Stock API (Go)

This project is a simple CRUD REST API for managing stock records in a PostgreSQL database.

## What the project is doing

- Starts an HTTP server on `:8080`
- Connects to PostgreSQL using `POSTGRES_URL` from `.env`
- Exposes endpoints to create, read, update, and delete stock records
- Uses JSON request/response payloads

The API works with a `stocks` table and a `Stock` model:

- `stockid` (int64)
- `name` (string)
- `price` (float64)
- `company` (string)

## Tech stack

- Go
- `github.com/gorilla/mux` (routing)
- `database/sql` (DB access)
- `github.com/joho/godotenv` (environment loading)
- PostgreSQL

## Project structure

- `main.go` - app entrypoint, starts server
- `router/router.go` - route definitions
- `middleware/handlers.go` - HTTP handlers + DB logic
- `models/models.go` - data models
- `.env` - environment variables

## API routes

- `GET /api/stock/{id}` - get one stock by ID
- `GET /api/stock` - get all stocks
- `POST /api/newstock` - create a stock
- `PUT /api/stock/{id}` - update stock by ID
- `DELETE /api/deletestock/{id}` - delete stock by ID

## Setup

1. Install Go and PostgreSQL.
2. Create a database (example: `stocksdb`).
3. Create a `stocks` table matching the model fields.
4. Set `POSTGRES_URL` in `.env`.
5. Install dependencies and run:

```bash
go mod tidy
go run main.go
```

Server runs at `http://localhost:8080`.

## Example request body

```json
{
  "name": "TCS",
  "price": 3845.5,
  "company": "Tata Consultancy Services"
}
```

## Current behavior notes

- The service opens a new DB connection for each DB operation.
- Error handling currently uses `log.Fatal` in handlers/DB functions, which exits the process on runtime errors.
- `GET /api/stock` currently maps to a query intended for all records.
