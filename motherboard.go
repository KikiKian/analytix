package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func getCPU() float64 {
	info, err := cpu.Percent(0, false)
	if err != nil {
		return 0.0
	}

	if len(info) == 0 {
		return 0.0
	}

	return info[0]
}

func getRAM() float64 {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		return 0.0
	}
	return virtualMemory.UsedPercent
}
