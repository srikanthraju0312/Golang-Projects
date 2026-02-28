# MongoGolang

A minimal REST API in Go for creating, retrieving, and deleting users in MongoDB.

## Overview

This project uses:
- `net/http` for the HTTP server
- `github.com/julienschmidt/httprouter` for route handling
- `gopkg.in/mgo.v2` (mgo) as the MongoDB driver

The server runs on `:8080` and connects to MongoDB at `localhost:27017`.

## Project Structure

- `main.go`: Application entrypoint, router setup, and MongoDB session creation.
- `controllers/user.go`: User handlers (`GET`, `POST`, `DELETE`).
- `models/user.go`: `User` model and BSON/JSON mappings.

## Data Model

User fields:
- `id` (`ObjectId`) -> stored as Mongo `_id`
- `name` (`string`)
- `age` (`int`)
- `gender` (`string`)

Example JSON response:

```json
{
  "id": "65f0a1b2c3d4e5f678901234",
  "name": "Leela",
  "age": 26,
  "gender": "female"
}
```

## Prerequisites

- Go installed
- MongoDB running locally on `localhost:27017`

## Run

From the `MongoGolang` directory:

```bash
go run .
```

Server starts at:

```text
http://localhost:8080
```

## API Endpoints

### Create user

- Method: `POST`
- Path: `/user`
- Body: JSON with `name`, `age`, `gender`

Example:

```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Leela","age":26,"gender":"female"}'
```

Returns `201 Created` with the created user including generated `id`.

### Get user by id

- Method: `GET`
- Path: `/user/:id`

Example:

```bash
curl http://localhost:8080/user/<object_id>
```

Returns `200 OK` with user JSON if found, otherwise `404`.

### Delete user by id

- Method: `DELETE`
- Path: `/user/:id`

Example:

```bash
curl -X DELETE http://localhost:8080/user/<object_id>
```

Returns `200 OK` with a text confirmation if deleted, otherwise `404`.

## MongoDB Details

This code currently uses:
- Database: `test`
- Collection: `users`

## Notes

- The app uses the legacy `mgo` driver (`gopkg.in/mgo.v2`).
- `id` in routes must be a valid MongoDB ObjectId hex string.
