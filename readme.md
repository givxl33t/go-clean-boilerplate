# Go Clean Architecture Boilerplate
This repository serves as a boilerplate for building a Go project using the **Clean Architecture** principles and a commonly adopted Golang project layout.

## Features
This boilerplate includes the following libraries and tools:
- [**Fiber**](https://github.com/gofiber/fiber) – Web framework for building APIs.
- [**Viper**](https://github.com/spf13/viper) – Configuration management library.
- [**GORM**](https://github.com/go-gorm/gorm) – ORM library for database operations.
- [**GoMock**](https://github.com/uber/mock) – Mocking framework for testing.
- [**Testify**](https://github.com/stretchr/testify) – Toolkit for assertions in tests.
- [**Validator**](https://github.com/go-playground/validator) – Struct and field validation library.
- [**Migrate**](https://github.com/golang-migrate/migrate) – Database migration tool.

## Project Structure
This project follows a modular and clean architecture design. Below is an explanation of its directory structure:

### Core Directories
- **`cmd/`**: Contains the main entry point(s) of the application, e.g., `main.go`.
- **`config/`**: Stores configuration files and related logic.
- **`docs/`**: Holds project documentation files.
- **`internal/`**: Contains internal-only application code.

### Application-Specific Directories
- **`infrastructure/`**: Frameworks and driver layer (e.g., database connection, logging).
- **`delivery/`**: Contains data delivery mechanisms such as HTTP handlers or RPC implementations.
- **`domain/`**: Defines core domain models and entities.
- **`repository/`**: Implements data access (data layer) logic.
- **`usecase/`**: Contains business logic and application use cases.
- **`mapper/`**: Handles mapping between different data structures.
- **`model/`**: Contains request/response models and shared data structures.
- **`exception/`**: Manages error handling across the application.

### Testing
- **`test/`**: Contains unit tests, integration tests, and test utilities.

## Installations
```
git clone https://github.com/givxl33t/go-clean-boilerplate.git
```

```
rm -rf .git/
```

```
cp .env.example .env
```

or using [gonew](https://pkg.go.dev/golang.org/x/tools/cmd/gonew)

```
gonew github.com/givxl33t/go-clean-boilerplate github.com/<username>/<your-repo>
```
## Run
```
go run cmd/web/main.go
```
## Testing

### Run Unit Test
```
make test.unit
```

### Run Integration Test
```
make test.integration
```

## Database Migration

### Create
```
make migrate.create name=create_table_users
```

### Up
```
make migrate.up
```

### Down
```
make migrate.down
```

### Force
```
make migrate.force version=20231216100
```