package requests

import (
    "go-auth-with-crud-api/server/utils"
)

// ResetPasswordFormRequest type.
type ResetPasswordFormRequest struct {
    Password             string `json:"password"`
    PasswordConfirmation string `json:"password_confirmation"`
}

// IsSubmittedResetPasswordFormValid is used to validate the passed request.
func IsSubmittedResetPasswordFormValid(form *ResetPasswordFormRequest) (map[string]interface{}, bool) {

    if len(form.Password) < 8 {
        return utils.Message(false, "Password must at least contain 8 characters"), false
    }
    if len(form.Password) > 32 {
        return utils.Message(false, "Password must not exceed 32 characters"), false
    }
    if len(form.PasswordConfirmation) < 8 {
        return utils.Message(false, "Password Confirmation must at least contain 8 characters"), false
    }
    if len(form.PasswordConfirmation) > 32 {
        return utils.Message(false, "Password Confirmation must not exceed 32 characters"), false
    }
    if form.Password != form.PasswordConfirmation {
        return utils.Message(false, "Password should match Password Confirmation"), false
    }
    return utils.Message(false, "Requirement passed"), true

}
