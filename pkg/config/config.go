// pkg/config/config.go

package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

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

// Note: ProjectConfig and DevOpsConfig are defined in types.go

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
		ProjectName:  defaultName,
		ModulePath:   GetConfig().DefaultModulePrefix + "/" + defaultName,
		OutputDir:    defaultOutputDir,
		Framework:    FrameworkFiber,     // Default to Fiber
		Database:     DatabasepostgresQL, // Default to postgresQL
		Tool:         ToolSqlx,           // Default to SQLX
		Architecture: ArchitectureSimple, // Default to Simple
		DevOps: DevOpsConfig{
			Enabled: false,
			Tools:   []string{},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// GetConfig returns the current application configuration
func GetConfig() *AppConfig {
	if appConfig == nil {
		appConfig = &AppConfig{}
		*appConfig = *defaultConfig // Copy defaults

		// Load from viper
		_ = viper.Unmarshal(appConfig)
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
		_ = viper.WriteConfig()
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
	v.Set("tools", cfg.Tool)
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
