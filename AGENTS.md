# AGENTS.md - GoBack TUI Backend Project Scaffolding Tool

## Overview

GoBack adalah TUI (Terminal User Interface) yang dibangun dengan Bubble Tea untuk memudahkan developer backend dalam membuat project backend dengan berbagai pilihan framework, database, ORM, arsitektur, dan DevOps tools.

## Project Structure

```
goback/
├── go.mod                  # Go module file
└── cmd/
    └── migrate.go          # CLI migrate command
```

## Application Flow

### 1. Main Menu State Flow

```
[Splash Screen Animation]
         ↓
[Main Menu Selection]
    ├── Mulai Proyek Baru → [Project Configuration Flow]
    ├── Buka Proyek → [Project Management]
    ├── Lihat Pengaturan → [Settings Menu]
    └── Keluar → [Exit Application]
```

### 2. Project Configuration Flow

#### 2.1 Framework Selection
```
[Framework Selection Screen]
    ├── Go Fiber
    ├── Go Gin  
    ├── Go Chi
    └── Go Echo
         ↓
[Database Selection Screen]
```

#### 2.2 Database Selection
```
[Database Selection Screen]
    ├── PostgreSQL
    ├── MySQL
    └── SQLite
         ↓
[ORM/Database Tool Selection Screen]
```

#### 2.3 ORM/Database Tool Selection
```
[ORM Tool Selection Screen]
    ├── GORM
    ├── SQLC
    └── SQLX
         ↓
[Architecture Pattern Selection Screen]
```

#### 2.4 Architecture Pattern Selection
```
[Architecture Pattern Selection Screen]
    ├── Simple Architecture
    ├── Domain-Driven Design (DDD)
    ├── Clean Architecture
    └── Hexagonal Architecture
         ↓
[DevOps Options Screen]
```

#### 2.5 DevOps Configuration
```
[DevOps Options Screen]
    ├── Tidak Menggunakan DevOps → [Project Details Screen]
    └── Menggunakan DevOps → [DevOps Tools Selection]
                                    ↓
                         [DevOps Tools Selection Screen]
                            ├── Kubernetes
                            ├── Helm
                            ├── Terraform
                            └── Ansible
                                    ↓
                         [Project Details Screen]
```

#### 2.6 Project Details
```
[Project Details Screen]
    ├── Project Name Input
    ├── Module Path Input
    ├── Description Input
    └── Output Directory Input
         ↓
[Configuration Review Screen]
```

#### 2.7 Configuration Review & Generation
```
[Configuration Review Screen]
    ├── Review All Selections
    ├── Edit Configuration → [Back to Previous Screens]
    └── Generate Project → [Project Generation Process]
                                    ↓
                         [Generation Progress Screen]
                                    ↓
                         [Success/Error Screen]
```

## TUI Components Architecture

### 1. Core Models (Bubble Tea)

#### Main Model
```go
type MainModel struct {
    State       AppState
    WindowSize  tea.WindowSizeMsg
    
    // Sub-models
    SplashModel      *SplashModel
    MenuModel        *MenuModel
    ConfigModel      *ConfigModel
    ProgressModel    *ProgressModel
    
    // Shared state
    Config          *ProjectConfig
    Error           error
}
```

#### Configuration Model
```go
type ConfigModel struct {
    Step            ConfigStep
    Framework       FrameworkChoice
    Database        DatabaseChoice
    ORM             ORMChoice
    Architecture    ArchitectureChoice
    DevOps          DevOpsConfig
    ProjectDetails  ProjectDetails
    
    // UI state
    Cursor          int
    Choices         []string
    InputValue      string
    InputMode       bool
}
```

### 2. Application States

```go
type AppState int

const (
    StateSplash AppState = iota
    StateMainMenu
    StateFrameworkSelection
    StateDatabaseSelection
    StateORMSelection
    StateArchitectureSelection
    StateDevOpsOptions
    StateDevOpsToolsSelection
    StateProjectDetails
    StateConfigReview
    StateGeneration
    StateProgress
    StateSuccess
    StateError
)
```

### 3. Configuration Steps

```go
type ConfigStep int

const (
    StepFramework ConfigStep = iota
    StepDatabase
    StepORM
    StepArchitecture
    StepDevOpsOptions
    StepDevOpsTools
    StepProjectDetails
    StepReview
)
```

## Project Generation System

### 1. Template System

#### Framework Templates
```
templates/frameworks/
├── fiber/
│   ├── main.go.tmpl
│   ├── routes.go.tmpl
│   ├── middleware.go.tmpl
│   └── config.go.tmpl
├── gin/
├── chi/
└── echo/
```

#### Architecture Templates
```
templates/architectures/
├── simple/
│   ├── handlers/
│   ├── models/
│   └── services/
├── ddd/
│   ├── domain/
│   ├── infrastructure/
│   ├── application/
│   └── interfaces/
├── clean/
│   ├── entities/
│   ├── usecases/
│   ├── interface_adapters/
│   └── frameworks_drivers/
└── hexagonal/
    ├── domain/
    ├── ports/
    └── adapters/
```

#### DevOps Templates
```
templates/devops/
├── kubernetes/
│   ├── deployment.yaml.tmpl
│   ├── service.yaml.tmpl
│   └── ingress.yaml.tmpl
├── helm/
│   ├── Chart.yaml.tmpl
│   ├── values.yaml.tmpl
│   └── templates/
├── terraform/
│   ├── main.tf.tmpl
│   ├── variables.tf.tmpl
│   └── outputs.tf.tmpl
└── ansible/
    ├── playbook.yml.tmpl
    └── inventory.ini.tmpl
```

### 2. Generator Engine

#### Template Generator
```go
type TemplateGenerator struct {
    Config      *ProjectConfig
    OutputDir   string
    Templates   map[string]*template.Template
}

func (tg *TemplateGenerator) Generate() error {
    // 1. Create directory structure
    // 2. Generate base files
    // 3. Generate framework-specific files
    // 4. Generate architecture-specific files
    // 5. Generate database configuration
    // 6. Generate DevOps files if needed
    // 7. Initialize go module
    // 8. Install dependencies
}
```

#### File Structure Generator
```go
type FileStructureGenerator struct {
    Architecture    ArchitectureChoice
    Framework      FrameworkChoice
    HasDevOps      bool
}

func (fsg *FileStructureGenerator) CreateDirectories(basePath string) error {
    // Create directory structure based on architecture pattern
}
```

### 3. Configuration Validation

#### Validator System
```go
type ConfigValidator struct {
    Rules []ValidationRule
}

type ValidationRule interface {
    Validate(config *ProjectConfig) error
}

// Example validators
type ProjectNameValidator struct{}
type ModulePathValidator struct{}
type DirectoryValidator struct{}
```

## Development Phases

### Phase 1: Core TUI Framework
- [ ] Setup Bubble Tea application structure
- [ ] Implement splash screen with animation
- [ ] Create main menu navigation
- [ ] Implement basic state management
- [ ] Add keyboard navigation and controls

### Phase 2: Configuration Flow
- [ ] Framework selection screen
- [ ] Database selection screen
- [ ] ORM selection screen
- [ ] Architecture pattern selection
- [ ] DevOps options configuration
- [ ] Project details input forms
- [ ] Configuration review screen

### Phase 3: Template System
- [ ] Design template structure
- [ ] Create framework templates (Fiber, Gin, Chi, Echo)
- [ ] Create architecture templates (Simple, DDD, Clean, Hexagonal)
- [ ] Create database configuration templates
- [ ] Create DevOps templates (Kubernetes, Helm, Terraform, Ansible)

### Phase 4: Project Generation
- [ ] Implement template rendering engine
- [ ] Create file structure generator
- [ ] Add configuration validation
- [ ] Implement project generation logic
- [ ] Add progress tracking and error handling

### Phase 5: Advanced Features
- [ ] Project management (open existing projects)
- [ ] Settings and preferences
- [ ] Template customization
- [ ] Plugin system for additional frameworks
- [ ] Export/import configuration
- [ ] CLI mode support

### Phase 6: Testing & Documentation
- [ ] Unit tests for core functionality
- [ ] Integration tests for project generation
- [ ] User documentation
- [ ] Developer documentation
- [ ] Example projects and tutorials

## Technical Specifications

### Dependencies
```go
// Core TUI
github.com/charmbracelet/bubbletea
github.com/charmbracelet/lipgloss
github.com/charmbracelet/bubbles

// CLI
github.com/spf13/cobra
github.com/spf13/viper

// Template Engine
text/template
html/template

// File Operations
os
path/filepath
io/ioutil

// Validation
github.com/go-playground/validator/v10

// YAML/JSON Processing
gopkg.in/yaml.v3
encoding/json
```

### Configuration Schema
```go
type ProjectConfig struct {
    ProjectName    string            `json:"project_name" validate:"required,min=1"`
    ModulePath     string            `json:"module_path" validate:"required"`
    Description    string            `json:"description"`
    OutputDir      string            `json:"output_dir" validate:"required"`
    Framework      FrameworkChoice   `json:"framework" validate:"required"`
    Database       DatabaseChoice    `json:"database" validate:"required"`
    ORM            ORMChoice         `json:"orm" validate:"required"`
    Architecture   ArchitectureChoice `json:"architecture" validate:"required"`
    DevOps         DevOpsConfig      `json:"devops"`
    CreatedAt      time.Time         `json:"created_at"`
}

type DevOpsConfig struct {
    Enabled     bool     `json:"enabled"`
    Tools       []string `json:"tools"`
    Kubernetes  bool     `json:"kubernetes"`
    Helm        bool     `json:"helm"`
    Terraform   bool     `json:"terraform"`
    Ansible     bool     `json:"ansible"`
}
```

### Error Handling Strategy
```go
type GoBackError struct {
    Code    ErrorCode
    Message string
    Details map[string]interface{}
    Cause   error
}

type ErrorCode int

const (
    ErrInvalidConfig ErrorCode = iota
    ErrTemplateNotFound
    ErrDirectoryCreation
    ErrFileGeneration
    ErrDependencyInstall
    ErrValidation
)
```

## User Experience Guidelines

### 1. Navigation
- Arrow keys or vim-style navigation (j/k)
- Enter to select, Escape/q to go back
- Tab/Shift+Tab for form navigation
- Ctrl+C to quit anytime

### 2. Visual Design
- Consistent color scheme using Lipgloss
- Clear visual hierarchy
- Animated transitions between states
- Progress indicators for long operations
- Error messages with helpful suggestions

### 3. Help System
- Context-sensitive help text
- Keyboard shortcuts displayed
- Tooltips for complex options
- Link to documentation

### 4. Accessibility
- Clear contrast for readability
- Consistent interaction patterns
- Informative error messages
- Screen reader friendly output

## Performance Considerations

### 1. Template Loading
- Lazy load templates when needed
- Cache compiled templates
- Efficient memory usage for large template sets

### 2. File Operations
- Batch file operations where possible
- Progress feedback for slow operations
- Graceful handling of I/O errors
- Cleanup on failure

### 3. UI Responsiveness
- Non-blocking operations where possible
- Smooth animations and transitions
- Efficient screen updates
- Minimal input lag

## Security Considerations

### 1. File System Operations
- Validate all file paths
- Prevent directory traversal attacks
- Safe file overwrite handling
- Proper file permissions

### 2. Template Security
- Sanitize template inputs
- Prevent code injection in templates
- Validate generated file contents

### 3. External Dependencies
- Verify downloaded dependencies
- Use secure HTTP for external calls
- Handle network timeouts gracefully

## Testing Strategy

### 1. Unit Tests
- Core business logic
- Template rendering
- Configuration validation
- Error handling

### 2. Integration Tests
- Full project generation flow
- Template integration
- File system operations
- CLI interface

### 3. End-to-End Tests
- Complete user workflows
- Generated project compilation
- DevOps template validity

### 4. Performance Tests
- Large project generation
- Memory usage optimization
- UI responsiveness under load

## Deployment & Distribution

### 1. Build System
- Cross-platform compilation
- Automated releases via GitHub Actions
- Binary optimization for size

### 2. Package Managers
- Homebrew formula
- Chocolatey package
- Debian/RPM packages
- Go install support

### 3. Documentation
- README with quick start
- Detailed user guide
- API documentation
- Video tutorials

This comprehensive flow provides a solid foundation for building the GoBack TUI application with all the requested features and considerations for scalability, maintainability, and user experience.