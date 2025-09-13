// pkg/scaffolding/generator/generator.go

package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

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
		{"Generating base files", tg.generateBaseFiles}, // Menggantikan createProjectStructure & setupDependencies
		{"Generating framework files", tg.generateFrameworkFiles},
		{"Generating database config", tg.generateDatabaseConfig},
		{"Generating ORM files", tg.generateORMFiles}, // Menambahkan langkah untuk ORM
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

// DIUBAH: Fungsi helper utama untuk memproses semua template
// Fungsi ini membaca file template, membuat direktori tujuan jika belum ada,
// mengeksekusi template dengan data config, dan menulis hasilnya.
func (tg *TemplateGenerator) generateFileFromTemplate(destPath, templatePath string) error {
	// Menghapus ekstensi .tmpl dari path tujuan
	if strings.HasSuffix(destPath, ".tmpl") {
		destPath = strings.TrimSuffix(destPath, ".tmpl")
	}

	fullDestPath := filepath.Join(tg.OutputDir, destPath)
	fullTemplatePath := filepath.Join("templates", templatePath)

	// Baca konten template
	templateContent, err := os.ReadFile(fullTemplatePath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", fullTemplatePath, err)
	}

	// Buat direktori tujuan jika belum ada
	if err := os.MkdirAll(filepath.Dir(fullDestPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", fullDestPath, err)
	}

	// Buat file tujuan
	outputFile, err := os.Create(fullDestPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", fullDestPath, err)
	}
	defer outputFile.Close()

	// Parse dan eksekusi template
	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	if err := tmpl.Execute(outputFile, tg.Config); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	return nil
}

func (tg *TemplateGenerator) validateConfiguration() error {
	if tg.Config.ProjectName == "" {
		return fmt.Errorf("project name is required")
	}
	if tg.Config.OutputDir == "" {
		return fmt.Errorf("output directory is required")
	}
	if tg.Config.ModulePath == "" {
		tg.Config.ModulePath = tg.Config.ProjectName
	}
	return nil
}

// BARU: Menggabungkan beberapa langkah menjadi satu
func (tg *TemplateGenerator) generateBaseFiles() error {
	baseTemplates := map[string]string{
		"go.mod":         "base/go.mod.tmpl",
		".gitignore":     "base/gitignore.tmpl",
		"README.md":      "base/README.md.tmpl",
		"Makefile":       "base/Makefile.tmpl",
		".env":           "base/.env.tmpl",
		".env.example":   "base/env.example.tmpl",
		"cmd/migrate.go": "base/cmd/migrate.go.tmpl",
	}

	for dest, src := range baseTemplates {
		if err := tg.generateFileFromTemplate(dest, src); err != nil {
			return err
		}
	}
	return nil
}

// DIISI: Menggunakan template dari direktori `frameworks`
func (tg *TemplateGenerator) generateFrameworkFiles() error {
	framework := strings.ToLower(tg.Config.Framework.String())
	if framework == "" {
		return nil // Tidak ada framework yang dipilih
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

// DIISI: Menggunakan template dari direktori `databases`
func (tg *TemplateGenerator) generateDatabaseConfig() error {
	database := strings.ToLower(tg.Config.Database.String())
	if database == "" {
		return nil // Tidak ada database yang dipilih
	}

	templatePath := filepath.Join("databases", database, "connection.go.tmpl")
	destPath := "internal/database/connection.go"

	// Cek jika file template ada
	if _, err := os.Stat(filepath.Join("templates", templatePath)); os.IsNotExist(err) {
		return nil // Abaikan jika template tidak ada
	}

	return tg.generateFileFromTemplate(destPath, templatePath)
}

// BARU: Fungsi untuk menangani file spesifik ORM
func (tg *TemplateGenerator) generateORMFiles() error {
	orm := strings.ToLower(tg.Config.ORM.String())
	if orm == "" {
		return nil
	}

	ormDir := filepath.Join("orm", orm)
	files, err := filepath.Glob(filepath.Join("templates", ormDir, "*.tmpl"))
	if err != nil || len(files) == 0 {
		return err // atau return nil jika tidak ada file adalah normal
	}

	for _, file := range files {
		templatePath := strings.TrimPrefix(file, "templates"+string(filepath.Separator))
		var destPath string
		switch filepath.Base(templatePath) {
		case "migration.go.tmpl":
			destPath = "internal/database/migration.go"
		case "model.go.tmpl":
			// Ini mungkin lebih cocok di dalam arsitektur
			destPath = "internal/models/base_model.go"
		default:
			destPath = strings.TrimPrefix(templatePath, ormDir+string(filepath.Separator))
		}
		if err := tg.generateFileFromTemplate(destPath, templatePath); err != nil {
			return fmt.Errorf("failed to generate ORM file from %s: %w", templatePath, err)
		}
	}
	return nil
}

// DIISI: Menggunakan template dari direktori `architectures` secara rekursif
func (tg *TemplateGenerator) generateArchitectureFiles() error {
	architecture := strings.ToLower(tg.Config.Architecture.String())
	if architecture == "" {
		return nil
	}

	templateRootDir := filepath.Join("templates", "architectures", architecture)

	return filepath.Walk(templateRootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		// Membuat path relatif untuk template dan tujuan
		relPath, err := filepath.Rel(templateRootDir, path)
		if err != nil {
			return err
		}

		destPath := relPath
		templatePath := filepath.ToSlash(filepath.Join("architectures", architecture, relPath))

		return tg.generateFileFromTemplate(destPath, templatePath)
	})
}

// DIISI: Menggunakan template dari direktori `devops` secara rekursif
func (tg *TemplateGenerator) generateDevOpsFiles() error {
	if !tg.Config.DevOps.Enabled {
		return nil
	}

	for _, tool := range tg.Config.DevOps.Tools {
		toolName := strings.ToLower(tool)
		templateRootDir := filepath.Join("templates", "devops", toolName)

		// Cek apakah direktori template untuk tool ini ada
		if _, err := os.Stat(templateRootDir); os.IsNotExist(err) {
			continue // Lanjut ke tool berikutnya jika tidak ada
		}

		err := filepath.Walk(templateRootDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || !strings.HasSuffix(path, ".tmpl") {
				return nil
			}

			// Buat path relatif
			relPath, err := filepath.Rel(templateRootDir, path)
			if err != nil {
				return err
			}

			// Simpan di dalam direktori `devops/<tool>`
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

// --- Fungsi Progress & Error Callback (Tidak berubah) ---

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
