package Template_Method

import "testing"

func TestOTPGenerationAndNotification(t *testing.T) {
	t.Run("Test SMS OTP Generation and Notification", func(t *testing.T) {
		smsOTP := &SMS{}
		otpProcessor := &OneTimePass{iOTP: smsOTP}
		err := otpProcessor.generateAndSendPassword()
		if err != nil {
			t.Errorf("Error generating and sending SMS OTP: %v", err)
		}
	})

	t.Run("Test Email OTP Generation and Notification", func(t *testing.T) {
		emailOTP := &Email{}
		otpProcessor := &OneTimePass{iOTP: emailOTP}
		err := otpProcessor.generateAndSendPassword()
		if err != nil {
			t.Errorf("Error generating and sending Email OTP: %v", err)
		}
	})

	t.Run("Test Phone Call OTP Generation and Notification", func(t *testing.T) {
		phoneCallOTP := &PhoneCall{}
		otpProcessor := &OneTimePass{iOTP: phoneCallOTP}
		err := otpProcessor.generateAndSendPassword()
		if err != nil {
			t.Errorf("Error generating and sending Phone Call OTP: %v", err)
		}
	})
}
