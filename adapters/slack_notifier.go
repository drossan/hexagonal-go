package adapters

import (
	"github.com/slack-go/slack"
	"log"
	"os"

	"github.com/drossan/core-api/domain/notification"
)

type SlackNotifier struct {
	Client *slack.Client
}

func NewSlackNotifier() *SlackNotifier {
	token := os.Getenv("SLACK_TOKEN")
	client := slack.New(token)
	return &SlackNotifier{
		Client: client,
	}
}

func (s *SlackNotifier) SendNotification(message string) error {
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	_, _, err := s.Client.PostMessage(channelID, slack.MsgOptionText(message, false))
	if err != nil {
		log.Printf("Failed to send notification to Slack: %v", err)
		return err
	}
	log.Printf("Notification sent to Slack channel %s", channelID)
	return nil
}

func (s *SlackNotifier) SendNotificationWithAttachments(attachments []notification.Attachment) error {
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	slackAttachments := make([]slack.Attachment, len(attachments))
	for i, attachment := range attachments {
		fields := make([]slack.AttachmentField, len(attachment.Fields))
		for j, field := range attachment.Fields {
			fields[j] = slack.AttachmentField{
				Title: field.Title,
				Value: field.Value,
				Short: field.Short,
			}
		}
		slackAttachments[i] = slack.Attachment{
			Title:  attachment.Title,
			Text:   attachment.Text,
			Color:  attachment.Color,
			Fields: fields,
		}
	}

	_, _, err := s.Client.PostMessage(channelID, slack.MsgOptionAttachments(slackAttachments...))
	if err != nil {
		log.Printf("Failed to send notification with attachments to Slack: %v", err)
		return err
	}
	log.Printf("Notification with attachments sent to Slack channel %s", channelID)
	return nil
}

// SendNotificationWithTemplate es un método no-op ya que Slack no usa plantillas.
func (s *SlackNotifier) SendNotificationWithTemplate(template string, data interface{}) error {
	return nil
}
