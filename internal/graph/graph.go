package graph

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

//go:embed graph/*
var staticFiles embed.FS

// startServer starts the HTTP server
func startServer(g Graph) {
	// Serve static files from the embedded file system
	http.Handle("/", http.FileServer(http.FS(staticFiles)))

	// API endpoint for graph data
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(g)
	})

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// openBrowser opens the default browser to the specified URL
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		log.Println("Unsupported platform. Please open the URL manually:", url)
		return
	}

	if err != nil {
		log.Printf("Failed to open browser: %v\n", err)
	}
}

// ShowGraph starts the server and opens the frontend in the browser
func ShowGraph(g Graph) {
	go startServer(g)                          // Start server in a goroutine (non-blocking)
	openBrowser("http://localhost:8080/graph") // Open the frontend in the browser
}
