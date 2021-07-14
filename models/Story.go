package models

type Story struct {
	ID           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	Name         string `json:"name"`
	Branch       string `json:"branch"`
	Clubs        string `json:"clubs"`
	Motto        string `json:"motto"`
	GithubLink   string `json:"github_link"`
	YoutubeLink  string `json:"youtube_link"`
	LinkedinLink string `json:"linkedin_link"`
	ImageURL     string `json:"image_url"`
	Journey      string `json:"journey"`
	Username     string `json:"username"`
	UserImageUrl string `json:"user_image"`
}

type Stories struct {
	Stories []Story `json:"stories"`
}
