package main

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/ntcharts/linechart/streamlinechart"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cpu      float64
	ram      float64
	download string
	upload   string
	network  string
	width    int
	height   int
	prevRecv uint64
	prevSent uint64
	cpuGraph streamlinechart.Model
	netGraph streamlinechart.Model
}

func (m model) Init() tea.Cmd { return tick() }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}

	case tickMsg:
		m.cpu = getCPU()
		m.ram = getRAM()
		var dl, ul uint64
		dl, ul, m.prevRecv, m.prevSent = getNetworkSpeed(m.prevRecv, m.prevSent)
		m.upload = formatSpeed(dl)
		m.download = formatSpeed(ul)
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf(
		"CPU: %.1f%%\nRAM: %.1f%%\nDownload: %s\nUpload: %s\n\npress q to quit",
		m.cpu,
		m.ram,
		m.download,
		m.upload,
	)
}

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	p.Run()
	os.Exit(0)
}
