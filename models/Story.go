package models

import "database/sql"

type Story struct {
	ID           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	Name         string `json:"name"`
	Branch       string `json:"branch"`
	Clubs        string `json:"clubs"`
	Motto        string `json:"motto"`
	GithubLink   sql.NullString `json:"github_link"`
	YoutubeLink  sql.NullString `json:"youtube_link"`
	LinkedinLink sql.NullString `json:"linkedin_link"`
	ImageURL     string `json:"image_url"`
	Journey      string `json:"journey"`
}

type Stories struct {
	Stories []Story `json:"stories"`
}
