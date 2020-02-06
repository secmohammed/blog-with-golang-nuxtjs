package migrations

import (
    "go-auth-with-crud-api/server/app/models"
)

func init() {
    database.AutoMigrate(models.Reminder{})
}

//RefreshReminders function is used to take the tables down form the database and refresh it
func RefreshReminders() {
    database.DropTableIfExists(&models.Reminder{})
    database.AutoMigrate(&models.Reminder{})
}
