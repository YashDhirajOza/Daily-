package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    // Default response for root (handles GET only for simplicity)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]string{"message": "Hello, this is my first Go server"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    })

    // Handler for /about with method checking
    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) { 
        // Check the HTTP method
		// this is for like if we want to send the out put 
        if r.Method == http.MethodGet {
            // Handle GET requests
            data := map[string]string{"message": "Hello from About section (GET)"}
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(data)
            return
			// this is like what if we got an input 
        } else if r.Method == http.MethodPost {
            // Handle POST requests: Read JSON from the body
            var input map[string]string
            err := json.NewDecoder(r.Body).Decode(&input)
            if err != nil {
                http.Error(w, "Invalid JSON", http.StatusBadRequest)
                return
            }

            // Add a confirmation to the input data
            input["status"] = "Received your POST data!"

            // Send back as JSON
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(input)
            return
        }

        // If method isn't GET or POST, return an error
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    })

    fmt.Println("Server starting on port 8080")
    http.ListenAndServe(":8080", nil)
}
