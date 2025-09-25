// internal/tui/models/splash.go

package models

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SplashModel struct {
	finished bool
	counter  int
}

func NewSplashModel() *SplashModel {
	return &SplashModel{
		finished: false,
		counter:  0,
	}
}

func (m *SplashModel) Init() tea.Cmd {
	return tea.Tick(200*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m *SplashModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		// Skip splash on any key
		m.finished = true
		return m, nil
	case TickMsg:
		m.counter++
		if m.counter >= 10 {
			m.finished = true
			return m, nil
		}
		return m, tea.Tick(200*time.Millisecond, func(t time.Time) tea.Msg {
			return TickMsg(t)
		})
	}
	return m, nil
}

func (m *SplashModel) View() string {
	if m.finished {
		return ""
	}

	logo := `
____       ____             _
  / ___| ___ | __ )  __ _  ___| | __
 | |  _ / _ \|  _ \ / _' |/ __| |/ /
 | |_| | (_) | |_) | (_| | (__|   <
  \____|\___/|____/ \__,_|\___|_|\_\

    Backend Project Scaffolding Tool by NarmadaWeb
`

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(5)

	loading := "Loading"
	for i := 0; i < m.counter%4; i++ {
		loading += "."
	}

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		style.Render(logo),
		"",
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(loading),
	)

	return content
}

func (m *SplashModel) Finished() bool {
	return m.finished
}

type TickMsg time.Time
