# Go Web API
![Go Web API](https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Aqua.svg)
This is a simple Go web API that demonstrates how to create a RESTful API using the `net/http` package. The API provides endpoints for managing a list of items, allowing you to create, read, update, and delete items.

## Features
- Create a new item
- Retrieve all items
- Retrieve a specific item by ID
- Update an existing item
- Delete an item
## Getting Started
To get started with this API, follow the instructions below.
### Prerequisites
- Go installed on your machine (version 1.16 or higher)
### Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/go-web-api.git
```
2. Navigate to the project directory:
```bash
cd go-web-api
```
3. Install dependencies (if any):
```bash
go mod tidy
```
### Running the API
To run the API, use the following command:
```bash
go run main.go
```
The API will start on `http://localhost:8080`.  You can use tools like `curl` or Postman to interact with the API endpoints.
- `GET /albums`: Retrieve all albums
- `GET /albums/{id}`: Retrieve a specific album by ID
- `POST /albums`: Create a new album
- `PUT /albums/{id}`: Update an existing album
- `DELETE /albums/{id}`: Delete an album

// create uml class diagram for the API

```plantuml
@startuml
class Album {
    +ID: int
    +Title: string
    +Artist: string
    +Price: float
}
class AlbumHandler {
    +GetAlbums(w http.ResponseWriter, r *http.Request)
    +GetAlbumByID(w http.ResponseWriter, r *http.Request)
    +CreateAlbum(w http.ResponseWriter, r *http.Request)
    +UpdateAlbum(w http.ResponseWriter, r *http.Request)
    +DeleteAlbum(w http.ResponseWriter, r *http.Request)
}
AlbumHandler --> Album : manages
@enduml
```This UML class diagram represents the structure of the Go web API. The `Album` class represents the data model for an album, while the `AlbumHandler` class contains methods for handling HTTP requests related to albums. The `AlbumHandler` class manages the `Album` class, indicating that it performs operations on the album data.