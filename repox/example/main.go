package main

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/repox"
)

func main() {

	var config = &gormx.Config{
		Usage:       "default",
		RunMode:     "debug",
		DSN:         "root:root@tcp(127.0.0.1:3307)",
		Database:    "worker_app",
		MaxIdleConn: 10,
		MaxOpenConn: 10,
		MaxLifeTime: 10,
	}

	manage := gormx.New([]*gormx.Config{config})

	m := NewUserRepo(manage.Default(), 100)

	fmt.Println(m.New())
	fmt.Println(m.New().TableName())

	var item = &UserModel{}
	err := m.First(item)
	fmt.Println(err)

	manageRepo := repox.Manager{
		NewDB: func() *gorm.DB {
			return manage.Default()
		},
	}
	manageRepo.Register(reflect.TypeOf((*UserRepo)(nil)), func(db *gorm.DB) repox.Repo {
		return NewUserRepo(db, 100)
	})

	repo := manageRepo.Load(reflect.TypeOf((*UserRepo)(nil)))
	fmt.Println(repo)
	if repo != nil {
		fmt.Println(repo.RepoName(), repo.(UserRepo).New())
	}

	_ = repo.(repox.Transaction[*UserModel]).Transaction(func(repo repox.ReaderWriterRepo[*UserModel]) error {
		var data = &UserModel{}
		err := repo.First(data)
		if err != nil {
			fmt.Println("t( ", err, ")")
		}

		return nil
	})
}
