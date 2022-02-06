package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"merge/api"
	"merge/cmd/interval"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// used components
var ()

// init is the reserved golang function that will initialize all components once.
func init() {

}

func main() {
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, os.Interrupt)
	res := make(chan error, 1)
	defer close(res)

	port := os.Getenv("SERVICE_PORT")
	log.Info().Msgf("Listening in port %s", port)

	s := http.Server{
		Addr:    ":" + port,
		Handler: createRootHandler(),
	}
	go func() {
		res <- s.ListenAndServe()
	}()

	select {
	case <-quit:
		log.Info().Msg("user initiated termination of server started")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := s.Shutdown(ctx)
		if err != nil {
			log.Error().Err(err).Msg("graceful shutdown failed")
		}
	case err := <-res:
		log.Error().Err(err).Msg("server stopped with error")
	}
}

func createRootHandler() http.Handler {
	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/merge", merge).Methods(http.MethodPost)
	return r
}

func merge(rw http.ResponseWriter, r *http.Request) {
	var input api.Request
	if err := decode(r, &input); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(input.Intervals) > api.MAX_SIZE_INTERVALS {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Too much intervals, c.f. Api Spec"))
		return
	}

	// do the merging
	out, err := interval.Merge(input.Intervals)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
	}

	// respond
	if err := encode(rw, &out); err != nil {
		return
	}
}

func decode(r *http.Request, poi interface{}) (err error) {
	if poi == nil {
		return errors.New("is nil")
	}
	err = json.NewDecoder(r.Body).Decode(poi)
	if err != nil {
		log.Warn().Err(err).Msg("json decoding failed:")
	}
	return
}

func encode(rw http.ResponseWriter, out interface{}) (err error) {
	if out == nil {
		return errors.New("out is nil")
	}
	err = json.NewEncoder(rw).Encode(out)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
	return
}
