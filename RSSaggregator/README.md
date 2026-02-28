# RSS Aggregator

This project is a Go-based backend service for an RSS aggregation platform.

Current implementation focuses on:
- HTTP API bootstrapping with `chi`
- PostgreSQL integration
- User creation with generated API keys
- SQL migrations (`goose` style) and type-safe queries via `sqlc`

## Tech Stack
- Go `1.24.3`
- Router: `github.com/go-chi/chi`
- CORS middleware: `github.com/go-chi/cors`
- PostgreSQL driver: `github.com/lib/pq`
- Env loading: `github.com/joho/godotenv`
- Query generation: `sqlc`

## Project Structure
- `main.go`: app entrypoint, router setup, DB connection
- `handlerReady.go`: health endpoint
- `handlerUser.go`: create-user endpoint
- `json.go`: shared JSON response helpers
- `internal/database/`: generated query + model code
- `sql/Schema/`: SQL migrations
- `sql/queries/`: SQL query definitions

## Environment Variables
Create a `.env` file:

```env
PORT=8080
DB_URL=postgres://<user>:<password>@localhost:5432/<db>?sslmode=disable
```

## Database
The `users` table currently includes:
- `id` (UUID, primary key)
- `created_at` (timestamp)
- `updated_at` (timestamp)
- `name` (text)
- `api_key` (unique, generated server-side)

## Run Locally
1. Start PostgreSQL and create the target database.
2. Apply migrations in `sql/Schema`.
3. Ensure `.env` is configured.
4. Run:

```bash
go run .
```

Server starts on `http://localhost:<PORT>` (default from `.env` is `8080`).

## API Endpoints (Current)
- `GET /v1/healthy`  
  Returns `200` with `{}`.

- `GET /v1/err`  
  Returns `400` with:

```json
{ "error": "Something went wrong" }
```

- `POST /v1/users`  
  Request body:

```json
{ "name": "Leela" }
```

  Returns created user object including generated `api_key`.

## Status
This is the foundation phase of the RSS aggregator. Core user and infrastructure pieces are in place; feed ingestion, follow/unfollow flows, and aggregation endpoints are expected next.
