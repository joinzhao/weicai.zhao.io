package option

import "gorm.io/gorm"

type Option interface {
	apply(*gorm.DB) *gorm.DB
}

// Apply use ops from db
func Apply(db *gorm.DB, ops ...Option) *gorm.DB {
	for i := 0; i < len(ops); i++ {
		db = ops[i].apply(db)
	}
	return db
}
