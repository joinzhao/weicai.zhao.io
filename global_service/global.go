package global_service

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"io"
	"reflect"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/repox"
)

type manage struct {
	mysqlManage    *gormx.Manager
	gormAndRepo    map[reflect.Type]string
	gormRepoManage map[string]*repox.ManageWithGorm
}

// RegisterGormRepo 注册 repo
func (m *manage) RegisterGormRepo(t reflect.Type, f repox.InitFuncWithGorm, usage string) {
	if m.gormAndRepo == nil {
		m.gormAndRepo = make(map[reflect.Type]string)
	}
	m.gormAndRepo[t] = usage

	if m.gormRepoManage == nil {
		m.gormRepoManage = make(map[string]*repox.ManageWithGorm)
	}

	if _, ok := m.gormRepoManage[usage]; !ok {
		repo := &repox.ManageWithGorm{
			NewDB: func() *gorm.DB {
				return m.mysqlManage.MustUseUsage(nil, usage)
			},
		}
		repo.Register(t, f)
		m.gormRepoManage[usage] = repo
	}
}

// GormRepo 获取已注册的 repo
func (m *manage) GormRepo(t reflect.Type) repox.Repo {
	if m.gormAndRepo == nil {
		fmt.Println(fmt.Errorf("this repo is not exists"))
		return nil
	}
	usage, ok := m.gormAndRepo[t]
	if !ok {
		fmt.Println(fmt.Errorf("this repo is not exists"))
		return nil
	}
	repoManage, ok := m.gormRepoManage[usage]
	if !ok {
		fmt.Println(fmt.Errorf("this repo is not exists"))
		return nil
	}

	repo := repoManage.Load(t)
	if repo == nil {
		fmt.Println(fmt.Errorf("this repo is not exists"))
		return nil
	}

	return repo
}

// RegisterMysqlManage 注册mysql资源，重复注册资源，将会实现资源覆盖
func (m *manage) RegisterMysqlManage(cfgs []*gormx.Config, ws ...io.Writer) error {
	if cfgs == nil || len(cfgs) == 0 {
		return fmt.Errorf("the database configuration item cannot be empty")
	}
	manage := gormx.New(cfgs)

	manage.SetWriters(ws...)

	for _, cfg := range cfgs {
		db, err := manage.Use(context.Background(), cfg.Usage)
		if err != nil {
			return err
		}
		sql, err := db.DB()
		if err != nil {
			return err
		}

		err = sql.Ping()
		if err != nil {
			return err
		}
	}

	m.mysqlManage = manage

	return nil
}

// Mysql mysql gorm object
func (m *manage) Mysql(ctx context.Context, usage string) (*gorm.DB, error) {
	return m.mysqlManage.Use(ctx, usage)
}

// DefaultMysql default mysql config
func (m *manage) DefaultMysql() *gorm.DB {
	return m.mysqlManage.Default()
}

// MustUseMysql the usage from config must exist
func (m *manage) MustUseMysql(ctx context.Context, usage string) *gorm.DB {
	return m.mysqlManage.MustUseUsage(ctx, usage)
}

// InitWithRepo 初始化目标参数， 并从 repo 池中取值
func (m *manage) InitWithRepo(target any) {
	// 初始化
	if target == nil {
		return
	}

	v := reflect.ValueOf(target)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		repo := m.GormRepo(t.Field(i).Type)
		if repo != nil {
			v.Field(i).Set(reflect.ValueOf(repo))
		}
	}
}
