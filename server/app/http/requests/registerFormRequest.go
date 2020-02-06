package requests

import (
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
    "strings"
)

// IsSubmittedRegisterFormValid is used to validate the passed request.
func IsSubmittedRegisterFormValid(user *models.User) (map[string]interface{}, bool) {
    if !strings.Contains(user.Email, "@") {
        return utils.Message(false, "Email address is required"), false
    }

    if len(user.Password) < 8 {
        return utils.Message(false, "Password must at least contain 8 characters"), false
    }
    if len(user.Password) > 32 {
        return utils.Message(false, "Password must not exceed 32 characters"), false
    }
    if len(user.Name) > 8 && !strings.Contains(user.Name, " ") {
        return utils.Message(false, "You should type your full name separated by a space."), false
    }

    //Email must be unique
    _, err := models.ByEmail(user.Email)
    if err == nil {
        return utils.Message(false, "Email address already in use by another user."), false
    }

    return utils.Message(false, "Requirement passed"), true

}
