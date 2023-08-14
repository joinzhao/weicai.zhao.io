package gormx

import "gorm.io/gorm"

type Where interface {
	Condition() string
	Args() []interface{}
}

type where struct {
	cond string
	args []interface{}
}

func NewWhere(cond string, args ...interface{}) Where {
	return &where{cond: cond, args: args}
}

func (w where) Condition() string {
	return w.cond
}

func (w where) Args() []interface{} {
	return w.args
}

func (w *where) SetCondition(cond string) {
	w.cond = cond
}

func (w *where) AddArg(v interface{}) {
	if w.args == nil {
		w.args = make([]interface{}, 0)
	}
	w.args = append(w.args, v)
}

type Wheres []Where

func (w Wheres) Build(db *gorm.DB) *gorm.DB {
	if w != nil && len(w) > 0 {
		for i := 0; i < len(w); i++ {
			db = db.Where(w[i].Condition(), w[i].Args()...)
		}
	}
	return db
}
