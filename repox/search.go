package repox

import (
	"gorm.io/gorm/schema"
	"weicai.zhao.io/gormx"
)

type SearchRepo[T schema.Tabler] interface {
	Tabler[T]
	Find[T]
	Counter[T]
}

type SearchParam[T schema.Tabler] struct {
	Condition gormx.Options
	Repo      SearchRepo[T]
	Page      int // page num
	Size      int // page size
}

func SearchPage[T schema.Tabler](param SearchParam[T]) ([]T, int64, error) {
	if param.Condition == nil {
		param.Condition = make([]gormx.Option, 0)
	}
	// 查询总量
	total, err := param.Repo.Count(param.Condition...)
	if err != nil {
		return nil, 0, err
	}
	// 查询总量为0
	if total == 0 {
		return []T{}, 0, err
	}
	var (
		limit  = param.Size
		offset = param.Size * (param.Page - 1) // 起始位置
	)
	// default offset number
	if offset < 0 {
		offset = 0
	}

	// 过滤 limit 条件
	if limit <= 0 {
		return []T{}, total, nil
	}
	// 查询起始条件, 起始位置已超出总量
	if total < int64(offset) {
		return []T{}, total, nil
	}

	param.Condition = append(param.Condition, gormx.Limit(limit), gormx.Offset(offset))

	items, _ := param.Repo.Find(param.Condition...)

	return items, total, nil
}
