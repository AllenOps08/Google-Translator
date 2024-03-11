
# Movie API

This is a simple Go application that implements a RESTful API for managing movies.

## Features

- Get a list of all movies
- Get details of a specific movie by ID
- Add a new movie
- Update an existing movie
- Delete a movie

## Requirements

- Go 1.11 or higher
- Gorilla Mux (router and dispatcher for Go HTTP services)

## Installation

1. Clone the repository:

```
git clone https://github.com/yourusername/movie-api.git
```

2. Navigate to the project directory:

```
cd movie-api
```

3. Install dependencies:

```
go get -u github.com/gorilla/mux
```

## Usage

1. Run the application:

```
go run main.go
```

2. The server will start running at `http://localhost:8000`.

## Endpoints

- **GET /movies**: Get a list of all movies.
- **GET /movies/{id}**: Get details of a specific movie by ID.
- **POST /movies**: Add a new movie.
- **PUT /movies/{id}**: Update an existing movie.
- **DELETE /movies/{id}**: Delete a movie.

## Sample Request and Response

### Request:

```
POST /movies

{
    "isbn": "1234567890",
    "title": "Sample Movie",
    "director": {
        "firstname": "John",
        "lastname": "Doe"
    }
}
```

### Response:

```
{
    "id": "123456789",
    "isbn": "1234567890",
    "title": "Sample Movie",
    "director": {
        "firstname": "John",
        "lastname": "Doe"
    }
}
```

