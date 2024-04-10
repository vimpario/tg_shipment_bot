package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	BotToken = "TG_API_BOT_TOKEN"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env not loaded")
	}
}

func GetValue(name string) string {
	return os.Getenv(name)
}
