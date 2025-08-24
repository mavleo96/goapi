# Go API Service

A simple REST API service built in Go for checking user coin balances with authentication middleware and health monitoring capabilities.

## Features

- **Coin Balance API** - Retrieve user coin balances with authentication
- **Health Check Endpoint** - Service health monitoring for load balancers
- **Authentication Middleware** - Token-based authentication system
- **Mock Database** - In-memory database for development and testing
- **Comprehensive Testing** - Unit tests for all components
- **Clean Architecture** - Well-structured Go project with clear separation of concerns

## Project Structure

```
goapi/
├── api/                    # API types and error handling
│   ├── api.go
│   └── api_test.go
├── cmd/api/               # Application entry point
│   └── main.go
├── internal/              # Internal application code
│   ├── handlers/          # HTTP request handlers
│   │   ├── api.go
│   │   ├── api_test.go
│   │   ├── get_coin_balance.go
│   │   ├── get_coin_balance_test.go
│   │   ├── health.go
│   │   └── health_test.go
│   ├── middleware/        # HTTP middleware
│   │   ├── authorization.go
│   │   └── authorization_test.go
│   └── tools/            # Database interface and utilities
│       ├── database.go
│       ├── database_test.go
│       └── mockdb.go
├── go.mod
├── go.sum
└── README.md
```

## Prerequisites

- Go 1.25.0 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/mavleo96/goapi.git
cd goapi
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

Start the API server:
```bash
go run cmd/api/main.go
```

The server will start on `localhost:8000` and display a cool ASCII art banner.

## API Endpoints

### Health Check

**GET** `/health`

Returns service health status. No authentication required.

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:45.123456Z",
  "service": "goapi",
  "version": "1.0.0"
}
```

**Example:**
```bash
curl http://localhost:8000/health
```

### Get Coin Balance

**GET** `/account/coins?username=<username>`

Retrieves the coin balance for a specific user. Requires authentication.

**Headers:**
- `Authorization: <token>` - User's authentication token

**Query Parameters:**
- `username` - The username to query balance for

**Response:**
```json
{
  "Code": 200,
  "Balance": 100
}
```

**Example:**
```bash
curl "http://localhost:8000/account/coins?username=alex" \
  -H "Authorization: 123ABC"
```

## Authentication

The API uses token-based authentication. Valid users and their tokens are:

| Username | Token  | Coin Balance |
|----------|--------|--------------|
| alex     | 123ABC | 100          |
| jason    | 456DEF | 200          |
| marie    | 789GHI | 300          |

## Error Responses

All errors follow a standardized format:

```json
{
  "Code": 400,
  "Message": "Error description"
}
```

Common error codes:
- `400` - Bad Request (invalid parameters, missing authentication)
- `500` - Internal Server Error (server-side errors)

## Testing

Run all tests:
```bash
go test ./...
```

Run tests for specific packages:
```bash
go test ./api
go test ./internal/handlers
go test ./internal/middleware
go test ./internal/tools
```

Run tests with verbose output:
```bash
go test -v ./...
```

## Development

### Project Dependencies

- **Chi** - HTTP router and middleware
- **Gorilla Schema** - Query parameter parsing
- **Logrus** - Structured logging

### Code Quality

Run code quality checks:
```bash
go vet ./...
go fmt ./...
```
