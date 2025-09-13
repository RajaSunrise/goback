// internal/tui/model.go

package tui

import (
	"fmt"

    "github.com/NarmadaWeb/goback/pkg/config"
	"github.com/NarmadaWeb/goback/internal/tui/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// AppState represents the current state of the application
type AppState int

const (
	StateSplash AppState = iota
	StateMainMenu
	StateFrameworkSelection
	StateDatabaseSelection
	StateToolSelection
	StateArchitectureSelection
	StateDevOpsOptions
	StateDevOpsToolsSelection
	StateProjectDetails
	StateConfigReview
	StateGeneration
	StateProgress
	StateSuccess
	StateError
	StateVersion
)

// ConfigStep represents the current step in the configuration process
type ConfigStep int

const (
	StepFramework ConfigStep = iota
	StepDatabase
	StepTool
	StepArchitecture
	StepDevOpsOptions
	StepDevOpsTools
	StepProjectDetails
	StepReview
)

// MainModel is the main model for the Bubble Tea application
type MainModel struct {
	State      AppState
	WindowSize tea.WindowSizeMsg

	// Sub-models
	SplashModel   *models.SplashModel
	MenuModel     *models.MenuModel
	VersionModel  *models.VersionModel
	ConfigModel   *models.ConfigModel
	ProgressModel *models.ProgressModel

	// Shared state
	Config *config.ProjectConfig
	Error  error
}

// NewMainModel creates a new MainModel
func NewMainModel() *MainModel {
	return &MainModel{
		State:         StateSplash,
		Config:        config.NewProjectConfig(),
		SplashModel:   models.NewSplashModel(),
		MenuModel:     models.NewMenuModel(),
		VersionModel:  models.NewVersionModel(),
		ConfigModel:   models.NewConfigModel(),
		ProgressModel: models.NewProgressModel(),
	}
}

// Init initializes the main model
func (m *MainModel) Init() tea.Cmd {
	return m.SplashModel.Init()
}

// Update handles incoming messages and updates the model
func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Handle window resize globally
	if windowSizeMsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.WindowSize = windowSizeMsg
	}

	switch m.State {
	case StateSplash:
		var model tea.Model
		model, cmd = m.SplashModel.Update(msg)
		m.SplashModel = model.(*models.SplashModel)

		if m.SplashModel.Finished() {
			m.State = StateMainMenu
		}

	case StateMainMenu:
		var model tea.Model
		model, cmd = m.MenuModel.Update(msg)
		m.MenuModel = model.(*models.MenuModel)

		if m.MenuModel.Selected() != "" {
			choice := m.MenuModel.Selected()
			m.MenuModel.ResetSelected()
			switch choice {
			case "Mulai Proyek Baru":
				m.State = StateFrameworkSelection
				m.ConfigModel.SetStep(models.StepFramework)
			case "Version":
				m.State = StateVersion
			case "Keluar":
				return m, tea.Quit
			}
		}

	case StateFrameworkSelection:
		var model tea.Model
		model, cmd = m.ConfigModel.Update(msg)
		m.ConfigModel = model.(*models.ConfigModel)

		if m.ConfigModel.IsStepComplete(models.StepFramework) {
			m.Config.Framework = m.ConfigModel.GetFrameworkChoice()
			m.State = StateDatabaseSelection
			// ConfigModel automatically moves to next step
		}

	case StateDatabaseSelection:
		var model tea.Model
		model, cmd = m.ConfigModel.Update(msg)
		m.ConfigModel = model.(*models.ConfigModel)

		if m.ConfigModel.IsStepComplete(models.StepDatabase) {
			m.Config.Database = m.ConfigModel.GetDatabaseChoice()
			m.State = StateToolSelection
			// ConfigModel automatically moves to next step
		}

	case StateToolSelection:
		var model tea.Model
		model, cmd = m.ConfigModel.Update(msg)
		m.ConfigModel = model.(*models.ConfigModel)

		if m.ConfigModel.IsStepComplete(models.StepTool) {
			m.Config.Tool = m.ConfigModel.GetToolChoice()
			m.State = StateArchitectureSelection
			// ConfigModel automatically moves to next step
		}

	case StateArchitectureSelection:
		var model tea.Model
		model, cmd = m.ConfigModel.Update(msg)
		m.ConfigModel = model.(*models.ConfigModel)

		if m.ConfigModel.IsStepComplete(models.StepArchitecture) {
			m.Config.Architecture = m.ConfigModel.GetArchitectureChoice()
			m.State = StateDevOpsOptions
			// ConfigModel automatically moves to next step
		}

	case StateDevOpsOptions:
		var model tea.Model
		model, cmd = m.ConfigModel.Update(msg)
		m.ConfigModel = model.(*models.ConfigModel)

		if m.ConfigModel.IsStepComplete(models.StepDevOpsOptions) {
			m.Config.DevOps.Enabled = m.ConfigModel.GetDevOpsEnabled()
			if m.Config.DevOps.Enabled {
				m.State = StateDevOpsToolsSelection
				// ConfigModel automatically moves to DevOps tools step
			} else {
				m.State = StateProjectDetails
				// ConfigModel automatically moves to project details step
			}
		}

	case StateDevOpsToolsSelection:
		var model tea.Model
		model, cmd = m.ConfigModel.Update(msg)
		m.ConfigModel = model.(*models.ConfigModel)

		if m.ConfigModel.IsStepComplete(models.StepDevOpsTools) {
			m.Config.DevOps = m.ConfigModel.GetDevOpsConfig()
			m.State = StateProjectDetails
			// ConfigModel automatically moves to project details step
		}

	case StateProjectDetails:
		var model tea.Model
		model, cmd = m.ConfigModel.Update(msg)
		m.ConfigModel = model.(*models.ConfigModel)

		if m.ConfigModel.IsStepComplete(models.StepProjectDetails) {
			m.Config.ProjectName = m.ConfigModel.GetProjectName()
			m.Config.ModulePath = m.ConfigModel.GetModulePath()
			m.Config.Description = m.ConfigModel.GetDescription()
			m.Config.OutputDir = m.ConfigModel.GetOutputDir()
			m.State = StateConfigReview
			// ConfigModel automatically moves to review step
		}

	case StateConfigReview:
		var model tea.Model
		model, cmd = m.ConfigModel.Update(msg)
		m.ConfigModel = model.(*models.ConfigModel)

		if m.ConfigModel.IsConfirmed() {
			m.State = StateGeneration
			return m, m.ProgressModel.StartGeneration(m.Config)
		}
		if m.ConfigModel.IsCancelled() {
			m.State = StateFrameworkSelection
			m.ConfigModel.SetStep(models.StepFramework)
		}

	case StateGeneration, StateProgress:
		var model tea.Model
		model, cmd = m.ProgressModel.Update(msg)
		m.ProgressModel = model.(*models.ProgressModel)

		if m.ProgressModel.IsFinished() {
			if m.ProgressModel.IsSuccess() {
				m.State = StateSuccess
			} else {
				m.State = StateError
				m.Error = m.ProgressModel.GetError()
			}
		}

	case StateSuccess:
		// Handle success state - could show success message and exit
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			if keyMsg.String() == "q" || keyMsg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		}

	case StateError:
		// Handle error state - could show error message and allow retry
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			if keyMsg.String() == "q" || keyMsg.String() == "ctrl+c" {
				return m, tea.Quit
			}
			if keyMsg.String() == "r" {
				// Retry - go back to main menu
				m.State = StateMainMenu
				m.Error = nil
			}
		}
	case StateVersion:
		var model tea.Model
		model, cmd = m.VersionModel.Update(msg)
		m.VersionModel = model.(*models.VersionModel)

		if m.VersionModel.ShouldClose() {
			m.State = StateMainMenu
			m.VersionModel.Reset()
		}
	}

	// Handle global quit
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	return m, cmd
}

// View renders the current view of the main model
func (m *MainModel) View() string {
	var view string

	switch m.State {
	case StateSplash:
		view = m.SplashModel.View()
	case StateMainMenu:
		view = m.MenuModel.View()
	case StateFrameworkSelection,
		StateDatabaseSelection,
		StateToolSelection,
		StateArchitectureSelection,
		StateDevOpsOptions,
		StateDevOpsToolsSelection,
		StateProjectDetails,
		StateConfigReview:
		view = m.ConfigModel.View()
	case StateGeneration, StateProgress:
		view = m.ProgressModel.View()
	case StateSuccess:
		view = m.renderSuccessView()
	case StateError:
		view = m.renderErrorView()
	case StateVersion:
		view = m.VersionModel.View()
	default:
		view = "State tidak dikenal"
	}

	// Apply window constraints if available
	if m.WindowSize.Width > 0 && m.WindowSize.Height > 0 {
		style := lipgloss.NewStyle().
			Width(m.WindowSize.Width).
			Height(m.WindowSize.Height).
			Align(lipgloss.Center, lipgloss.Center)
		view = style.Render(view)
	}

	return view
}

// renderSuccessView renders the success state view
func (m *MainModel) renderSuccessView() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true).
		Margin(1).
		Padding(1).
		Border(lipgloss.RoundedBorder())

	content := fmt.Sprintf(`
✅ Proyek berhasil dibuat!

📁 Nama: %s
📂 Lokasi: %s
🏗️  Framework: %s
🗄️  Database: %s
🏛️  Arsitektur: %s

Langkah selanjutnya:
  cd %s
  go mod tidy
  go run main.go

Tekan 'q' untuk keluar
`, m.Config.ProjectName, m.Config.OutputDir, m.Config.Framework.String(),
		m.Config.Database.String(), m.Config.Architecture.String(), m.Config.OutputDir)

	return style.Render(content)
}

// renderErrorView renders the error state view
func (m *MainModel) renderErrorView() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Bold(true).
		Margin(1).
		Padding(1).
		Border(lipgloss.RoundedBorder())

	errorMsg := "Terjadi kesalahan yang tidak diketahui"
	if m.Error != nil {
		errorMsg = m.Error.Error()
	}

	content := fmt.Sprintf(`
❌ Gagal membuat proyek!

Error: %s

Tekan 'r' untuk kembali ke menu utama
Tekan 'q' untuk keluar
`, errorMsg)

	return style.Render(content)
}
