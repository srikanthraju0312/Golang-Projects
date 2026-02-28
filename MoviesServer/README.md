# MoviesServer

MoviesServer is a simple Go REST API for managing a small in-memory list of movies.

It exposes CRUD endpoints using Gorilla Mux:
- `GET /movies` to list all movies
- `GET /movies/{id}` to fetch one movie
- `POST /movies` to create a movie
- `PUT /movies/{id}` to update a movie
- `DELETE /movies/{id}` to delete a movie

The server starts on port `8000` and seeds a couple of sample movies at startup.
