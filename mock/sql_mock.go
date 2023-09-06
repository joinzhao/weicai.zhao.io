package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type gormMock struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func (m *gormMock) Mock() sqlmock.Sqlmock {
	return m.mock
}
func (m *gormMock) DB() *gorm.DB {
	return m.db
}

func NewGormMock() *gormMock {
	sql, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal("init sql-mock fail: ", err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sql,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("init gorm connect fail: ", err)
	}

	return &gormMock{
		db:   db,
		mock: mock,
	}
}
