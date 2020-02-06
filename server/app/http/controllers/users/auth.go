package users

import (
    "encoding/json"
    "net/http"

    "go-auth-with-crud-api/server/app/http/requests"
    "go-auth-with-crud-api/server/app/mails"
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"

    "github.com/gorilla/mux"
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
        w.WriteHeader(http.StatusForbidden)

        utils.Respond(w, utils.Message(false, err.Error()))
        return
    }
    isActive, err := user.IsActivated()
    if !isActive {
        w.WriteHeader(http.StatusForbidden)
        utils.Respond(w, utils.Message(false, "You have not activated your account yet."))
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

//ParseChangePasswordForm is used to parse the user's password.
func ParseChangePasswordForm(w http.ResponseWriter, r *http.Request) {
    var form requests.ChangePasswordFormRequest
    err := json.NewDecoder(r.Body).Decode(&form) //decode the request body into struct and failed if any error occur
    if err != nil {
        utils.Respond(w, utils.Message(false, "Invalid request"))
        return
    }
    messages, status := requests.IsSubmittedChangePasswordFormValid(&form)
    if !status {
        w.WriteHeader(http.StatusUnprocessableEntity)
        utils.Respond(w, messages)
        return
    }
    userID := r.Context().Value(utils.ContextKeyAuthToken).(uint)
    user, err := models.ByID(userID)
    if err != nil {
        w.WriteHeader(http.StatusForbidden)
        utils.Respond(w, utils.Message(false, "Could not find this user throughout its token."))
        return
    }
    user, err = user.ChangePassword(form.CurrentPassword, form.Password)
    if err != nil {
        w.WriteHeader(http.StatusForbidden)
        utils.Respond(w, utils.Message(false, err.Error()))
        return
    }
    user.Password = ""
    response := utils.Message(true, "Password has been changed")
    response["user"] = user
    utils.Respond(w, response)
}

//ActivateRegisteredAccount function is used to activate the account through out the passed token.
func ActivateRegisteredAccount(w http.ResponseWriter, r *http.Request) {
    token := mux.Vars(r)["token"]
    if len(token) != 32 {
        utils.Respond(w, utils.Message(false, "Invalid Token format."))
        return
    }

    activation, err := models.ByActivationToken(token)
    if err != nil {
        utils.Respond(w, utils.Message(false, err.Error()))
        return
    }
    activation, err = activation.Activate()
    if err != nil {
        utils.Respond(w, utils.Message(false, err.Error()))

        return
    }
    utils.Respond(w, utils.Message(true, "User has been activated successfully."))
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
        w.WriteHeader(http.StatusUnprocessableEntity)
        utils.Respond(w, messages)
        return
    }
    user, err = user.Create()
    activation, err := user.GenerateToken()
    if err != nil {
        utils.Respond(w, utils.Message(false, "Couldn't create an activation token, try again."))
        return
    }
    mails.SendActivationMail(user, activation)
    user.Password = ""
    response := utils.Message(true, "Account has been created")
    response["user"] = user
    utils.Respond(w, response)
}
