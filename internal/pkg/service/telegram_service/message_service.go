package telegram_service

import (
	tgbotapi "github.com/crocone/tg-bot"
	"log"
)

func SendMessage(msg tgbotapi.Chattable, bot *tgbotapi.BotAPI) {

	if _, err := bot.Send(msg); err != nil {
		log.Panic("Send message error: %v", err)
	}

}
