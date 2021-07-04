package models

import (
	"time"
)

type Answer struct {
	ID         uint   `json:"id"`
	QId        uint   `json:"q_id"`
	UserId     uint   `json:"user_id"`
	AnswerText string `json:"a_text"`
}

type Answers struct {
	Answers []Answer `json:"ans"`
}

type Question struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Question  string    `json:"q_text"`
	CreatedOn time.Time `json:"created_on"`
}

type Questions struct {
	Questions []Question `json:"qs"`
}
