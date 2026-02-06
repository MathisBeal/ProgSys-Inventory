package proc

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/process"
)

type Proc struct {
	Pid    int32   `json:"pid"`
	Name   string  `json:"name"`
	User   string  `json:"user"`
	CPU    float64 `json:"cpu"`
	Memory float32 `json:"memory"`
	Status string  `json:"status"`
}

// Chargement des Processus
func ReadProc(user string) (*[]Proc, error) {
	var myProcs []Proc
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, proc := range procs {
		var myProc Proc
		myProc.User, _ = proc.Username()
		if user == "" || user == myProc.User {
			myProc.Pid = proc.Pid
			myProc.Name, _ = proc.Name()
			myProc.CPU, _ = proc.CPUPercent()
			myProc.Memory, _ = proc.MemoryPercent()
			status, _ := proc.Status()
			myProc.Status = status[0]
			myProcs = append(myProcs, myProc)
		}
	}

	fmt.Printf("Loaded %d processes\n", len(myProcs))

	return &myProcs, nil
}

func KillProcess(pid int32) error {
	myProc, err := process.NewProcess(pid)
	if err != nil {
		return err
	}

	return myProc.Kill()
}
