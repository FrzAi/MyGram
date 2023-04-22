package repository

import "gorm.io/gorm"

type Repo struct {
	db *gorm.DB
}

type RepoInterface interface {
	UserRepo
	PhotoRepo
	CommentRepo
	SosialMediaRepo
}

// constructor function
func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db} // handle dependencies
}
