package notification

type Notifier interface {
	SendNotification(message string) error
	SendNotificationWithAttachments(attachments []Attachment) error
	SendNotificationWithTemplate(template string, data interface{}) error
}

type Attachment struct {
	Title  string
	Text   string
	Color  string
	Fields []Field
}

type Field struct {
	Title string
	Value string
	Short bool
}
