package responsex

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weicai.zhao.io/consts"
	"weicai.zhao.io/errorx"
)

type ginResponse struct {
	mode consts.Mode
	ctx  *gin.Context
}

func NewGinJsonResponse(mode consts.Mode, ctx *gin.Context) Response {
	return &ginResponse{mode: mode, ctx: ctx}
}

func (r *ginResponse) Fail(err error) {
	var (
		resp = make(map[string]interface{})
	)

	switch err.(type) {
	case errorx.Error:
		resp["code"] = err.(errorx.Error).Code()
		resp["msg"] = err.(errorx.Error).Msg()
		if r.mode == consts.ProductionMode || r.mode == consts.ReleaseMode {
		} else {
			errs := err.(errorx.Error).Errors()
			debugs := make([]string, 0)
			if errs != nil && len(errs) > 0 {
				for i := 0; i < len(errs); i++ {
					debugs = append(debugs, errs[i].Error())
				}
			}
			resp["debug"] = debugs
		}
	default:
		resp["code"] = -1
		resp["msg"] = "unknown error"
	}

	r.ctx.JSON(http.StatusOK, resp)
}
func (r *ginResponse) Ok(v any) {
	var resp = make(map[string]interface{})
	resp["code"] = 0
	resp["msg"] = "successful"
	resp["data"] = v

	r.ctx.JSON(http.StatusOK, resp)
}
func (r *ginResponse) OkPage(v any, total int64) {
	var resp = make(map[string]interface{})
	resp["code"] = 0
	resp["msg"] = "successful"
	resp["data"] = v
	resp["total"] = total

	r.ctx.JSON(http.StatusOK, resp)
}
