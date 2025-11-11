# Go Web Server

<br>

A low-level Go web server for performance. It uses minimal dependencies, enabling small application size and faster deployments. Due to less reliability on third-party code, we get more control and flexibility over the entire codebase. That means more test coverage, less bugs and faster APIs.

<br>

### Features

- JWT: Create, verify and refresh JWT tokens.
- sqlc: Interact with the database using familiar & performant SQL queries.
- Authorization: Manage access to sensitive resources.
- Middlewares: For request monitoring, logging & access control.

<br>

### Project Structure

```
- cmd/
    - server/
        - main.go              // Application entry point
- config/
    - constants.go             // Constants
    - environments.go          // Environment variables
- internal/
    - apis/
        - routes.go            // API routes
    - core/                    // Core functionalities
        - api.go               // API helpers
        - api_test.go          // Unit test cases for API helpers
    - handler/                 // API handlers
        - auth_handler.go
        - core_handler.go
        - product_handler.go
    - jwt/                     // JWT feature
        - model.go
        - service.go
        - repository.go
        - jwt_repository.go    // JWT repository implementation
    - product/                 // Product feature
        - model.go
        - service.go
        - repository.go
        - sql_repository.go    // Repository implementation using SQL
    - user/                    // User feature
        - model.go
        - service.go
        - repository.go
        - sql_repository.go    // Repository implementation using SQL
    - utils/
        - handlers.go          // Initialize API handlers
- .env.example                 // Sample .env file
- Dockerfile                   // Multi-stage dockerfile
- .gitignore
- go.mod
- go.sum
- README.md
```

<br>

### Command to start server

```bash
go run ./cmd/server/
```
