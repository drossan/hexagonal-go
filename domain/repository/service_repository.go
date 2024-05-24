package repository

import (
	"github.com/drossan/core-api/domain/notification"
)

type NotificationServiceInterface interface {
	RegisterNotifier(name string, notifier notification.Notifier)
	SendNotification(notifierNames []string, message string) error
	SendNotificationWithAttachments(notifierNames []string, attachments []notification.Attachment) error
	SendNotificationWithTemplate(notifierNames []string, templatePath string, data interface{}) error
}
