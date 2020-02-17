package users

import (
    "encoding/json"
    "go-auth-with-crud-api/server/app/http/requests"
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
    "net/http"
    "os"
    "strconv"

    "github.com/gorilla/mux"
)

// Update function is used to update the user's profile that's authenticated.
func Update(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value(utils.ContextKeyAuthToken).(uint)
    user, err := models.ByID(userID)
    maxUploadSize, err := strconv.ParseInt(os.Getenv("MAX_IMAGE_UPLOAD_SIZE_IN_BYTES"), 10, 64)
    if err != nil {
        utils.Respond(w, utils.Message(false, "Couldn't parse this user id"))
        return
    }
    r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
    if err := r.ParseMultipartForm(maxUploadSize); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        utils.Respond(w, utils.Message(false, err.Error()))
        return
    }
    err = utils.ParseForm(r, user)
    if err != nil {
        utils.Respond(w, utils.Message(false, "Invalid request "+err.Error()))
        return
    }

    messages, status := requests.IsSubmittedUpdateUserProfileFormValid(user, r)
    if !status {
        w.WriteHeader(http.StatusUnprocessableEntity)
        utils.Respond(w, messages)
        return
    }
    messages, status = utils.UploadFile(r, "avatar", os.Getenv("AVATAR_STORAGE_PATH"))
    if !status {
        utils.Respond(w, messages)
        return
    }
    // if the avatar status while uploading is true, it means that
    // user uploaded a photo, therefore we will be updating
    // the avatar attribute. unless that, we won't update the attribute
    if messages["status"].(bool) {
        user.Avatar = messages["message"].(string)
    }
    err = models.Update(&user)
    if err != nil {
        utils.Respond(w, utils.Message(false, "Couldn't update your info."))
        return
    }
    response := utils.Message(true, "Profile has been updated successfully")
    user.Password = ""
    response["user"] = user
    utils.Respond(w, response)
}

// Show function is used to show the user's profile.
func Show(w http.ResponseWriter, r *http.Request) {
    userID, err := strconv.ParseUint(mux.Vars(r)["user"], 10, 32)
    if err != nil {
        utils.Respond(w, utils.Message(false, "Couldn't parse this user id"))
        return
    }
    user, err := models.ByID(uint(userID))
    user.Password = ""
    user.Token = ""
    if err != nil {
        utils.Respond(w, utils.Message(false, err.Error()))
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}
