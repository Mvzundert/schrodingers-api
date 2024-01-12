package main

import (
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	scriptPath := filepath.Join("rocks", "pong.rock")
	cmd := exec.Command("./bin/chirp", scriptPath)

	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Error executing script", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
