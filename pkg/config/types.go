// pkg/config/types.go

package config

import (
	"strings"
	"time"
)

// ProjectConfig holds all the configuration for the project to be generated
type ProjectConfig struct {
	ProjectName  string             `json:"project_name" validate:"required,min=1"`
	ModulePath   string             `json:"module_path" validate:"required,modulepath"`
	Description  string             `json:"description"`
	OutputDir    string             `json:"output_dir" validate:"required"`
	Framework    FrameworkChoice    `json:"framework" validate:"required"`
	Database     DatabaseChoice     `json:"database" validate:"required"`
	Tool         ToolChoice         `json:"tool" validate:"required"`
	Architecture ArchitectureChoice `json:"architecture" validate:"required"`
	DevOps       DevOpsConfig       `json:"devops"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

// DevOpsConfig holds the DevOps tool configuration
type DevOpsConfig struct {
	Enabled    bool     `json:"enabled"`
	Tools      []string `json:"tools"`
	Kubernetes bool     `json:"kubernetes"`
	Helm       bool     `json:"helm"`
	Terraform  bool     `json:"terraform"`
	Ansible    bool     `json:"ansible"`
}

// Choice types for project configuration
type (
	FrameworkChoice    string
	DatabaseChoice     string
	ToolChoice         string
	ArchitectureChoice string
)

// Framework choices
const (
	FrameworkFiber FrameworkChoice = "fiber"
	FrameworkGin   FrameworkChoice = "gin"
	FrameworkChi   FrameworkChoice = "chi"
	FrameworkEcho  FrameworkChoice = "echo"
)

// Database choices
const (
	DatabasePostgreSQL DatabaseChoice = "postgresql"
	DatabaseMySQL      DatabaseChoice = "mysql"
	DatabaseSQLite     DatabaseChoice = "sqlite"
)

// Tool choices
const (
	ToolSqlx ToolChoice = "sqlx"
	ToolSqlc ToolChoice = "sqlc"
)

// Architecture choices
const (
	ArchitectureSimple    ArchitectureChoice = "simple"
	ArchitectureDDD       ArchitectureChoice = "ddd"
	ArchitectureClean     ArchitectureChoice = "clean"
	ArchitectureHexagonal ArchitectureChoice = "hexagonal"
)

// DevOps tool choices
const (
	DevOpsKubernetes = "kubernetes"
	DevOpsHelm       = "helm"
	DevOpsTerraform  = "terraform"
	DevOpsAnsible    = "ansible"
)

// Validation functions

// IsValidFramework checks if framework choice is valid
func IsValidFramework(framework FrameworkChoice) bool {
	validFrameworks := []FrameworkChoice{
		FrameworkFiber,
		FrameworkGin,
		FrameworkChi,
		FrameworkEcho,
	}

	for _, valid := range validFrameworks {
		if framework == valid {
			return true
		}
	}
	return false
}

// IsValidDatabase checks if database choice is valid
func IsValidDatabase(database DatabaseChoice) bool {
	validDatabases := []DatabaseChoice{
		DatabasePostgreSQL,
		DatabaseMySQL,
		DatabaseSQLite,
	}

	for _, valid := range validDatabases {
		if database == valid {
			return true
		}
	}
	return false
}

// IsValidTool checks if Tool choice is valid
func IsValidTool(tool ToolChoice) bool {
	validTools := []ToolChoice{
		ToolSqlx,
		ToolSqlc,
	}

	for _, valid := range validTools {
		if tool == valid {
			return true
		}
	}
	return false
}

// IsValidArchitecture checks if architecture choice is valid
func IsValidArchitecture(architecture ArchitectureChoice) bool {
	validArchitectures := []ArchitectureChoice{
		ArchitectureSimple,
		ArchitectureDDD,
		ArchitectureClean,
		ArchitectureHexagonal,
	}

	for _, valid := range validArchitectures {
		if architecture == valid {
			return true
		}
	}
	return false
}

// IsValidDevOpsTool checks if DevOps tool is valid
func IsValidDevOpsTool(tool string) bool {
	validTools := []string{
		DevOpsKubernetes,
		DevOpsHelm,
		DevOpsTerraform,
		DevOpsAnsible,
	}

	for _, valid := range validTools {
		if tool == valid {
			return true
		}
	}
	return false
}

// GetValidFrameworks returns list of valid framework choices
func GetValidFrameworks() []FrameworkChoice {
	return []FrameworkChoice{
		FrameworkFiber,
		FrameworkGin,
		FrameworkChi,
		FrameworkEcho,
	}
}

// GetValidDatabases returns list of valid database choices
func GetValidDatabases() []DatabaseChoice {
	return []DatabaseChoice{
		DatabasePostgreSQL,
		DatabaseMySQL,
		DatabaseSQLite,
	}
}

// GetValidTools returns list of valid Tool choices
func GetValidTools() []ToolChoice {
	return []ToolChoice{
		ToolSqlx,
		ToolSqlc,
	}
}

// GetValidArchitectures returns list of valid architecture choices
func GetValidArchitectures() []ArchitectureChoice {
	return []ArchitectureChoice{
		ArchitectureSimple,
		ArchitectureDDD,
		ArchitectureClean,
		ArchitectureHexagonal,
	}
}

// GetValidDevOpsTools returns list of valid DevOps tools
func GetValidDevOpsTools() []string {
	return []string{
		DevOpsKubernetes,
		DevOpsHelm,
		DevOpsTerraform,
		DevOpsAnsible,
	}
}

// String methods for better display

func (f FrameworkChoice) String() string {
	switch f {
	case FrameworkFiber:
		return "Go Fiber"
	case FrameworkGin:
		return "Go Gin"
	case FrameworkChi:
		return "Go Chi"
	case FrameworkEcho:
		return "Go Echo"
	default:
		return string(f)
	}
}

func (d DatabaseChoice) String() string {
	switch d {
	case DatabasePostgreSQL:
		return "PostgreSQL"
	case DatabaseMySQL:
		return "MySQL"
	case DatabaseSQLite:
		return "SQLite"
	default:
		return string(d)
	}
}

func (t ToolChoice) String() string {
	switch t {
	case ToolSqlc:
		return "SQLC"
	case ToolSqlx:
		return "SQLX"
	default:
		return strings.ToUpper(string(t))
	}
}

func (a ArchitectureChoice) String() string {
	switch a {
	case ArchitectureSimple:
		return "Simple Architecture"
	case ArchitectureDDD:
		return "Domain-Driven Design (DDD)"
	case ArchitectureClean:
		return "Clean Architecture"
	case ArchitectureHexagonal:
		return "Hexagonal Architecture"
	default:
		return string(a)
	}
}

// Description methods for detailed information

func (f FrameworkChoice) Description() string {
	switch f {
	case FrameworkFiber:
		return "Fast HTTP web framework inspired by Express"
	case FrameworkGin:
		return "High-performance HTTP web framework"
	case FrameworkChi:
		return "Lightweight, idiomatic HTTP router"
	case FrameworkEcho:
		return "High performance, extensible web framework"
	default:
		return ""
	}
}

func (d DatabaseChoice) Description() string {
	switch d {
	case DatabasePostgreSQL:
		return "Advanced open-source relational database"
	case DatabaseMySQL:
		return "Popular open-source relational database"
	case DatabaseSQLite:
		return "Lightweight embedded database"
	default:
		return ""
	}
}

func (t ToolChoice) Description() string {
	switch t {
	case ToolSqlc:
		return "Generate type-safe code from SQL"
	case ToolSqlx:
		return "Extensions on database/sql for easier usage"
	default:
		return ""
	}
}

func (a ArchitectureChoice) Description() string {
	switch a {
	case ArchitectureSimple:
		return "Simple layered architecture with handlers, services, and models"
	case ArchitectureDDD:
		return "Domain-Driven Design with domain, infrastructure, application layers"
	case ArchitectureClean:
		return "Clean Architecture with entities, use cases, and adapters"
	case ArchitectureHexagonal:
		return "Hexagonal Architecture with ports and adapters pattern"
	default:
		return ""
	}
}

// GetDevOpsToolDescription returns description for DevOps tools
func GetDevOpsToolDescription(tool string) string {
	switch tool {
	case DevOpsKubernetes:
		return "Container orchestration platform"
	case DevOpsHelm:
		return "Kubernetes package manager"
	case DevOpsTerraform:
		return "Infrastructure as code tool"
	case DevOpsAnsible:
		return "IT automation and configuration management"
	default:
		return ""
	}
}

// Configuration compatibility matrix

// IsCompatible checks if choices are compatible with each other
func IsCompatible(framework FrameworkChoice, database DatabaseChoice, tool ToolChoice) bool {
	// All current combinations are compatible
	// This can be extended in the future if certain combinations are not supported
	return IsValidFramework(framework) && IsValidDatabase(database) && IsValidTool(tool)
}

// GetRecommendedTool returns recommended Tool for given database
func GetRecommendedTool(database DatabaseChoice) ToolChoice {
	switch database {
	case DatabasePostgreSQL, DatabaseMySQL:
		return ToolSqlx // SQLX works well with PostgreSQL and MySQL
	case DatabaseSQLite:
		return ToolSqlx // SQLX also supports SQLite
	default:
		return ToolSqlx // Default to SQLX
	}
}

// GetRecommendedArchitecture returns recommended architecture for project complexity
func GetRecommendedArchitecture(complexity string) ArchitectureChoice {
	switch complexity {
	case "simple", "small":
		return ArchitectureSimple
	case "medium":
		return ArchitectureClean
	case "large", "enterprise":
		return ArchitectureDDD
	default:
		return ArchitectureSimple
	}
}

// Feature flags for choices

// HasMigrations checks if the Tool supports migrations
func (t ToolChoice) HasMigrations() bool {
	switch t {
	case ToolSqlc:
		return false // SQLC uses external migration tools
	case ToolSqlx:
		return false // SQLX uses external migration tools
	default:
		return false
	}
}

// HasCodeGeneration checks if the Tool generates code
func (t ToolChoice) HasCodeGeneration() bool {
	switch t {
	case ToolSqlc:
		return true // SQLC generates Go code from SQL
	case ToolSqlx:
		return false // SQLX uses struct tags
	default:
		return false
	}
}

// SupportsRelations checks if database supports relations
func (d DatabaseChoice) SupportsRelations() bool {
	switch d {
	case DatabasePostgreSQL, DatabaseMySQL:
		return true
	case DatabaseSQLite:
		return true // SQLite supports foreign keys
	default:
		return false
	}
}

// RequiresServer checks if database requires external server
func (d DatabaseChoice) RequiresServer() bool {
	switch d {
	case DatabasePostgreSQL, DatabaseMySQL:
		return true
	case DatabaseSQLite:
		return false // SQLite is embedded
	default:
		return true
	}
}
