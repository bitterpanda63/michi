package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

type authModel struct {
	token  string
	cursor int
}

func initialAuthModel() authModel {
	return authModel{
		token:  "",
		cursor: 0,
	}
}

func (m authModel) Init() tea.Cmd {
	return nil
}

func (m authModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			viper.Set("api_token", m.token)
			viper.WriteConfig()
			return m, tea.Quit
		case "backspace":
			if len(m.token) > 0 {
				m.token = m.token[:len(m.token)-1]
			}
		default:
			m.token += msg.String()
		}
	}
	return m, nil
}

func (m authModel) View() string {
	return fmt.Sprintf(
		"API TOKEN:\n%s",
		m.token,
	)
}

func RunAuthTUI() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.michi")
	viper.SafeWriteConfig()

	p := tea.NewProgram(initialAuthModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
