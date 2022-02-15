package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/stats", StatsWeb)
	http.ListenAndServe(":8080", nil)
}

func StatsWeb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stats, _ := Statistics()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(stats)
}

func WebPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stats, _ := Statistics()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(stats)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Stats</h1>"+"<a href='%s'>Stats</a>", "/stats")
}

const (
	defaultDockerPath        string = "/usr/local/bin/docker" //my docker installation dir
	defaultDockerCommand     string = "stats"
	defaultDockerNoStreamArg string = "--no-stream"
	defaultDockerFormatArg   string = "--format"
	defaultDockerFormat      string = `{"container":"{{.Container}}","memory":{"raw":"{{.MemUsage}}","percent":"{{.MemPerc}}"},"cpu":"{{.CPUPerc}}","io":{"network":"{{.NetIO}}","block":"{{.BlockIO}}"},"pids":{{.PIDs}}}`
)

func Statistics() ([]Stats, error) {
	out, err := exec.Command(defaultDockerPath, defaultDockerCommand, defaultDockerNoStreamArg, defaultDockerFormatArg, defaultDockerFormat).Output()
	if err != nil {
		return nil, err
	}

	containers := strings.Split(string(out), "\n")
	stats := make([]Stats, 0)
	for _, con := range containers {
		if len(con) == 0 {
			continue
		}

		var s Stats
		if err := json.Unmarshal([]byte(con), &s); err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	return stats, nil
}

type Stats struct {
	Container string      `json:"container"`
	Memory    MemoryStats `json:"memory"`
	CPU       string      `json:"cpu"`
	IO        IOStats     `json:"io"`
	PIDs      int         `json:"pids"`
}

type MemoryStats struct {
	Raw     string `json:"raw"`
	Percent string `json:"percent"`
}

type IOStats struct {
	Network string `json:"network"`
	Block   string `json:"block"`
}
