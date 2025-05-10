package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ZerologInit() {

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339},
	).With().Caller().Logger()

}

var (
	checkingMessageCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "checking_message_counter",
			Help: "Number of messages that were sent to a checking server",
		})
)

func main() {
	ZerologInit()

	prometheus.MustRegister(checkingMessageCounter)

	log.Info().Msg("Program Started!")

	go func() {
		log.Info().Msg("Starting metrics server on port 2112")
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal().Err(http.ListenAndServe(":2112", nil)).Msg("Metrics server failed")

	}()

	signs := make(chan os.Signal, 1)
	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signs
		log.Info().Msg("Program Stopped!")
		os.Exit(0)
	}()

	for {
		log.Info().Msg("Checking...")
		checkingMessageCounter.Inc()

		time.Sleep(5 * time.Second)
		go func() {
			log.Info().Msg("All good")
			time.Sleep(20 * time.Second)
		}()
	}
}
