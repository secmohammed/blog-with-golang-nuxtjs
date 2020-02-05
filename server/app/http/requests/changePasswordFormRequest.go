package requests

import (
    "go-auth-with-crud-api/server/utils"
)

// ChangePasswordFormRequest type.
type ChangePasswordFormRequest struct {
    CurrentPassword      string `json:"current_password"`
    Password             string `json:"password"`
    PasswordConfirmation string `json:"password_confirmation"`
}

// IsSubmittedChangePasswordFormValid is used to validate the passed request.
func IsSubmittedChangePasswordFormValid(form *ChangePasswordFormRequest) (map[string]interface{}, bool) {
    if len(form.CurrentPassword) > 32 {
        return utils.Message(false, "Current Password must not exceed 32 characters."), false
    }
    if len(form.CurrentPassword) < 8 {
        return utils.Message(false, "Current Password must be at least 8 characters."), false
    }
    if len(form.Password) > 32 {
        return utils.Message(false, "Password must not exceed 32 characters."), false
    }
    if len(form.Password) < 8 || len(form.CurrentPassword) < 8 {
        return utils.Message(false, "Password must be at least 8 characters."), false
    }
    if form.Password != form.PasswordConfirmation {
        return utils.Message(false, "Passwords don't match."), false
    }
    return utils.Message(false, "Requirement passed"), true

}
