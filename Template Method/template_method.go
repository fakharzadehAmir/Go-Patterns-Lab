package Template_Method

import (
	"fmt"
	"math/rand"
)

// IOneTimePassword is the interface that defines the methods for generating,
// sending, and notifying a one-time password.
type IOneTimePassword interface {
	getRandomPass() string
	getMessage(string) string
	Notify(string) error
	cacheOTP(string) string
}

// OneTimePass is the struct that contains the OTP generation and notification logic.
type OneTimePass struct {
	iOTP IOneTimePassword
}

// generateAndSendPassword generates an OTP, caches it, sends a message,
// and notifies the user.
func (otp *OneTimePass) generateAndSendPassword() error {
	createdPass := otp.iOTP.getRandomPass()
	otp.iOTP.cacheOTP(createdPass)
	message := otp.iOTP.getMessage(createdPass)
	err := otp.iOTP.Notify(message)
	if err != nil {
		return err
	}
	return nil
}

// SMS is a concrete implementation of IOneTimePassword for sending OTP via SMS.
type SMS struct {
	OneTimePass
}

// getRandomPass generates a random 4-digit OTP.
func (s *SMS) getRandomPass() string {
	pass := 1000 + rand.Intn(8999)
	return fmt.Sprintf("%v", pass)
}

// cacheOTP caches the OTP and returns a log message.
func (s *SMS) cacheOTP(otp string) string {
	return fmt.Sprintf("save password (via SMS), %v", otp)
}

// getMessage formats the OTP message.
func (s *SMS) getMessage(otp string) string {
	return "password (via SMS): " + otp
}

// Notify sends the OTP via SMS and logs the action.
func (s *SMS) Notify(message string) error {
	fmt.Printf("SMS is sent: %v\n", message)
	return nil
}

// Email is a concrete implementation of IOneTimePassword for sending OTP via email.
type Email struct {
	OneTimePass
}

// getRandomPass generates a random 4-digit OTP.
func (e *Email) getRandomPass() string {
	pass := 1000 + rand.Intn(8999)
	return fmt.Sprintf("%v", pass)
}

// cacheOTP caches the OTP and returns a log message.
func (e *Email) cacheOTP(otp string) string {
	return fmt.Sprintf("save password (via email), %v", otp)
}

// getMessage formats the OTP message.
func (e *Email) getMessage(otp string) string {
	return "password (via email): " + otp
}

// Notify sends the OTP via email and logs the action.
func (e *Email) Notify(message string) error {
	fmt.Printf("Email is sent: %v\n", message)
	return nil
}

// PhoneCall is a concrete implementation of IOneTimePassword for sending OTP via phone call.
type PhoneCall struct {
	OneTimePass
}

// getRandomPass generates a random 4-digit OTP.
func (pc *PhoneCall) getRandomPass() string {
	pass := 1000 + rand.Intn(8999)
	return fmt.Sprintf("%v", pass)
}

// cacheOTP caches the OTP and returns a log message.
func (pc *PhoneCall) cacheOTP(otp string) string {
	return fmt.Sprintf("save password (via phone call), %v", otp)
}

// getMessage formats the OTP message.
func (pc *PhoneCall) getMessage(otp string) string {
	return "password (via phone call): " + otp
}

// Notify sends the OTP via phone call and logs the action.
func (pc *PhoneCall) Notify(message string) error {
	fmt.Printf("Phone call is made: %v\n", message)
	return nil
}
