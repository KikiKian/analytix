package main

import (
	"github.com/NimbleMarkets/ntcharts/canvas/runes"
	"github.com/NimbleMarkets/ntcharts/linechart/streamlinechart"
	"github.com/charmbracelet/lipgloss"
)

func newDualGraph(width, height int, color1, color2 string) streamlinechart.Model {
	style1 := lipgloss.NewStyle().Foreground(lipgloss.Color(color1))
	style2 := lipgloss.NewStyle().Foreground(lipgloss.Color(color2))

	slc := streamlinechart.New(width, height,
		streamlinechart.WithDataSetStyles("line1", runes.ArcLineStyle, style1),
		streamlinechart.WithDataSetStyles("line2", runes.ArcLineStyle, style2),
	)
	return slc
}

func newGraph(width, height int, color string) streamlinechart.Model {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
	slc := streamlinechart.New(width, height,
		streamlinechart.WithStyles(runes.ArcLineStyle, style),
	)
	return slc
}

func pushDual(slc *streamlinechart.Model, val1, val2 float64) {
	slc.PushDataSet("line1", val1)
	slc.PushDataSet("line2", val2)
	slc.DrawAll()
}

func pushSingle(slc *streamlinechart.Model, value float64) {
	slc.Push(value)
	slc.Draw()
}

func graphColor(percent float64) string {
	switch {
	case percent > 80:
		return "#E24B4A"
	case percent > 60:
		return "#EF9F27"
	default:
		return "#5DCAA5"
	}
}
