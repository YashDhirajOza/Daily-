package models

import (
    "time"
)

type URL struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    OriginalURL string    `json:"original_url" gorm:"not null"`
    ShortCode   string    `json:"short_code" gorm:"unique;not null;index"`
    CustomCode  string    `json:"custom_code" gorm:"unique"`
    Clicks      int64     `json:"clicks" gorm:"default:0"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    ExpiresAt   *time.Time `json:"expires_at"`
    IsActive    bool      `json:"is_active" gorm:"default:true"`
    UserIP      string    `json:"user_ip"`
}

type ClickAnalytics struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    URLID     uint      `json:"url_id"`
    UserIP    string    `json:"user_ip"`
    UserAgent string    `json:"user_agent"`
    Referer   string    `json:"referer"`
    Country   string    `json:"country"`
    ClickedAt time.Time `json:"clicked_at"`
}
