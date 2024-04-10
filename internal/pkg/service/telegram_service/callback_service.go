package telegram_service

import (
	"awesomeProject/internal/pkg/model"
	"awesomeProject/internal/pkg/service/api_tracker_service"
	"awesomeProject/internal/pkg/service/menu_telegram_service"
	tgbotapi "github.com/crocone/tg-bot"
)

func Callbacks(update tgbotapi.Update, bot *tgbotapi.BotAPI, urlApi string) {
	userProfile := model.Profile{
		update.CallbackQuery.From.UserName,
		update.CallbackQuery.From.FirstName,
		update.CallbackQuery.From.LastName,
	}
	data := update.CallbackQuery.Data
	chatId := update.CallbackQuery.From.ID
	//firstName := update.CallbackQuery.From.FirstName
	//lastName := update.CallbackQuery.From.LastName
	bodyData := UnmarshData(api_tracker_service.GetTrackInfo(urlApi)) //maybe struct about message
	var text string
	switch data {
	case "profile":
		text = menu_telegram_service.ProfileService(userProfile, text)
	case "package":
		text = menu_telegram_service.PackageService(userProfile, bodyData)
	case "add":
		text = menu_telegram_service.AddService()
	case "info":
		text = menu_telegram_service.InfoService()
	default:
		text = "Unknown command"
	}

	msg := tgbotapi.NewMessage(chatId, text)
	if text != "" {
		SendMessage(msg, bot)
	}
}
