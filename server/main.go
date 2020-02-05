package main

import (
    "fmt"
    "os"

    // load dotenv variables
    _ "github.com/joho/godotenv/autoload"
    // load migrations

    _ "go-auth-with-crud-api/server/migrations"
    "go-auth-with-crud-api/server/routes"
)

func main() {
    fmt.Printf("server started on :%s", os.Getenv("APP_PORT"))
    routes.RegisterAPIRoutes()
}
