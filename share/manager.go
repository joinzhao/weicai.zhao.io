package share

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"weicai.zhao.io/consts"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/repox"
	"weicai.zhao.io/responsex"
)

type Manager struct {
	mode         consts.Mode
	mysqlManager *gormx.Manager
	repoManager  map[string]*repox.Manager
}

func (m *Manager) RepoManager(usage string) *repox.Manager {
	return m.repoManager[usage]
}
func (m *Manager) GormManager() *gormx.Manager {
	return m.mysqlManager
}

func (m *Manager) GinResponse(ctx *gin.Context) responsex.Response {
	return responsex.NewGinJsonResponse(m.mode, ctx)
}

func (m *Manager) NewWithRepo(target any) {
	// 初始化
	if target == nil {
		return
	}

	v := reflect.ValueOf(target)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		for _, manager := range m.repoManager {
			repo := manager.Load(v.Field(i).Type())
			if repo != nil {
				v.Field(i).Set(reflect.ValueOf(repo))
				// 找到之后退出循环
				break
			}
		}
	}
}
