package logger

import (
	"io"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	logger = newLogger()
	once   sync.Once
)

type Logger struct {
	logger *zerolog.Logger
}

func newLogger() *Logger {
	var zlog zerolog.Logger

	once.Do(func() {
		var writer io.Writer = zerolog.ConsoleWriter{Out: os.Stderr}

		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zlog = zerolog.New(writer).With().Timestamp().Logger()
	})

	return &Logger{logger: &zlog}
}

func New() *zerolog.Logger {
	return logger.logger
}
