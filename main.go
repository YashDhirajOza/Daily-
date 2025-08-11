package main

import (
    "fmt"
    "github.com/kelindar/search"  // Import the library
)

func main() {
    // Load the free local model (path to your downloaded GGUF file)
    m, err := search.NewVectorizer("models/MiniLM-L6-v2.Q8_0.gguf", 0)  // 0 = CPU only
    if err != nil {
        fmt.Println("Error loading model:", err)
        return
    }
    defer m.Close()  // Clean up

    // Generate embedding for a text input
    text := "Hello, this is a test sentence about Go programming."
    embedding, err := m.EmbedText(text)
    if err != nil {
        fmt.Println("Error embedding text:", err)
        return
    }

    // Print the vector (a slice of floats)
    fmt.Println("Embedding vector:", embedding)
}
