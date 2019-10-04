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

type ErrorResponse struct {
    ok bool
    err string `json:error`
}

var jwtAuthentication = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	notAuth = []string{"/api/user/new", "/api/user/login"}
	requestPath := r.URL.Path
	for _, value := range notAuth {
	    if value == requestPath {
		next.ServeHTTP(w, r)
		return
	    }
	}
	
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" { // No authentication header
	    w.WriteHeader(http.StatusForbidden)
	    w.Header().Add("Content-Type", "application/json")
	    response := ErrorResponse{ok: false, err: "No authorization token"}
	    json.NewEncoder(w).encode(response)
	    return
	}

	splitToken := strings.Split(tokenHeader, " ")
	if len(splitToken) != 2 {
	    w.WriteHeader(http.StatusForbidden)
	    w.Header().Add("Content-Type", "application/json")
	    response := ErrorResponse{ok: false, err: "Malformed auth token"}
	    json.NewEncoder(w).encode(response)
	    return
	}


    })
}
