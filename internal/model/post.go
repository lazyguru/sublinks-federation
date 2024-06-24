package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id          string    `json:"id" gorm:"primary_key"`
	UrlStub     string    `json:"url_stub"`
	Title       string    `json:"title"`
	Author      Person    `json:"author"`
	AuthorId    string    ``
	Community   Group     `json:"community"`
	CommunityId string    ``
	Nsfw        bool      `json:"nsfw"`
	Published   time.Time `json:"published"`
	Content     string    `json:"content"`
	Comments    []Comment `json:"comments,omitempty" gorm:"foreignKey:PostId"`
}
