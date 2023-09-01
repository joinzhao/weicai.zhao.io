package repox

import (
	"gorm.io/gorm"
	"reflect"
	"sync"
)

// InitFuncWithGorm 初始化 gorm repo 接口方法
type InitFuncWithGorm func(db *gorm.DB) Repo

// NewDB 实例化 gorm db 对象
type NewDB func() *gorm.DB

// Manager 资源管理
type Manager struct {
	pool  sync.Map
	NewDB NewDB
}

func New(f NewDB) *Manager {
	return &Manager{
		pool:  sync.Map{},
		NewDB: f,
	}
}

// Register 注册
func (m *Manager) Register(t reflect.Type, f InitFuncWithGorm) {
	m.pool.Store(t, f)
}

// Load 加载
func (m *Manager) Load(t reflect.Type) Repo {
	v, ok := m.pool.Load(t)
	if ok {
		return v.(InitFuncWithGorm)(m.NewDB())
	}
	return nil
}
