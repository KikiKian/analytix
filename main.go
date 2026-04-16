package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "hello\n\npress q to quit"
}

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	p.Run()
	os.Exit(0)
}
