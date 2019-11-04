package main

import (
    "net/http"
    mux "github.com/gorilla/mux"
    handlers "github.com/gorilla/handlers"
    "os"
    spotify "github.com/zmb3/spotify"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
    db *gorm.DB
    spotifyAuth *spotify.Authenticator
}

func main() {
    db, err := gorm.Open("postgres",
		"host=postgres port=5432 user=postgres dbname=gorm password=password sslmode=disable")
    if err != nil {
	panic(err)
    }
    // Test the db with a ping
    err = db.DB().Ping()
    if err != nil {
	panic(err)
    }

    spotifyAuth := spotify.NewAuthenticator("http://localhost:8080/spotify_callback", spotify.ScopeUserReadPrivate, spotify.ScopeUserReadEmail, spotify.ScopeUserLibraryRead, spotify.ScopeUserFollowRead)

    app := App{db: db,
	       spotifyAuth: &spotifyAuth}

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Artist{})

    r := mux.NewRouter()
    r.Use(app.AddHeaders)
    r.Use(app.JwtAuthentication)
    r.Use(app.GetUserMiddleware)
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/login", app.SpotifyLoginHandler)
    api.HandleFunc("/getUserInfo", app.GetThisUserInfo)
    api.HandleFunc("/healthCheck", app.HealthCheckHandler)
    api.HandleFunc("/getListenTo", app.GetListenTo)
    api.HandleFunc("/getCallbackURL", app.CallbackURL)

    // Log all requests and responses to stdout for debugging
    loggedRouter := handlers.LoggingHandler(os.Stdout, r)
    // Run
    err = http.ListenAndServe(":5000", loggedRouter)
    if err != nil {
	panic(err)
    }

}

