package api

type Request struct {
	Intervals [][]uint64 `json:"intervals"`
}

const MAX_SIZE_INTERVALS int = 100000
