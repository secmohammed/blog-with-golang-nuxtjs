package models

import (
    "os"

    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
)

// Token type.
type Token struct {
    UserID uint
    jwt.StandardClaims
}

// User type
type User struct {
    BaseGorm
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string
    Token    string `json:"token"`
    Avatar   string `json:"avatar"`
}

//Create function is used to create a users record
func (user *User) Create() (*User, error) {

    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user.Password = string(hashedBytes)
    err = Create(&user)
    return user, err
}

// ByID will look up the users using the given id.
func ByID(id uint) (*User, error) {
    var user User
    err := db.Where("id = ?", id).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// ResetPassword function is used to reset the user's password.
func (user *User) ResetPassword(password string) (*User, error) {

    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user.Password = string(hashedBytes)
    err = db.Save(&user).Error
    return user, err

}

//ChangePassword function is used to change the current logged in user's password.
func (user *User) ChangePassword(currentPassword, password string) (*User, error) {

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword))
    if err != nil {
        switch err {
        case bcrypt.ErrMismatchedHashAndPassword:
            return nil, ErrorInvalidPassword
        default:
            return nil, err
        }
    }
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user.Password = string(hashedBytes)
    err = db.Save(&user).Error
    return user, err
}

// Authenticate function is used to authorize user throughout his credentials.
func (user *User) Authenticate() (*User, error) {
    foundUser, err := ByEmail(user.Email)
    if err != nil {
        return nil, err
    }
    err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))

    if err != nil {
        switch err {
        case bcrypt.ErrMismatchedHashAndPassword:
            return nil, ErrorInvalidPassword
        default:
            return nil, err
        }
    }
    //Create JWT token
    tk := &Token{UserID: foundUser.ID}
    token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
    tokenString, _ := token.SignedString([]byte(os.Getenv("APP_KEY")))
    foundUser.Token = tokenString
    err = Update(foundUser)
    if err != nil {
        return nil, err
    }
    return foundUser, nil
}

//FindByAPIToken function is used to retrieve the current authenticated user via its token.
func FindByAPIToken(token string) (*User, error) {
    var user User
    err := db.Where("token = ?", token).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// ByEmail function will lok up the users using the given email.
func ByEmail(email string) (*User, error) {
    var user User
    err := db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
