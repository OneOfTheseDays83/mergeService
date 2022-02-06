package metric

import (
	"github.com/rs/zerolog/log"
	"runtime"
	"strconv"
)

// from https://golangcode.com/print-the-current-memory-usage/

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	log.Info().Str("Alloc", strconv.FormatUint(m.Alloc/1024, 10)).Msg("in kB")
}
