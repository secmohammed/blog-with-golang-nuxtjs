package utils

import (
    "encoding/json"
    "net/http"
)

//Message function is used to respond with a message with the status of the message.
func Message(status bool, message string) map[string]interface{} {
    return map[string]interface{}{"status": status, "message": message}
}

//Respond function is used to respond in json format.
func Respond(w http.ResponseWriter, data map[string]interface{}) {
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
