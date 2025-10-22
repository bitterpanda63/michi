package tui

import (
	"fmt"
	"os"

	"github.com/bitterpanda63/michi/config"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
)

type authModel struct {
	provider   string
	providers  []string
	selected   int
	tokenInput textinput.Model
	err        error
}

func initialAuthModel() authModel {
	ti := textinput.New()
	ti.Placeholder = "api-token"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 32

	return authModel{
		providers:  []string{"Mistral", "Claude"},
		selected:   0,
		tokenInput: ti,
		err:        nil,
	}
}

func (m authModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m authModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyUp, tea.KeyDown:
			m.selected = (m.selected + 1) % len(m.providers)
			if msg.Type == tea.KeyUp {
				m.selected = (m.selected - 2 + len(m.providers)) % len(m.providers)
			}
			m.provider = m.providers[m.selected]
		case tea.KeyEnter:
			token := m.tokenInput.Value()
			if token == "" {
				m.err = fmt.Errorf("token cannot be empty")
				return m, nil
			}
			switch m.provider {
			case "Mistral":
				config.SetMistralAPIToken(token)
			case "Claude":
				config.SetClaudeAPIToken(token)
			}
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, nil
	}

	m.tokenInput, cmd = m.tokenInput.Update(msg)
	return m, cmd
}

func (m authModel) View() string {
	s := "Select API Provider:\n"
	for i, provider := range m.providers {
		cursor := " "
		if m.selected == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, provider)
	}
	s += "\n"
	s += fmt.Sprintf("Token for %s: %s\n", m.provider, m.tokenInput.View())
	if m.err != nil {
		s += fmt.Sprintf("\nError: %v\n", m.err)
	}
	return s
}

func RunAuthTUI() {
	p := tea.NewProgram(initialAuthModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
