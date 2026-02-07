# Go Web API

A lightweight RESTful API server built with Go and the Gin framework. This project demonstrates how to build a production-ready web API with proper structure, configuration management, and CRUD operations.

## Features

- **RESTful Endpoints**: Complete CRUD operations for album management
- **Gin Framework**: Fast and efficient HTTP routing and middleware
- **Configuration Management**: Centralized config loading for easy deployment
- **JSON Support**: Automatic JSON marshalling and unmarshalling
- **Error Handling**: Proper HTTP status codes and error responses

## Architecture

The application follows a clean, modular structure:

- `main.go` - Application entry point
- `webapi/config.go` - Configuration management
- `webapi/album.go` - Data model definition
- `webapi/handlers.go` - HTTP request handlers and route registration
- `webapi/handlers_test.go` - Unit tests for handlers

## Getting Started

### Prerequisites

- Go 1.25 or higher
- Gin framework (automatically handled by `go mod`)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/abishen/webapi.git
cd webapi
```

2. Download dependencies:
```bash
go mod download
```

### Running the API

Start the server:
```bash
go run main.go
```

The API will start on the configured server address (default: `http://localhost:8080`).

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/album` | Retrieve all albums |
| `GET` | `/album/:id` | Get a specific album by ID |
| `POST` | `/album` | Create a new album |
| `DELETE` | `/album/:id` | Delete an album |

### Example Requests

**Get all albums:**
```bash
curl http://localhost:8080/album
```

**Get album by ID:**
```bash
curl http://localhost:8080/album/1
```

**Create a new album:**
```bash
curl -X POST http://localhost:8080/album \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Blue","artist":"Joni Mitchell","price":19.99}'
```

**Delete an album:**
```bash
curl -X DELETE http://localhost:8080/album/1
```

## Project Structure

```
webapi/
├── main.go              # Entry point
├── go.mod              # Module definition
├── go.sum              # Dependency checksums
├── README.md           # This file
└── webapi/
    ├── album.go        # Album model
    ├── handlers.go     # Route handlers
    ├── handlers_test.go # Handler tests
    └── config.go       # Configuration
```

## Testing

Run the test suite:
```bash
go test ./...
```

## UML Class Diagram

The API follows a layered architecture with clear separation of concerns:

```mermaid
classDiagram
    class Album {
        -ID: string
        -Title: string
        -Artist: string
        -Price: float64
    }
    
    class Handlers {
        +getAlbums(c *gin.Context)
        +getAlbumByID(c *gin.Context)
        +postAlbum(c *gin.Context)
        +deleteAlbum(c *gin.Context)
        +RegisterRoutes(router *gin.Engine)
    }
    
    class Config {
        -ServerAddr: string
        +LoadConfig() Config
    }
    
    class App {
        -config: Config
        -router: gin.Engine
        +main()
    }
    
    Handlers --> Album: manages
    App --> Config: loads
    App --> Handlers: registers routes
    Handlers --> "gin.Context": processes
```

## Component Interaction Flow

1. **Application Start**: `main.go` loads configuration and initializes the Gin router
2. **Route Registration**: `RegisterRoutes()` sets up HTTP endpoints
3. **Request Handling**: Gin routes incoming requests to appropriate handler functions
4. **Data Processing**: Handlers manipulate the in-memory album slice
5. **Response**: Results are JSON-encoded and returned to the client

## Dependencies

- **github.com/gin-gonic/gin** - High-performance HTTP web framework
- **github.com/stretchr/testify** - Testing assertions and utilities

## Building and Running

### Build the application:
```bash
go build -o webapi
./webapi
```

### Run tests with coverage:
```bash
go test -cover ./...
```

### Build cross-platform:
```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o webapi-linux

# macOS
GOOS=darwin GOARCH=amd64 go build -o webapi-macos

# Windows
GOOS=windows GOARCH=amd64 go build -o webapi.exe
```

## Development

The current implementation stores albums in memory. For a production system, consider:

- Adding a database layer (PostgreSQL, MongoDB, etc.)
- Implementing proper error handling and logging
- Adding authentication and authorization
- Implementing caching strategies
- Adding request validation and sanitization

## License

This project is open source and available under the MIT License.
