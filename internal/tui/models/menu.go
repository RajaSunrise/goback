// internal/tui/models/menu.go

package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuModel struct {
	choices  []string
	cursor   int
	selected string
}

func NewMenuModel() *MenuModel {
	return &MenuModel{
		choices: []string{
			"Mulai Proyek Baru",
			"Version",
			"Keluar",
		},
	}
}

func (m *MenuModel) Init() tea.Cmd {
	return nil
}

func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case keyUp, keyK:
			if m.cursor > 0 {
				m.cursor--
			}
		case keyDown, keyJ:
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case keyEnter, keySpace:
			m.selected = m.choices[m.cursor]
		case keyQ, keyCtrlC:
			m.selected = "Keluar"
		}
	}
	return m, nil
}

func (m *MenuModel) View() string {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86")).
		MarginBottom(2).
		Render("GoBack - Backend Project Scaffolding Tool")

	var menu string
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			choice = lipgloss.NewStyle().
				Foreground(lipgloss.Color("86")).
				Bold(true).
				Render(choice)
		}
		menu += cursor + " " + choice + "\n"
	}

	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		MarginTop(2).
		Render("↑/↓: navigate • enter: select • q: quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		menu,
		help,
	)
}

func (m *MenuModel) Selected() string {
	return m.selected
}

func (m *MenuModel) ResetSelected() {
	m.selected = ""
}
