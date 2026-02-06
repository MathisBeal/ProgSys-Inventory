package disk

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/disk"
)

type DiskFS struct {
	Partition disk.PartitionStat `json:"partition"`
	Usage     *disk.UsageStat    `json:"usage,omitempty"`
}

func ReadDisk() (*[]DiskFS, error) {
	parts, err := disk.Partitions(true) // true = toutes les partitions
	if err != nil {
		return nil, err
	}

	var out []DiskFS

	for _, p := range parts {
		fs := DiskFS{Partition: p}

		// Usage peut échouer (ex: pseudo-fs, permissions)
		if u, err := disk.Usage(p.Mountpoint); err == nil {
			fs.Usage = u
		}

		// Exclusion des FS spécifique au système
		switch fs.Partition.Fstype {
		case "none", "proc", "tmpfs", "overlay", "sysfs", "cgroup2", "mqueue":
			// ignoré
		default:
			out = append(out, fs)
		}
	}

	fmt.Printf("Loaded %d disk partitions\n", len(out))

	return &out, nil
}
