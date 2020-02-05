package migrations

import (
    "go-auth-with-crud-api/server/app/models"
)

func init() {
    database.AutoMigrate(models.Activation{})
}

//RefreshAcivations function is used to take the tables down form the database and refresh it
func RefreshAcivations() {
    database.DropTableIfExists(&models.Activation{})
    database.AutoMigrate(&models.Activation{})
}
