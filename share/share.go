package share

import (
	"github.com/gin-gonic/gin"
	"weicai.zhao.io/consts"
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

// ------------------------ MODE ----------------------
var mode consts.Mode = consts.DebugMode

func SetMode(m consts.Mode) {
	mode = m
}
func GinResponse(ctx *gin.Context) responsex.Response {
	return responsex.NewGinJsonResponse(mode, ctx)
}
