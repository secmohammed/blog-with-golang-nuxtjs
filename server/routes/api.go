package routes

import (
    "fmt"
    "net/http"
    "os"

    "go-auth-with-crud-api/server/app/http/controllers/users"
    "go-auth-with-crud-api/server/app/http/middlewares"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

// RegisterAPIRoutes is used to register the routes we need for the web application.
func RegisterAPIRoutes() {
    router := mux.NewRouter()
    // auth routes.
    userGuestArea := router.PathPrefix("/api/auth").Subrouter()
    userGuestArea.Use(middlewares.Guest)
    userGuestArea.HandleFunc("/register", users.ParseRegisterForm).Methods("POST")
    userGuestArea.HandleFunc("/login", users.ParseLoginForm).Methods("POST")
    userGuestArea.HandleFunc("/activate/{token}", users.ActivateRegisteredAccount).Methods("GET")
    userAuthArea := router.PathPrefix("/api/auth").Subrouter()
    userAuthArea.Use(middlewares.Authenticate)
    userAuthArea.HandleFunc("/user", users.GetAuthenticatedUser).Methods("GET")
    userAuthArea.HandleFunc("/change-password", users.ParseChangePasswordForm).Methods("POST")
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        AllowCredentials: true,
        // Enable Debugging for testing, consider disabling in production
        Debug: true,
    })

    // Insert the middleware
    handler := c.Handler(router)

    http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), handler)

}
