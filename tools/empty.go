package tools

import (
	"reflect"
	"weicai.zhao.io/typex"
)

type CanEmpty interface {
	string | typex.Number
}

func IsEmpty[T CanEmpty](v T) bool {
	var t = reflect.TypeOf(v).Kind()
	var val = reflect.ValueOf(v)
	switch t {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return val.Int() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.String:
		return val.String() == ""
	}
	return false
}

func IsNil[T any](v []T) bool {
	return v == nil || len(v) == 0
}
