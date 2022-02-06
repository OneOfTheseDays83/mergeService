package metric

import (
	"github.com/rs/zerolog/log"
	"time"
)

func TimerStart(msg string) (string, time.Time) {
	return msg, time.Now()
}

func TimerDuration(msg string, start time.Time) {
	log.Info().Str("duration", time.Since(start).String()).Msg(msg)
}
