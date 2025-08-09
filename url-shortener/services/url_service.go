package services

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
    "strings"
    "time"
    "url-shortener/database"
    "url-shortener/models"
    "url-shortener/config"
    
    "gorm.io/gorm"
)

func GenerateShortCode() string {
    bytes := make([]byte, 6)
    rand.Read(bytes)
    return strings.TrimRight(base64.URLEncoding.EncodeToString(bytes), "=")
}

func ShortenURL(originalURL, customCode, userIP string, expiresAt *time.Time) (*models.URL, error) {
    var shortCode string
    
    if customCode != "" {
        // Check if custom code already exists
        var existingURL models.URL
        if err := database.DB.Where("custom_code = ?", customCode).First(&existingURL).Error; err == nil {
            return nil, errors.New("custom code already exists")
        }
        shortCode = customCode
    } else {
        // Generate unique short code
        for {
            shortCode = GenerateShortCode()
            var existingURL models.URL
            if err := database.DB.Where("short_code = ?", shortCode).First(&existingURL).Error; errors.Is(err, gorm.ErrRecordNotFound) {
                break
            }
        }
    }
    
    url := &models.URL{
        OriginalURL: originalURL,
        ShortCode:   shortCode,
        CustomCode:  customCode,
        UserIP:      userIP,
        ExpiresAt:   expiresAt,
    }
    
    if err := database.DB.Create(url).Error; err != nil {
        return nil, err
    }
    
    // Cache in Redis for faster retrieval
    config.RedisClient.Set(config.Ctx, shortCode, originalURL, time.Hour*24)
    
    return url, nil
}

func GetOriginalURL(shortCode string) (string, error) {
    // Try Redis first
    originalURL, err := config.RedisClient.Get(config.Ctx, shortCode).Result()
    if err == nil {
        return originalURL, nil
    }
    
    // Fallback to database
    var url models.URL
    if err := database.DB.Where("short_code = ? AND is_active = ?", shortCode, true).First(&url).Error; err != nil {
        return "", errors.New("URL not found")
    }
    
    // Check expiration
    if url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt) {
        return "", errors.New("URL has expired")
    }
    
    // Cache in Redis for next time
    config.RedisClient.Set(config.Ctx, shortCode, url.OriginalURL, time.Hour*24)
    
    return url.OriginalURL, nil
}

func RecordClick(shortCode, userIP, userAgent, referer string) error {
    var url models.URL
    if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
        return err
    }
    
    // Increment click count
    database.DB.Model(&url).Update("clicks", gorm.Expr("clicks + ?", 1))
    
    // Record analytics
    analytics := &models.ClickAnalytics{
        URLID:     url.ID,
        UserIP:    userIP,
        UserAgent: userAgent,
        Referer:   referer,
        ClickedAt: time.Now(),
    }
    
    return database.DB.Create(analytics).Error
}
