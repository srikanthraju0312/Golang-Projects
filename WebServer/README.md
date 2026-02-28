# Go Web Form Server

This project is a simple Go web server that serves an HTML form and stores submitted user data in both JSON and CSV formats.

## What the project does

- Serves a form at `/` from `static/index.html`
- Accepts `POST` submissions at `/submit`
- Captures `name` and `email` fields
- Persists each submission to:
  - `data.json` (pretty-formatted JSON array)
  - `data.csv` (CSV rows with header)
- Uses a mutex to keep file writes safe for concurrent requests

## Project structure

```text
.
|- main.go
|- static/
|  `- index.html
|- data.json
`- data.csv
```

## How to run

1. Make sure Go is installed.
2. From the project root, run:

```bash
go run main.go
```

3. Open your browser and visit:

```text
http://localhost:8080
```

## HTTP routes

- `GET /` -> Serves the input form
- `POST /submit` -> Saves form data to JSON and CSV, then returns a success page

## Data format

JSON entry example:

```json
{
  "name": "Alice",
  "email": "alice@example.com"
}
```

CSV row example:

```csv
Name,Email
Alice,alice@example.com
```

## Current status

The project currently demonstrates end-to-end form submission and file persistence, suitable as a starter web server/data capture example in Go.
