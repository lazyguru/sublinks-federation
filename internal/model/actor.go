package model

import "gorm.io/gorm"

type Actor struct {
	ActorType    string `json:"actor_type"`
	Id           string `json:"id"`
	Username     string `json:"username,omitempty"`
	Name         string `json:"name,omitempty"`
	Bio          string `json:"bio"`
	MatrixUserId string `json:"matrix_user_id,omitempty"`
	PublicKey    string `json:"public_key"`
	PrivateKey   string `json:"private_key"`
}

type Person struct {
	gorm.Model
	Id           string `json:"id" gorm:"primarykey"`
	Username     string `json:"username,omitempty" gorm:"not null"`
	Name         string `json:"name,omitempty" gorm:"nullable"`
	Bio          string `json:"bio"`
	MatrixUserId string `json:"matrix_user_id,omitempty"`
	PublicKey    string `json:"public_key"`
	PrivateKey   string `json:"private_key"`
	Posts        []Post `json:"posts,omitempty" gorm:"foreignKey:AuthorId"`
}

type Group struct {
	gorm.Model
	Id         string `json:"id" gorm:"primarykey"`
	Username   string `json:"username,omitempty" gorm:"not null"`
	Name       string `json:"name,omitempty" gorm:"nullable"`
	Bio        string `json:"bio"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	Posts      []Post `json:"posts,omitempty" gorm:"foreignKey:CommunityId"`
}
