package models

import (
    "errors"
    "go-auth-with-crud-api/server/utils"
    "time"

    "github.com/jinzhu/gorm"
)

//BaseGorm is a type of struct to express the attributes in json format.
type BaseGorm struct {
    ID        uint       `json:"id" gorm:"primary_key"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

var (
    // db variable to get the database connection.
    db *gorm.DB = utils.GetDatabaseConnection()

    // ErrorNotFound is returned when a resource cannot be found.
    ErrorNotFound = errors.New("resource not found")

    //ErrorInvaildID will be thrown in case of the id is invalid or equal to zero.
    ErrorInvaildID = errors.New("ID provided was invalid")

    //ErrorInvalidPassword will be thrown in case of password mismatch
    ErrorInvalidPassword = errors.New("incorrect password provided")
)

// Update will update the provided struct with all of the data passed through the struct object.
func Update(dst interface{}) error {
    defer db.Close()

    return db.Save(dst).Error
}

// Create will create the provided struct with all of the passed data.
func Create(dst interface{}) error {
    defer db.Close()

    return db.Create(dst).Error
}

// Delete will delete the provided struct.
func Delete(dst interface{}, key, value string) error {
    defer db.Close()

    return db.Where(key+" = ?", value).Delete(dst).Error
}
