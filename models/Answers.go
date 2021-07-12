package models

type Answer struct {
	ID         uint   `json:"id"`
	QId        uint   `json:"q_id"`
	UserId     uint   `json:"user_id"`
	AnswerText string `json:"a_text"`
	Username   string `json:"username"`
}

type Answers struct {
	Answers []Answer `json:"ans"`
}
