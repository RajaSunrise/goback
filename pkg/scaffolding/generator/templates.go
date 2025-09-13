// pkg/scaffolding/generator/templates.go

package generator

// Basic project templates
const goModTemplate = `module {{.ModulePath}}

go 1.21

require (
{{- if eq .Framework "fiber"}}
	github.com/gofiber/fiber/v2 v2.52.0
{{- else if eq .Framework "gin"}}
	github.com/gin-gonic/gin v1.9.1
{{- else if eq .Framework "chi"}}
	github.com/go-chi/chi/v5 v5.0.10
{{- else if eq .Framework "echo"}}
	github.com/labstack/echo/v4 v4.11.4
{{- end}}
{{- if eq .Database "postgresql"}}
	github.com/lib/pq v1.10.9
{{- else if eq .Database "mysql"}}
	github.com/go-sql-driver/mysql v1.7.1
{{- else if eq .Database "sqlite"}}
	github.com/mattn/go-sqlite3 v1.14.18
{{- end}}
{{- if eq .ORM "gorm"}}
	gorm.io/gorm v1.25.5
	{{- if eq .Database "postgresql"}}
	gorm.io/driver/postgres v1.5.4
	{{- else if eq .Database "mysql"}}
	gorm.io/driver/mysql v1.5.2
	{{- else if eq .Database "sqlite"}}
	gorm.io/driver/sqlite v1.5.4
	{{- end}}
{{- else if eq .ORM "sqlx"}}
	github.com/jmoiron/sqlx v1.3.5
{{- end}}
	github.com/joho/godotenv v1.4.0
	gopkg.in/yaml.v3 v3.0.1
)
`

const readmeTemplate = `# {{.ProjectName}}

{{if .Description}}{{.Description}}{{else}}A backend API built with GoBack scaffolding tool{{end}}

## Features

- **Framework**: {{.Framework}}
- **Database**: {{.Database}}
- **ORM**: {{.ORM}}
- **Architecture**: {{.Architecture}}
{{- if .DevOps.Enabled}}
- **DevOps**: {{range .DevOps.Tools}}{{.}} {{end}}
{{- end}}

## Getting Started

### Prerequisites

- Go 1.21 or higher
{{- if eq .Database "postgresql"}}
- PostgreSQL
{{- else if eq .Database "mysql"}}
- MySQL
{{- end}}

### Installation

1. Clone the repository:
   ` + "`" + `` + "`" + `` + "`" + `bash
   git clone <your-repo-url>
   cd {{.ProjectName}}
   ` + "`" + `` + "`" + `` + "`" + `

2. Install dependencies:
   ` + "`" + `` + "`" + `` + "`" + `bash
   go mod tidy
   ` + "`" + `` + "`" + `` + "`" + `

3. Copy environment file:
   ` + "`" + `` + "`" + `` + "`" + `bash
   cp .env.example .env
   ` + "`" + `` + "`" + `` + "`" + `

4. Configure your environment variables in ` + "`" + `.env` + "`" + `

### Running the Application

#### Development
` + "`" + `` + "`" + `` + "`" + `bash
make run
` + "`" + `` + "`" + `` + "`" + `

#### Production
` + "`" + `` + "`" + `` + "`" + `bash
make build
./bin/{{.ProjectName}}
` + "`" + `` + "`" + `` + "`" + `

## API Documentation

The API will be available at ` + "`" + `http://localhost:8080` + "`" + `

### Health Check
- **GET** ` + "`" + `/health` + "`" + ` - Health check endpoint

## Project Structure

` + "`" + `` + "`" + `` + "`" + `
{{.ProjectName}}/
├── cmd/                    # Application entrypoints
├── internal/               # Private application code
{{- if eq .Architecture "simple"}}
│   ├── handlers/          # HTTP handlers
│   ├── services/          # Business logic
│   ├── models/            # Data models
│   └── database/          # Database configuration
{{- else if eq .Architecture "ddd"}}
│   ├── domain/            # Domain layer
│   ├── application/       # Application layer
│   ├── infrastructure/    # Infrastructure layer
│   └── interfaces/        # Interface layer
{{- else if eq .Architecture "clean"}}
│   ├── entities/          # Business entities
│   ├── usecases/          # Business logic
│   ├── interface_adapters/ # Interface adapters
│   └── frameworks_drivers/ # Frameworks and drivers
{{- else if eq .Architecture "hexagonal"}}
│   ├── domain/            # Domain models and services
│   ├── ports/             # Input and output ports
│   └── adapters/          # Input and output adapters
{{- end}}
├── pkg/                   # Public library code
├── configs/               # Configuration files
├── scripts/               # Build and deployment scripts
{{- if .DevOps.Enabled}}
├── deployments/           # Deployment configurations
{{- if .DevOps.Kubernetes}}
│   ├── kubernetes/        # Kubernetes manifests
{{- end}}
{{- if .DevOps.Helm}}
│   ├── helm/              # Helm charts
{{- end}}
{{- if .DevOps.Terraform}}
│   ├── terraform/         # Terraform configurations
{{- end}}
{{- if .DevOps.Ansible}}
│   └── ansible/           # Ansible playbooks
{{- end}}
{{- end}}
├── Makefile               # Build automation
└── README.md              # This file
` + "`" + `` + "`" + `` + "`" + `

## Development

### Make Commands

- ` + "`" + `make run` + "`" + ` - Run the application in development mode
- ` + "`" + `make build` + "`" + ` - Build the application
- ` + "`" + `make test` + "`" + ` - Run tests
- ` + "`" + `make clean` + "`" + ` - Clean build artifacts
- ` + "`" + `make lint` + "`" + ` - Run linter

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

Generated with ❤️ by GoBack v0.1.0
`

const makefileTemplate = `.PHONY: run build test clean lint help

# Variables
BINARY_NAME={{.ProjectName}}
BINARY_PATH=./bin/$(BINARY_NAME)
MAIN_PATH=./main.go

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run the application in development mode
	@echo "Running {{.ProjectName}}..."
	go run $(MAIN_PATH)

build: ## Build the application
	@echo "Building {{.ProjectName}}..."
	@mkdir -p bin
	go build -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Binary built at $(BINARY_PATH)"

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run golangci-lint
	@echo "Running linter..."
	golangci-lint run

format: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

install: build ## Install the binary
	@echo "Installing $(BINARY_NAME)..."
	go install $(MAIN_PATH)

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):latest .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(BINARY_NAME):latest

# Development helpers
dev: ## Run with hot reload (requires air)
	@echo "Running with hot reload..."
	air

watch: ## Watch for changes and run tests
	@echo "Watching for changes..."
	fswatch -o . | xargs -n1 -I{} make test

# Database commands (if applicable)
{{- if ne .Database "sqlite"}}
db-up: ## Start database (Docker)
	@echo "Starting database..."
	docker-compose up -d database

db-down: ## Stop database (Docker)
	@echo "Stopping database..."
	docker-compose down

db-logs: ## Show database logs
	@echo "Database logs..."
	docker-compose logs -f database
{{- end}}

# Deployment commands
{{- if .DevOps.Enabled}}
{{- if .DevOps.Kubernetes}}
k8s-deploy: ## Deploy to Kubernetes
	@echo "Deploying to Kubernetes..."
	kubectl apply -f deployments/kubernetes/

k8s-delete: ## Delete Kubernetes deployment
	@echo "Deleting Kubernetes deployment..."
	kubectl delete -f deployments/kubernetes/
{{- end}}

{{- if .DevOps.Helm}}
helm-install: ## Install Helm chart
	@echo "Installing Helm chart..."
	helm install {{.ProjectName}} deployments/helm/

helm-upgrade: ## Upgrade Helm release
	@echo "Upgrading Helm release..."
	helm upgrade {{.ProjectName}} deployments/helm/

helm-uninstall: ## Uninstall Helm release
	@echo "Uninstalling Helm release..."
	helm uninstall {{.ProjectName}}
{{- end}}
{{- end}}
`

const gitignoreTemplate = `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
*.test
*.out

# Go workspace file
go.work

# Dependency directories
vendor/

# Go modules
go.sum

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool
*.out
coverage.html

# IDE files
.vscode/
.idea/
*.swp
*.swo
*~

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Environment files
.env
.env.local
.env.*.local

# Log files
*.log
logs/

# Database files
{{- if eq .Database "sqlite"}}
*.db
*.sqlite
*.sqlite3
{{- end}}

# Temporary files
tmp/
temp/
*.tmp

# Build artifacts
dist/
build/

# Docker
.dockerignore

# Kubernetes secrets
secrets.yaml
*-secret.yaml

# Terraform
*.tfstate
*.tfstate.*
.terraform/
.terraform.lock.hcl

# Node modules (if any frontend tools)
node_modules/

# Certificates
*.pem
*.key
*.crt

# Local configuration overrides
config.local.yaml
settings.local.json
`

const envExampleTemplate = `# Application Configuration
APP_NAME={{.ProjectName}}
APP_VERSION=1.0.0
APP_ENV=development
APP_PORT=8080
APP_HOST=localhost

# Database Configuration
{{- if eq .Database "postgresql"}}
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME={{.ProjectName}}_dev
DB_SSLMODE=disable
DATABASE_URL=postgres://postgres:password@localhost:5432/{{.ProjectName}}_dev?sslmode=disable
{{- else if eq .Database "mysql"}}
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME={{.ProjectName}}_dev
DATABASE_URL=root:password@tcp(localhost:3306)/{{.ProjectName}}_dev?parseTime=true
{{- else if eq .Database "sqlite"}}
DB_PATH=./{{.ProjectName}}.db
DATABASE_URL=file:./{{.ProjectName}}.db
{{- end}}

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRES_IN=24h

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=Content-Type,Authorization

# Rate Limiting
RATE_LIMIT_REQUESTS_PER_MINUTE=60

# Logging
LOG_LEVEL=info
LOG_FORMAT=json

# Security
BCRYPT_COST=12

# External Services (examples)
REDIS_URL=redis://localhost:6379
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASSWORD=your-app-password

# Development
DEBUG=true
PRETTY_LOG=true
`

// Framework-specific templates

// Generic main template (fallback)
const genericMainTemplate = `package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("{{.ProjectName}} server starting on port %s", port)
	// TODO: Initialize your chosen framework here
}
`

// Fiber templates
const fiberMainTemplate = `package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"{{.ModulePath}}/internal/handlers"
	"{{.ModulePath}}/internal/database"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "{{.ProjectName}}",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Routes
	setupRoutes(app, db)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("{{.ProjectName}} server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func setupRoutes(app *fiber.App, db interface{}) {
	// Health check
	app.Get("/health", handlers.HealthCheck)

	// API routes
	api := app.Group("/api/v1")

	// Add your routes here
	_ = api
}
`

const fiberHandlerTemplate = `package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheck returns the health status of the application
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
		"service": "{{.ProjectName}}",
		"version": "1.0.0",
	})
}
`

const fiberRoutesTemplate = `package routes

import (
	"github.com/gofiber/fiber/v2"
	"{{.ModulePath}}/internal/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// Health check
	app.Get("/health", handlers.HealthCheck)

	// API v1 routes
	api := app.Group("/api/v1")

	// User routes example
	users := api.Group("/users")
	users.Get("/", handlers.GetUsers)      // GET /api/v1/users
	users.Post("/", handlers.CreateUser)   // POST /api/v1/users
	users.Get("/:id", handlers.GetUser)    // GET /api/v1/users/:id
	users.Put("/:id", handlers.UpdateUser) // PUT /api/v1/users/:id
	users.Delete("/:id", handlers.DeleteUser) // DELETE /api/v1/users/:id
}
`

const fiberCorsTemplate = `package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupCORS configures CORS middleware
func SetupCORS() fiber.Handler {
	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}

	allowedMethods := os.Getenv("CORS_ALLOWED_METHODS")
	if allowedMethods == "" {
		allowedMethods = "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"
	}

	allowedHeaders := os.Getenv("CORS_ALLOWED_HEADERS")
	if allowedHeaders == "" {
		allowedHeaders = "Origin,Content-Type,Accept,Authorization"
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     allowedMethods,
		AllowHeaders:     allowedHeaders,
		AllowCredentials: true,
	})
}
`

// Gin templates
const ginMainTemplate = `package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"{{.ModulePath}}/internal/handlers"
	"{{.ModulePath}}/internal/database"
	"{{.ModulePath}}/internal/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Set Gin mode
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create Gin router
	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.SetupCORS())

	// Routes
	setupRoutes(r, db)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("{{.ProjectName}} server starting on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

func setupRoutes(r *gin.Engine, db interface{}) {
	// Health check
	r.GET("/health", handlers.HealthCheck)

	// API routes
	api := r.Group("/api/v1")
	{
		// Add your routes here
		_ = api
	}
}
`

const ginHandlerTemplate = `package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns the health status of the application
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "{{.ProjectName}}",
		"version": "1.0.0",
	})
}
`

const ginRoutesTemplate = `package routes

import (
	"github.com/gin-gonic/gin"
	"{{.ModulePath}}/internal/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", handlers.HealthCheck)

	// API v1 routes
	api := r.Group("/api/v1")
	{
		// User routes example
		users := api.Group("/users")
		{
			users.GET("", handlers.GetUsers)      // GET /api/v1/users
			users.POST("", handlers.CreateUser)   // POST /api/v1/users
			users.GET("/:id", handlers.GetUser)   // GET /api/v1/users/:id
			users.PUT("/:id", handlers.UpdateUser) // PUT /api/v1/users/:id
			users.DELETE("/:id", handlers.DeleteUser) // DELETE /api/v1/users/:id
		}
	}
}
`

const ginCorsTemplate = `package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupCORS configures CORS middleware
func SetupCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
		if allowedOrigins == "" {
			allowedOrigins = "*"
		}

		allowedMethods := os.Getenv("CORS_ALLOWED_METHODS")
		if allowedMethods == "" {
			allowedMethods = "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"
		}

		allowedHeaders := os.Getenv("CORS_ALLOWED_HEADERS")
		if allowedHeaders == "" {
			allowedHeaders = "Origin,Content-Type,Accept,Authorization"
		}

		origin := c.Request.Header.Get("Origin")
		if allowedOrigins == "*" || contains(strings.Split(allowedOrigins, ","), origin) {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		c.Header("Access-Control-Allow-Methods", allowedMethods)
		c.Header("Access-Control-Allow-Headers", allowedHeaders)
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.TrimSpace(s) == item {
			return true
		}
	}
	return false
}
`

// Chi templates
const chiMainTemplate = `package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"{{.ModulePath}}/internal/handlers"
	"{{.ModulePath}}/internal/database"
	chimiddleware "{{.ModulePath}}/internal/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create Chi router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(chimiddleware.SetupCORS())

	// Routes
	setupRoutes(r, db)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Printf("{{.ProjectName}} server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}

func setupRoutes(r chi.Router, db interface{}) {
	// Health check
	r.Get("/health", handlers.HealthCheck)

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		// Add your routes here
	})
}
`

const chiHandlerTemplate = `package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthCheck returns the health status of the application
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  "ok",
		"service": "{{.ProjectName}}",
		"version": "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
`

const chiRoutesTemplate = `package routes

import (
	"github.com/go-chi/chi/v5"
	"{{.ModulePath}}/internal/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes(r chi.Router) {
	// Health check
	r.Get("/health", handlers.HealthCheck)

	// API v1 routes
	r.Route("/api/v1", func(r chi.Router) {
		// User routes example
		r.Route("/users", func(r chi.Router) {
			r.Get("/", handlers.GetUsers)      // GET /api/v1/users
			r.Post("/", handlers.CreateUser)   // POST /api/v1/users
			r.Get("/{id}", handlers.GetUser)   // GET /api/v1/users/{id}
			r.Put("/{id}", handlers.UpdateUser) // PUT /api/v1/users/{id}
			r.Delete("/{id}", handlers.DeleteUser) // DELETE /api/v1/users/{id}
		})
	})
}
`

const chiCorsTemplate = `package middleware

import (
	"net/http"
	"os"
	"strings"
)

// SetupCORS configures CORS middleware
func SetupCORS() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
			if allowedOrigins == "" {
				allowedOrigins = "*"
			}

			allowedMethods := os.Getenv("CORS_ALLOWED_METHODS")
			if allowedMethods == "" {
				allowedMethods = "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"
			}

			allowedHeaders := os.Getenv("CORS_ALLOWED_HEADERS")
			if allowedHeaders == "" {
				allowedHeaders = "Origin,Content-Type,Accept,Authorization"
			}

			origin := r.Header.Get("Origin")
			if allowedOrigins == "*" || contains(strings.Split(allowedOrigins, ","), origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.TrimSpace(s) == item {
			return true
		}
	}
	return false
}
`

// Echo templates
const echoMainTemplate = `package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/joho/godotenv"

	"{{.ModulePath}}/internal/handlers"
	"{{.ModulePath}}/internal/database"
	echomiddleware "{{.ModulePath}}/internal/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echomiddleware.SetupCORS())

	// Routes
	setupRoutes(e, db)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Graceful shutdown
	go func() {
		log.Printf("{{.ProjectName}} server starting on port %s", port)
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}

func setupRoutes(e *echo.Echo, db interface{}) {
	// Health check
	e.GET("/health", handlers.HealthCheck)

	// API routes
	api := e.Group("/api/v1")

	// Add your routes here
	_ = api
}
`

const echoHandlerTemplate = `package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck returns the health status of the application
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"service": "{{.ProjectName}}",
		"version": "1.0.0",
	})
}
`

const echoRoutesTemplate = `package routes

import (
	"github.com/labstack/echo/v4"
	"{{.ModulePath}}/internal/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes(e *echo.Echo) {
	// Health check
	e.GET("/health", handlers.HealthCheck)

	// API v1 routes
	api := e.Group("/api/v1")

	// User routes example
	users := api.Group("/users")
	users.GET("", handlers.GetUsers)        // GET /api/v1/users
	users.POST("", handlers.CreateUser)     // POST /api/v1/users
	users.GET("/:id", handlers.GetUser)     // GET /api/v1/users/:id
	users.PUT("/:id", handlers.UpdateUser)  // PUT /api/v1/users/:id
	users.DELETE("/:id", handlers.DeleteUser) // DELETE /api/v1/users/:id
}
`

const echoCorsTemplate = `package middleware

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupCORS configures CORS middleware
func SetupCORS() echo.MiddlewareFunc {
	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}

	allowedMethods := os.Getenv("CORS_ALLOWED_METHODS")
	if allowedMethods == "" {
		allowedMethods = "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"
	}

	allowedHeaders := os.Getenv("CORS_ALLOWED_HEADERS")
	if allowedHeaders == "" {
		allowedHeaders = "Origin,Content-Type,Accept,Authorization"
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(allowedOrigins, ","),
		AllowMethods:     strings.Split(allowedMethods, ","),
		AllowHeaders:     strings.Split(allowedHeaders, ","),
		AllowCredentials: true,
	})
}
`

// Database configuration templates
const postgresConfigTemplate = `package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/lib/pq"
	{{- if eq .ORM "gorm"}}
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	{{- else if eq .ORM "sqlx"}}
	"github.com/jmoiron/sqlx"
	{{- end}}
)

{{- if eq .ORM "gorm"}}
func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_NAME", "{{.ProjectName}}_dev"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_SSLMODE", "disable"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	return db, nil
}
{{- else if eq .ORM "sqlx"}}
func Connect() (*sqlx.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_NAME", "{{.ProjectName}}_dev"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_SSLMODE", "disable"),
		)
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	return db, nil
}
{{- else}}
func Connect() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_NAME", "{{.ProjectName}}_dev"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_SSLMODE", "disable"),
		)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	return db, nil
}
{{- end}}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
`

const mysqlConfigTemplate = `package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	{{- if eq .ORM "gorm"}}
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	{{- else if eq .ORM "sqlx"}}
	"github.com/jmoiron/sqlx"
	{{- end}}
)

{{- if eq .ORM "gorm"}}
func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			getEnv("DB_USER", "root"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "3306"),
			getEnv("DB_NAME", "{{.ProjectName}}_dev"),
		)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	return db, nil
}
{{- else if eq .ORM "sqlx"}}
func Connect() (*sqlx.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			getEnv("DB_USER", "root"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "3306"),
			getEnv("DB_NAME", "{{.ProjectName}}_dev"),
		)
	}

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	return db, nil
}
{{- else}}
func Connect() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			getEnv("DB_USER", "root"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "3306"),
			getEnv("DB_NAME", "{{.ProjectName}}_dev"),
		)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	return db, nil
}
{{- end}}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
`

const sqliteConfigTemplate = `package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	{{- if eq .ORM "gorm"}}
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	{{- else if eq .ORM "sqlx"}}
	"github.com/jmoiron/sqlx"
	{{- end}}
)

{{- if eq .ORM "gorm"}}
func Connect() (*gorm.DB, error) {
	dbPath := getEnv("DB_PATH", "./{{.ProjectName}}.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite: %w", err)
	}

	return db, nil
}
{{- else if eq .ORM "sqlx"}}
func Connect() (*sqlx.DB, error) {
	dbPath := getEnv("DB_PATH", "./{{.ProjectName}}.db")

	db, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite: %w", err)
	}

	return db, nil
}
{{- else}}
func Connect() (*sql.DB, error) {
	dbPath := getEnv("DB_PATH", "./{{.ProjectName}}.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping SQLite: %w", err)
	}

	return db, nil
}
{{- end}}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
`

// Configuration templates
const postgresYamlTemplate = `database:
  driver: postgres
  host: ${DB_HOST:localhost}
  port: ${DB_PORT:5432}
  user: ${DB_USER:postgres}
  password: ${DB_PASSWORD:password}
  name: ${DB_NAME:{{.ProjectName}}_dev}
  sslmode: ${DB_SSLMODE:disable}
  max_open_conns: 25
  max_idle_conns: 25
  max_lifetime: 5m
`

const mysqlYamlTemplate = `database:
  driver: mysql
  host: ${DB_HOST:localhost}
  port: ${DB_PORT:3306}
  user: ${DB_USER:root}
  password: ${DB_PASSWORD:password}
  name: ${DB_NAME:{{.ProjectName}}_dev}
  charset: utf8mb4
  parse_time: true
  max_open_conns: 25
  max_idle_conns: 25
  max_lifetime: 5m
`

const sqliteYamlTemplate = `database:
  driver: sqlite3
  path: ${DB_PATH:./{{.ProjectName}}.db}
  max_open_conns: 1
  max_idle_conns: 1
`

const configYamlTemplate = `app:
  name: {{.ProjectName}}
  version: 1.0.0
  env: ${APP_ENV:development}
  port: ${APP_PORT:8080}
  host: ${APP_HOST:localhost}

server:
  read_timeout: 15s
  write_timeout: 15s
  idle_timeout: 60s
  max_header_bytes: 1048576

cors:
  allowed_origins:
    - ${CORS_ALLOWED_ORIGINS:http://localhost:3000}
  allowed_methods:
    - GET
    - POST
    - PUT
    - DELETE
    - OPTIONS
  allowed_headers:
    - Origin
    - Content-Type
    - Accept
    - Authorization
  allow_credentials: true

jwt:
  secret: ${JWT_SECRET:your-super-secret-jwt-key}
  expires_in: ${JWT_EXPIRES_IN:24h}

logging:
  level: ${LOG_LEVEL:info}
  format: ${LOG_FORMAT:json}
  output: stdout

rate_limit:
  requests_per_minute: ${RATE_LIMIT_REQUESTS_PER_MINUTE:60}
  burst: 10

security:
  bcrypt_cost: ${BCRYPT_COST:12}
`

// Simple Architecture Templates
const simpleUserModelTemplate = `package models

import (
	"time"
	{{- if eq .ORM "gorm"}}
	"gorm.io/gorm"
	{{- end}}
)

// User represents a user in the system
type User struct {
	{{- if eq .ORM "gorm"}}
	ID        uint           ` + "`gorm:\"primarykey\" json:\"id\"`" + `
	CreatedAt time.Time      ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time      ` + "`json:\"updated_at\"`" + `
	DeletedAt gorm.DeletedAt ` + "`gorm:\"index\" json:\"deleted_at,omitempty\"`" + `
	{{- else}}
	ID        int       ` + "`db:\"id\" json:\"id\"`" + `
	CreatedAt time.Time ` + "`db:\"created_at\" json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`db:\"updated_at\" json:\"updated_at\"`" + `
	{{- end}}
	Name      string    ` + "`{{if eq .ORM \"gorm\"}}gorm:\"not null\"{{else}}db:\"name\"{{end}} json:\"name\" validate:\"required\"`" + `
	Email     string    ` + "`{{if eq .ORM \"gorm\"}}gorm:\"uniqueIndex;not null\"{{else}}db:\"email\"{{end}} json:\"email\" validate:\"required,email\"`" + `
	Password  string    ` + "`{{if eq .ORM \"gorm\"}}gorm:\"not null\"{{else}}db:\"password\"{{end}} json:\"-\" validate:\"required,min=6\"`" + `
	Active    bool      ` + "`{{if eq .ORM \"gorm\"}}gorm:\"default:true\"{{else}}db:\"active\"{{end}} json:\"active\"`" + `
}

// UserCreateRequest represents the request body for creating a user
type UserCreateRequest struct {
	Name     string ` + "`json:\"name\" validate:\"required\"`" + `
	Email    string ` + "`json:\"email\" validate:\"required,email\"`" + `
	Password string ` + "`json:\"password\" validate:\"required,min=6\"`" + `
}

// UserUpdateRequest represents the request body for updating a user
type UserUpdateRequest struct {
	Name   string ` + "`json:\"name,omitempty\"`" + `
	Email  string ` + "`json:\"email,omitempty\" validate:\"omitempty,email\"`" + `
	Active *bool  ` + "`json:\"active,omitempty\"`" + `
}

// UserResponse represents the response body for user operations
type UserResponse struct {
	ID        {{if eq .ORM "gorm"}}uint{{else}}int{{end}}      ` + "`json:\"id\"`" + `
	Name      string    ` + "`json:\"name\"`" + `
	Email     string    ` + "`json:\"email\"`" + `
	Active    bool      ` + "`json:\"active\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
`

const simpleUserServiceTemplate = `package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"{{.ModulePath}}/internal/models"
	{{- if eq .ORM "gorm"}}
	"gorm.io/gorm"
	{{- else if eq .ORM "sqlx"}}
	"github.com/jmoiron/sqlx"
	{{- else}}
	"database/sql"
	{{- end}}
)

// UserService handles user business logic
type UserService struct {
	{{- if eq .ORM "gorm"}}
	db *gorm.DB
	{{- else if eq .ORM "sqlx"}}
	db *sqlx.DB
	{{- else}}
	db *sql.DB
	{{- end}}
}

// NewUserService creates a new UserService
func NewUserService({{if eq .ORM "gorm"}}db *gorm.DB{{else if eq .ORM "sqlx"}}db *sqlx.DB{{else}}db *sql.DB{{end}}) *UserService {
	return &UserService{db: db}
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User

	{{- if eq .ORM "gorm"}}
	if err := s.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	if err := s.db.Select(&users, "SELECT * FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC"); err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	{{- else}}
	rows, err := s.db.Query("SELECT id, name, email, active, created_at, updated_at FROM users ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	{{- end}}

	return users, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) (*models.User, error) {
	var user models.User

	{{- if eq .ORM "gorm"}}
	if err := s.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	if err := s.db.Get(&user, "SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- else}}
	if err := s.db.QueryRow("SELECT id, name, email, active, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- end}}

	return &user, nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(req models.UserCreateRequest) (*models.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Active:   true,
	}

	{{- if eq .ORM "gorm"}}
	if err := s.db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	query := ` + "`INSERT INTO users (name, email, password, active, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, created_at, updated_at`" + `
	if err := s.db.QueryRowx(query, user.Name, user.Email, user.Password, user.Active).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	{{- else}}
	query := ` + "`INSERT INTO users (name, email, password, active) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`" + `
	if err := s.db.QueryRow(query, user.Name, user.Email, user.Password, user.Active).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	{{- end}}

	return &user, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}, req models.UserUpdateRequest) (*models.User, error) {
	user, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Active != nil {
		user.Active = *req.Active
	}

	{{- if eq .ORM "gorm"}}
	if err := s.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	query := ` + "`UPDATE users SET name = $1, email = $2, active = $3, updated_at = NOW() WHERE id = $4`" + `
	if _, err := s.db.Exec(query, user.Name, user.Email, user.Active, user.ID); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	{{- else}}
	query := ` + "`UPDATE users SET name = $1, email = $2, active = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4`" + `
	if _, err := s.db.Exec(query, user.Name, user.Email, user.Active, user.ID); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	{{- end}}

	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) error {
	{{- if eq .ORM "gorm"}}
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	if _, err := s.db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	{{- else}}
	if _, err := s.db.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	{{- end}}

	return nil
}
`

const simpleUserHandlerTemplate = `package handlers

import (
	"net/http"
	"strconv"

	{{- if eq .Framework "fiber"}}
	"github.com/gofiber/fiber/v2"
	{{- else if eq .Framework "gin"}}
	"github.com/gin-gonic/gin"
	{{- else if eq .Framework "chi"}}
	"github.com/go-chi/chi/v5"
	"encoding/json"
	{{- else if eq .Framework "echo"}}
	"github.com/labstack/echo/v4"
	{{- end}}

	"{{.ModulePath}}/internal/models"
	"{{.ModulePath}}/internal/services"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

{{- if eq .Framework "fiber"}}
// GetUsers returns all users
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var responses []models.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	return c.JSON(responses)
}

// GetUser returns a user by ID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.userService.GetUserByID({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user.ToResponse())
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user.ToResponse())
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var req models.UserUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userService.UpdateUser({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}}, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user.ToResponse())
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	if err := h.userService.DeleteUser({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

{{- else if eq .Framework "gin"}}
// GetUsers returns all users
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responses []models.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	c.JSON(http.StatusOK, responses)
}

// GetUser returns a user by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user.ToResponse())
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user.ToResponse())
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req models.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.userService.UpdateUser({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}}, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user.ToResponse())
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userService.DeleteUser({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

{{- else if eq .Framework "echo"}}
// GetUsers returns all users
func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var responses []models.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	return c.JSON(http.StatusOK, responses)
}

// GetUser returns a user by ID
func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, err := h.userService.GetUserByID({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}})
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user.ToResponse())
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c echo.Context) error {
	var req models.UserCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user.ToResponse())
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var req models.UserUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	user, err := h.userService.UpdateUser({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}}, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user.ToResponse())
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := h.userService.DeleteUser({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}}); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

{{- else if eq .Framework "chi"}}
// GetUsers returns all users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var responses []models.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

// GetUser returns a user by ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ToResponse())
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user.ToResponse())
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.UpdateUser({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}}, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ToResponse())
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.userService.DeleteUser({{if eq .ORM "gorm"}}uint(id){{else}}id{{end}}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
{{- end}}
`

// DDD Architecture Templates
const dddUserEntityTemplate = `package entities

import (
	"time"
	{{- if eq .ORM "gorm"}}
	"gorm.io/gorm"
	{{- end}}
)

// User represents the User domain entity
type User struct {
	{{- if eq .ORM "gorm"}}
	ID        uint           ` + "`gorm:\"primarykey\" json:\"id\"`" + `
	CreatedAt time.Time      ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time      ` + "`json:\"updated_at\"`" + `
	DeletedAt gorm.DeletedAt ` + "`gorm:\"index\" json:\"-\"`" + `
	{{- else}}
	ID        int       ` + "`json:\"id\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
	{{- end}}
	Name     string ` + "`json:\"name\"`" + `
	Email    string ` + "`json:\"email\"`" + `
	Password string ` + "`json:\"-\"`" + `
	Active   bool   ` + "`json:\"active\"`" + `
}

// NewUser creates a new User entity
func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Active:   true,
	}
}

// IsValid validates the user entity
func (u *User) IsValid() bool {
	return u.Name != "" && u.Email != "" && u.Password != ""
}

// Activate activates the user
func (u *User) Activate() {
	u.Active = true
}

// Deactivate deactivates the user
func (u *User) Deactivate() {
	u.Active = false
}
`

const dddUserRepositoryTemplate = `package repositories

import "{{.ModulePath}}/internal/domain/entities"

// UserRepository defines the interface for user persistence operations
type UserRepository interface {
	GetAll() ([]*entities.User, error)
	GetByID(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) error
	Exists(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) bool
	Count() (int64, error)
}
`

const dddUserServiceTemplate = `package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"{{.ModulePath}}/internal/domain/entities"
	"{{.ModulePath}}/internal/domain/repositories"
)

// UserService handles user domain logic
type UserService struct {
	repo repositories.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]*entities.User, error) {
	return s.repo.GetAll()
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) (*entities.User, error) {
	return s.repo.GetByID(id)
}

// CreateUser creates a new user with business rules
func (s *UserService) CreateUser(name, email, password string) (*entities.User, error) {
	// Check if user already exists
	existing, _ := s.repo.GetByEmail(email)
	if existing != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user entity
	user := entities.NewUser(name, email, string(hashedPassword))
	if !user.IsValid() {
		return nil, fmt.Errorf("invalid user data")
	}

	// Save to repository
	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}, name, email string, active *bool) (*entities.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if name != "" {
		user.Name = name
	}
	if email != "" {
		// Check email uniqueness
		existing, _ := s.repo.GetByEmail(email)
		if existing != nil && existing.ID != user.ID {
			return nil, fmt.Errorf("email %s is already taken", email)
		}
		user.Email = email
	}
	if active != nil {
		user.Active = *active
	}

	if err := s.repo.Update(user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) error {
	if !s.repo.Exists(id) {
		return fmt.Errorf("user not found")
	}

	return s.repo.Delete(id)
}

// VerifyPassword verifies user password
func (s *UserService) VerifyPassword(email, password string) (*entities.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if !user.Active {
		return nil, fmt.Errorf("user account is deactivated")
	}

	return user, nil
}
`

const dddUserPersistenceTemplate = `package persistence

import (
	"fmt"

	"{{.ModulePath}}/internal/domain/entities"
	"{{.ModulePath}}/internal/domain/repositories"
	{{- if eq .ORM "gorm"}}
	"gorm.io/gorm"
	{{- else if eq .ORM "sqlx"}}
	"github.com/jmoiron/sqlx"
	{{- else}}
	"database/sql"
	{{- end}}
)

// UserRepositoryImpl implements UserRepository interface
type UserRepositoryImpl struct {
	{{- if eq .ORM "gorm"}}
	db *gorm.DB
	{{- else if eq .ORM "sqlx"}}
	db *sqlx.DB
	{{- else}}
	db *sql.DB
	{{- end}}
}

// NewUserRepository creates a new UserRepository implementation
func NewUserRepository({{if eq .ORM "gorm"}}db *gorm.DB{{else if eq .ORM "sqlx"}}db *sqlx.DB{{else}}db *sql.DB{{end}}) repositories.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// GetAll retrieves all users
func (r *UserRepositoryImpl) GetAll() ([]*entities.User, error) {
	var users []*entities.User

	{{- if eq .ORM "gorm"}}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	if err := r.db.Select(&users, "SELECT * FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC"); err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	{{- else}}
	rows, err := r.db.Query("SELECT id, name, email, password, active, created_at, updated_at FROM users ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &entities.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Active, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	{{- end}}

	return users, nil
}

// GetByID retrieves a user by ID
func (r *UserRepositoryImpl) GetByID(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) (*entities.User, error) {
	var user entities.User

	{{- if eq .ORM "gorm"}}
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	if err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- else}}
	if err := r.db.QueryRow("SELECT id, name, email, password, active, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Active, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- end}}

	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepositoryImpl) GetByEmail(email string) (*entities.User, error) {
	var user entities.User

	{{- if eq .ORM "gorm"}}
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	if err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL", email); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- else}}
	if err := r.db.QueryRow("SELECT id, name, email, password, active, created_at, updated_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Active, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	{{- end}}

	return &user, nil
}

// Create creates a new user
func (r *UserRepositoryImpl) Create(user *entities.User) error {
	{{- if eq .ORM "gorm"}}
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	query := ` + "`INSERT INTO users (name, email, password, active, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, created_at, updated_at`" + `
	if err := r.db.QueryRowx(query, user.Name, user.Email, user.Password, user.Active).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	{{- else}}
	query := ` + "`INSERT INTO users (name, email, password, active) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`" + `
	if err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Active).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	{{- end}}

	return nil
}

// Update updates a user
func (r *UserRepositoryImpl) Update(user *entities.User) error {
	{{- if eq .ORM "gorm"}}
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	query := ` + "`UPDATE users SET name = $1, email = $2, active = $3, updated_at = NOW() WHERE id = $4`" + `
	if _, err := r.db.Exec(query, user.Name, user.Email, user.Active, user.ID); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	{{- else}}
	query := ` + "`UPDATE users SET name = $1, email = $2, active = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4`" + `
	if _, err := r.db.Exec(query, user.Name, user.Email, user.Active, user.ID); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	{{- end}}

	return nil
}

// Delete deletes a user
func (r *UserRepositoryImpl) Delete(id {{if eq .ORM "gorm"}}uint{{else}}int{{end}}) error {
	{{- if eq .ORM "gorm"}}
	if err := r.db.Delete(&entities.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	{{- else if eq .ORM "sqlx"}}
	if _, err := r.db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	{{- else}}

`
