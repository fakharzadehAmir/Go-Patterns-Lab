package Chain_of_Responsibility

import (
	"testing"
)

func TestLoggerChain(t *testing.T) {
	// Create logger instances
	consoleLogger := &ConsoleLogger{}
	fileLogger := &FileLogger{}

	// Build the chain of responsibility
	consoleLogger.setNextLogger(fileLogger)

	// Test cases
	t.Run("Test INFO log", func(t *testing.T) {
		message := "This is an INFO message"
		consoleLogger.LogMessage(INFO, message)
	})

	t.Run("Test WARN log", func(t *testing.T) {
		message := "This is a WARN message"
		consoleLogger.LogMessage(WARN, message)
	})

	t.Run("Test ERR log", func(t *testing.T) {
		message := "This is an ERR message"
		consoleLogger.LogMessage(ERR, message)
	})

}
