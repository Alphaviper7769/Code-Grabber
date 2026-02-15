package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init(logFilePath string) error {
	file, err := os.OpenFile(
		logFilePath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)
	if err != nil {
		return err
	}

	Logger = log.New(file, "[CodeGrabber] ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}
