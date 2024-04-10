package model

type PackageInfo struct {
	Code    int `json:"code"`
	Payload struct {
		TrackNumber string `json:"track_number"`
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		LastStatus  string `json:"last_status"`
	} `json:"payload"`
}

type Button struct {
	Name string
	Data string
}
