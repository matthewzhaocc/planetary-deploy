package internal

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logStatement struct {
	message string
}

func SetupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func LogMessage(message string) {
	log.Print(message)
}
