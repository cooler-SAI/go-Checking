package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

func main() {
	ZerologInit()

	log.Info().Msg("Program Started!")

	signs := make(chan os.Signal, 1)
	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signs
		log.Info().Msg("Program Stopped!")
		os.Exit(0)
	}()

	for {
		log.Info().Msg("Checking...")
		time.Sleep(5 * time.Second)
		go func() {
			log.Info().Msg("All good")
			time.Sleep(20 * time.Second)
		}()
	}
}
