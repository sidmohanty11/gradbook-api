package models

type Message struct {
	ID      uint   `json:"id"`
	UserId  uint   `json:"user_id"`
	Message string `json:"message"`
}

type Messages struct {
	Messages []Message `json:"messages"`
}
