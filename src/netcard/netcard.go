package netcard

import (
	"github.com/shirou/gopsutil/v4/net"
)

type NetCard struct {
	Interface net.InterfaceStat   `json:"interface"`
	IO        *net.IOCountersStat `json:"io,omitempty"`
}

func ReadNetwork(nom string) (*[]NetCard, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	counters, err := net.IOCounters(true) // true => stats par interface
	if err != nil {
		return nil, err
	}

	// index des compteurs par nom d'interface
	byName := make(map[string]net.IOCountersStat, len(counters))
	for _, c := range counters {
		byName[c.Name] = c
	}

	// fusion
	out := make([]NetCard, 0, len(ifaces))
	for _, itf := range ifaces {
		card := NetCard{Interface: itf}
		if c, ok := byName[itf.Name]; ok {
			tmp := c
			card.IO = &tmp
		}
		if nom == "" || nom == card.Interface.Name {
			out = append(out, card)
		}
	}
	return &out, nil
}
