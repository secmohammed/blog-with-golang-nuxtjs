package mails

import (
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
    msg := "From: " + from + "\n" +
        "To: " + user.Email + "\n" +
        "Subject: Reset Your Password\n\n" +
        "Hello, " + user.Name + " \n\n here is your password reset link" + os.Getenv("FRONTEND_APP_URL") + "auth/reset-password?token=" + reminder.Token
    err := smtp.SendMail(host+":"+port,
        smtp.PlainAuth("", from, password, host),
        from, []string{user.Email}, []byte(msg))
    if err != nil {
        log.Printf("smtp error: %s", err)
        return
    }
}
