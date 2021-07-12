package models

import (
	"time"
)

type Question struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Question  string    `json:"q_text"`
	CreatedOn time.Time `json:"created_on"`
	Username  string    `json:"username"`
	ImageURL  string    `json:"image_url"`
}

type Questions struct {
	Questions []Question `json:"qs"`
}
