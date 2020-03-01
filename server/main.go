package main

import (
    "fmt"
    "os"

    // load dotenv variables
    "github.com/jinzhu/gorm"
    _ "github.com/joho/godotenv/autoload"
    // load migrations

    _ "go-auth-with-crud-api/server/migrations"
    "go-auth-with-crud-api/server/routes"
    "go-auth-with-crud-api/server/utils"
)

var db *gorm.DB = utils.GetDatabaseConnection()

func main() {
    fmt.Printf("server started on :%s", os.Getenv("APP_PORT"))
    routes.RegisterAPIRoutes()
    defer db.Close()
}
