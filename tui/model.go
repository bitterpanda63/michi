package tui

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/bitterpanda63/michi/config"
	"github.com/bitterpanda63/michi/mistral"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	promptInput textinput.Model
	response    strings.Builder
	renderer    *glamour.TermRenderer
	isStreaming bool
	mu          sync.Mutex
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 600
	ti.Width = 100
	ti.Cursor.Style = lipgloss.NewStyle().Underline(true).Foreground(lipgloss.Color("147"))
	ti.TextStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("147"))
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return model{
		promptInput: ti,
		renderer:    renderer,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if !m.isStreaming && m.promptInput.Value() != "" {
				m.isStreaming = true
				// Start streaming in a goroutine
				go func() {
					mistral.ChatStream(
						config.GetMistralAPIToken(),
						m.promptInput.Value(),
						&m.response,
						&m.mu,
					)
					m.isStreaming = false
				}()
			}
		}
	}

	// Only update the prompt input if we're not streaming
	if !m.isStreaming {
		m.promptInput, cmd = m.promptInput.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	var sb strings.Builder

	// Prompt input
	sb.WriteString(m.promptInput.View())
	sb.WriteString("\n\n")

	// Lock the mutex while reading the response
	m.mu.Lock()
	defer m.mu.Unlock()

	// Response (rendered as Markdown)
	if m.response.Len() > 0 {
		rendered, err := m.renderer.Render(m.response.String())
		if err != nil {
			sb.WriteString(fmt.Sprintf("Error rendering response: %v\n", err))
		} else {
			sb.WriteString(rendered)
		}
	} else if m.isStreaming {
		sb.WriteString("Thinking...")
	}

	return sb.String()
}

func RunTUI() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
