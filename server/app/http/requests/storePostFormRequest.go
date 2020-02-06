package requests

import (
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
)

// IsSubmittedStorePostFormValid is used to validate the passed request.
func IsSubmittedStorePostFormValid(post *models.Post) (map[string]interface{}, bool) {
    if post.Title == "" {
        return utils.Message(false, "Title is required"), false
    }
    if post.Body == "" {
        return utils.Message(false, "Body is required"), false
    }
    return utils.Message(false, "Requirement passed"), true

}
