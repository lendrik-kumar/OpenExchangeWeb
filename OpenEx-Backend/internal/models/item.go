package models

import (
    "time"
)

type Item struct {
    ID          uint   `gorm:"primaryKey"`
    UserID      uint   `gorm:"not null"`
    User        User   `gorm:"foreignKey:UserID"`
    HostelID    uint   `gorm:"not null"`
    Hostel      Hostel `gorm:"foreignKey:HostelID"`
    Title       string `gorm:"not null"`
    Description string `gorm:"not null"`
    Price       float64
    Image       string
    Status      string `gorm:"default:'pending'"`
    Type        string `gorm:"not null"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}