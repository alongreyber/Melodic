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
    SpotifyID string `gorm:"unique"`
    Name string
    URI string
    Endpoint string
    Images []SpotifyImage `gorm:"gorm:association_autoupdate"`

    UsersFollowing []User `gorm:"many2many:user_following_artist"`
    UsersRecentlyFollowed []User `gorm:"many2many:user_recentlyfollowed_artist"`
    UsersRecentlyListened []User `gorm:"many2many:user_recentlylistened_artist"`
}

type Tag struct {
    gorm.Model
    Name string
}

type ArtistReview struct {
    gorm.Model
    ArtistID uint
    UserID uint
    OverallScore int
    Notes string

    LyricsScore int
    VocalsScore int
    SonicsScore int
    ProductionScore int
    ImpactScore int

    Tags []Tag
}

type User struct {
    gorm.Model
    SpotifyID string `gorm:"unique"`
    SpotifyTokenAccess string
    SpotifyTokenRefresh string
    SpotifyTokenExpiry time.Time
    SpotifyTokenType string

    ArtistsFollowing []Artist `gorm:"many2many:user_following_artist"`
    ArtistsRecentlyFollowed []Artist `gorm:"many2many:user_recentlyfollowed_artist"`
    ArtistsRecentlyListened []Artist `gorm:"many2many:user_recentlylistened_artist"`
    ArtistReviews []ArtistReview
}
