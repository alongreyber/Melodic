package main

import (
    "net/http"
    mux "github.com/gorilla/mux"
    prisma "melodic-backend/prisma-client"
)

type App struct {
    db *prisma.Client
}

func main() {
    db := prisma.New(nil)
    app := App{db: db}

    r := mux.NewRouter()
    r.Use(app.JwtAuthentication)
    r.Use(app.AddHeaders)
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/login", app.SpotifyLoginHandler)

    // Run
    err := http.ListenAndServe(":5000", r)
    if err != nil {
	panic(err)
    }

}

