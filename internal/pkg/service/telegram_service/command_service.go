package telegram_service

import (
	tgbotapi "github.com/crocone/tg-bot"
)

func Commands(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	command := update.Message.Command()
	switch command {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберимте действие")
		msg.ReplyMarkup = StartMenu()
		msg.ParseMode = "Markdown"
		SendMessage(msg, bot)
	case "profile":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Profile Case")
		msg.ReplyMarkup = ProfileMenu()
		//log.Println("Passed start command")
	}
}
