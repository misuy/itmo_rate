package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Role    string
	Reviews []Review
}

func NewUser(name string) User {
	return User{
		Name: name,
	}
}

func (user *User) AddReviews(db *gorm.DB, reviews []Review) error {
	return db.Model(user).Association("Reviews").Append(reviews)
}
