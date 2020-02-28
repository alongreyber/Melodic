package main

import (
    "github.com/jinzhu/gorm"
    "time"
)

// Reference to a spotify image
type SpotifyImage struct {
    gorm.Model
    ArtistID uint

    Height int
    Width int
    URL string
}

type Artist struct {
    gorm.Model
    SpotifyID string
    Name string
    URI string
    Endpoint string
    Images []SpotifyImage `gorm:"gorm:association_autoupdate"`

    UsersFollowing []User `gorm:"many2many:user_following_artist"`
    UsersListenTo []User `gorm:"many2many:user_listento_artist"`
    UsersToReview []Artist `gorm:"many2many:user_toreview_artist"`
}

type User struct {
    gorm.Model
    SpotifyID string
    SpotifyTokenAccess string
    SpotifyTokenRefresh string
    SpotifyTokenExpiry time.Time
    SpotifyTokenType string

    ArtistsFollowing []Artist `gorm:"many2many:user_following_artist"`
    ArtistsListenTo []Artist `gorm:"many2many:user_listento_artist"`
    ArtistsToReview []Artist `gorm:"many2many:user_toreview_artist"`
}
