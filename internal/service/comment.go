package service

import (
	"sublinks/sublinks-federation/internal/db"
	"sublinks/sublinks-federation/internal/model"
)

type CommentService struct {
	db db.Database
}

func NewCommentService(db db.Database) *CommentService {
	return &CommentService{db}
}

func (p CommentService) GetById(id string) *model.Comment {
	comment := &model.Comment{}
	err := p.db.Preload("Post").Find(comment, model.Comment{Id: id}).Error
	if err != nil {
		return nil
	}
	return comment
}

func (p CommentService) Save(comment interface{}) bool {
	err := p.db.Save(comment)
	return err == nil
}
