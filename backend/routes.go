package main

import (
    "encoding/json"
    "net/http"
    jwt "github.com/dgrijalva/jwt-go"
    "time"
    "context"
    oauth2 "golang.org/x/oauth2"
    "fmt"
)

// Contains any claims stored in the JWT
// that is associated with the current session
type Claims struct {
    ID uint
    jwt.StandardClaims
}

// Add claims to the context variable. Used in the 
// JwtAuthentication method
func addClaims(ctx context.Context, c Claims) (context.Context) {
    return context.WithValue(ctx, "claims", c)
}

// Pull claims out of the current context
// If they don't exist, return false
func getClaims(ctx context.Context) (*Claims, bool) {
    claimsUntyped := ctx.Value("claims")
    if claimsUntyped == nil {
	return nil, false
    }
    claims, ok := claimsUntyped.(Claims)
    return &claims, ok
}

// Add the user object to context
func addUser(ctx context.Context, u User) (context.Context) {
    return context.WithValue(ctx, "user", u)
}

// Pull the user object out of the context variable
// if it exists
func getUser(ctx context.Context) User {
    userUntyped := ctx.Value("user")
    if userUntyped == nil {
	panic("No user in context")
    }
    user, ok := userUntyped.(User)
    if !ok {
	panic("Failed to coerce user type")
    }
    return user
}


func (app App) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    okResponse(w, "Check success!")
}

func (app App) SpotifyLoginHandler(w http.ResponseWriter, r *http.Request) {
    // For spotify login we are given a code and need to exchange it for a token
    state, ok := r.URL.Query()["state"]
    if !ok || len(state) != 1 {
	errorResponse(w, "No state given", http.StatusBadRequest)
	return
    }
    // We're not validating state here because it has been verified on the frontend
    spotify_token, err := app.spotifyAuth.Token(state[0], r)
    if err != nil {
	errorResponse(w, "Could not get token", http.StatusBadRequest)
	return
    }
    // create a client using the specified token
    client := app.spotifyAuth.NewClient(spotify_token)

    // See if a user exists
    // If not, create one
    spotify_user, err := client.CurrentUser()
    if err != nil {
	errorResponse(w, "Could not get spotify user info", http.StatusBadRequest)
	return
    }
    var user User
    app.db.Where(&User{SpotifyID: spotify_user.ID}).First(&user)
    if app.db.Error != nil {
	user = User{
	    SpotifyID: spotify_user.ID,
	    SpotifyTokenAccess: spotify_token.AccessToken,
	    SpotifyTokenRefresh: spotify_token.RefreshToken,
	    SpotifyTokenExpiry: spotify_token.Expiry,
	    SpotifyTokenType: spotify_token.TokenType,
	}
	app.db.Create(user)
	if app.db.Error != nil {
	    fmt.Println(err)
	    errorResponse(w, "Could not create new user", http.StatusInternalServerError)
	    return
	}
    }
    // Create JWT for app login
    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &Claims{
	ID: user.ID,
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

}

func (app App) GetListenTo(w http.ResponseWriter, r *http.Request) {
    user := getUser(r.Context())
    token, ok := getToken(user)
    if !ok {
	panic("Couldn't make token")
    }
    spotifyClient := app.spotifyAuth.NewClient(token)

    // Get followed artists and compare to existing artistsFollowing
    _, err := spotifyClient.CurrentUsersFollowedArtists()
    if err != nil {
	panic(err)
    }

    // If we haven't initialized this list, do so now
    /*
    userArtistsFollowing, err := app.db.Artists(
	&prisma.ArtistsParams{
	    Where: &prisma.UserWhere

	}).Exec( r.Context() )
    if err != nil {
	panic(err)
    }
    if len(userArtistsFollowing) == 0 {
	for apiArtist := range(apiFollowedAritsts) {
	    // Create artist
	    newArtists, err := app.db.CreateArtist(
		prisma.ArtistCreateInput{
		    SpotifyID: apiArtist.ID,
		    Name: apiArtist.Name,
		    Uri: apiAritst.URI,
		    Endpoint: apiArtist.Endpoint,
		},
	    ).Exec( r.Context() )
	    // TODO fix
	    if err != nil {
		panic(err)
	    }
	    // Add connection to user
	    _, err := app.db.UpdateUser(
		prisma.UserUpdateInput{
		    ArtistsFollowing: prisma.ArtistUpdateManyInput{
			Connect: newArtist.ID,
		    },
		},
	    ).Exec( r.Context() )
	    if err != nil {
		panic(err)
	    }
	}
    }
    */

}

// TODO This should return error not bool
func getToken(user User) (*oauth2.Token, bool) {
    token := oauth2.Token{
	AccessToken:  user.SpotifyTokenAccess,
	RefreshToken: user.SpotifyTokenRefresh,
	Expiry:       user.SpotifyTokenExpiry,
	TokenType:    user.SpotifyTokenType,
    }
    return &token, true
}

func (app App) GetThisUserInfo(w http.ResponseWriter, r *http.Request) {
    user := getUser(r.Context())
    token, ok := getToken(user)
    if !ok {
	panic("Couldn't make token")
    }
    client := app.spotifyAuth.NewClient(token)
    spotifyUser, err := client.CurrentUser()
    if err != nil {
	errorResponse(w, "Could not get current user info", http.StatusInternalServerError)
	return
    }


    okResponse(w, spotifyUser)
    
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
    json.NewEncoder(w).Encode(response)
}

var jwtKey = []byte("my_secret_key")

func (app App) RefreshHandler(w http.ResponseWriter, r *http.Request) {
    claimsContext := r.Context().Value("claims")
    claims := claimsContext.(*Claims)
    if time.Unix(claims.ExpiresAt, 0).Sub(time.Now())  > 30*time.Second {
	errorResponse(w, "Token not expired", http.StatusBadRequest)
	return
    } 

    expirationTime := time.Now().Add(1 * time.Hour)
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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept")
	next.ServeHTTP(w, r)
    })
}

func (app App) GetUserMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	claims, ok := getClaims(r.Context())
	// If we don't have claims just call the next handler
	if !ok {
	    next.ServeHTTP(w, r)
	    return
	}
	
	var user User
	userQuery := &User{}
	userQuery.ID = claims.ID
	app.db.Where(userQuery).First(&user)
	if app.db.Error != nil {
	    errorResponse(w, "Couldn't get user", http.StatusInternalServerError)
	    return
	}
	// Add user to context
	ctx := addUser(r.Context(), user)
	next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func (app App) JwtAuthentication(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Check if this route needs auth
	notAuth := []string{"/api/login", "/api/healthCheck"}
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
	ctx := addClaims(r.Context(), *claims)
	next.ServeHTTP(w, r.WithContext(ctx))
    })
}
