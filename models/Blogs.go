package models

import (
	"time"
)

type Blog struct {
	ID      uint   `json:"id"`
	UserId  uint   `json:"user_id"`
	BlogTitle string `json:"blog_title"`
	BlogText string `json:"blog_text"`
	CreatedOn time.Time `json:"created_on"`
}

type Blogs struct {
	Blogs []Blog `json:"blogs"`
}