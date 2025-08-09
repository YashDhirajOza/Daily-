package main

import (
    "log"
    "url-shortener/config"
    "url-shortener/database"
    "url-shortener/handlers"
    
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    
    // Connect to database and Redis
    database.ConnectDatabase()
    config.ConnectRedis()
    
    // Setup Gin router
    r := gin.Default()
    
    // API routes
    api := r.Group("/api/v1")
    {
        api.POST("/shorten", handlers.ShortenURL)
        api.GET("/analytics/:shortCode", handlers.GetAnalytics)
    }
    
    // Redirect route
    r.GET("/:shortCode", handlers.RedirectURL)
    
    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy"})
    })
    
    log.Println("Server starting on :8080")
    r.Run(":8080")
}
