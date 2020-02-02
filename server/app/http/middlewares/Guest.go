package middlewares

import (
    u "go-auth-with-crud-api/server/utils"
    "net/http"
)

// Guest is used to check if user has a valid jwt token or not.
var Guest = func(next http.Handler) http.Handler {

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        response := make(map[string]interface{})
        tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

        if tokenHeader != "" { //Token is missing, returns with error code 403 Unauthorized
            response = u.Message(false, "You cannot visit here as you are already authenticated.")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            u.Respond(w, response)
            return
        }

        next.ServeHTTP(w, r)
    })
}
