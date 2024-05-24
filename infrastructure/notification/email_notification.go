package notification

import (
	"fmt"
	"net/smtp"
)

type EmailNotification struct {
	smtpHost string
	smtpPort string
	auth     smtp.Auth
}

func NewEmailNotification(smtpHost, smtpPort, smtpUser, smtpPassword string) *EmailNotification {
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	return &EmailNotification{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		auth:     auth,
	}
}

func (e *EmailNotification) Send(to string, message string) error {
	from := "youremail@example.com"
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: Notification\n\n%s", from, to, message)
	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, e.auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
