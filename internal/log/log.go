package log

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/matthewzhaocc/planetary-deploy/internal/util"
)

type Logger struct {
}

func Setup() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (l *Logger) LogMessage(message string) {
	log.Print(message)
}

func (l *Logger) LogRequest(r *http.Request) {
	logStr := util.JoinString(r.Method, r.URL.Path)
	log.Print(logStr)
}
