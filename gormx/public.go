package gormx

import (
	"errors"
	"gorm.io/gorm"
)

func SessionWithNewDB(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true})
}

func IsRecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
