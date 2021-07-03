package models

import (
	"time"
)

type Message struct {
	ID      uint   `json:"id"`
	FromUserId  uint   `json:"from_user_id"`
	ToUserId  uint   `json:"to_user_id"`
	Content string `json:"content"`
	CreatedOn time.Time `json:"created_on"`
}

type Messages struct {
	Messages []Message `json:"messages"`
}
