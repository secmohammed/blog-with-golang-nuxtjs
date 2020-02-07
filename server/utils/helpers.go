package utils

import (
    "fmt"
    "math/rand"
    "net/http"

    "github.com/gorilla/schema"
)

var pool = "aXn31obYZ"

// GenerateRandomString a random string of A-Z chars with len = l
func GenerateRandomString(l int) string {
    bytes := make([]byte, l)
    rand.Read(bytes)
    return fmt.Sprintf("%x", bytes)
}

// ParseForm function is used to parse the inputs from the form.
func ParseForm(r *http.Request, form interface{}) error {
    // parseForm must be called in order to fill the postForm with the data coming from the input form data.
    if err := r.ParseForm(); err != nil {
        return err
    }
    decoder := schema.NewDecoder()

    if err := decoder.Decode(form, r.PostForm); err != nil {
        return err
    }
    return nil

}
