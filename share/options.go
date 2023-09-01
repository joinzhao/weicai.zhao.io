package share

import (
	"weicai.zhao.io/consts"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/repox"
)

func New(ops ...Option) (*Manager, error) {
	var m = &Manager{
		mode: consts.DebugMode,
	}
	for i := 0; i < len(ops); i++ {

		if err := ops[i](m); err != nil {
			return nil, err
		}
	}
	return m, nil
}

type Option func(m *Manager) error

func WithMysqlManager(configs ...*gormx.Config) Option {
	return func(m *Manager) error {
		m.mysqlManager = gormx.New(configs)
		return nil
	}
}

// WithRepoManager repox.Manager 依赖 gormx.Manager，必须先执行 gormx.Manager 的初始化方法
func WithRepoManager(f func(m *gormx.Manager) *repox.Manager) Option {
	return func(m *Manager) error {
		m.repoManager = f(m.mysqlManager)
		return nil
	}
}

func WithMode(mode consts.Mode) Option {
	return func(m *Manager) error {
		m.mode = mode
		return nil
	}
}
