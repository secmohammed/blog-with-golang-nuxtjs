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
    Active bool   `json:"active"; gorm:"default:false"`
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

//Activate function is used to activate the current authenticated user.
func (user *User) Activate() (*Activation, error) {
    activation, err := ByUserID(user.ID)
    if err != nil {
        return nil, err
    }
    db.Model(&activation).Updates(Activation{
        Token:  "",
        Active: true,
    })
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
