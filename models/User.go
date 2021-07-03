package models

import "time"

type User struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	HashPassword string    `json:"hash"`
	ImageURL     string    `json:"image_url"`
	CreatedOn    time.Time `json:"created_on"`
	LastLogin    time.Time `json:"last_login"`
}

type Users struct {
	Users []User `json:"users"`
}
