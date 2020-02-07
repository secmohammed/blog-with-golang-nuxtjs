package models

import (
    "errors"
    "go-auth-with-crud-api/server/utils"

    "github.com/jinzhu/gorm"
)

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
    return db.Save(dst).Error
}

// Create will create the provided struct with all of the passed data.
func Create(dst interface{}) error {
    return db.Create(dst).Error
}

// Delete will delete the provided struct.
func Delete(dst interface{}, key, value string) error {
    return db.Where(key+" = ?", value).Delete(dst).Error
}
