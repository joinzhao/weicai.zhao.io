package main

import (
	"fmt"
	"gorm.io/gorm"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/repox"
	"weicai.zhao.io/share"
)

func Example() {
	var config = &gormx.Config{
		Usage:       "default",
		RunMode:     "debug",
		DSN:         "root:root@tcp(127.0.0.1:3307)",
		Database:    "worker_app",
		MaxIdleConn: 10,
		MaxOpenConn: 10,
		MaxLifeTime: 10,
	}

	manager, err := share.New(share.WithMysqlManager(config), share.WithRepoManager(func(m *gormx.Manager) *repox.Manager {
		return &repox.Manager{
			NewDB: func() *gorm.DB {
				return m.Default()
			},
		}
	}))
	if err != nil {
		panic(err)
	}

	fmt.Println(manager)
}
