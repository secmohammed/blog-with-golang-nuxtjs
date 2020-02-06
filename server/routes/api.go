package routes

import (
    "fmt"
    "net/http"
    "os"

    "go-auth-with-crud-api/server/app/http/controllers/posts"
    "go-auth-with-crud-api/server/app/http/controllers/users"
    "go-auth-with-crud-api/server/app/http/middlewares"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

// RegisterAPIRoutes is used to register the routes we need for the web application.
func RegisterAPIRoutes() {
    router := mux.NewRouter()
    postPublicArea := router.PathPrefix("/api/posts").Subrouter()
    postPublicArea.HandleFunc("", posts.Index).Methods("GET")
    postPublicArea.HandleFunc("/{post}", posts.Show).Methods("GET")
    postAuthArea := router.PathPrefix("/api/posts").Subrouter()
    postAuthArea.Use(middlewares.Authenticate)
    postAuthArea.HandleFunc("", posts.Store).Methods("POST")
    postAuthArea.HandleFunc("/{post}", posts.Destroy).Methods("DELETE")
    // auth routes.
    userGuestArea := router.PathPrefix("/api/auth").Subrouter()
    userGuestArea.Use(middlewares.Guest)
    userGuestArea.HandleFunc("/register", users.ParseRegisterForm).Methods("POST")
    userGuestArea.HandleFunc("/login", users.ParseLoginForm).Methods("POST")
    userGuestArea.HandleFunc("/activate/{token}", users.ActivateRegisteredAccount).Methods("GET")
    userGuestArea.HandleFunc("/forget-password", users.ParseForgetPasswordForm).Methods("POST")
    userGuestArea.HandleFunc("/reset-password/{token}", users.ParseResetPassword).Methods("POST")
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
