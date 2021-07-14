package models

import (
	"time"
)

type Blog struct {
	ID            uint      `json:"id"`
	UserId        uint      `json:"user_id"`
	BlogTitle     string    `json:"blog_title"`
	BlogThumbnail string    `json:"blog_thumbnail"`
	BlogText      string    `json:"blog_text"`
	CreatedOn     time.Time `json:"created_on"`
	Username      string    `json:"username"`
	ImageURL      string    `json:"image_url"`
}

type Blogs struct {
	Blogs []Blog `json:"blogs"`
}
