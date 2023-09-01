package share

import (
	"github.com/gin-gonic/gin"
	"weicai.zhao.io/consts"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/repox"
	"weicai.zhao.io/responsex"
)

type Manager struct {
	mode         consts.Mode
	mysqlManager *gormx.Manager
	repoManager  *repox.Manager
}

func (m *Manager) RepoManager() *repox.Manager {
	return m.repoManager
}
func (m *Manager) GormManager() *gormx.Manager {
	return m.mysqlManager
}

func (m *Manager) GinResponse(ctx *gin.Context) responsex.Response {
	return responsex.NewGinJsonResponse(m.mode, ctx)
}
