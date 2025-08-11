package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/qdrant/go-client/qdrant"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
)

const (
    qdrantURL = "https://your-cluster-id.qdrant.io:6334"  // From dashboard
    qdrantAPIKey = "your-api-key"                         // From dashboard
    collectionName = "embeddings"                         // Your collection
)

func main() {
    // Connect to Qdrant
    conn, err := grpc.Dial(qdrantURL, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(qdrant.NewApiKeyCredentials(qdrantAPIKey)))
    if err != nil {
        fmt.Println("Qdrant connection failed:", err)
        return
    }
    client := qdrant.NewCollectionsClient(conn)
    pointsClient := qdrant.NewPointsClient(conn)

    // Create collection if not exists (run once)
    _, _ = client.Create(context.Background(), &qdrant.CreateCollection{
        CollectionName: collectionName,
        VectorsConfig: &qdrant.VectorsConfig{Config: &qdrant.VectorsConfig_Params{Params: &qdrant.VectorParams{Size: 384, Distance: qdrant.Distance_Cosine}}},
    })

    http.HandleFunc("/embed", func(w http.ResponseWriter, r *http.Request) {
        handleEmbed(w, r, pointsClient)
    })

    fmt.Println("Server starting on port 8080")
    http.ListenAndServe(":8080", nil)
}

func handleEmbed(w http.ResponseWriter, r *http.Request, client qdrant.PointsClient) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input struct { Text string }
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Generate embedding (replace with your local free method, e.g., Ollama)
    embedding := generateEmbedding(input.Text)

    // Upsert to Qdrant
    _, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
        CollectionName: collectionName,
        Points: []*qdrant.PointStruct{{
            Id:     &qdrant.PointId{Uuid: "unique-id-for-this-point"},  // Generate a unique ID
            Vectors: &qdrant.Vectors{VectorsOptions: &qdrant.Vectors_Vector{Vector: &qdrant.Vector{Data: embedding}}},
            Payload: map[string]*qdrant.Value{"content": {Kind: &qdrant.Value_StringValue{StringValue: input.Text}}},
        }},
    })
    if err != nil {
        http.Error(w, "Upsert failed", http.StatusInternalServerError)
        return
    }

    // Query similar points
    resp, _ := client.Search(context.Background(), &qdrant.SearchPoints{
        CollectionName: collectionName,
        Vector:         embedding,
        Limit:          3,  // Top 3 similar
    })

    var results []string
    for _, point := range resp.Result {
        results = append(results, point.Payload["content"].GetStringValue())
    }

    // Respond
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string][]string{"similar": results})
}

// Placeholder for embedding (use Ollama or similar free tool)
func generateEmbedding(text string) []float32 {
    return []float32{0.1, 0.2 /* ... 384 values */}
}
