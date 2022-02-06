package interval

import (
	"github.com/rs/zerolog/log"
	"merge/cmd/metric"
	"os"
	"strconv"
)

func Merge(input [][]uint64) ([][]uint64, error) {

	metrics := os.Getenv("PRINT_METRICS")

	if len(metrics) > 0 {
		defer metric.TimerDuration(metric.TimerStart("Merge processing time"))
		defer metric.PrintMemUsage()
	}

	log.Info().Str("intervals", strconv.Itoa(len(input))).Msg("requested to merge")

	intervals, err := NewIntervals(input)

	if err != nil {
		return nil, err
	}

	intervals.Sort()

	// create the output with the first interval as starting point
	out := Intervals{intervals[0]}

	// iterate through the remaining intervals and check for overlapping with the already merged one
	for i := 1; i < len(intervals); i++ {
		// check if the current interval overlaps the last one of the already merged
		overlaps := out[len(out)-1].Overlaps(intervals[i])

		if overlaps {
			// extend the current interval
			out[len(out)-1].Merge(intervals[i])
		} else {
			// add a new interval
			out = append(out, intervals[i])
		}
	}

	return out.Convert(), nil
}
