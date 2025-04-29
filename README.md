# Approval API

A Golang CRUD API for managing approval workflows with PostgreSQL database.

## Features

- RESTful API endpoints for approval management
- PostgreSQL database integration
- Database migrations using Goose
- Proper layered architecture (handlers, repositories, models)
- UUID primary keys
- JSONB column support for comments
- Timestamp tracking (created_at, updated_at)

## Prerequisites

- Go 1.20+
- PostgreSQL 12+
- Make (optional)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/approval-api.git
cd approval-api
```

### 2. Set up environment variables
Create a .env file in the root directory:

```bash
env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=approval_db
SERVER_PORT=8080
```

### 3. Install dependencies
```bash
go mod download
```

### 4. Database Setup
Ensure PostgreSQL is running, then create the database:

```bash
psql -U postgres -c "CREATE DATABASE approval_db;"
```

### 5. Run migrations
```bash
go build -o migrate ./migrations
./migrate up
```

