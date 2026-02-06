package cpu

import (
	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUCore struct {
	Info  cpu.InfoStat   `json:"info"`
	Usage float64        `json:"usage_percent"`
	Times *cpu.TimesStat `json:"times,omitempty"`
}

type CPUInfo struct {
	Cores []CPUCore `json:"cores"`
}

func ReadCPU() (*CPUInfo, error) {
	infos, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	usage, err := cpu.Percent(0, true) // usage par core
	if err != nil {
		return nil, err
	}

	times, err := cpu.Times(true) // stats par core
	if err != nil {
		return nil, err
	}

	out := make([]CPUCore, 0, len(infos))
	for i, info := range infos {
		core := CPUCore{
			Info: info,
		}

		if i < len(usage) {
			core.Usage = usage[i]
		}
		if i < len(times) {
			tmp := times[i] // adresse stable
			core.Times = &tmp
		}

		out = append(out, core)
	}

	return &CPUInfo{Cores: out}, nil
}
