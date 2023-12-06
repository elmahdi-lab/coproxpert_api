package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LocalLogger struct{}

func (l *LocalLogger) LogInfo(message string) {
	l.logToFile("INFO", message)
}

func (l *LocalLogger) LogError(message string) {
	l.logToFile("ERROR", message)
}

func (l *LocalLogger) logToFile(logLevel, message string) {
	logFilePath := os.Getenv("LOG_FILE_PATH")
	if logFilePath == "" {
		log.Fatal("LOG_FILE_PATH environment variable is not set")
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close log file: %v", err)
		}
	}(file)

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, logLevel, message)
	_, err = file.WriteString(logMessage)
	if err != nil {
		log.Fatalf("failed to write to log file: %v", err)
	}
}
