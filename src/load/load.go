package load

import (
	"github.com/shirou/gopsutil/v4/load"
)

type LoadInfo struct {
	Avg  *load.AvgStat  `json:"avg"`
	Misc *load.MiscStat `json:"misc"`
}

func ReadLoad() (*LoadInfo, error) {
	avg, err := load.Avg()
	if err != nil {
		return nil, err
	}

	misc, err := load.Misc()
	if err != nil {
		return nil, err
	}

	return &LoadInfo{
		Avg:  avg,
		Misc: misc,
	}, nil
}
