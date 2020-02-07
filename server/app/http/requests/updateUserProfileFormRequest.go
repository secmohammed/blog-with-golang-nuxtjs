package requests

import (
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
    "io/ioutil"
    "net/http"
    "strings"

    "github.com/h2non/filetype"
)

// UpdateUserProfileFormRequest type.
type UpdateUserProfileFormRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// IsSubmittedUpdateUserProfileFormValid is used to validate the passed request.
func IsSubmittedUpdateUserProfileFormValid(form *models.User, r *http.Request) (map[string]interface{}, bool) {
    file, _, err := r.FormFile("avatar")
    if err != nil && err != http.ErrMissingFile {
        return utils.Message(false, err.Error()), false
    }
    if err != http.ErrMissingFile {
        fileBytes, err := ioutil.ReadAll(file)
        if err != nil {
            return utils.Message(false, "Invalid file"), false
        }

        kind, _ := filetype.Match(fileBytes)
        if kind == filetype.Unknown {
            return utils.Message(false, "Unknown file type"), false
        }
        fileMimeType := kind.MIME.Value
        if fileMimeType != "image/jpeg" && fileMimeType != "image/jpg" &&
            fileMimeType != "image/png" {
            return utils.Message(false, "File isn't an image."), false
        }

    }
    if form.Email != "" && !strings.Contains(form.Email, "@") {
        return utils.Message(false, "Invalid email address format."), false
    }
    if form.Name != "" && !strings.Contains(form.Name, " ") {
        return utils.Message(false, "You should type your full name separated by a space."), false
    }

    return utils.Message(false, "Requirement passed"), true

}
