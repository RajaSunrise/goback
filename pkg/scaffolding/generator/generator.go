// pkg/scaffolding/generator/generator.go

package generator

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/NarmadaWeb/goback/pkg/config"
)

// TemplateGenerator handles project generation from templates
type TemplateGenerator struct {
	Config           *config.ProjectConfig
	OutputDir        string
	progressCallback func(step int, message string)
	errorCallback    func(step int, err error)
	currentStep      int
	totalSteps       int
}

// NewTemplateGenerator creates a new template generator
func NewTemplateGenerator(cfg *config.ProjectConfig) *TemplateGenerator {
	return &TemplateGenerator{
		Config:     cfg,
		OutputDir:  cfg.OutputDir,
		totalSteps: 7,
	}
}

// SetProgressCallback sets the progress callback
func (tg *TemplateGenerator) SetProgressCallback(callback func(step int, message string)) {
	tg.progressCallback = callback
}

// SetErrorCallback sets the error callback
func (tg *TemplateGenerator) SetErrorCallback(callback func(step int, err error)) {
	tg.errorCallback = callback
}

// Generate generates the project structure and files
func (tg *TemplateGenerator) Generate() error {
	steps := []struct {
		name    string
		handler func() error
	}{
		{"Validating configuration", tg.validateConfiguration},
		{"Generating base files", tg.generateBaseFiles},
		{"Generating framework files", tg.generateFrameworkFiles},
		{"Generating database config", tg.generateDatabaseConfig},
		{"Generating Tool files", tg.generateToolFiles},
		{"Generating architecture files", tg.generateArchitectureFiles},
		{"Generating DevOps files", tg.generateDevOpsFiles},
	}

	for i, step := range steps {
		tg.currentStep = i
		tg.reportProgress(i, fmt.Sprintf("Step %d/%d: %s", i+1, len(steps), step.name))

		if err := step.handler(); err != nil {
			tg.reportError(i, err)
			return fmt.Errorf("step %d (%s) failed: %w", i+1, step.name, err)
		}
	}

	tg.reportProgress(len(steps), "Project generation completed successfully!")
	return nil
}

// generateFileFromTemplate is the main helper function for processing templates.
// It reads a template file, creates the destination directory if it doesn't exist,
// executes the template with the config data, and writes the result.
func (tg *TemplateGenerator) generateFileFromTemplate(destPath, templatePath string) error {
	// Remove .tmpl extension from destination path
	destPath = strings.TrimSuffix(destPath, ".tmpl")

	fullDestPath := filepath.Join(tg.OutputDir, destPath)
	fullTemplatePath := filepath.Join("templates", templatePath)

	// Read template content
	templateContent, err := os.ReadFile(fullTemplatePath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", fullTemplatePath, err)
	}

	// Create destination directory if it doesn't exist
	if err2 := os.MkdirAll(filepath.Dir(fullDestPath), 0755); err2 != nil {
		return fmt.Errorf("failed to create directory for %s: %w", fullDestPath, err2)
	}

	// Create destination file
	outputFile, err := os.Create(fullDestPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", fullDestPath, err)
	}
	defer outputFile.Close()

	// Custom template functions
	funcMap := template.FuncMap{
		"title":      strings.ToTitle,
		"toTitle":    strings.ToTitle,
		"snakeCase":  strcase.ToSnake,
		"kebabCase":  strcase.ToKebab,
		"upper":      strings.ToUpper,
		"replaceAll": func(s, old, new string) string { return strings.Replace(s, old, new, -1) },
		"b64enc":     func(s string) string { return base64.StdEncoding.EncodeToString([]byte(s)) },
		"default": func(val string, def string) string {
			if val == "" {
				return def
			}
			return val
		},
	}

	// Parse and execute template
	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	if err := tmpl.Execute(outputFile, tg.Config); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	return nil
}

func (tg *TemplateGenerator) validateConfiguration() error {
	errors := config.ValidateProjectConfig(tg.Config)
	if len(errors) > 0 {
		// Join the errors into a single string to return as an error
		return fmt.Errorf("configuration validation failed: %s", strings.Join(errors, ", "))
	}
	return nil
}

// generateBaseFiles generates the base project files.
func (tg *TemplateGenerator) generateBaseFiles() error {
	baseTemplates := map[string]string{
		"go.mod":         "base/go.mod.tmpl",
		".gitignore":     "base/gitignore.tmpl",
		"README.md":      "base/README.md.tmpl",
		"Makefile":       "base/Makefile.tmpl",
		".env":           "base/.env.tmpl",
		".env.example":   "base/env.example.tmpl",
		"cmd/migrate.go": "base/cmd/migrate.go.tmpl",
		".golangci.yml":  "base/.golangci.yml.tmpl",
	}

	for dest, src := range baseTemplates {
		if err := tg.generateFileFromTemplate(dest, src); err != nil {
			return err
		}
	}
	return nil
}

// generateFrameworkFiles generates the framework-specific files.
func (tg *TemplateGenerator) generateFrameworkFiles() error {
	framework := string(tg.Config.Framework)
	if framework == "" {
		return nil // No framework selected
	}

	frameworkDir := filepath.Join("frameworks", framework)
	files, err := filepath.Glob(filepath.Join("templates", frameworkDir, "*.tmpl"))
	if err != nil {
		return err
	}

	for _, file := range files {
		templatePath := strings.TrimPrefix(file, "templates"+string(filepath.Separator))

		var destPath string
		switch filepath.Base(templatePath) {
		case "main.go.tmpl":
			destPath = "cmd/api/main.go"
		case "routes.go.tmpl":
			destPath = "internal/routes/routes.go"
		case "config.go.tmpl":
			destPath = "internal/config/framework.go"
		case "handlers.go.tmpl":
			destPath = "internal/handlers/handlers.go"
		case "middleware.go.tmpl":
			destPath = "internal/middleware/middleware.go"
		default:
			destPath = strings.TrimPrefix(templatePath, frameworkDir+string(filepath.Separator))
		}

		if err := tg.generateFileFromTemplate(destPath, templatePath); err != nil {
			return fmt.Errorf("failed to generate framework file from %s: %w", templatePath, err)
		}
	}

	return nil
}

// generateDatabaseConfig generates the database configuration files.
func (tg *TemplateGenerator) generateDatabaseConfig() error {
	database := string(tg.Config.Database)
	if database == "" {
		return nil // No database selected
	}

	var templatePath string
	if tg.Config.Tool == config.ToolGorm {
		templatePath = filepath.Join("tools", "gorm", "connection.go.tmpl")
	} else {
		templatePath = filepath.Join("databases", database, "connection.go.tmpl")
	}

	destPath := "internal/database/connection.go"

	// Check if the template file exists
	if _, err := os.Stat(filepath.Join("templates", templatePath)); os.IsNotExist(err) {
		return nil // Ignore if template doesn't exist
	}

	return tg.generateFileFromTemplate(destPath, templatePath)
}

// generateToolFiles generates the Tool-specific files.
func (tg *TemplateGenerator) generateToolFiles() error {
	tool := strings.ToLower(tg.Config.Tool.String())
	if tool == "" {
		return nil
	}

	toolDir := filepath.Join("tools", tool)
	files, err := filepath.Glob(filepath.Join("templates", toolDir, "*.tmpl"))
	if err != nil || len(files) == 0 {
		return err // atau return nil jika tidak ada file adalah normal
	}

	for _, file := range files {
		templatePath := strings.TrimPrefix(file, "templates"+string(filepath.Separator))
		var destPath string
		switch filepath.Base(templatePath) {
		case "model.go.tmpl":
			// This might be better suited within the architecture templates
			destPath = "internal/models/base_model.go"
		case "sqlc.yaml.tmpl":
			destPath = "sqlc.yaml"
		default:
			destPath = strings.TrimPrefix(templatePath, toolDir+string(filepath.Separator))
		}
		if err := tg.generateFileFromTemplate(destPath, templatePath); err != nil {
			return fmt.Errorf("failed to generate Tool file from %s: %w", templatePath, err)
		}
	}
	return nil
}

// generateArchitectureFiles generates the architecture-specific files recursively.
func (tg *TemplateGenerator) generateArchitectureFiles() error {
	architecture := string(tg.Config.Architecture)
	if architecture == "" {
		return nil
	}

	templateRootDir := filepath.Join("templates", "architectures", architecture)

	return filepath.Walk(templateRootDir, func(path string, info os.FileInfo, err_ error) error {
		if err_ != nil {
			return err_
		}
		if info.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		// Create relative paths for the template and destination
		relPath, err := filepath.Rel(templateRootDir, path)
		if err != nil {
			return err
		}

		destPath := relPath
		templatePath := filepath.ToSlash(filepath.Join("architectures", architecture, relPath))

		return tg.generateFileFromTemplate(destPath, templatePath)
	})
}

// generateDevOpsFiles generates the DevOps-specific files recursively.
func (tg *TemplateGenerator) generateDevOpsFiles() error {
	if !tg.Config.DevOps.Enabled {
		return nil
	}

	for _, tool := range tg.Config.DevOps.Tools {
		toolName := strings.ToLower(tool)
		templateRootDir := filepath.Join("templates", "devops", toolName)

		// Check if the template directory for this tool exists
		if _, err := os.Stat(templateRootDir); os.IsNotExist(err) {
			continue // Continue to the next tool if it doesn't exist
		}

		err := filepath.Walk(templateRootDir, func(path string, info os.FileInfo, err_ error) error {
			if err_ != nil {
				return err_
			}
			if info.IsDir() || !strings.HasSuffix(path, ".tmpl") {
				return nil
			}

			// Create relative path
			relPath, err := filepath.Rel(templateRootDir, path)
			if err != nil {
				return err
			}

			// Store in the 'devops/<tool>' directory
			destPath := filepath.Join("devops", toolName, relPath)
			templatePath := filepath.ToSlash(filepath.Join("devops", toolName, relPath))

			return tg.generateFileFromTemplate(destPath, templatePath)
		})

		if err != nil {
			return fmt.Errorf("failed to generate files for DevOps tool %s: %w", toolName, err)
		}
	}
	return nil
}

// reportProgress sends a progress update via the callback.
func (tg *TemplateGenerator) reportProgress(step int, message string) {
	if tg.progressCallback != nil {
		tg.progressCallback(step, message)
	}
}

func (tg *TemplateGenerator) reportError(step int, err error) {
	if tg.errorCallback != nil {
		tg.errorCallback(step, err)
	}
}
