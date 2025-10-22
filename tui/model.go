package tui

import (
	"fmt"
	"os"

	"github.com/bitterpanda63/michi/config"
	"github.com/bitterpanda63/michi/mistral"
	"github.com/charmbracelet/bubbletea"
)

type model struct {
	prompt   string
	response string
	cursor   int
}

func initialModel() model {
	return model{
		prompt:   "",
		response: "",
		cursor:   0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			// TODO: Handle prompt submission
			mistral.ChatStream(config.GetMistralAPIToken(), m.prompt)
			return m, nil
		case "backspace":
			if len(m.prompt) > 0 {
				m.prompt = m.prompt[:len(m.prompt)-1]
			}
		default:
			m.prompt += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf(
		"> %s\n%s",
		m.prompt,
		m.response,
	)
}

func RunTUI() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
