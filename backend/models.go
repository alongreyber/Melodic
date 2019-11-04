package main

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Artist struct {
    gorm.Model
    SpotifyID string
    Name string
    URI string
    Endpoint string
}

type User struct {
    gorm.Model
    SpotifyID string
    SpotifyTokenAccess string
    SpotifyTokenRefresh string
    SpotifyTokenExpiry time.Time
    SpotifyTokenType string

    ArtistsFollowing []Artist `gorm:"many2many:user_artists_following`
    ArtistsListenTo []Artist `gorm:"many2many:user_artist_listento"`
}

