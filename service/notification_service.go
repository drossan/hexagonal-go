package service

import (
	"github.com/drossan/core-api/domain/notification"
	"github.com/drossan/core-api/domain/repository"
)

type NotificationService struct {
	notifiers map[string]notification.Notifier
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		notifiers: make(map[string]notification.Notifier),
	}
}

func (s *NotificationService) RegisterNotifier(name string, notifier notification.Notifier) {
	s.notifiers[name] = notifier
}

func (s *NotificationService) SendNotification(notifierNames []string, message string) error {
	for _, name := range notifierNames {
		if notifier, ok := s.notifiers[name]; ok {
			if err := notifier.SendNotification(message); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *NotificationService) SendNotificationWithAttachments(notifierNames []string, attachments []notification.Attachment) error {
	for _, name := range notifierNames {
		if notifier, ok := s.notifiers[name]; ok {
			if err := notifier.SendNotificationWithAttachments(attachments); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *NotificationService) SendNotificationWithTemplate(notifierNames []string, templatePath string, data interface{}) error {
	for _, name := range notifierNames {
		if notifier, ok := s.notifiers[name]; ok {
			if err := notifier.SendNotificationWithTemplate(templatePath, data); err != nil {
				return err
			}
		}
	}
	return nil
}

// Asegúrate de que NotificationService implemente NotificationServiceInterface
var _ repository.NotificationServiceInterface = &NotificationService{}
