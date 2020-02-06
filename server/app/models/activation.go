package models

import (
    "go-auth-with-crud-api/server/utils"

    "github.com/jinzhu/gorm"
)

// Activation type
type Activation struct {
    gorm.Model
    UserID uint   `json:"user_id"`
    Token  string `json:"token"`
    Active bool   `json:"active"`
}

// ByUserID function is used to fetch the record by the user id.
func ByUserID(id uint) (*Activation, error) {
    var activation Activation
    err := db.Where("user_id = ?", id).First(&activation).Error
    if err != nil {
        return nil, err
    }
    return &activation, nil
}

//ByActivationToken function is used to fetch the record by the token passed.
func ByActivationToken(token string) (*Activation, error) {
    var activation Activation
    err := db.Where("token = ?", token).First(&activation).Error
    if err != nil {
        return nil, err
    }
    return &activation, nil
}

//IsActivated function is used to check on the user if activated or not.
func (user *User) IsActivated() (bool, error) {
    var activation Activation
    err := db.Where("user_id = ? ", user.ID).First(&activation).Error
    if err != nil {
        return false, err
    }
    if activation.Token != "" || !activation.Active {
        return false, nil
    }
    return true, nil
}

//Activate function is used to activate the current authenticated user.
func (activation *Activation) Activate() (*Activation, error) {
    activation.Token = ""
    activation.Active = true
    err := db.Save(&activation).Error
    if err != nil {
        return nil, err
    }
    return activation, nil
}

//GenerateToken function is used to create a token for the associated user.
func (user *User) GenerateToken() (*Activation, error) {
    activation := &Activation{
        UserID: user.ID,
        Token:  utils.GenerateRandomString(32),
    }
    err := db.Create(&activation).Error
    return activation, err
}
