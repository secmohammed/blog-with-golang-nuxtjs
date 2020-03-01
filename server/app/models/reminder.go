package models

import (
    "go-auth-with-crud-api/server/utils"
)

// Reminder type
type Reminder struct {
    BaseGorm
    UserID uint   `json:"user_id"`
    Token  string `json:"token"`
}

//RevokeReminderToken is used to delete the associated record for the passed reminder.
func (reminder *Reminder) RevokeReminderToken() error {

    return db.Delete(&reminder).Error
}

//ByReminderToken function is used to fetch the record by the token passed.
func ByReminderToken(token string) (*Reminder, error) {

    var reminder Reminder
    err := db.Where("token = ?", token).First(&reminder).Error
    if err != nil {
        return nil, err
    }
    return &reminder, nil
}

//GenerateReminderToken function is used to create a token for the associated user.
func (user *User) GenerateReminderToken() (*Reminder, error) {
    reminder := &Reminder{
        UserID: user.ID,
        Token:  utils.GenerateRandomString(32),
    }

    err := db.Create(&reminder).Error
    return reminder, err
}
