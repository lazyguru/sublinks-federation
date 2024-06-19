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
	group := model.Group{}
	return a.load(&group, id)
}

func (a CommunityService) GetByUsername(username string) *model.Group {
	person := model.Group{}
	err := a.db.Preload("Posts").Find(&person, model.Group{Username: username}).Error
	if err != nil {
		return nil
	}
	return &person
}

func (a CommunityService) load(group *model.Group, id string) *model.Group {
	err := a.db.Preload("Posts").Find(group, id).Error
	if err != nil {
		return nil
	}
	return group
}

func (a CommunityService) Save(group *model.Group) bool {
	err := a.db.Save(group)
	return err == nil
}
