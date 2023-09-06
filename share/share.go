package share

import (
	"github.com/gin-gonic/gin"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/repox"
	"weicai.zhao.io/responsex"
)

var global *Manager

func init() {
	global, _ = New()
}

func Init(ops ...Option) (err error) {
	global, err = New(ops...)
	return
}

func RepoManager(usage string) *repox.Manager {
	return global.RepoManager(usage)
}

func GormManager() *gormx.Manager {
	return global.mysqlManager
}

func GinResponse(ctx *gin.Context) responsex.Response {
	return global.GinResponse(ctx)
}
