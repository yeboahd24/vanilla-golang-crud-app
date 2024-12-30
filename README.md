# Go CRUD Application

A simple CRUD (Create, Read, Update, Delete) application built with Go, using clean architecture principles and PostgreSQL as the database.

## Features

- Clean Architecture (Handler → Service → Repository)
- PostgreSQL Database
- SQLC for type-safe database operations
- Interface-based design
- Configuration using Viper
- Database migrations
- RESTful API endpoints

## Project Structure

```
crud-app/
├── config/         # Configuration management
├── handler/        # HTTP handlers
├── service/        # Business logic
├── repository/     # Data access layer
├── model/          # Domain models
├── db/
│   ├── migration/  # Database migrations
│   ├── query/      # SQLC queries
│   └── sqlc/       # Generated SQLC code
└── main.go         # Application entry point
```

## Prerequisites

- Go 1.16 or higher
- PostgreSQL
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [sqlc](https://sqlc.dev/)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd crud-app
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory:
```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://postgres:your_password@localhost:5432/crud?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
```

4. Create the database:
```bash
createdb crud
```

5. Run database migrations:
```bash
migrate -path db/migration -database "postgresql://postgres:your_password@localhost:5432/crud?sslmode=disable" up
```

## Running the Application

```bash
go run main.go
```

## API Endpoints

### Users

- **Create User**
  ```http
  POST /users
  Content-Type: application/json

  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```

- **Get User**
  ```http
  GET /users?id=1
  ```

- **Update User**
  ```http
  PUT /users?id=1
  Content-Type: application/json

  {
    "name": "John Updated",
    "email": "john.updated@example.com"
  }
  ```

- **Delete User**
  ```http
  DELETE /users?id=1
  ```

- **List Users**
  ```http
  GET /users
  ```

## Development

### Generate SQLC Code

After modifying queries in `db/query/`:
```bash
sqlc generate
```

### Database Migrations

Create a new migration:
```bash
migrate create -ext sql -dir db/migration -seq <migration_name>
```

Run migrations:
```bash
migrate -path db/migration -database "postgresql://postgres:your_password@localhost:5432/crud?sslmode=disable" up
```

Rollback migrations:
```bash
migrate -path db/migration -database "postgresql://postgres:your_password@localhost:5432/crud?sslmode=disable" down
```

## Architecture

This project follows clean architecture principles:

1. **Handler Layer** (`handler/`)
   - Handles HTTP requests and responses
   - Input validation
   - Route management

2. **Service Layer** (`service/`)
   - Business logic
   - Use case implementation
   - Domain rules

3. **Repository Layer** (`repository/`)
   - Data access abstraction
   - Database operations
   - Uses SQLC for type-safe queries

## Testing

Run tests:
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
