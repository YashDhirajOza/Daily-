package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"  // New import for simulating delays
)

func main() {
    // Default response for root (handles GET only for simplicity)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]string{"message": "Hello, this is my first Go server"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    })

    // Handler for /about with method checking and goroutine
    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            // Handle GET requests
            data := map[string]string{"message": "Hello from About section (GET)"}
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(data)
            return
        } else if r.Method == http.MethodPost {
            // Handle POST requests: Read JSON from the body
            var input map[string]string
            err := json.NewDecoder(r.Body).Decode(&input)
            if err != nil {
                http.Error(w, "Invalid JSON", http.StatusBadRequest)
                return
            }

            // Start a goroutine for background processing
            go func(data map[string]string) {
                // Simulate a time-consuming task (e.g., logging or DB write)
                fmt.Println("Background processing started for:", data)
                time.Sleep(2 * time.Second)  // Fake delay
                fmt.Println("Background processing complete!")
            }(input)  // Pass the input to the goroutine

            // Immediately respond to the client
            response := map[string]string{"status": "POST received, processing in background!"}
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(response)
            return
        }

        // If method isn't GET or POST, return an error
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    })

    fmt.Println("Server starting on port 8080")
    http.ListenAndServe(":8080", nil)
}
