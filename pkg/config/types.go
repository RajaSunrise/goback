// pkg/config/types.go

package config

import (
	"strings"
)

// Choice types for project configuration
type (
	FrameworkChoice    string
	DatabaseChoice     string
	ORMChoice          string
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

// ORM choices
const (
	ORMGorm ORMChoice = "gorm"
	ORMSqlc ORMChoice = "sqlc"
	ORMSqlx ORMChoice = "sqlx"
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

// IsValidORM checks if ORM choice is valid
func IsValidORM(orm ORMChoice) bool {
	validORMs := []ORMChoice{
		ORMGorm,
		ORMSqlc,
		ORMSqlx,
	}

	for _, valid := range validORMs {
		if orm == valid {
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

// GetValidORMs returns list of valid ORM choices
func GetValidORMs() []ORMChoice {
	return []ORMChoice{
		ORMGorm,
		ORMSqlc,
		ORMSqlx,
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

func (o ORMChoice) String() string {
	switch o {
	case ORMGorm:
		return "GORM"
	case ORMSqlc:
		return "SQLC"
	case ORMSqlx:
		return "SQLX"
	default:
		return strings.ToUpper(string(o))
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

func (o ORMChoice) Description() string {
	switch o {
	case ORMGorm:
		return "Feature-rich ORM with associations, hooks, and migrations"
	case ORMSqlc:
		return "Generate type-safe code from SQL"
	case ORMSqlx:
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
func IsCompatible(framework FrameworkChoice, database DatabaseChoice, orm ORMChoice) bool {
	// All current combinations are compatible
	// This can be extended in the future if certain combinations are not supported
	return IsValidFramework(framework) && IsValidDatabase(database) && IsValidORM(orm)
}

// GetRecommendedORM returns recommended ORM for given database
func GetRecommendedORM(database DatabaseChoice) ORMChoice {
	switch database {
	case DatabasePostgreSQL, DatabaseMySQL:
		return ORMGorm // GORM works well with PostgreSQL and MySQL
	case DatabaseSQLite:
		return ORMGorm // GORM also supports SQLite
	default:
		return ORMGorm // Default to GORM
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

// HasMigrations checks if the ORM supports migrations
func (o ORMChoice) HasMigrations() bool {
	switch o {
	case ORMGorm:
		return true
	case ORMSqlc:
		return false // SQLC uses external migration tools
	case ORMSqlx:
		return false // SQLX uses external migration tools
	default:
		return false
	}
}

// HasCodeGeneration checks if the ORM generates code
func (o ORMChoice) HasCodeGeneration() bool {
	switch o {
	case ORMGorm:
		return false // GORM uses reflection
	case ORMSqlc:
		return true // SQLC generates Go code from SQL
	case ORMSqlx:
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
