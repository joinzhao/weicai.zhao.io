package option

import "gorm.io/gorm"

type Limit int

func (op Limit) apply(db *gorm.DB) *gorm.DB {
	return db.Limit(int(op))
}

type Offset int

func (op Offset) apply(db *gorm.DB) *gorm.DB {
	return db.Offset(int(op))
}
