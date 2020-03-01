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
    if len(user.Password) < 8 {
        return utils.Message(false, "Password must at least contain 8 characters"), false
    }
    if len(user.Password) > 32 {
        return utils.Message(false, "Password must not exceed 32 characters"), false
    }
    return utils.Message(false, "Requirement passed"), true

}
