package cmd

import (
	"log"

	"github.com/joho/godotenv"
)

func StartupEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

}
