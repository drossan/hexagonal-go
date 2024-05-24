package adapters

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/drossan/core-api/domain/notification"
)

type EmailNotifier struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	FromEmail    string
	auth         smtp.Auth
}

func NewEmailNotifier(smtpHost, smtpPort, smtpUser, smtpPassword, fromEmail string) *EmailNotifier {
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	return &EmailNotifier{
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SMTPUser:     smtpUser,
		SMTPPassword: smtpPassword,
		FromEmail:    fromEmail,
		auth:         auth,
	}
}

func (e *EmailNotifier) SendNotification(message string) error {
	to := os.Getenv("TO_EMAIL")
	subject := "Notification"
	body := message

	msg := "From: " + e.FromEmail + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(e.SMTPHost+":"+e.SMTPPort, e.auth, e.FromEmail, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}
	log.Printf("Email sent to %s", to)
	return nil
}

func (e *EmailNotifier) SendNotificationWithAttachments(attachments []notification.Attachment) error {
	// Implementar si es necesario
	return nil
}

func (e *EmailNotifier) SendNotificationWithTemplate(templatePath string, data interface{}) error {
	to := os.Getenv("TO_EMAIL")
	subject := "Notification"

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Failed to parse template: %v", err)
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Printf("Failed to execute template: %v", err)
		return err
	}

	msg := "From: " + e.FromEmail + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body.String()

	err = smtp.SendMail(e.SMTPHost+":"+e.SMTPPort, e.auth, e.FromEmail, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("Failed to send email with template: %v", err)
		return err
	}
	log.Printf("Email sent to %s with template", to)
	return nil
}
