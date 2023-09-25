package share

import (
	"context"
	"gorm.io/gorm"
	"weicai.zhao.io/consts"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/repox"
)

func New(ops ...Option) (*Manager, error) {
	var m = &Manager{
		mode:        consts.DebugMode,
		repoManager: make(map[string]*repox.Manager),
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

func WithMysqlManagerBy(manager *gormx.Manager) Option {
	return func(m *Manager) error {
		m.mysqlManager = manager
		return nil
	}
}

func WithRepoManagerByConfigs(configs ...*gormx.Config) Option {
	return func(m *Manager) error {
		for _, config := range configs {
			m.repoManager[config.Usage] = repox.New(func() *gorm.DB {
				return m.mysqlManager.MustUseUsage(context.Background(), config.Usage)
			})
		}
		return nil
	}
}

func WithRepoManagerByUsages(usages ...string) Option {
	return func(m *Manager) error {
		for _, usage := range usages {
			m.repoManager[usage] = repox.New(func() *gorm.DB {
				return m.mysqlManager.MustUseUsage(context.Background(), usage)
			})
		}
		return nil
	}
}

// WithRepoManager repox.Manager 依赖 gormx.Manager，必须先执行 gormx.Manager 的初始化方法
func WithRepoManager(usage string) Option {
	return func(m *Manager) error {
		m.repoManager[usage] = repox.New(func() *gorm.DB {
			return m.mysqlManager.MustUseUsage(context.Background(), usage)
		})
		return nil
	}
}

func WithMode(mode consts.Mode) Option {
	return func(m *Manager) error {
		m.mode = mode
		return nil
	}
}
