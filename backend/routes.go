package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    prisma "melodic-backend/prisma"
    jwt "github.com/dgrijalva/jwt-go"
    "strings"
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

func respond(w http.ResponseWriter, data map[string] interface{}) {
}

type responseStruct struct {
    ok: bool
    data: map[string] interface{}
}

func errorRespond(w http.ResponseWriter, err string, status int) {
    response := responseStruct{ok: false, data: err}
    w.WriteHeader(status)
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder.Encode(response)
}

func okRespond(w http.ResponseWriter, data map[string] interface{}, status int = http.StatusOK) {
    response := responseStruct{ok: true, data: err}
    w.WriteHeader(status)
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder.Encode(response)
}

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
    "user1": "password1",
    "user2": "password2"
}

type Credentials struct {
    Password string `json:"password"`
    Username string `json:"username"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func (app App) SignInHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials

    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
	errorResponse("Bad request", http.StatusBadRequest)
	return
    }

    expectedPassword, ok := users[creds.Username]

    if !ok || expectedPassword != creds.Password {
	errorResponse("Not Authorized", http.StatusUnauthorized)
	return
    }

    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
	Username: creds.Username,
	StandardClaims: jwt.StandardClaims{
	    // In jwt, the expiry time is expressed in unix milliseconds
	    ExpiresAt: expirationTime.Unix()
	}
    }

    // Delcare the token 
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)

    if err != nil {
	errorResponse("Error creating JWT", http.StatusInternalServiceError)
	return
    }
    http.SetCookie(w, &http.Cookie{
	Name: "token",
	Value: tokenString,
	Expires: expirationTime
    })
}

var JwtAuthentication = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
	    if err == http.ErrNoCookie {
		errorResponse("No token found", http.StatusUnauthorized)
	    }
	    errorResponse("Bad request", http.StatusBadRequest)
	}

	tokenString := c.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
	    if err == jwt.ErrSignatureInvalid {
		errorResponse("Invailid JWT Signature", http.StatusUnauthorized)
		return
	    }
	    errorResponse("Bad request", http.StatusBadRequest)
	    return
	}
	if !token.Valid {
	    errorResponse("Invalid Token", http.StatusUnauthorized)
	    return
	}

	next.ServeHTTP(w, r)
    })
}
