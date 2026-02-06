package memory

import (
	"github.com/shirou/gopsutil/v4/mem"
)

type MemInfo struct {
	Virtual *mem.VirtualMemoryStat `json:"virtual"`
	Swap    *mem.SwapMemoryStat    `json:"swap"`
}

func ReadMemory() (*MemInfo, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	sm, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}

	return &MemInfo{
		Virtual: vm,
		Swap:    sm,
	}, nil
}
