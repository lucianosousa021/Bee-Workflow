package model

type VideoMessage struct {
	Type   string `json:"type"`
	Data   string `json:"data"`
	UserID string `json:"userId"`
}
