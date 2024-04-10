package menu_telegram_service

import (
	"awesomeProject/internal/pkg/model"
	"fmt"
)

func ProfileService(userProfile model.Profile, text string) string {
	text = fmt.Sprintf("Login: %v\nName: %v", userProfile.Login, userProfile.Name)
	return text
}
