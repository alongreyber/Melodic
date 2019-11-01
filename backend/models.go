package main

import (
    "github.com/jinzhu/gorm"
    "time"
)

type User struct {
    gorm.Model
    SpotifyID string
    SpotifyTokenAccess string
    SpotifyTokenRefresh string
    SpotifyTokenExpiry time.Time
    SpotifyTokenType string
}
