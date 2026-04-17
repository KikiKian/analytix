package model

import (
	"fmt"

	"analytix/system"
	"analytix/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Cpu        float64
	Ram        float64
	Download   string
	Upload     string
	Network    string
	Width      int
	Height     int
	PrevRecv   uint64
	PrevSent   uint64
	DlRaw      uint64
	UlRaw      uint64
	CpuHistory []float64
	DlHistory  []float64
	UlHistory  []float64
	CmdMode    bool
	CmdInput   string
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.CmdMode {
			switch msg.String() {
			case "esc":
				m.CmdMode = false
				m.CmdInput = ""
			case "enter":
				m.CmdMode = false
				m.CmdInput = ""
			case "backspace":
				if len(m.CmdInput) > 0 {
					m.CmdInput = m.CmdInput[:len(m.CmdInput)-1]
				}
			default:
				m.CmdInput += msg.String()
			}
			return m, nil
		}
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case ":":
			m.CmdMode = true
			m.CmdInput = ""
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	case tickMsg:
		m.Cpu = system.GetCPU()
		m.Ram = system.GetRAM()
		var dl, ul uint64
		dl, ul, m.PrevRecv, m.PrevSent = system.GetNetworkSpeed(m.PrevRecv, m.PrevSent)
		m.DlRaw = dl
		m.UlRaw = ul
		m.Download = system.FormatSpeed(dl)
		m.Upload = system.FormatSpeed(ul)
		m.CpuHistory = ui.AppendHistory(m.CpuHistory, m.Cpu)
		m.DlHistory = ui.AppendHistory(m.DlHistory, float64(m.DlRaw/1024))
		m.UlHistory = ui.AppendHistory(m.UlHistory, float64(m.UlRaw/1024))
		return m, tick()
	}
	return m, nil
}

func (m Model) View() string {
	graphWidth := m.Width - 12
	if graphWidth < 10 {
		graphWidth = 10
	}

	var cmdBar string
	if m.CmdMode {
		cmdBar = ":" + m.CmdInput + "█"
	} else {
		cmdBar = "press : to enter command  |  q to quit"
	}

	main := fmt.Sprintf(
		"CPU: %.1f%%\n%s\n\nDownload: %s\n%s\n\nUpload: %s\n%s",
		m.Cpu,
		ui.RenderCPUGraph(m.CpuHistory, graphWidth, 8),
		m.Download,
		ui.RenderDownloadGraph(m.DlHistory, graphWidth, 6),
		m.Upload,
		ui.RenderUploadGraph(m.UlHistory, graphWidth, 6),
	)

	return main + "\n\n" + cmdBar
}
