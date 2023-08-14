package global_service

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"weicai.zhao.io/repox"
)

var (
	_svc *manage
)

func Init(options ...Option) error {
	if _svc != nil {
		return fmt.Errorf("the variable has been initialized")
	}
	_svc = &manage{
		mysqlManage:    nil,
		gormAndRepo:    nil,
		gormRepoManage: nil,
	}

	for _, option := range options {
		err := option(_svc)
		if err != nil {
			return err
		}
	}

	return nil
}

//func New(options ...Option) (*manage, error) {
//	if err := Init(options...); err != nil {
//		return nil, err
//	}
//	return _svc, nil
//}

func RegisterGormRepo(t reflect.Type, f repox.InitFuncWithGorm, usage string) {
	_svc.RegisterGormRepo(t, f, usage)
}

func GormRepo(t reflect.Type) repox.Repo {
	return _svc.GormRepo(t)
}

func Mysql(ctx context.Context, usage string) (*gorm.DB, error) {
	return _svc.Mysql(ctx, usage)
}

func DefaultMysql() *gorm.DB {
	return _svc.DefaultMysql()
}

func MustUseMysql(ctx context.Context, usage string) *gorm.DB {
	return _svc.MustUseMysql(ctx, usage)
}

func InitWithRepo(target any) {
	_svc.InitWithRepo(target)
}
