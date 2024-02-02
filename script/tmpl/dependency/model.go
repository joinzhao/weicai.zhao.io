package dependency

import "weicai.zhao.io/script/tmpl/po"

type ModelRepo interface {
	GetBy(id int64) (item po.Model, err error)

	Find(item po.Model) (items []po.Model, err error)

	Cnt(item po.Model) (c int64, err error)

	Create(item *po.Model) error

	Update(id int64, item *po.Model) error

	Delete(id int64) error
}
