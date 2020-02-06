package models

import (
    "errors"
    "os"

    "go-auth-with-crud-api/server/utils"

    "github.com/dgrijalva/jwt-go"
    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
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

// Token type.
type Token struct {
    UserID uint
    jwt.StandardClaims
}

// User type
type User struct {
    gorm.Model
    Name       string `json:"name"`
    Email      string `json:"email"`
    Password   string `json:"password"`
    Token      string `json:"token" ;sql:"-"`
    Activation Activation
}

//Create function is used to create a users record
func (user *User) Create() (*User, error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user.Password = string(hashedBytes)
    err = db.Create(&user).Error
    return user, err
}

// UpdateUser will update the provided user with all of the data passed through the user object.
func UpdateUser(user *User) error {
    return db.Save(user).Error
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

//Delete function will delete the user with the provided id.
func Delete(id uint) error {
    // gorm will delete all of the records if the id equals to zero.
    if id == 0 {
        return ErrorInvaildID
    }
    // Gorm delete needs the primary key with the reference of the object to understand which table we are deleting from.
    user := User{
        Model: gorm.Model{
            ID: id,
        },
    }
    return db.Delete(&user).Error
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
    foundUser.Password = ""
    foundUser.Token = tokenString
    return foundUser, nil
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
