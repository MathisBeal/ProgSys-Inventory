package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evoliatis/buildup/cpu"
	"github.com/evoliatis/buildup/disk"
	"github.com/evoliatis/buildup/load"
	"github.com/evoliatis/buildup/memory"
	"github.com/evoliatis/buildup/netcard"
	"github.com/evoliatis/buildup/proc"
)

func HealthHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// Liste des processus
func PSHandler(w http.ResponseWriter, r *http.Request) {
	myProcs, err := proc.ReadProc("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myProcs)
}

// Liste des processus par user
func PSUserHandler(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	myProcs, err := proc.ReadProc(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myProcs)
}

func NetHandler(w http.ResponseWriter, r *http.Request) {
	out, err := netcard.ReadNetwork("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(out)
}

func NetNameHandler(w http.ResponseWriter, r *http.Request) {
	card := r.PathValue("card")
	out, err := netcard.ReadNetwork(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(out)
}

func MemHandler(w http.ResponseWriter, r *http.Request) {
	out, err := memory.ReadMemory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(out)
}

func DiskHandler(w http.ResponseWriter, r *http.Request) {
	out, err := disk.ReadDisk()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(out)
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	out, err := load.ReadLoad()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(out)
}

func CPUHandler(w http.ResponseWriter, r *http.Request) {
	out, err := cpu.ReadCPU()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(out)
}
