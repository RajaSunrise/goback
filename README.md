# GoBack - TUI Backend Project Scaffolding Tool

GoBack adalah TUI (Terminal User Interface) yang dibangun dengan Bubble Tea untuk memudahkan developer backend dalam membuat project backend dengan berbagai pilihan framework, database, ORM, arsitektur, dan DevOps tools.

## Features

✨ **Interactive TUI Mode**: Beautiful terminal interface untuk konfigurasi proyek step-by-step  
🚀 **CLI Mode**: Command-line interface untuk otomasi dan scripting  
🏗️ **Multiple Frameworks**: Fiber, Gin, Chi, Echo  
🗄️ **Database Support**: PostgreSQL, MySQL, SQLite  
⚙️ **ORM Options**: GORM, SQLC, SQLX  
🏛️ **Architecture Patterns**: Simple, DDD, Clean Architecture, Hexagonal  
🔧 **DevOps Integration**: Kubernetes, Helm, Terraform, Ansible  
📝 **Auto-generated**: Go module, README, Makefile, .gitignore  

## Installation

### From Source

```bash
git clone https://github.com/NarmadaWeb/goback.git
cd goback
go build .
sudo mv goback /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/NarmadaWeb/goback@latest
```

## Quick Start

### Interactive TUI Mode (Recommended)

```bash
goback
# or
goback tui
```

This will launch the interactive TUI where you can:
1. Select your framework (Fiber, Gin, Chi, Echo)
2. Choose database (PostgreSQL, MySQL, SQLite)
3. Pick ORM/database tool (GORM, SQLC, SQLX)
4. Select architecture pattern (Simple, DDD, Clean, Hexagonal)
5. Configure DevOps tools (optional)
6. Fill in project details
7. Review and generate

### CLI Mode

```bash
# Basic usage
goback new my-api

# With specific options
goback new my-api \
  --framework=fiber \
  --database=postgresql \
  --orm=gorm \
  --architecture=clean \
  --output=./my-project \
  --module=github.com/myorg/my-api

# With DevOps tools
goback new my-api \
  --devops \
  --devops-tools=kubernetes,helm
```

## Available Options

### Frameworks
- **Fiber**: Fast HTTP web framework inspired by Express
- **Gin**: High-performance HTTP web framework
- **Chi**: Lightweight, idiomatic HTTP router
- **Echo**: High performance, extensible web framework

### Databases
- **PostgreSQL**: Advanced open-source relational database
- **MySQL**: Popular open-source relational database
- **SQLite**: Lightweight embedded database

### ORMs
- **GORM**: Feature-rich ORM with associations, hooks, and migrations
- **SQLC**: Generate type-safe code from SQL
- **SQLX**: Extensions on database/sql for easier usage

### Architecture Patterns
- **Simple**: Simple layered architecture with handlers, services, and models
- **DDD**: Domain-Driven Design with domain, infrastructure, application layers
- **Clean**: Clean Architecture with entities, use cases, and adapters
- **Hexagonal**: Hexagonal Architecture with ports and adapters pattern

### DevOps Tools
- **Kubernetes**: Container orchestration platform
- **Helm**: Kubernetes package manager
- **Terraform**: Infrastructure as code tool
- **Ansible**: IT automation and configuration management

## Commands

### Interactive Mode
```bash
goback              # Launch TUI
goback tui          # Launch TUI explicitly
```

### CLI Commands
```bash
goback new [name]   # Create new project
goback list         # List available options
goback config       # Manage configuration
goback version      # Show version
```

### Configuration Management
```bash
goback config show                    # Show current config
goback config set key value          # Set config value
```

## Project Structure

Generated projects follow the selected architecture pattern:

### Simple Architecture
```
my-project/
├── main.go
├── go.mod
├── internal/
│   ├── handlers/          # HTTP handlers
│   ├── services/          # Business logic
│   ├── models/            # Data models
│   └── config/            # Configuration
├── pkg/                   # Public packages
├── scripts/               # Build scripts
├── docs/                  # Documentation
├── Makefile
├── README.md
└── .gitignore
```

### Clean Architecture
```
my-project/
├── main.go
├── internal/
│   ├── entities/          # Business entities
│   ├── usecases/          # Business logic
│   ├── interface_adapters/
│   │   ├── controllers/   # HTTP controllers
│   │   ├── presenters/    # Response formatters
│   │   └── gateways/      # Data access
│   └── frameworks_drivers/
│       ├── web/           # Web framework
│       └── database/      # Database drivers
└── ...
```

## Generated Files

Each generated project includes:

- **main.go**: Application entry point with chosen framework
- **go.mod**: Go module with required dependencies
- **README.md**: Project documentation with setup instructions
- **Makefile**: Build automation with common tasks
- **.gitignore**: Git ignore rules for Go projects
- **Project structure**: Directories based on architecture pattern

## Configuration

GoBack stores configuration in `~/.goback.yaml`:

```yaml
default_output_dir: "./"
default_module_prefix: "github.com/user"
default_author: ""
animation_speed: 100
show_splash_screen: true
auto_save: true
theme: "default"
```

## Examples

### Creating a Fiber API with PostgreSQL and GORM
```bash
goback new ecommerce-api \
  --framework=fiber \
  --database=postgresql \
  --orm=gorm \
  --architecture=ddd \
  --output=./ecommerce
```

### Creating a Gin API with Clean Architecture
```bash
goback new user-service \
  --framework=gin \
  --database=mysql \
  --orm=gorm \
  --architecture=clean
```

### Creating a Chi API with DevOps
```bash
goback new payment-service \
  --framework=chi \
  --database=postgresql \
  --orm=sqlx \
  --devops \
  --devops-tools=kubernetes,helm,terraform
```

## Development

### Building from Source
```bash
git clone https://github.com/NarmadaWeb/goback.git
cd goback
go mod tidy
go build .
```

### Testing
```bash
go test ./...
```

### Running TUI in Development
```bash
go run . tui
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Architecture

GoBack follows a modular architecture:

- **cmd/**: CLI commands and application entry point
- **internal/tui/**: Terminal User Interface components
- **internal/scaffolding/**: Project generation logic
- **internal/config/**: Configuration management
- **templates/**: Project templates for different combinations
- **assets/**: Static assets and logos

## Requirements

- Go 1.21 or higher
- Terminal with 256 colors support (recommended)

## Roadmap

- [ ] More framework support (Echo variations, FastHTTP)
- [ ] Additional database drivers (MongoDB, Redis)
- [ ] Plugin system for custom templates
- [ ] Web-based UI companion
- [ ] Project templates marketplace
- [ ] Integration with popular IDEs

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI
- Styling with [Lipgloss](https://github.com/charmbracelet/lipgloss)
- CLI powered by [Cobra](https://github.com/spf13/cobra)
- Configuration management with [Viper](https://github.com/spf13/viper)

---

Made with ❤️ for the Go community