package migrations

import (
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
)

var database = utils.GetDatabaseConnection()

func init() {
    database.AutoMigrate(models.User{})
}

//Refresh function is used to take the tables down form the database and refresh it
func Refresh() {
    database.DropTableIfExists(&models.User{})
    database.AutoMigrate(&models.User{})
}
