package option

import (
	"fmt"
	"gorm.io/gorm"
)

type sort string

const (
	AscSort  sort = "ASC"
	DescSort sort = "DESC"
)

type Order struct {
	Column string
	Sort   sort
}

func (op Order) apply(db *gorm.DB) *gorm.DB {
	return db.Order(fmt.Sprintf("%s %s", op.Column, op.Sort))
}
