package Decorator

import (
	"fmt"
	"time"
)

// INotification is the interface representing the core notification functionality.
type INotification interface {
	Send()
}

// EmailNotification is a concrete notification type for sending emails.
type EmailNotification struct {
	recipientEmail string
}

// Send sends an email notification to the recipient.
func (en *EmailNotification) Send() {
	fmt.Printf("Sending email to %s\n", en.recipientEmail)
}

// SMSNotification is a concrete notification type for sending SMS messages.
type SMSNotification struct {
	recipientPhoneNo string
}

// Send sends an SMS notification to the recipient.
func (sn *SMSNotification) Send() {
	fmt.Printf("Sending SMS to %s\n", sn.recipientPhoneNo)
}

// IDecoratorNotification is the interface for notification decorators.
type IDecoratorNotification interface {
	Send()
	UpdateNotificationParams(data string)
}

// MessageDecorator is a decorator that adds messages to the notification.
type MessageDecorator struct {
	notification INotification
	messages     []string
}

// UpdateNotificationParams updates the decorator's messages.
func (md *MessageDecorator) UpdateNotificationParams(data string) {
	md.messages = append(md.messages, data)
}

// Send sends the notification with added messages.
func (md *MessageDecorator) Send() {
	md.notification.Send()
	for _, message := range md.messages {
		fmt.Println(message)
	}
}

// ScheduleDecorator is a decorator that adds scheduling information to the notification.
type ScheduleDecorator struct {
	notification INotification
	time         time.Time
}

// UpdateNotificationParams updates the decorator's scheduling time.
func (sd *ScheduleDecorator) UpdateNotificationParams(data string) {
	sd.time, _ = time.Parse("2006-01-02 15:04", data)
}

// Send sends the notification with added scheduling information.
func (sd *ScheduleDecorator) Send() {
	sd.notification.Send()
	fmt.Printf("This notification is scheduled for %v\n", sd.time.Format(time.RFC822))
}
