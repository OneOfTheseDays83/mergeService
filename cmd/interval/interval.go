package interval

import (
	"errors"
	"github.com/rs/zerolog/log"
	"sort"
	"strconv"
)

type Interval struct {
	start uint64
	end   uint64
}

type Intervals []Interval

func NewIntervals(intervals [][]uint64) (Intervals, error) {
	ret := make([]Interval, 0, len(intervals))

	for _, value := range intervals {
		if len(value) != 2 {
			log.Error().Str("count", strconv.Itoa(len(value))).Msg("wrong number of elements")
			return ret, errors.New("interval had not exactly two elements")
		}

		// sort start and end in case someone gave us wrong order
		ret = append(ret, Interval{
			start: min(value[0], value[1]),
			end:   max(value[0], value[1]),
		})
	}

	return ret, nil
}

func (i *Intervals) Sort() {
	sort.Slice(*i, func(j, k int) bool {
		return (*i)[j].start < (*i)[k].start
	})
}

func (i *Intervals) Convert() [][]uint64 {
	var res [][]uint64

	for _, value := range *i {
		res = append(res, []uint64{value.start, value.end})
	}

	return res
}

// Overlaps checks if the intervals overlap
func (i *Interval) Overlaps(other Interval) (ret bool) {

	// other starts within given interval
	// first condition not necessary if the Intervals are sorted (then next interval is always greater or equal)
	if (other.start >= i.start) && (other.start <= i.end) {
		ret = true
	}

	return
}

func (i *Interval) Merge(other Interval) {
	i.end = max(i.end, other.end)
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
