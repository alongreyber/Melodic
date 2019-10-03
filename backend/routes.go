package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    prisma "melodic-backend/prisma"
)

func (app App) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    users, _ := app.db.Users(nil).Exec( r.Context() )

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(users)
}

func (app App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    // Create a user
    name := "Alice"
    newUser, err := app.db.CreateUser(prisma.UserCreateInput{
	Name: name,
    }).Exec( r.Context() )
    if err != nil {
	panic(err)
    }
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "{status: 'Created user %d'}", newUser.ID)
}
