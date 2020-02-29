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

    spotifyAuth := spotify.NewAuthenticator("http://localhost:8080/", spotify.ScopeUserReadPrivate, spotify.ScopeUserReadEmail, spotify.ScopeUserLibraryRead, spotify.ScopeUserFollowRead)

    app := App{db: db,
	       spotifyAuth: &spotifyAuth}

   //db.LogMode(true)

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Artist{})
    db.AutoMigrate(&SpotifyImage{})

    r := mux.NewRouter()
    r.Use(app.AddHeaders)
    r.Use(app.JwtAuthentication)
    r.Use(app.GetUserMiddleware)
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/login", app.SpotifyLoginHandler)
    api.HandleFunc("/getUserInfo", app.GetThisUserInfo)
    api.HandleFunc("/healthCheck", app.HealthCheckHandler)
    api.HandleFunc("/getCallbackURL", app.CallbackURL).Methods("GET")

    api.HandleFunc("/listenTo", app.GetListenTo).Methods("GET")
    api.HandleFunc("/toReview", app.GetToReview).Methods("GET")

    api.HandleFunc("/moveToReview/{id:[0-9]+}", app.MoveToReview).Methods("GET")
    api.HandleFunc("/moveToListen/{id:[0-9]+}", app.MoveToListenTo).Methods("GET")

    // Log all requests and responses to stdout for debugging
    loggedRouter := handlers.LoggingHandler(os.Stdout, r)
    // Run
    err = http.ListenAndServe(":80", loggedRouter)
    if err != nil {
	panic(err)
    }

}

