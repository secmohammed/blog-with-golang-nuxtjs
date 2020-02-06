package utils

import (
    "math/rand"
)

var pool = "aXn31obYZ"

// GenerateRandomString a random string of A-Z chars with len = l
func GenerateRandomString(l int) string {
    bytes := make([]byte, l)
    for i := 0; i < l; i++ {
        bytes[i] = pool[rand.Intn(len(pool))]
    }
    return string(bytes)
}
