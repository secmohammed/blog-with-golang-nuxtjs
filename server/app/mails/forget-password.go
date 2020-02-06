package mails

import (
    "fmt"
    "go-auth-with-crud-api/server/app/models"
    "log"
    "net/smtp"
    "os"
)

//SendForgetPasswordEmail function is used to send an reset password token for the passed email.
func SendForgetPasswordEmail(user *models.User, reminder *models.Reminder) {
    from := os.Getenv("MAIL_USERNAME")
    password := os.Getenv("MAIL_PASSWORD")
    host := os.Getenv("MAIL_HOST")
    port := os.Getenv("MAIL_PORT")
    fmt.Println(port, from)
    // Choose auth method and set it up
    msg := "From: " + from + "\n" +
        "To: " + user.Email + "\n" +
        "Subject: Welcome on board, Activate your account\n\n" +
        "Hello, " + user.Name + " \n\n We are glad that you have joined with us, here is your activation link " + os.Getenv("APP_URL") + "/activate/" + reminder.Token
    err := smtp.SendMail(host+":"+port,
        smtp.PlainAuth("", from, password, host),
        from, []string{user.Email}, []byte(msg))
    if err != nil {
        log.Printf("smtp error: %s", err)
        return
    }
}
