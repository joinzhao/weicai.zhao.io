package main

import (
	"gorm.io/gorm"
	"weicai.zhao.io/repox"
)

type UserModel struct {
	ID int
}

func (UserModel) TableName() string {
	return "admin_user"
}

func NewUserRepo(db *gorm.DB, batchSize int) UserRepo {
	return repox.New[*UserModel]("admin_user_repo", db, batchSize, func() *UserModel { return &UserModel{} })
}

type UserRepo = repox.ReaderWriterRepo[*UserModel]
