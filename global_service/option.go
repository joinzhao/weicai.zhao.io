package global_service

import (
	"io"
	"weicai.zhao.io/gormx"
)

type Option func(*manage) error

// GormOption 初始化 gorm 配置项
func GormOption(cfgs []*gormx.Config, ws ...io.Writer) Option {
	return func(m *manage) error {
		return m.RegisterMysqlManage(cfgs, ws...)
	}
}
