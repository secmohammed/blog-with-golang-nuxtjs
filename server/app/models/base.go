package models

// Update will update the provided struct with all of the data passed through the struct object.
func Update(dst interface{}) error {
    return db.Save(dst).Error
}

// Create will create the provided struct with all of the passed data.
func Create(dst interface{}) error {
    return db.Create(dst).Error
}
