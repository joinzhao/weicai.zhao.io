package gormx

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Option interface {
	Do(db *gorm.DB) *gorm.DB
}

type order struct {
	fields  []string
	orderBy string
}

func NewOrderOption(by string, fields ...string) Option {
	return &order{orderBy: by, fields: fields}
}

func (o *order) Do(db *gorm.DB) *gorm.DB {
	if len(o.fields) == 0 {
		return db
	}

	return db.Order(fmt.Sprintf("%s %s", strings.Join(o.fields, ","), o.orderBy))
}

type where struct {
	cond string
	args []interface{}
}

func NewWhereOption(cond string, args ...interface{}) Option {
	return &where{cond: cond, args: args}
}

func (o *where) Do(db *gorm.DB) *gorm.DB {
	return db.Where(o.cond, o.args...)
}

type Limit int

func (o Limit) Do(db *gorm.DB) *gorm.DB {
	return db.Limit(int(o))
}

type Offset int

func (o Offset) Do(db *gorm.DB) *gorm.DB {
	return db.Offset(int(o))
}

type Options []Option

func (w Options) Build(db *gorm.DB) *gorm.DB {
	if w != nil && len(w) > 0 {
		for i := 0; i < len(w); i++ {
			db = w[i].Do(db)
		}
	}
	return db
}

type PreloadString string

func (s PreloadString) Do(tx *gorm.DB) *gorm.DB {
	return tx.Preload(string(s))
}
