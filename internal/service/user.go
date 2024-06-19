package service

import (
	"sublinks/sublinks-federation/internal/db"
	"sublinks/sublinks-federation/internal/model"
)

type UserService struct {
	db db.Database
}

func NewUserService(db db.Database) *UserService {
	return &UserService{db}
}

func (a UserService) GetById(id string) *model.Person {
	person := model.Person{}
	return a.Load(&person, id)
}

func (a UserService) GetByUsername(username string) *model.Person {
	person := model.Person{}
	err := a.db.Preload("Posts").Find(&person, model.Person{Username: username}).Error
	if err != nil {
		return nil
	}
	return &person
}

func (a UserService) Load(person *model.Person, id string) *model.Person {
	err := a.db.Preload("Posts").Find(person, id).Error
	if err != nil {
		return nil
	}
	return person
}

func (a UserService) Save(person *model.Person) bool {
	err := a.db.Save(person)
	return err == nil
}
