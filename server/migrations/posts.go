package migrations

import (
    "go-auth-with-crud-api/server/app/models"
)

func init() {
    database.AutoMigrate(models.Post{})
}

//RefreshPosts function is used to take the tables down form the database and refresh it
func RefreshPosts() {
    database.DropTableIfExists(&models.Post{})
    database.AutoMigrate(&models.Post{})
}
