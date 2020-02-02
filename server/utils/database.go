package utils

import (
    "fmt"

    // setting up gorm with postgres connection
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "go-auth-with-crud-api/server/config"

    "github.com/jinzhu/gorm"
)

var databaseCredentials = config.GetDatabase()
var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    databaseCredentials["host"],
    databaseCredentials["port"],
    databaseCredentials["username"],
    databaseCredentials["password"],
    databaseCredentials["database"],
)
var db, err = gorm.Open("postgres", psqlInfo)

//GetDatabaseConnection function to return the current database connection.
func GetDatabaseConnection() *gorm.DB {
    db.LogMode(true)
    return db
}
