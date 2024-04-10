package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/pkg/service/telegram_service"
	tgbotapi "github.com/crocone/tg-bot"
	"log"
)

var urlApi = "https://api.pkge.net/v1/packages?trackNumber=LK263936204CN"
var bot *tgbotapi.BotAPI

func main() {

	botToken := config.GetValue(config.BotToken)

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram bot API: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Failed to start listening for updates: %v", err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			telegram_service.Callbacks(update, bot, urlApi)
		} else if update.Message.IsCommand() {
			telegram_service.Commands(update, bot)
		} else {
			log.Fatalf("Failed in Callack Query update")
		}

	}
}
