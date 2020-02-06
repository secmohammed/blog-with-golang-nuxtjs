package requests

import (
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
    "strings"
)

// ForgetPasswordFormRequest type.
type ForgetPasswordFormRequest struct {
    Email string `json:"email"`
}

// IsSubmittedForgetPasswordFormValid is used to validate the passed request.
func IsSubmittedForgetPasswordFormValid(form *ForgetPasswordFormRequest) (map[string]interface{}, bool) {
    if !strings.Contains(form.Email, "@") {
        return utils.Message(false, "Email address is required"), false
    }
    //Email must be unique
    _, err := models.ByEmail(form.Email)
    if err != nil {
        return utils.Message(false, "Email Doesn't exist"), false
    }

    return utils.Message(false, "Requirement passed"), true

}
