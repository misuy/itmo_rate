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

func (user *User) AddReview(db *gorm.DB, review *Review) error {
	return db.Model(user).Association("Reviews").Append(review)
}
