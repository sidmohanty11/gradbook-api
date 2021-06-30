package models

type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Email          string `json:"email" gorm:"unique"`
	Password       []byte `json:"-"`
	Branch         string `json:"branch"`
	RegistrationNo string `json:"reg_no"`
}

type Users struct {
	Users User `json:"users"`
}
