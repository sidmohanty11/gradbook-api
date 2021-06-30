package models

type Story struct {
	ID     uint   `json:"id"`
	UserId uint   `json:"user_id"`
	Story  string `json:"story_text"`
}

type Stories struct {
	Stories []Story `json:"stories"`
}
