# Car Listing Service

A RESTful microservice for managing car listings, built with Go Fiber, GORM, and PostgreSQL.

## Features

- **Fiber v2** - Fast HTTP framework for Go
- **GORM** - Powerful ORM for database operations
- **PostgreSQL** - Robust relational database
- **Hot Reload** - Air for development live-reloading
- **Environment Config** - Flexible configuration via .env
- **RESTful API** - Standard REST endpoints for CRUD operations
- **Middleware** - CORS, logging, and recovery built-in

## Tech Stack

- **Language**: Go
- **Framework**: Fiber v2
- **ORM**: GORM
- **Database**: PostgreSQL
- **Dev Tool**: Air (live reload)

## Environment Variables

Create a `.env` file in the root directory (use `.env.example` as a template):

```env
# Server Configuration
PORT=3002

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=car_listing
DB_SSLMODE=disable
```

## Installation

```bash
# Install dependencies
go mod download

# Install Air for live reload (optional)
go install github.com/air-verse/air@latest

# Copy environment variables
cp .env.example .env

# Update .env with your database credentials
```

## Database Setup

1. Create a PostgreSQL database:
```sql
CREATE DATABASE car_listing;
```

2. The application will automatically connect using the credentials in your `.env` file

3. You can run migrations to set up your schema (add migration files as needed)

## Development

Run with live reload using Air:

```bash
~/go/bin/air
```

Or run directly with Go:

```bash
go run main.go
```

The server will start on port 3002 (or your configured PORT).

## Production Build

```bash
# Build the binary
go build -o car-listing-service

# Run the binary
./car-listing-service
```

## API Endpoints

Base path: `/api/cars`

### Health Check
- `GET /health` - Service health status

### Car Listings (Example - extend as needed)
- `GET /api/cars` - List all cars
- `GET /api/cars/:id` - Get a specific car
- `POST /api/cars` - Create a new car listing
- `PUT /api/cars/:id` - Update a car listing
- `DELETE /api/cars/:id` - Delete a car listing

## Project Structure

```
car-listing-service/
├── database/
│   └── database.go       # Database connection & config
├── models/               # Database models (add as needed)
├── handlers/             # HTTP request handlers (add as needed)
├── routes/               # Route definitions (add as needed)
├── main.go               # Application entry point
├── go.mod                # Go module file
├── .air.toml             # Air configuration
├── .env.example          # Environment variables template
├── .gitignore
└── README.md
```

## Database Connection

The database package (`database/database.go`) provides:
- Automatic connection pooling
- Environment-based configuration
- Error handling and logging
- Connection instance export

Access the database instance:
```go
import "car-listing-service/database"

db := database.GetDB()
```

## Integration

This service is designed to work with the API Gateway. All requests should be routed through:

```
Client → API Gateway → Car Listing Service
```

The gateway proxies requests from `/api/car-listing/*` to this service.

## Adding Models

Create your car model in `models/car.go`:

```go
package models

import "gorm.io/gorm"

type Car struct {
    gorm.Model
    Make        string  `json:"make"`
    Model       string  `json:"model"`
    Year        int     `json:"year"`
    Price       float64 `json:"price"`
    Description string  `json:"description"`
}
```

Then migrate in `main.go`:
```go
database.DB.AutoMigrate(&models.Car{})
```

## Error Handling

The service includes:
- Panic recovery middleware
- Structured error responses
- Database connection error handling
- Request validation (add as needed)

## Logging

Request logging is enabled by default using Fiber's logger middleware. All requests are logged with:
- HTTP method
- Path
- Status code
- Response time
