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

// ManageWithGorm 资源管理
type ManageWithGorm struct {
	pool  sync.Map
	NewDB NewDB
}

// Register 注册
func (m *ManageWithGorm) Register(t reflect.Type, f InitFuncWithGorm) {
	m.pool.Store(t, f)
}

// Load 加载
func (m *ManageWithGorm) Load(t reflect.Type) Repo {
	v, ok := m.pool.Load(t)
	if ok {
		return v.(InitFuncWithGorm)(m.NewDB())
	}
	return nil
}
