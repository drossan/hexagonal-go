package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SlackNotification struct {
	webhookURL string
}

func NewSlackNotification(webhookURL string) *SlackNotification {
	return &SlackNotification{webhookURL: webhookURL}
}

func (s *SlackNotification) Send(to string, message string) error {
	payload := map[string]string{
		"channel": to,
		"text":    message,
	}
	payloadBytes, _ := json.Marshal(payload)

	resp, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send slack notification, status code: %d", resp.StatusCode)
	}

	return nil
}
