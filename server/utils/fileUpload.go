package utils

import (
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"

    "github.com/h2non/filetype"
)

// UploadFile is used to upload files.
func UploadFile(r *http.Request, fieldName, path string) (map[string]interface{}, bool) {
    file, _, _ := r.FormFile(fieldName)
    // if file not found, it means that user didn't upload a file(optional to upload)
    if file == nil {
        return Message(false, "Couldn't parse file from request."), true
    }
    defer file.Close()
    fileBytes, _ := ioutil.ReadAll(file)
    fileName := GenerateRandomString(12)
    kind, _ := filetype.Match(fileBytes)

    newPath := filepath.Join(path, fileName+"."+kind.Extension)
    newFile, err := os.Create(newPath)
    if err != nil {
        return Message(false, "Couldn't create file"+err.Error()), false
    }
    defer newFile.Close() // idempotent, okay to call twice
    // writing image content.
    if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
        return Message(false, "Couldn't write file"+err.Error()), false
    }
    return Message(true, newPath), true
}
