package main

import (
    "encoding/json"
    "net/http"
    jwt "github.com/dgrijalva/jwt-go"
    "time"
    "context"
    oauth2 "golang.org/x/oauth2"
    "fmt"
    "github.com/jinzhu/gorm"
    uuid "github.com/google/uuid"
    spotify "github.com/zmb3/spotify"
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

func (app App) CallbackURL(w http.ResponseWriter, r *http.Request) {
    state := uuid.New()
    url := app.spotifyAuth.AuthURL(state.String())
    response := make(map[string]string)
    response["url"] = url
    response["state"] = state.String()
    okResponse(w, response)
}

func (app App) LogoutHandler(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
	Name: "token",
	Value: "",
	MaxAge: -1,
    })
    okResponse(w, "Logged Out")
}

func (app App) SpotifyLoginHandler(w http.ResponseWriter, r *http.Request) {
    // For spotify login we are given a code and need to exchange it for a token
    state, ok := r.URL.Query()["state"]
    if !ok || len(state) != 1 {
	errorResponse(w, fmt.Errorf("No state given"), http.StatusUnauthorized)
	return
    }
    // We're not validating state here because it has been verified on the frontend
    spotify_token, err := app.spotifyAuth.Token(state[0], r)
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not get token: %v", err), http.StatusInternalServerError)
	return
    }
    // create a client using the specified token
    client := app.spotifyAuth.NewClient(spotify_token)

    // See if a user exists
    // If not, create one
    spotify_user, err := client.CurrentUser()
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not get spotify user info: %v", err), http.StatusInternalServerError )
	return
    }
    var user User
    err = app.db.Where(&User{SpotifyID: spotify_user.ID}).First(&user).Error
    if gorm.IsRecordNotFoundError(err) {
	user = User{}
	err = app.db.Create(&user).Error
	if err != nil {
	    errorResponse(w, fmt.Errorf("Could not create new user: %v", err), http.StatusInternalServerError )
	    return
	}

    }
    user.SpotifyID           = spotify_user.ID
    user.SpotifyTokenAccess  = spotify_token.AccessToken
    user.SpotifyTokenRefresh = spotify_token.RefreshToken
    user.SpotifyTokenExpiry  = spotify_token.Expiry
    user.SpotifyTokenType    = spotify_token.TokenType

    err = app.db.Save(&user).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not save user: %v", err), http.StatusInternalServerError)
	return
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
	errorResponse(w, fmt.Errorf("Error creating JWT: %v", err) , http.StatusInternalServerError)
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

func GetAllFollowingArtists(spotifyClient spotify.Client) (*[]spotify.FullArtist, error) {
    allArtists := make([]spotify.FullArtist, 1)
    cursor := ""
    for { // Iterate over the pages
	spotifyFollowedArtistsData, err := spotifyClient.CurrentUsersFollowedArtistsOpt(50, cursor)
	if err != nil {
	    return nil, fmt.Errorf("Error from spotify: %v", err)
	}
	allArtists = append(allArtists, spotifyFollowedArtistsData.Artists...)
	// Move to next page
	cursor = spotifyFollowedArtistsData.Cursor.After
	if cursor == "" {
	    break
	}
    }
    return &allArtists, nil

}

func MakeArtist(spotifyArtist spotify.FullArtist) Artist {
    images := make([]SpotifyImage, 1)
    for _, im := range(spotifyArtist.Images) {
	images = append(images, SpotifyImage{
	    Height: im.Height,
	    Width: im.Width,
	    URL: im.URL,
	})
    }
    artist := Artist{
	SpotifyID: spotifyArtist.ID.String(),
	Name: spotifyArtist.Name,
	URI: string(spotifyArtist.URI),
	Endpoint: spotifyArtist.Endpoint,
	Images: images,
    }
    return artist

}

func (app App) InitializeFollowing(w http.ResponseWriter, r *http.Request) {
    user := getUser(r.Context())
    token, ok := getToken(user)
    if !ok {
	panic("Couldn't make token")
    }
    spotifyClient := app.spotifyAuth.NewClient(token)

    // TODO use associations
    // Load following and listento artists
    err := app.db.Set("gorm:auto_preload", true).Preload("ArtistsRecentlyFollowed.Images").First(&user, user.ID).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Couldn't load user's artists from db: %v", err), http.StatusInternalServerError)
	return
    }

    if len(user.ArtistsFollowing) != 0 {
	errorResponse(w, fmt.Errorf("User already initialized"), http.StatusBadRequest)
	return
    }

    // Get all artists in DB
    var allArtists []Artist
    err = app.db.Find(&allArtists).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not get artists from db: %v", err), http.StatusInternalServerError)
	return
    }

    // Get current artists from spotify
    SpotifyFollowingArtists, err := GetAllFollowingArtists(spotifyClient)
    if err != nil {
	errorResponse(w, err, http.StatusInternalServerError)
	return
    }

    // spotifyFollowedArtistsData is a paged data structure
    for _, spotifyArtist := range(*SpotifyFollowingArtists) {
	// Get artist from DB. If it doesn't exist, create it
	var artist *Artist
	for _, a := range(allArtists) {
	    if(spotifyArtist.ID.String() == a.SpotifyID) {
		artist = &a
		break
	    }
	}
	if artist == nil {
	    // Going to need to create it
	    artistVal := MakeArtist(spotifyArtist)
	    artist = &artistVal
	}
	// Make sure artist is not already in list
	// Add artist to user follows list
	user.ArtistsFollowing = append(user.ArtistsFollowing, *artist)
    }
    // Save associations
    err = app.db.Save(&user).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not save user: %v", err), http.StatusInternalServerError)
	return
    }

    okResponse(w, "Done")

}

func (app App) SearchArtists(w http.ResponseWriter, r *http.Request) {
    user := getUser(r.Context())
    token, ok := getToken(user)
    if !ok {
	panic("Couldn't make token")
    }
    spotifyClient := app.spotifyAuth.NewClient(token)

    queryArray, ok := r.URL.Query()["q"]
    if !ok || len(queryArray) != 1 || queryArray[0] == "" {
	// Just return no results
	okResponse(w, []string{})
    }
    query := queryArray[0]

    results, err := spotifyClient.Search(query, spotify.SearchTypeArtist)
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not get search result: %v", err), http.StatusInternalServerError)
	return
    }

    okResponse(w, results.Artists.Artists[:5])
}

func (app App) RefreshRecentlyFollowed(w http.ResponseWriter, r *http.Request) {

    user := getUser(r.Context())
    token, ok := getToken(user)
    if !ok {
	panic("Couldn't make token")
    }
    spotifyClient := app.spotifyAuth.NewClient(token)

    // Get all artists in DB
    var allArtists []Artist
    err := app.db.Find(&allArtists).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not get artists from db: %v", err), http.StatusInternalServerError)
	return
    }

    // Get current artists from spotify
    SpotifyFollowingArtists, err := GetAllFollowingArtists(spotifyClient)
    if err != nil {
	errorResponse(w, err, http.StatusInternalServerError)
	return
    }

    // Load list of currently followed artists
    err = app.db.Model(&user).Association("ArtistsFollowing").Find(&user.ArtistsFollowing).Error
    if err != nil {
	errorResponse(w, err, http.StatusInternalServerError)
	return
    }

    // Now what we do is check for any new artists
    // If we've followed new artists add them to the "listen to" list

    // Loop over artists and see if any have been added
    for _, spotifyArtist := range(*SpotifyFollowingArtists) {
	alreadyFollowing := false
	for _, a := range(user.ArtistsFollowing) {
	    if a.SpotifyID == spotifyArtist.ID.String() {
		alreadyFollowing = true
		break
	    }
	}
	if alreadyFollowing {
	    continue
	}

	// Add the artist to the DB if it doesn't exist
	var artist *Artist
	for _, a := range(allArtists) {
	    if spotifyArtist.ID.String() == "" {
		fmt.Printf("Blank Artist: %s\n", spotifyArtist.Name)
	    }
	    if spotifyArtist.ID.String() == a.SpotifyID  {
		artist = &a
		break
	    }
	}
	if artist == nil {
	    // Going to need to create it
	    artistVal := MakeArtist(spotifyArtist)
	    artist = &artistVal
	    err = app.db.Set("gorm:assocation_autocreate", true).Create(artist).Error
	    if err != nil {
		errorResponse(w, fmt.Errorf("Could not create artist: %v", err), http.StatusInternalServerError)
		return
	    }
	}
	// Add to 
	err = app.db.Model(&user).Association("ArtistsFollowing").Append(*artist).Error
	if err != nil {
	    errorResponse(w, fmt.Errorf("Could not save user: %v", err), http.StatusInternalServerError)
	    return
	}
	err = app.db.Model(&user).Association("ArtistsRecentlyFollowed").Append(*artist).Error
	if err != nil {
	    errorResponse(w, fmt.Errorf("Could not save user: %v", err), http.StatusInternalServerError)
	    return
	}
    }
    okResponse(w, "Done")
}

func (app App) GetRecentlyFollowed(w http.ResponseWriter, r *http.Request) {
    user := getUser(r.Context())
    // Load following and listento artists
    err := app.db.Set("gorm:auto_preload", true).Preload("ArtistsRecentlyFollowed.Images").First(&user, user.ID).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Couldn't load user's artists from db: %v", err), http.StatusInternalServerError)
	return
    }
    // Send user the ArtistsRecentlyFollowed
    okResponse(w, user.ArtistsRecentlyFollowed)
}

func (app App) RefreshRecentlyListened(w http.ResponseWriter, r *http.Request) {
    user := getUser(r.Context())
    token, ok := getToken(user)
    if !ok {
	panic("Couldn't make token")
    }

    spotifyClient := app.spotifyAuth.NewClient(token)

    // Get all artists in DB
    var allArtists []Artist
    err := app.db.Find(&allArtists).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not get artists from db: %v", err), http.StatusInternalServerError)
	return
    }

    // Clear existing list of RecentlyListened
    err = app.db.Model(&user).Association("ArtistsRecentlyListened").Clear().Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not clear associations: %v", err), http.StatusInternalServerError)
	return
    }

    recentlyPlayedSpotifyList, err := spotifyClient.PlayerRecentlyPlayed()
    if err != nil {
	errorResponse(w, fmt.Errorf("Failed to load recently played: %v", err), http.StatusInternalServerError)
	return
    }

    for _, recentlyPlayedSpotify := range(recentlyPlayedSpotifyList) {
	// Get the first artist associated with this track
	spotifyArtist := recentlyPlayedSpotify.Track.Artists[0]
	var artist *Artist
	for _, a := range(allArtists) {
	    if(spotifyArtist.ID.String() == a.SpotifyID) {
		artist = &a
		break
	    }
	}
	if artist == nil {
	    fmt.Printf("Artist not in DB: %s %s", 
		       spotifyArtist.Name, spotifyArtist.ID.String())
	    // Look up the full artist
	    fullArtist, err := spotifyClient.GetArtist(spotifyArtist.ID)
	    if err != nil {
		errorResponse(w, fmt.Errorf("Failed to get artist: %v", err), http.StatusInternalServerError)
		return
	    }

	    // Make into object to insert into DB
	    artistVal := MakeArtist(*fullArtist)
	    artist = &artistVal
	}
	
	// Check if this is already in list
	exists := false
	for _, a := range(user.ArtistsFollowing) {
	    if(a.SpotifyID == artist.SpotifyID) {
		exists = true
	    }
	}
	if !exists {
	    user.ArtistsRecentlyListened = append(user.ArtistsRecentlyListened, *artist)
	}
    }

    err = app.db.Save(&user).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Could not save user: %v", err), http.StatusInternalServerError)
	return
    }
    okResponse(w, "Done")
}

func (app App) GetRecentlyListened(w http.ResponseWriter, r *http.Request) {
    user := getUser(r.Context())

    // Load following and listento artists
    err := app.db.Set("gorm:auto_preload", true).Preload("ArtistsRecentlyListened.Images").First(&user, user.ID).Error
    if err != nil {
	errorResponse(w, fmt.Errorf("Couldn't load user's artists from db: %v", err), http.StatusInternalServerError)
	return
    }
    // Send user the ArtistsRecentlyFollowed
    okResponse(w, user.ArtistsRecentlyListened)
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
	errorResponse(w, fmt.Errorf("Could not get current user info: %v", err), http.StatusInternalServerError) 
	return
    }
    okResponse(w, spotifyUser)
    
}

func okResponse(w http.ResponseWriter, data interface{}) {
    response := make(map[string]interface{})
    response["ok"] = true
    response["data"] = data
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func errorResponse(w http.ResponseWriter, err error, status int) {
    response := make(map[string]interface{})
    response["ok"] = false
    response["error"] = err.Error()
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)
}

var jwtKey = []byte("my_secret_key")

func (app App) RefreshHandler(w http.ResponseWriter, r *http.Request) {
    claimsContext := r.Context().Value("claims")
    claims := claimsContext.(*Claims)
    if time.Unix(claims.ExpiresAt, 0).Sub(time.Now())  > 30*time.Second {
	errorResponse(w, fmt.Errorf("Token not expired"), http.StatusBadRequest)
	return
    } 

    expirationTime := time.Now().Add(1 * time.Hour)
    claims.ExpiresAt = expirationTime.Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
	errorResponse(w, fmt.Errorf("Error creating new token: %v", err), http.StatusInternalServerError)
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
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
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
	    errorResponse(w, fmt.Errorf("Couldn't get user: %v", app.db.Error), http.StatusInternalServerError)
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
	notAuth := []string{"/api/login", "/api/logout", "/api/healthCheck", "/api/getCallbackURL"}
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
		errorResponse(w, fmt.Errorf("No token found"), http.StatusUnauthorized)
		return
	    } else {
		errorResponse(w, fmt.Errorf("Bad request: %v", err), http.StatusInternalServerError )
		return
	    }
	}

	// Parse token and validate
	tokenString := c.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
	    if err == jwt.ErrSignatureInvalid {
		errorResponse(w, fmt.Errorf("Invalid JWT Signature"), http.StatusUnauthorized)
		return
	    } else {
		errorResponse(w, fmt.Errorf("Bad request: %v", err), http.StatusInternalServerError )
		return
	    }
	}
	if !token.Valid {
	    errorResponse(w, fmt.Errorf("Invalid Token"), http.StatusUnauthorized)
	    return
	}

	// Add claims to context
	ctx := addClaims(r.Context(), *claims)
	next.ServeHTTP(w, r.WithContext(ctx))
    })
}
