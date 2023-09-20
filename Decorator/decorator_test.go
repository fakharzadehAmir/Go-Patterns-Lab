package Decorator

import (
	"testing"
	"time"
)

func TestDecoratorPattern(t *testing.T) {
	// Create email notification
	emailNotification := &EmailNotification{recipientEmail: "test@example.com"}

	// Create SMS notification
	smsNotification := &SMSNotification{recipientPhoneNo: "+123456789"}

	// Create message decorator for email
	emailMessageDecorator := &MessageDecorator{notification: emailNotification}
	emailMessageDecorator.UpdateNotificationParams("Email: Hello from Decorator Pattern!")
	emailMessageDecorator.UpdateNotificationParams("Email: This is another message.")

	// Create message decorator for SMS
	smsMessageDecorator := &MessageDecorator{notification: smsNotification}
	smsMessageDecorator.UpdateNotificationParams("SMS: Hello from Decorator Pattern!")
	smsMessageDecorator.UpdateNotificationParams("SMS: This is another message.")

	// Create schedule decorator for email
	emailScheduleDecorator := &ScheduleDecorator{notification: emailMessageDecorator, time: time.Now().Add(2 * time.Hour)}

	// Create schedule decorator for SMS
	smsScheduleDecorator := &ScheduleDecorator{notification: smsMessageDecorator, time: time.Now().Add(3 * time.Hour)}

	// Test Email Notification
	t.Run("Email Notification", func(t *testing.T) {
		emailNotification.Send()
	})

	// Test SMS Notification
	t.Run("SMS Notification", func(t *testing.T) {
		smsNotification.Send()
	})

	// Test Message Decorator for Email
	t.Run("Message Decorator for Email", func(t *testing.T) {
		emailMessageDecorator.Send()
	})

	// Test Message Decorator for SMS
	t.Run("Message Decorator for SMS", func(t *testing.T) {
		smsMessageDecorator.Send()
	})

	// Test Schedule Decorator for Email
	t.Run("Schedule Decorator for Email", func(t *testing.T) {
		emailScheduleDecorator.Send()
	})

	// Test Schedule Decorator for SMS
	t.Run("Schedule Decorator for SMS", func(t *testing.T) {
		smsScheduleDecorator.Send()
	})
}
