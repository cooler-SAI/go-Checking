package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
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

	for {
		log.Info().Msg("Checking...")

		time.Sleep(5 * time.Second)

	}
}
