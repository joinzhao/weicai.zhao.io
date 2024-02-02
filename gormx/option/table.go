package option

import "gorm.io/gorm"

type Table string

func (op Table) apply(db *gorm.DB) *gorm.DB {
	return db.Table(string(op))
}
