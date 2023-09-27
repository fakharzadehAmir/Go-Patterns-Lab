package Chain_of_Responsibility

import "fmt"

// LogLevel represents the severity levels for logging.
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERR
)

// ILogger is the interface for loggers in the chain.
type ILogger interface {
	LogMessage(level LogLevel, message string)
	setNextLogger(next ILogger)
}

// ConsoleLogger represents a logger that logs messages to the console.
type ConsoleLogger struct {
	Next ILogger // Next logger in the chain
}

func (cl *ConsoleLogger) setNextLogger(next ILogger) {
	cl.Next = next
}

func (cl *ConsoleLogger) LogMessage(logLev LogLevel, message string) {
	if logLev == INFO || logLev == WARN {
		fmt.Printf("Console Log: %s", message)
	} else if cl.Next != nil {
		cl.Next.LogMessage(logLev, message)
	}
}

// FileLogger represents a logger that logs messages to a file.
type FileLogger struct {
	Next ILogger // Next logger in the chain
}

func (fl *FileLogger) setNextLogger(next ILogger) {
	fl.Next = next
}

func (fl *FileLogger) LogMessage(logLev LogLevel, message string) {
	if logLev == ERR {
		fmt.Printf("File Log: %s", message)
	} else if fl.Next != nil {
		fl.Next.LogMessage(logLev, message)
	}
}
