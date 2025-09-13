// internal/tui/models/progress.go

package models

import (
	"fmt"
	"time"

	"github.com/NarmadaWeb/goback/pkg/config"
	"github.com/NarmadaWeb/goback/pkg/scaffolding/generator"
	"github.com/NarmadaWeb/goback/internal/tui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ProgressModel handles project generation progress
type ProgressModel struct {
	config     *config.ProjectConfig
	generator  *generator.TemplateGenerator
	finished   bool
	success    bool
	error      error
	progress   float64
	currentMsg string
	steps      []string
	stepIndex  int
	startTime  time.Time
}

// Progress messages
type progressMsg struct {
	step    int
	message string
	done    bool
	err     error
}

type generationCompleteMsg struct {
	success bool
	err     error
}

// NewProgressModel creates a new progress model
func NewProgressModel() *ProgressModel {
	return &ProgressModel{
		steps: []string{
			"Validating configuration...",
			"Creating project structure...",
			"Generating framework files...",
			"Setting up database configuration...",
			"Applying architecture pattern...",
			"Installing dependencies...",
			"Generating DevOps files...",
			"Finalizing project...",
		},
	}
}

// Init initializes the progress model
func (m *ProgressModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the progress model
func (m *ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case progressMsg:
		m.stepIndex = msg.step
		m.currentMsg = msg.message
		m.progress = float64(msg.step) / float64(len(m.steps))

		if msg.err != nil {
			m.finished = true
			m.success = false
			m.error = msg.err
			return m, nil
		}

		if msg.done {
			m.finished = true
			m.success = true
			return m, nil
		}

		return m, nil

	case generationCompleteMsg:
		m.finished = true
		m.success = msg.success
		m.error = msg.err
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the progress screen
func (m *ProgressModel) View() string {
	if m.finished {
		if m.success {
			return m.renderSuccess()
		}
		return m.renderError()
	}

	return m.renderProgress()
}

// renderProgress renders the ongoing progress
func (m *ProgressModel) renderProgress() string {
	title := styles.TitleStyle.Render("🚧 Generating Project...")

	// Progress bar
	barWidth := 50
	filled := int(m.progress * float64(barWidth))
	bar := "█"
	empty := "░"

	filledBar := ""
	for i := 0; i < filled; i++ {
		filledBar += bar
	}
	emptyBar := ""
	for i := 0; i < barWidth-filled; i++ {
		emptyBar += empty
	}

	progressBar := lipgloss.NewStyle().
		Foreground(styles.AccentColor).
		Render(fmt.Sprintf("[%s%s] %.0f%%",
			lipgloss.NewStyle().Render(filledBar),
			lipgloss.NewStyle().Faint(true).Render(emptyBar),
			m.progress*100))

	// Current step
	currentStep := ""
	if m.stepIndex < len(m.steps) {
		currentStep = m.steps[m.stepIndex]
	}
	if m.currentMsg != "" {
		currentStep = m.currentMsg
	}

	// Steps list
	var stepsList []string
	for i, step := range m.steps {
		var status string
		var style lipgloss.Style

		if i < m.stepIndex {
			status = "✅"
			style = styles.SuccessStyle
		} else if i == m.stepIndex {
			status = "🔄"
			style = styles.AccentStyle
		} else {
			status = "⏳"
			style = styles.MutedStyle
		}

		stepsList = append(stepsList, fmt.Sprintf("%s %s", status, style.Render(step)))
	}

	// Time elapsed
	elapsed := ""
	if !m.startTime.IsZero() {
		elapsed = fmt.Sprintf("Elapsed: %s", time.Since(m.startTime).Round(time.Second))
	}

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		progressBar,
		"",
		styles.AccentStyle.Render("Current: "+currentStep),
		"",
		lipgloss.JoinVertical(lipgloss.Left, stepsList...),
		"",
		styles.MutedStyle.Render(elapsed),
		"",
		styles.HelpStyle.Render("Press Ctrl+C to cancel"),
	)

	// Center the content
	return lipgloss.Place(
		80, 25,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}

// renderSuccess renders the success screen
func (m *ProgressModel) renderSuccess() string {
	if m.config == nil {
		return "Generation completed successfully!"
	}

	title := styles.SuccessStyle.Render("✅ Project Generated Successfully!")

	projectInfo := fmt.Sprintf(`
📁 Project Name: %s
📂 Location: %s
🚀 Framework: %s
🗄️  Database: %s
⚙️  ORM: %s
🏗️  Architecture: %s`,
		m.config.ProjectName,
		m.config.OutputDir,
		m.config.Framework,
		m.config.Database,
		m.config.Tool,
		m.config.Architecture,
	)

	if m.config.DevOps.Enabled {
		devopsInfo := fmt.Sprintf("\n🔧 DevOps: %v", m.config.DevOps.Tools)
		projectInfo += devopsInfo
	}

	nextSteps := `
Next Steps:
1. cd ` + m.config.OutputDir + `
2. go mod tidy
3. go run main.go

Your project is ready to use!`

	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		"",
		styles.InfoStyle.Render(projectInfo),
		"",
		styles.DescriptionStyle.Render(nextSteps),
		"",
		styles.HelpStyle.Render("Press any key to exit"),
	)

	return lipgloss.Place(
		80, 25,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}

// renderError renders the error screen
func (m *ProgressModel) renderError() string {
	title := styles.ErrorStyle.Render("❌ Generation Failed!")

	errorMsg := "An unknown error occurred"
	if m.error != nil {
		errorMsg = m.error.Error()
	}

	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		"",
		styles.ErrorStyle.Render("Error: "+errorMsg),
		"",
		styles.DescriptionStyle.Render("The project generation failed. Please check your configuration and try again."),
		"",
		styles.HelpStyle.Render("Press 'r' to retry or 'q' to quit"),
	)

	return lipgloss.Place(
		80, 25,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}

// StartGeneration starts the project generation process
func (m *ProgressModel) StartGeneration(config *config.ProjectConfig) tea.Cmd {
	m.config = config
	m.startTime = time.Now()
	m.generator = generator.NewTemplateGenerator(config)

	return tea.Batch(
		func() tea.Msg {
			return progressMsg{step: 0, message: "Starting generation..."}
		},
		m.runGeneration(),
	)
}

// runGeneration runs the actual generation process
func (m *ProgressModel) runGeneration() tea.Cmd {
	return func() tea.Msg {
		if m.generator == nil {
			return generationCompleteMsg{
				success: false,
				err:     fmt.Errorf("generator not initialized"),
			}
		}

		// Set up progress callback
		m.generator.SetProgressCallback(func(step int, message string) {
			// This would ideally send a message back to the model
			// For now, we'll track internally
		})

		// Run the generation
		if err := m.generator.Generate(); err != nil {
			return generationCompleteMsg{
				success: false,
				err:     err,
			}
		}

		return generationCompleteMsg{
			success: true,
			err:     nil,
		}
	}
}

// Getter methods for state checking
func (m *ProgressModel) IsFinished() bool {
	return m.finished
}

func (m *ProgressModel) IsSuccess() bool {
	return m.success
}

func (m *ProgressModel) GetError() error {
	return m.error
}

// Simulate progress updates - this would be called by the generator
func (m *ProgressModel) updateProgress(step int, message string) tea.Cmd {
	return func() tea.Msg {
		return progressMsg{
			step:    step,
			message: message,
		}
	}
}
