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

type Qa struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Question  string    `json:"question"`
	CreatedOn time.Time `json:"created_on"`
	Answers   []Answer  `json:"ans"`
}

type Qas struct {
	Qas []Qa `json:"qas"`
}
