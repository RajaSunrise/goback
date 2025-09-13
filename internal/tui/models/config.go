// internal/tui/models/config.go

package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NarmadaWeb/goback/internal/tui/styles"
	"github.com/NarmadaWeb/goback/pkg/config"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ConfigStep represents the current step in configuration
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

// ConfigModel handles the project configuration flow
type ConfigModel struct {
	Step         ConfigStep
	cursor       int
	choices      []string
	confirmed    bool
	cancelled    bool
	stepComplete map[ConfigStep]bool

	inputs     []textinput.Model
	focusIndex int

	framework           config.FrameworkChoice
	database            config.DatabaseChoice
	tool                config.ToolChoice
	architecture        config.ArchitectureChoice
	devopsEnabled       bool
	devopsTools         []string
	devopsToolsSelected map[string]bool

	validationErrors []string
}

// NewConfigModel creates a new configuration model
func NewConfigModel() *ConfigModel {
	m := &ConfigModel{
		Step:                StepFramework,
		stepComplete:        make(map[ConfigStep]bool),
		devopsToolsSelected: make(map[string]bool),
		inputs:              make([]textinput.Model, 4),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = styles.InputStyle
		t.CharLimit = 256
		t.Prompt = "  "

		switch i {
		case 0:
			t.Placeholder = "proyek-backend-saya"
			t.Focus()
			t.Prompt = "❯ "
		case 1:
			t.Placeholder = "github.com/user/proyek-backend-saya"
		case 2:
			t.Placeholder = "Deskripsi singkat proyek (opsional)"
		case 3:
			t.Placeholder = "./proyek-backend-saya"
		}
		m.inputs[i] = t
	}

	m.initializeProjectDetailsDefaults()
	return m
}

// Init initializes the configuration model
func (m *ConfigModel) Init() tea.Cmd {
	m.setupStep()
	return textinput.Blink
}

// Update handles incoming messages
func (m *ConfigModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Step == StepProjectDetails {
		return m.updateProjectDetailsInputs(msg)
	}
	if m.Step == StepReview {
		return m.updateReview(msg)
	}
	return m.updateChoices(msg)
}

func (m *ConfigModel) updateReview(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch strings.ToLower(msg.String()) {
		case "y", "enter":
			m.confirmed = true
			m.stepComplete[StepReview] = true
			return m, nil
		case "n", "e", "esc":
			m.Step = StepProjectDetails
			m.setupStep()
			return m, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *ConfigModel) updateProjectDetailsInputs(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m.goToPreviousStep()
		case tea.KeyEnter:
			if m.focusIndex == len(m.inputs)-1 {
				if m.validateInputs() {
					m.completeStep()
				}
				return m, nil
			}
			m.focusIndex = (m.focusIndex + 1) % len(m.inputs)
		case tea.KeyTab, tea.KeyDown:
			m.focusIndex = (m.focusIndex + 1) % len(m.inputs)
		case tea.KeyShiftTab, tea.KeyUp:
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}
		}
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		if i == m.focusIndex {
			m.inputs[i].Focus()
			m.inputs[i].Prompt = "❯ "
		} else {
			m.inputs[i].Blur()
			m.inputs[i].Prompt = "  "
		}
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	// Auto-update module path and output dir based on project name
	if m.focusIndex == 0 {
		projectName := m.inputs[0].Value()
		m.inputs[1].SetValue("github.com/user/" + projectName)
		m.inputs[3].SetValue("./" + projectName)
	}

	return m, tea.Batch(cmds...)
}

func (m *ConfigModel) updateChoices(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			return m.handleSelection()
		case " ": // Space for multi-select
			if m.Step == StepDevOpsTools {
				return m.handleSelection()
			}
		case "c":
			if m.Step == StepDevOpsTools && len(m.devopsTools) > 0 {
				m.completeStep()
				return m, nil
			}
		case "esc", "q":
			if m.Step == StepFramework {
				m.cancelled = true
				return m, nil
			}
			return m.goToPreviousStep()
		}
	}
	return m, nil
}

func (m *ConfigModel) handleSelection() (tea.Model, tea.Cmd) {
	if m.cursor >= len(m.choices) {
		return m, nil
	}
	selected := m.choices[m.cursor]

	switch m.Step {
	case StepFramework:
		m.framework = m.getFrameworkFromString(selected)
		m.completeStep()
	case StepDatabase:
		m.database = m.getDatabaseFromString(selected)
		m.completeStep()
	case StepTool:
		m.tool = m.getToolFromString(selected)
		m.completeStep()
	case StepArchitecture:
		m.architecture = m.getArchitectureFromString(selected)
		m.completeStep()
	case StepDevOpsOptions:
		m.devopsEnabled = (selected == "Ya, gunakan DevOps tools")
		m.completeStep()
	case StepDevOpsTools:
		tool := m.getDevOpsToolFromString(selected)
		if m.devopsToolsSelected[tool] {
			delete(m.devopsToolsSelected, tool)
		} else {
			m.devopsToolsSelected[tool] = true
		}
		// Rebuild the tools slice to maintain order
		m.devopsTools = []string{}
		for _, choice := range []string{"Kubernetes", "Helm", "Terraform", "Ansible"} {
			toolKey := m.getDevOpsToolFromString(choice)
			if m.devopsToolsSelected[toolKey] {
				m.devopsTools = append(m.devopsTools, toolKey)
			}
		}
	}
	return m, nil
}

func (m *ConfigModel) goToPreviousStep() (tea.Model, tea.Cmd) {
	delete(m.stepComplete, m.Step) // Mark current step as incomplete
	prevStep := m.Step - 1
	if prevStep < StepFramework {
		m.cancelled = true
		return m, nil
	}
	// Skip devops tools if devops was not enabled
	if prevStep == StepDevOpsTools && !m.devopsEnabled {
		prevStep = StepDevOpsOptions
	}
	m.Step = prevStep
	m.setupStep()
	return m, nil
}

// View renders the current configuration step
func (m *ConfigModel) View() string {
	switch m.Step {
	case StepFramework:
		return m.renderFrameworkSelection()
	case StepDatabase:
		return m.renderDatabaseSelection()
	case StepTool:
		return m.renderToolSelection()
	case StepArchitecture:
		return m.renderArchitectureSelection()
	case StepDevOpsOptions:
		return m.renderDevOpsOptions()
	case StepDevOpsTools:
		return m.renderDevOpsToolsSelection()
	case StepProjectDetails:
		return m.renderProjectDetails()
	case StepReview:
		return m.renderConfigReview()
	default:
		return "Unknown step"
	}
}

func (m *ConfigModel) setupStep() {
	m.cursor = 0
	switch m.Step {
	case StepFramework:
		m.choices = []string{"Go Fiber", "Go Gin", "Go Chi", "Go Echo"}
	case StepDatabase:
		m.choices = []string{"PostgreSQL", "MySQL", "SQLite"}
	case StepTool:
		m.choices = []string{"SQLX", "SQLC"}
	case StepArchitecture:
		m.choices = []string{"Simple Architecture", "Domain-Driven Design (DDD)", "Clean Architecture", "Hexagonal Architecture"}
	case StepDevOpsOptions:
		m.choices = []string{"Ya, gunakan DevOps tools", "Tidak menggunakan DevOps tools"}
	case StepDevOpsTools:
		m.choices = []string{"Kubernetes", "Helm", "Terraform", "Ansible"}
	case StepProjectDetails:
		m.choices = []string{} // No choices for input fields
	case StepReview:
		m.choices = []string{} // No choices for review
	}
}

func (m *ConfigModel) initializeProjectDetailsDefaults() {
	cwd, err := os.Getwd()
	defaultProjectName := "my-backend-project"
	if err == nil {
		defaultProjectName = filepath.Base(cwd)
	}
	m.inputs[0].SetValue(defaultProjectName)
	m.inputs[1].SetValue("github.com/user/" + defaultProjectName)
	m.inputs[3].SetValue("./" + defaultProjectName)
}

func (m *ConfigModel) validateInputs() bool {
	cfg := &config.ProjectConfig{
		ProjectName:  m.GetProjectName(),
		ModulePath:   m.GetModulePath(),
		OutputDir:    m.GetOutputDir(),
		Framework:    m.framework,
		Database:     m.database,
		Tool:         m.tool,
		Architecture: m.architecture,
	}
	m.validationErrors = config.ValidateProjectConfig(cfg)
	return len(m.validationErrors) == 0
}

func (m *ConfigModel) renderConfigReview() string {
	title := styles.TitleStyle.Render("🔍 Review Konfigurasi Proyek")
	subtitle := styles.SubtitleStyle.Render("Pastikan semua detail sudah benar sebelum melanjutkan.")

	var content strings.Builder
	addRow := func(icon, label, value string) {
		content.WriteString(fmt.Sprintf("%s %s: %s\n", icon, styles.FieldLabelStyle.Render(label), styles.FieldValueStyle.Render(value)))
	}

	addRow("📁", "Nama Proyek", m.GetProjectName())
	addRow("📦", "Module Path", m.GetModulePath())
	if desc := m.GetDescription(); desc != "" {
		addRow("📝", "Deskripsi", desc)
	}
	addRow("📂", "Direktori Output", m.GetOutputDir())
	content.WriteString("\n")
	addRow("🏗️ ", "Framework", m.framework.String())
	addRow("🗄️ ", "Database", m.database.String())
	addRow("🔗", "Tool", m.tool.String())
	addRow("🏛️ ", "Arsitektur", m.architecture.String())

	if m.devopsEnabled && len(m.devopsTools) > 0 {
		content.WriteString("\n")
		addRow("🚀", "DevOps Tools", strings.Join(m.devopsTools, ", "))
	}

	separator := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(strings.Repeat("─", 50))
	content.WriteString("\n" + separator + "\n\n")

	content.WriteString(styles.ConfirmStyle.Render("Apakah Anda siap untuk membuat proyek ini?"))
	box := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("86")).Padding(1, 2).Render(content.String())
	help := styles.HelpStyle.Render("✅ y/enter: Ya, Lanjutkan  |  ✏️ e/esc: Edit/Kembali  |  ❌ ctrl+c: Keluar")

	return lipgloss.JoinVertical(lipgloss.Left, title, subtitle, box, "\n", help)
}

func (m *ConfigModel) renderProjectDetails() string {
	title := styles.TitleStyle.Render("📝 Detail Proyek")
	subtitle := styles.SubtitleStyle.Render("Masukkan detail untuk proyek Anda.")
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n%s\n\n", title, subtitle)
	labels := []string{"Nama Proyek *", "Go Module Path *", "Deskripsi (Opsional)", "Direktori Output *"}

	for i := range m.inputs {
		b.WriteString(labels[i] + "\n")
		b.WriteString(m.inputs[i].View())
		b.WriteString("\n\n")
	}

	if len(m.validationErrors) > 0 {
		var errorContent strings.Builder
		errorContent.WriteString(styles.ErrorTitleStyle.Render("Validation Errors:"))
		for _, err := range m.validationErrors {
			errorContent.WriteString("\n" + styles.ErrorStyle.Render("• "+err))
		}
		b.WriteString("\n" + errorContent.String() + "\n")
	}

	help := styles.HelpStyle.Render("↑/↓/tab: navigasi • enter: lanjut • esc: kembali • ctrl+c: keluar")
	b.WriteString("\n" + help)
	return b.String()
}

func (m *ConfigModel) renderGenericChoiceView(titleText, subtitleText string) string {
	title := styles.TitleStyle.Render(titleText)
	subtitle := styles.SubtitleStyle.Render(subtitleText)
	var options strings.Builder
	for i, choice := range m.choices {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
			choice = styles.SelectedStyle.Render(choice)
		} else {
			choice = styles.OptionStyle.Render(choice)
		}
		options.WriteString(fmt.Sprintf("%s%s\n\n", cursor, choice))
	}
	help := styles.HelpStyle.Render("↑/↓: navigate • enter: select • esc: back • ctrl+c: quit")
	return lipgloss.JoinVertical(lipgloss.Left, title, subtitle, "\n", options.String(), help)
}

func (m *ConfigModel) renderFrameworkSelection() string {
	return m.renderGenericChoiceView("🏗️  Pilih Framework Backend", "Framework web yang akan digunakan untuk proyek Anda.")
}

func (m *ConfigModel) renderDatabaseSelection() string {
	return m.renderGenericChoiceView("🗄️  Pilih Database", "Database yang akan digunakan untuk proyek Anda.")
}

func (m *ConfigModel) renderToolSelection() string {
	return m.renderGenericChoiceView("🔗 Pilih Database Tool", "Tool untuk berinteraksi dengan database.")
}

func (m *ConfigModel) renderArchitectureSelection() string {
	return m.renderGenericChoiceView("🏛️  Pilih Arsitektur Proyek", "Pola arsitektur untuk struktur proyek Anda.")
}

func (m *ConfigModel) renderDevOpsOptions() string {
	return m.renderGenericChoiceView("🚀 Konfigurasi DevOps", "Apakah Anda ingin menambahkan file konfigurasi DevOps?")
}

func (m *ConfigModel) renderDevOpsToolsSelection() string {
	title := styles.TitleStyle.Render("🚀 Pilih DevOps Tools")
	subtitle := styles.SubtitleStyle.Render("Gunakan spasi untuk memilih/membatalkan, 'c' untuk melanjutkan.")
	var options strings.Builder
	for i, choice := range m.choices {
		cursor := "  "
		checkbox := "[ ]"
		tool := m.getDevOpsToolFromString(choice)
		if m.devopsToolsSelected[tool] {
			checkbox = "[x]"
		}
		if i == m.cursor {
			cursor = "> "
			choice = styles.SelectedStyle.Render(checkbox + " " + choice)
		} else {
			choice = styles.OptionStyle.Render(checkbox + " " + choice)
		}
		options.WriteString(cursor + choice + "\n\n")
	}
	helpText := "↑/↓: navigate • spasi: toggle • c: continue • esc: back • ctrl+c: quit"
	if len(m.devopsTools) == 0 {
		helpText = "↑/↓: navigate • spasi: toggle • esc: back • ctrl+c: quit"
	}
	help := styles.HelpStyle.Render(helpText)
	return lipgloss.JoinVertical(lipgloss.Left, title, subtitle, "\n", options.String(), help)
}

func (m *ConfigModel) completeStep() {
	m.stepComplete[m.Step] = true
	nextStep := m.Step + 1
	// Skip devops tools step if devops is not enabled
	if nextStep == StepDevOpsTools && !m.devopsEnabled {
		nextStep = StepProjectDetails
	}
	if nextStep > StepReview {
		nextStep = StepReview
	}
	m.Step = nextStep
	m.setupStep()
}

// Getters and other helper functions
func (m *ConfigModel) getFrameworkFromString(s string) config.FrameworkChoice {
	switch s {
	case "Go Fiber": return config.FrameworkFiber
	case "Go Gin": return config.FrameworkGin
	case "Go Chi": return config.FrameworkChi
	case "Go Echo": return config.FrameworkEcho
	default: return ""
	}
}

func (m *ConfigModel) getDatabaseFromString(s string) config.DatabaseChoice {
	switch s {
	case "PostgreSQL": return config.DatabasePostgreSQL
	case "MySQL": return config.DatabaseMySQL
	case "SQLite": return config.DatabaseSQLite
	default: return ""
	}
}

func (m *ConfigModel) getToolFromString(s string) config.ToolChoice {
	switch s {
	case "SQLC":
		return config.ToolSqlc
	case "SQLX":
		return config.ToolSqlx
	default:
		return ""
	}
}

func (m *ConfigModel) getArchitectureFromString(s string) config.ArchitectureChoice {
	switch s {
	case "Simple Architecture": return config.ArchitectureSimple
	case "Domain-Driven Design (DDD)": return config.ArchitectureDDD
	case "Clean Architecture": return config.ArchitectureClean
	case "Hexagonal Architecture": return config.ArchitectureHexagonal
	default: return ""
	}
}

func (m *ConfigModel) getDevOpsToolFromString(s string) string {
	return strings.ToLower(s)
}

func (m *ConfigModel) SetStep(step ConfigStep) {
	m.Step = step
	m.setupStep()
}

func (m *ConfigModel) IsStepComplete(step ConfigStep) bool { return m.stepComplete[step] }
func (m *ConfigModel) IsConfirmed() bool { return m.confirmed }
func (m *ConfigModel) IsCancelled() bool { return m.cancelled }
func (m *ConfigModel) GetFrameworkChoice() config.FrameworkChoice { return m.framework }
func (m *ConfigModel) GetDatabaseChoice() config.DatabaseChoice { return m.database }
func (m *ConfigModel) GetToolChoice() config.ToolChoice { return m.tool }
func (m *ConfigModel) GetArchitectureChoice() config.ArchitectureChoice { return m.architecture }
func (m *ConfigModel) GetDevOpsEnabled() bool { return m.devopsEnabled }
func (m *ConfigModel) GetDevOpsConfig() config.DevOpsConfig {
	cfg := config.DevOpsConfig{
		Enabled: m.devopsEnabled,
		Tools:   m.devopsTools,
	}
	for _, tool := range m.devopsTools {
		switch tool {
		case "kubernetes":
			cfg.Kubernetes = true
		case "helm":
			cfg.Helm = true
		case "terraform":
			cfg.Terraform = true
		case "ansible":
			cfg.Ansible = true
		}
	}
	return cfg
}

func (m *ConfigModel) GetProjectName() string { return m.inputs[0].Value() }
func (m *ConfigModel) GetModulePath() string { return m.inputs[1].Value() }
func (m *ConfigModel) GetDescription() string { return m.inputs[2].Value() }
func (m *ConfigModel) GetOutputDir() string { return m.inputs[3].Value() }
