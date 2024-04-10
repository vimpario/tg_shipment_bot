package telegram_service

import (
	"awesomeProject/internal/pkg/model"
	tgbotapi "github.com/crocone/tg-bot"
)

func StartMenu() tgbotapi.InlineKeyboardMarkup {

	buttonContents := model.StartMenu()
	buttons := make([][]tgbotapi.InlineKeyboardButton, len(buttonContents))

	for i, buttonContent := range buttonContents {
		buttons[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonContent.Name, buttonContent.Data))
	}

	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func ProfileMenu() tgbotapi.InlineKeyboardMarkup {
	buttonContents := model.ProfileMenu()
	buttons := make([][]tgbotapi.InlineKeyboardButton, len(buttonContents))

	for i, buttonContent := range buttonContents {
		buttons[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonContent.Name, buttonContent.Data))
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}
