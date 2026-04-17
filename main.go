package main

import (
	"fmt"
	"os"

	"analytix/system"
	"analytix/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cpu        float64
	ram        float64
	download   string
	upload     string
	network    string
	width      int
	height     int
	prevRecv   uint64
	prevSent   uint64
	dlRaw      uint64
	ulRaw      uint64
	cpuHistory []float64
	dlHistory  []float64
	ulHistory  []float64
}

func (m model) Init() tea.Cmd { return tick() }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tickMsg:
		m.cpu = system.GetCPU()
		m.ram = system.GetRAM()
		var dl, ul uint64
		dl, ul, m.prevRecv, m.prevSent = system.GetNetworkSpeed(m.prevRecv, m.prevSent)
		m.dlRaw = dl
		m.ulRaw = ul
		m.download = system.FormatSpeed(dl)
		m.upload = system.FormatSpeed(ul)
		m.cpuHistory = ui.AppendHistory(m.cpuHistory, m.cpu)
		m.dlHistory = ui.AppendHistory(m.dlHistory, float64(m.dlRaw/1024))
		m.ulHistory = ui.AppendHistory(m.ulHistory, float64(m.ulRaw/1024))
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	graphWidth := m.width - 12
	if graphWidth < 10 {
		graphWidth = 10
	}
	return fmt.Sprintf(
		"CPU: %.1f%%\n%s\n\nDownload: %s\n%s\n\nUpload: %s\n%s\n\npress q to quit",
		m.cpu,
		ui.RenderCPUGraph(m.cpuHistory, graphWidth, 8),
		m.download,
		ui.RenderDownloadGraph(m.dlHistory, graphWidth, 6),
		m.upload,
		ui.RenderUploadGraph(m.ulHistory, graphWidth, 6),
	)
}

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
