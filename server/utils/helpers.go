package utils

import (
    "net/http"

    "github.com/gorilla/schema"
)

// Must function is used to parse if there is an error being thrown or not.
func Must(err error) {
    if err != nil {
        panic(err)
    }
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
