package menu_telegram_service

import (
	"awesomeProject/internal/pkg/model"
	"fmt"
)

func PackageService(userProfile model.Profile, bodyData string) string {
	text := fmt.Sprintf("Привет %v %v\nТвоя инфа:\n %v", userProfile.Name, userProfile.LastName, string(bodyData))
	return text
}
