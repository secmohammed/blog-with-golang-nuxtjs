package requests

import (
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
    "strings"
)

// IsSubmittedLoginFormValid is used to validate the passed request.
func IsSubmittedLoginFormValid(user *models.User) (map[string]interface{}, bool) {
    if !strings.Contains(user.Email, "@") {
        return utils.Message(false, "Email address is required"), false
    }
    if len(user.Password) < 6 {
        return utils.Message(false, "Password is less than 6 characters"), false
    }

    return utils.Message(false, "Requirement passed"), true

}
