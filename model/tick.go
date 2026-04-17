package model

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
