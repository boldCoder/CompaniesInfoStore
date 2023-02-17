package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger() zerolog.Logger {
	writer := zerolog.ConsoleWriter{Out: os.Stdout}
	logger := zerolog.New(writer).With().Timestamp().Logger()

	return logger
}
