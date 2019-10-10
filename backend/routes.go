package main

import (
    "encoding/json"
    "net/http"
    prisma "melodic-backend/prisma-client"
    jwt "github.com/dgrijalva/jwt-go"
    "time"
    "context"
    spotify "github.com/zmb3/spotify"
)

type SpotifyAuthCode struct {
    Code string `json:code`
}

func (app App) SpotifyLoginHandler(w http.ResponseWriter, r *http.Request) {
    auth := spotify.NewAuthenticator("localhost:8080/login_redirect", spotify.ScopeUserReadPrivate, spotify.ScopeUserReadEmail)
    // For spotify login we are given a code and need to exchange it for a token
    // use the same state string here that you used to generate the URL
    spotify_token, err := auth.Token("123456789", r)
    if err != nil {
	errorResponse(w, "Could not get token", http.StatusBadRequest)
	return
    }
    // create a client using the specified token
    client := auth.NewClient(spotify_token)

    // See if a user exists
    // If not, create one
    spotify_user, err := client.CurrentUser()
    if err != nil {
	errorResponse(w, "Could not get spotify user info", http.StatusBadRequest)
	return
    }
    user, err := app.db.User(prisma.UserWhereUniqueInput{
	SpotifyId: &spotify_user.ID,
    }).Exec( r.Context() )
    if err != nil {
	// Create a new user
	user, err = app.db.CreateUser(prisma.UserCreateInput{
	    SpotifyId: spotify_user.ID,
	    // Automatically serializes to json
	    AccessToken: map[string]interface{}{"token": spotify_token},
	}).Exec( r.Context() )
	if err != nil {
	    errorResponse(w, "Could not create new user", http.StatusInternalServerError)
	    return
	}
    }
    // Create JWT for app login
    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
	Id: user.ID,
	StandardClaims: jwt.StandardClaims{
	    // In jwt, the expiry time is expressed in unix milliseconds
	    ExpiresAt: expirationTime.Unix(),
	},
    }

    // Delcare the token 
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)

    if err != nil {
	errorResponse(w, "Error creating JWT", http.StatusInternalServerError)
	return
    }
    // Return the token to the user
    http.SetCookie(w, &http.Cookie{
	Name: "token",
	Value: tokenString,
	Expires: expirationTime,
    })
    // Return success
    okResponse(w, "Logged In")
    return

}

func (app App) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    users, _ := app.db.Users(nil).Exec( r.Context() )

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(users)
}

func errorResponse(w http.ResponseWriter, err string, status int) {
    // Create this for a json response
    response := make(map[string]interface{})
    response["ok"] = false
    response["error"] = err
    w.WriteHeader(status)
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func okResponse(w http.ResponseWriter, data interface{}) {
    response := make(map[string]interface{})
    response["ok"] = true
    response["data"] = data
    w.WriteHeader(http.StatusOK)
    w.Header().Add("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
    Id string `json:"id"`
    jwt.StandardClaims
}

func (app App) RefreshHandler(w http.ResponseWriter, r *http.Request) {
    claimsContext := r.Context().Value("claims")
    claims := claimsContext.(*Claims)
    if time.Unix(claims.ExpiresAt, 0).Sub(time.Now())  > 30*time.Second {
	errorResponse(w, "Token not expired", http.StatusBadRequest)
	return
    } 

    expirationTime := time.Now().Add(5 * time.Minute)
    claims.ExpiresAt = expirationTime.Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
	errorResponse(w, "Error creating new token", http.StatusInternalServerError)
	return
    }
    http.SetCookie(w, &http.Cookie{
	Name: "token",
	Value: tokenString,
	Expires: expirationTime,
    })
}

func (app App) AddHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    })
}

func (app App) JwtAuthentication(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Check if this route needs auth
	notAuth := []string{"/login"}
	requestPath := r.URL.Path
	for _, value := range notAuth {
	    if value == requestPath {
		next.ServeHTTP(w, r)
		return
	    }
	}

	// Check if token is in cookie
	c, err := r.Cookie("token")
	if err != nil {
	    if err == http.ErrNoCookie {
		errorResponse(w, "No token found", http.StatusUnauthorized)
		return
	    }
	    errorResponse(w, "Bad request", http.StatusBadRequest)
	    return
	}

	// Parse token and validate
	tokenString := c.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
	    if err == jwt.ErrSignatureInvalid {
		errorResponse(w, "Invailid JWT Signature", http.StatusUnauthorized)
		return
	    }
	    errorResponse(w, "Bad request", http.StatusBadRequest)
	    return
	}
	if !token.Valid {
	    errorResponse(w, "Invalid Token", http.StatusUnauthorized)
	    return
	}

	// Add claims to context
	ctx := context.WithValue(r.Context(), "claims", claims)
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)
    })
}
