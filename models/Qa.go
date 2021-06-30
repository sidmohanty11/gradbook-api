package models

type Answer struct {
	ID     uint   `json:"id"`
	Answer string `json:"answer"`
}

type Qa struct {
	ID       uint     `json:"id"`
	UserId   uint     `json:"user_id"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Qas struct {
	Qas []Qa `json:"qas"`
}
