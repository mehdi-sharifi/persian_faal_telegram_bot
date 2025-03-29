package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"persian_faal_bot/internal/services"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	poems, err := services.LoadPoems("data/hafez_poems.csv")
	if err != nil {
		log.Fatal("Error loading poems", err)
	}
	fmt.Println("Loading poems...")
	services.StartBot(botToken, poems)
}
