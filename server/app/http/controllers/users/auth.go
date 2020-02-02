package users

import (
    "encoding/json"
    "net/http"

    "go-auth-with-crud-api/server/app/http/requests"
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
)

//ParseLoginForm to parse the login form when submitted.
func ParseLoginForm(w http.ResponseWriter, r *http.Request) {
    user := &models.User{}
    err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
    if err != nil {
        utils.Respond(w, utils.Message(false, "Invalid request"))
        return
    }
    messages, status := requests.IsSubmittedLoginFormValid(user)
    if !status {
        utils.Respond(w, messages)
        return
    }

    user, err = user.Authenticate()
    if err != nil {
        switch err {
        case models.ErrorNotFound:
            utils.Respond(w, utils.Message(false, "Invalid Email address"))
            break
        case models.ErrorInvalidPassword:
            utils.Respond(w, utils.Message(false, "Invalid password"))
            break
        default:
            utils.Respond(w, utils.Message(false, err.Error()))
            break
        }
        return
    }
    response := utils.Message(true, "Logged in successfully")
    response["user"] = user
    utils.Respond(w, response)
}

//GetAuthenticatedUser is used to retrieve the authenticated user via token.
func GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value(utils.ContextKeyAuthToken).(uint)
    user, err := models.ByID(userID)
    if err != nil {
        utils.Respond(w, utils.Message(false, "Couldn't retrieve the user through token."))
        return
    }
    json.NewEncoder(w).Encode(user)
}

//ParseRegisterForm to parse the registration form when submitted.
func ParseRegisterForm(w http.ResponseWriter, r *http.Request) {

    user := &models.User{}
    err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
    if err != nil {
        utils.Respond(w, utils.Message(false, "Invalid request"))
        return
    }
    messages, status := requests.IsSubmittedRegisterFormValid(user)
    if !status {
        utils.Respond(w, messages)
        return
    }
    user, err = user.Create()
    user.Password = ""
    response := utils.Message(true, "Account has been created")
    response["user"] = user
    utils.Respond(w, response)
}
