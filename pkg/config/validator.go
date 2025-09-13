package config

import (
	"fmt"
	"path/filepath"
	"regexp"
	"errors"

	"github.com/go-playground/validator/v10"
)

// ValidateProjectConfig validates the project configuration using struct tags
func ValidateProjectConfig(config *ProjectConfig) []string {
	var error []string
	validate := validator.New()

	// Register custom validation for module path
	_ = validate.RegisterValidation("modulepath", validateModulePath)

	err := validate.Struct(config)
	if err != nil {
		for _, err := range func() validator.ValidationErrors {
			var target validator.ValidationErrors
			_ = errors.As(err, &target)
			return target
		}() {
			error = append(error, formatValidationError(err))
		}
	}

	// Additional custom validations
	if !IsValidFramework(config.Framework) {
		error = append(error, "Invalid framework choice.")
	}
	if !IsValidDatabase(config.Database) {
		error = append(error, "Invalid database choice.")
	}
	if !IsValidTool(config.Tool) {
		error = append(error, "Invalid Tool choice.")
	}
	if !IsValidArchitecture(config.Architecture) {
		error = append(error, "Invalid architecture choice.")
	}

	if config.DevOps.Enabled && len(config.DevOps.Tools) == 0 {
		error = append(error, "At least one DevOps tool must be selected when DevOps is enabled.")
	}

	if !isValidPath(config.OutputDir) {
		error = append(error, "Invalid output directory path.")
	}

	return error
}

// formatValidationError creates a user-friendly error message
func formatValidationError(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()

	switch field {
	case "ProjectName":
		return "Project name is required and cannot be empty."
	case "ModulePath":
		if tag == "required" {
			return "Go module path is required."
		}
		if tag == "modulepath" {
			return "Invalid Go module path format. (e.g., github.com/user/project)"
		}
	case "OutputDir":
		return "Output directory is required."
	case "Framework":
		return "A framework selection is required."
	case "Database":
		return "A database selection is required."
	case "Tool":
		return "A Tool selection is required."
	case "Architecture":
		return "An architecture selection is required."
	}
	return fmt.Sprintf("Validation failed on field '%s' with tag '%s'", field, tag)
}

// validateModulePath is a custom validator for Go module paths
func validateModulePath(fl validator.FieldLevel) bool {
	// A simple regex to validate module path format.
	// This is not exhaustive but covers common cases like github.com/user/repo
	re := regexp.MustCompile(`^([\w-]+\.)+[\w-]+(/[\w-]+)+$`)
	return re.MatchString(fl.Field().String())
}

// isValidPath checks if a path is valid
func isValidPath(path string) bool {
	if path == "" {
		return false
	}

	// Basic path validation - can be enhanced
	if filepath.IsAbs(path) {
		return true
	}

	// Relative paths are also valid
	return !filepath.IsAbs(path)
}
