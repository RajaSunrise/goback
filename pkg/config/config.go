// pkg/config/config.go

package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// ProjectConfig represents the complete project configuration
type ProjectConfig struct {
	ProjectName  string             `json:"project_name" yaml:"project_name" validate:"required,min=1"`
	ModulePath   string             `json:"module_path" yaml:"module_path" validate:"required"`
	Description  string             `json:"description" yaml:"description"`
	OutputDir    string             `json:"output_dir" yaml:"output_dir" validate:"required"`
	Framework    FrameworkChoice    `json:"framework" yaml:"framework" validate:"required"`
	Database     DatabaseChoice     `json:"database" yaml:"database" validate:"required"`
	ORM          ORMChoice          `json:"orm" yaml:"orm" validate:"required"`
	Architecture ArchitectureChoice `json:"architecture" yaml:"architecture" validate:"required"`
	DevOps       DevOpsConfig       `json:"devops" yaml:"devops"`
	CreatedAt    time.Time          `json:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" yaml:"updated_at"`
}

// DevOpsConfig represents DevOps tools configuration
type DevOpsConfig struct {
	Enabled    bool     `json:"enabled" yaml:"enabled"`
	Tools      []string `json:"tools" yaml:"tools"`
	Kubernetes bool     `json:"kubernetes" yaml:"kubernetes"`
	Helm       bool     `json:"helm" yaml:"helm"`
	Terraform  bool     `json:"terraform" yaml:"terraform"`
	Ansible    bool     `json:"ansible" yaml:"ansible"`
}

// AppConfig represents application-level configuration
type AppConfig struct {
	DefaultOutputDir    string `json:"default_output_dir" yaml:"default_output_dir"`
	DefaultModulePrefix string `json:"default_module_prefix" yaml:"default_module_prefix"`
	DefaultAuthor       string `json:"default_author" yaml:"default_author"`
	AnimationSpeed      int    `json:"animation_speed" yaml:"animation_speed"`
	ShowSplashScreen    bool   `json:"show_splash_screen" yaml:"show_splash_screen"`
	AutoSave            bool   `json:"auto_save" yaml:"auto_save"`
	Theme               string `json:"theme" yaml:"theme"`
}

// Default configuration values
var (
	defaultConfig = &AppConfig{
		DefaultOutputDir:    "./",
		DefaultModulePrefix: "github.com/user",
		DefaultAuthor:       "",
		AnimationSpeed:      100,
		ShowSplashScreen:    true,
		AutoSave:            true,
		Theme:               "default",
	}

	appConfig *AppConfig
)

// NewProjectConfig creates a new project configuration with defaults
func NewProjectConfig() *ProjectConfig {
	now := time.Now()

	// Get current directory name as default project name
	currentDir, _ := os.Getwd()
	defaultName := filepath.Base(currentDir)

	// Get default output directory from app config
	defaultOutputDir := GetConfig().DefaultOutputDir
	if defaultOutputDir == "./" {
		defaultOutputDir = "./" + defaultName
	}

	return &ProjectConfig{
		ProjectName: defaultName,
		ModulePath:  GetConfig().DefaultModulePrefix + "/" + defaultName,
		OutputDir:   defaultOutputDir,
		DevOps:      DevOpsConfig{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// GetConfig returns the current application configuration
func GetConfig() *AppConfig {
	if appConfig == nil {
		appConfig = &AppConfig{}
		*appConfig = *defaultConfig // Copy defaults

		// Load from viper
		if err := viper.Unmarshal(appConfig); err == nil {
			// Successfully loaded from config file
		}
	}
	return appConfig
}

// InitDefaults initializes default configuration values in viper
func InitDefaults() {
	// Set default values in viper
	viper.SetDefault("default_output_dir", defaultConfig.DefaultOutputDir)
	viper.SetDefault("default_module_prefix", defaultConfig.DefaultModulePrefix)
	viper.SetDefault("default_author", defaultConfig.DefaultAuthor)
	viper.SetDefault("animation_speed", defaultConfig.AnimationSpeed)
	viper.SetDefault("show_splash_screen", defaultConfig.ShowSplashScreen)
	viper.SetDefault("auto_save", defaultConfig.AutoSave)
	viper.SetDefault("theme", defaultConfig.Theme)

	// Try to create config file if it doesn't exist
	createDefaultConfigFile()
}

// createDefaultConfigFile creates a default configuration file if it doesn't exist
func createDefaultConfigFile() {
	configHome, err := os.UserHomeDir()
	if err != nil {
		return
	}

	configPath := filepath.Join(configHome, ".goback.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default config file
		viper.SetConfigFile(configPath)
		viper.WriteConfig()
	}
}

// SaveConfig saves the current configuration to file
func SaveConfig() error {
	return viper.WriteConfig()
}

// UpdateConfig updates application configuration
func UpdateConfig(key string, value interface{}) error {
	viper.Set(key, value)
	return SaveConfig()
}

// ValidateProjectConfig validates a project configuration
func ValidateProjectConfig(cfg *ProjectConfig) []string {
	var errors []string

	// Required fields validation
	if cfg.ProjectName == "" {
		errors = append(errors, "Project name is required")
	}

	if cfg.ModulePath == "" {
		errors = append(errors, "Module path is required")
	}

	if cfg.OutputDir == "" {
		errors = append(errors, "Output directory is required")
	}

	if cfg.Framework == "" {
		errors = append(errors, "Framework selection is required")
	}

	if cfg.Database == "" {
		errors = append(errors, "Database selection is required")
	}

	if cfg.ORM == "" {
		errors = append(errors, "ORM selection is required")
	}

	if cfg.Architecture == "" {
		errors = append(errors, "Architecture selection is required")
	}

	// Validate choices are valid
	if !IsValidFramework(cfg.Framework) {
		errors = append(errors, "Invalid framework choice")
	}

	if !IsValidDatabase(cfg.Database) {
		errors = append(errors, "Invalid database choice")
	}

	if !IsValidORM(cfg.ORM) {
		errors = append(errors, "Invalid ORM choice")
	}

	if !IsValidArchitecture(cfg.Architecture) {
		errors = append(errors, "Invalid architecture choice")
	}

	// DevOps validation
	if cfg.DevOps.Enabled && len(cfg.DevOps.Tools) == 0 {
		errors = append(errors, "At least one DevOps tool must be selected when DevOps is enabled")
	}

	// Path validation
	if !isValidPath(cfg.OutputDir) {
		errors = append(errors, "Invalid output directory path")
	}

	return errors
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

// GetProjectConfigDefaults returns default values for project configuration
func GetProjectConfigDefaults() map[string]interface{} {
	return map[string]interface{}{
		"framework":      "fiber",
		"database":       "postgresql",
		"orm":            "gorm",
		"architecture":   "simple",
		"devops_enabled": false,
	}
}

// LoadProjectConfig loads project configuration from file
func LoadProjectConfig(filepath string) (*ProjectConfig, error) {
	v := viper.New()
	v.SetConfigFile(filepath)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg ProjectConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// SaveProjectConfig saves project configuration to file
func SaveProjectConfig(cfg *ProjectConfig, filepath string) error {
	v := viper.New()
	v.Set("project_name", cfg.ProjectName)
	v.Set("module_path", cfg.ModulePath)
	v.Set("description", cfg.Description)
	v.Set("output_dir", cfg.OutputDir)
	v.Set("framework", cfg.Framework)
	v.Set("database", cfg.Database)
	v.Set("orm", cfg.ORM)
	v.Set("architecture", cfg.Architecture)
	v.Set("devops", cfg.DevOps)
	v.Set("created_at", cfg.CreatedAt)
	v.Set("updated_at", cfg.UpdatedAt)

	v.SetConfigFile(filepath)
	return v.WriteConfig()
}

// GetConfigDir returns the configuration directory
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, ".config", "goback")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

// GetRecentProjects returns list of recently created projects
func GetRecentProjects() []string {
	projects := viper.GetStringSlice("recent_projects")
	if projects == nil {
		return []string{}
	}
	return projects
}

// AddRecentProject adds a project to recent projects list
func AddRecentProject(projectPath string) error {
	projects := GetRecentProjects()

	// Remove if already exists
	for i, p := range projects {
		if p == projectPath {
			projects = append(projects[:i], projects[i+1:]...)
			break
		}
	}

	// Add to front
	projects = append([]string{projectPath}, projects...)

	// Limit to 10 recent projects
	if len(projects) > 10 {
		projects = projects[:10]
	}

	viper.Set("recent_projects", projects)
	return SaveConfig()
}

// ResetConfig resets configuration to defaults
func ResetConfig() error {
	*appConfig = *defaultConfig

	// Clear viper and set defaults
	for key := range viper.AllSettings() {
		viper.Set(key, nil)
	}

	InitDefaults()
	return SaveConfig()
}

// GetConfigSummary returns a summary of current configuration
func GetConfigSummary() map[string]interface{} {
	cfg := GetConfig()
	return map[string]interface{}{
		"default_output_dir":    cfg.DefaultOutputDir,
		"default_module_prefix": cfg.DefaultModulePrefix,
		"default_author":        cfg.DefaultAuthor,
		"theme":                 cfg.Theme,
		"show_splash_screen":    cfg.ShowSplashScreen,
		"animation_speed":       cfg.AnimationSpeed,
		"auto_save":             cfg.AutoSave,
		"recent_projects_count": len(GetRecentProjects()),
	}
}
