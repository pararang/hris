# Attendance and Payslip Generation System

A Go backend service for payslip generation following Clean Architecture principles.

## Project Structure

```
hris/
├── cmd/                  # Application entry points
│   └── dbseed/           # Database seeder
├── config/               # Configuration files
├── dbase/                # Database related files
│   ├── migrations/       # Database migrations
│   └── seeders/          # Database seeders
├── domain/               # Enterprise business rules
│   ├── repository/       # Repository interfaces
│   └── usecase/          # Use case interfaces
├── dto/                  # Data Transfer Objects
├── entity/               # Business entities
├── libs/                 # Shared libraries
│   ├── auth/             # Authentication utilities
│   ├── httpresp/         # HTTP response utilities
│   └── logger/           # Logging utilities
├── mocks/                # Mock implementations for testing
├── repository/           # Repository implementations
├── rest/                 # REST API
│   ├── handler/          # HTTP handlers
│   └── middleware/       # HTTP middlewares
└── usecase/              # Application business rules
```

## Getting Started

### Prerequisites

- Go 1.23.11+
- PostgreSQL 14+
- Goose for migrations db schema https://github.com/pressly/goose
- air https://github.com/cosmtrek/air for development hot reloading - optional

### Environment Setup

1. Copy the example environment file and adjust as needed:
   ```bash
   cp .env.example .env
   ```

2. Configure your environment variables in the `.env` file:
   ```
   ENV=development

   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=password
   DB_NAME=hris

   APP_PORT=8080
   API_KEY=your_api_key

   JWT_SECRET=your_jwt_secret
   JWT_EXPIRY_MINUTES=60

   GOOSE_DRIVER=postgres
   GOOSE_DBSTRING=postgres://postgres:password@localhost:5432/hris
   GOOSE_MIGRATION_DIR=./dbase/migrations
   GOOSE_TABLE=goose_migrations
   ```

### Database Setup

1. Create a PostgreSQL database:
   ```bash
   createdb hris
   ```

2. Run database migrations:

   ```bash
   goose up
   ```

3. (Optional) Seed the database with initial data:
   ```bash
   make seed
   ```

### Running the Application

#### Using Make

```bash
# Install dependencies
go mod tidy

# Run the application
make run

# (Optional) Run with hot reloading
air
```

#### Using Go Directly

```bash
# Install dependencies
go mod tidy

# Run the application
go run cmd/main.go
```

The API will be available at `http://localhost:8080`.

### API Endpoints

#### Authentication
- `POST /v1/login` - User login (employees and admins)

#### Employee Endpoints
- `POST /v1/attendances/clockin` - Clock in for the day
- `POST /v1/attendances/clockout` - Clock out for the day
- `POST /v1/overtimes` - Submit overtime request
- `POST /v1/reimbursements` - Submit reimbursement request
- `GET /v1/payslips` - List employee's payslips
- `GET /v1/payslips/:id` - Get specific payslip details

#### Admin Endpoints
- `POST /v1/attendances/periods` - Create a new attendance period
- `POST /v1/payrolls` - Process payroll and generate payslips
- `GET /v1/payrolls/periods/:id` - Get payroll summary for a specific period

#### Health Check
- `GET /v1/health` - API health check endpoint

### Postman Collection

A Postman collection is included in this repository to help you test the API endpoints. You can import the collection into Postman using the following file:

- [HRIS Postman Collection](hris.postman_collection.json)

To use the collection:
1. Download [Postman](https://www.postman.com/downloads/)
2. Import the collection file (`hris.postman_collection.json`)
3. Set up environment variables for:
   - `base_url`: The base URL of your API (e.g., `http://localhost:8080`)
   - `token`: Authentication token (will be automatically set after login)

## Architecture

This project follows Clean Architecture principles with clear separation of concerns:

### Layer Structure

1. **Domain Layer** (`/domain`): Contains business entities and interfaces (contracts)
   - Repository Interfaces: Define data access contracts (`/domain/repository`)
   - Use Case Interfaces: Define business logic contracts (`/domain/usecase`)

2. **Entity Layer** (`/entity`): Core business objects
   - Employee, Attendance, Overtime, Reimbursement, Payslip, etc.

3. **Use Case Layer** (`/usecase`): Implements application-specific business rules
   - Implements the use case interfaces defined in the domain layer
   - Orchestrates the flow of data to and from entities
   - Contains application-specific business rules

4. **Repository Layer** (`/repository`): Provides concrete implementations for repository interfaces
   - Database implementations for UserRepository, AttendanceRepository, etc.
   - Handles data persistence and retrieval

5. **Interface Layer** (`/rest`): Handles external concerns
   - HTTP handlers (`/rest/handler`): Process HTTP requests and responses
   - Middleware (`/rest/middleware`): Authentication, logging, etc.

6. **DTO Layer** (`/dto`): Data Transfer Objects
   - Request/response structures for API communication
   - Validation rules for incoming data

7. **Libraries** (`/libs`): Shared utilities and services
   - Authentication services (`/libs/auth`)
   - HTTP response utilities (`/libs/httpresp`)
   - Logging utilities (`/libs/logger`)

### Flow of Control

1. HTTP request comes in through the router (`/rest/router.go`)
2. Router directs the request to the appropriate handler
3. Handler parses the request and calls the appropriate use case
4. Use case implements business logic, using repositories to access data
5. Repositories interact with the database to persist or retrieve data
6. Response flows back through the layers to the client

### Key Components

#### Entities
- **Employee**: User with employee role
- **AttendancePeriod**: Period for which attendance is tracked
- **Attendance**: Employee's attendance record for a specific date
- **Overtime**: Overtime record for an employee
- **Reimbursement**: Reimbursement request from an employee
- **Payslip**: Employee's payslip for a specific period

#### Use Cases
- **UserUseCase**: Authentication and user management
- **AttendanceUseCase**: Clock in/out and attendance period management
- **OvertimeUseCase**: Overtime request submission and approval
- **ReimbursementUseCase**: Reimbursement request submission and approval
- **PayslipUseCase**: Payslip generation and retrieval

#### API Handlers
- **UserHandler**: Login and user management
- **AttendanceHandler**: Clock in/out and attendance period creation
- **OvertimeHandler**: Overtime request submission
- **ReimbursementHandler**: Reimbursement request submission
- **PayslipHandler**: Payslip generation and retrieval


## Database Schema

The following Entity Relationship Diagram (ERD) illustrates the database schema used in this application:

![HRIS Entity Relationship Diagram](hris-erd.png)

## Development

### Makefile Commands

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test

# Format code
make fmt

# Tidy dependencies
make tidy
```
