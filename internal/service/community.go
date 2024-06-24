package service

import (
	"sublinks/sublinks-federation/internal/db"
	"sublinks/sublinks-federation/internal/model"
)

type CommunityService struct {
	db db.Database
}

func NewCommunityService(db db.Database) *CommunityService {
	return &CommunityService{db}
}

func (a CommunityService) GetById(id string) *model.Group {
	group := &model.Group{}
	err := a.db.Preload("Posts").Find(group, model.Group{Id: id}).Error
	if err != nil {
		return nil
	}
	return group
}

func (a CommunityService) GetByUsername(username string) *model.Group {
	group := &model.Group{}
	err := a.db.Preload("Posts").Find(group, model.Group{Username: username}).Error
	if err != nil {
		return nil
	}
	return group
}

func (a CommunityService) Save(group *model.Group) bool {
	err := a.db.Save(group)
	return err == nil
}
