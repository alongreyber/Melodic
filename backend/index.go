package main

import (
    "net/http"
    mux "github.com/gorilla/mux"
    prisma "melodic-backend/prisma"
)

type App struct {
    db *prisma.Client
}

func main() {
    db := prisma.New(nil)
    app := App{db: db}

    r := mux.NewRouter()
    r.HandleFunc("/", app.GetUserHandler)
    r.HandleFunc("/createUser", app.CreateUserHandler)

    // Run
    err := http.ListenAndServe(":8080", r)
    if err != nil {
	panic(err)
    }

}
