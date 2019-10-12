package main

import (
    "net/http"
    mux "github.com/gorilla/mux"
    handlers "github.com/gorilla/handlers"
    prisma "melodic-backend/prisma-client"
    "os"
    spotify "github.com/zmb3/spotify"
)

type App struct {
    db *prisma.Client
    spotifyAuth *spotify.Authenticator
}

func main() {
    options := prisma.Options{
	Endpoint: "http://prisma:4466",
	Secret: "my-secret-0000",
    }
    db := prisma.New(&options)
    app := App{db: db}

    spotifyAuth := spotify.NewAuthenticator("http://localhost:8080/spotify_callback", spotify.ScopeUserReadPrivate, spotify.ScopeUserReadEmail)
    app.spotifyAuth = &spotifyAuth

    r := mux.NewRouter()
    r.Use(app.AddHeaders)
    r.Use(app.JwtAuthentication)
    r.Use(app.GetUserMiddleware)
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/login", app.SpotifyLoginHandler)
    api.HandleFunc("/getUserInfo", app.GetThisUserInfo)
    api.HandleFunc("/healthCheck", app.HealthCheckHandler)

    // Log all requests and responses to stdout for debugging
    loggedRouter := handlers.LoggingHandler(os.Stdout, r)
    // Run
    err := http.ListenAndServe(":5000", loggedRouter)
    if err != nil {
	panic(err)
    }

}

