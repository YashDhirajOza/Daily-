package handlers

import (
    "net/http"
    "time"
    "url-shortener/models"
    "url-shortener/services"
    "url-shortener/database"
    
    "github.com/gin-gonic/gin"
)

type ShortenRequest struct {
    URL        string     `json:"url" binding:"required"`
    CustomCode string     `json:"custom_code"`
    ExpiresAt  *time.Time `json:"expires_at"`
}

func ShortenURL(c *gin.Context) {
    var req ShortenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    userIP := c.ClientIP()
    
    url, err := services.ShortenURL(req.URL, req.CustomCode, userIP, req.ExpiresAt)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "short_url":    "http://localhost:8080/" + url.ShortCode,
        "short_code":   url.ShortCode,
        "original_url": url.OriginalURL,
        "expires_at":   url.ExpiresAt,
    })
}

func RedirectURL(c *gin.Context) {
    shortCode := c.Param("shortCode")
    
    originalURL, err := services.GetOriginalURL(shortCode)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found or expired"})
        return
    }
    
    // Record the click asynchronously
    go services.RecordClick(
        shortCode,
        c.ClientIP(),
        c.GetHeader("User-Agent"),
        c.GetHeader("Referer"),
    )
    
    c.Redirect(http.StatusMovedPermanently, originalURL)
}

func GetAnalytics(c *gin.Context) {
    shortCode := c.Param("shortCode")
    
    var url models.URL
    if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }
    
    var analytics []models.ClickAnalytics
    database.DB.Where("url_id = ?", url.ID).Find(&analytics)
    
    c.JSON(http.StatusOK, gin.H{
        "url":           url,
        "total_clicks":  len(analytics),
        "click_history": analytics,
    })
}
