package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint         `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	ImageURL  string       `json:"image_url"`
	CreatedOn time.Time    `json:"created_on"`
	LastLogin sql.NullTime `json:"last_login"`
}

type Users struct {
	Users []User `json:"users"`
}
