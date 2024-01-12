package main

import (
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a GET request
	if r.Method != "GET" {
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set up the Rockstar script execution
	scriptPath := filepath.Join("rocks", "ping.rock")
	cmd := exec.Command("rockstar", scriptPath)

	// Execute the script and capture the output
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Error executing script", http.StatusInternalServerError)
		return
	}

	// Write the output to the response
	w.Write(output)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
