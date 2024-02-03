/*
Copyright © 2024 Anurag Yadav <contact@anuragyadav.in>
*/

package main

import (
	"go-jwt/cmd"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func main() {

	// Start CMD
	cmd.Execute()
}

func init() {
	// Setup logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading .env file: %s", err.Error())
	}

}
