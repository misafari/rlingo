# Rlingo - Translation Management System (TMS)

A modern, multi-tenant Translation Management System built with Go, providing REST APIs, background processing, and a web-based admin panel.

## Project Structure

```
├── cmd/
│   ├── api/           → main entry for REST API
│   └── worker/        → background jobs, async processing
│
├── internal/
│   ├── app/           → use-cases (application logic)
│   ├── domain/        → entities, interfaces (pure Go)
│   ├── infra/         → adapters: db, cache, storage, search, auth
│   ├── http/          → handlers, routes, middleware
│   ├── config/        → env + config management
│   ├── security/      → JWT, RBAC, tenant context
│   ├── job/           → async workers, cron, redis streams
│   ├── tenant/        → multi-tenancy middleware + resolver
│   ├── tests/         → integration & e2e tests
│   └── utils/         → logging, errors, tracing, validation
│
├── pkg/               → reusable libraries (optional)
├── api/               → OpenAPI spec, JSON schemas
├── web/               → React admin panel (frontend)
├── migrations/        → SQL migrations
├── docker/            → Docker & Compose files
└── Makefile
```

## Architecture

This project follows **Clean Architecture** principles with a clear separation of concerns:

- **Domain Layer** (`internal/domain/`): Core business entities and interfaces (pure Go, no external dependencies)
- **Application Layer** (`internal/app/`): Use cases and business logic orchestration
- **Infrastructure Layer** (`internal/infra/`): External adapters (database, cache, storage, search engines, authentication providers)
- **Interface Layer** (`internal/http/`): HTTP handlers, routes, and middleware

## Features

- 🏢 **Multi-tenancy**: Tenant isolation and context management
- 🔐 **Security**: JWT-based authentication, RBAC authorization
- ⚡ **Async Processing**: Background workers for translation jobs
- 🌐 **REST API**: Comprehensive API for translation management
- 🎨 **Web Admin**: React-based admin panel
- 🔄 **Database Migrations**: Version-controlled schema changes

## Getting Started

### Prerequisites

- Go 1.25.3 or later
- Docker and Docker Compose (for local development)
- PostgreSQL (for database)
- Redis (for caching and job queues)

### Installation

```bash
# Clone the repository
git clone https://github.com/misafari/rlingo.git
cd rlingo

# Install dependencies
go mod download

# Run database migrations
make migrate-up

# Start services with Docker Compose
docker-compose up -d

# Run the API server
make run-api

# Run the worker
make run-worker
```

## Development

### Running Tests

```bash
make test
```

### Building

```bash
# Build API
make build-api

# Build Worker
make build-worker
```

## License

See [LICENSE](LICENSE) file for details.

