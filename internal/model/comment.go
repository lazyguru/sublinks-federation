package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id        string    `json:"id" gorm:"primary_key"`
	UrlStub   string    `json:"url_stub"`
	Post      Post      `json:"post_id"`
	PostId    string    ``
	Author    Person    `json:"author_id"`
	AuthorId  string    ``
	Nsfw      bool      `json:"nsfw"`
	Published time.Time `json:"published"`
	Content   string    `json:"content"`
}
