package internal

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logStatement struct {
	message string
}

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func logMessage(message string) {
	log.Print(message)
}
