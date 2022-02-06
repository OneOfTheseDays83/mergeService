package main

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"math/rand"
	"merge/api"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		log.Error().Msg("Wrong number of args given")
		return
	}

	intervals, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		log.Error().Msg("Invalid number of elements given")
		return
	}

	input, err := randomInput(intervals)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/v1/merge", bytes.NewBuffer(input))

	log.Info().Str("count", strconv.Itoa(intervals)).Msg("intervals sent")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func randomInput(numberOfIntervals int) (out []byte, err error) {
	input := api.Request{}

	for i := 0; i < numberOfIntervals; i++ {
		input.Intervals = append(input.Intervals, getRandomInterval())
	}

	log.Info().Msgf("intervals sent10 ^ 4", input.Intervals)

	out, err = json.Marshal(input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal")
		return
	}

	return
}

func getRandomInterval() []uint64 {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 2 ^ 31
	return []uint64{uint64(rand.Intn(max-min) + min), uint64(rand.Intn(max-min) + min)}
}
