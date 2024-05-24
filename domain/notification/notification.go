package notification

type Notification interface {
	Send(to string, message string) error
}
