package telegram_service

import (
	"awesomeProject/internal/pkg/model"
	"encoding/json"
	"log"
)

func UnmarshData(data []byte) string {
	var packageInfo model.PackageInfo
	var infoMsg string
	err := json.Unmarshal(data, &packageInfo)
	if err != nil {
		log.Fatalf("Error get JSON data")
	}

	infoMsg = packageInfo.Payload.TrackNumber + "\n" + packageInfo.Payload.LastStatus

	return infoMsg
}
