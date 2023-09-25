package validatorx

import (
	"github.com/go-playground/validator/v10"
	"weicai.zhao.io/errorx"
)

func New() *validator.Validate {
	return validator.New()
}

func Struct(v interface{}) errorx.Error {
	if err := New().Struct(v); err != nil {
		return errorx.ParamVerifyError.WithError(err)
	}
	return nil
}
