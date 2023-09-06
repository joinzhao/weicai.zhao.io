package repox

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"weicai.zhao.io/gormx"
)

type templateRepo[T schema.Tabler] struct {
	repoName  string
	db        *gorm.DB
	batchSize int
	tableFunc NewTabler[T]
}

func NewRepo[T schema.Tabler](repoName string, db *gorm.DB, batchSize int, f NewTabler[T]) ReaderWriterRepo[T] {
	return &templateRepo[T]{
		repoName:  repoName,
		db:        db,
		batchSize: batchSize,
		tableFunc: f,
	}
}

type NewTabler[T schema.Tabler] func() T

func (m *templateRepo[T]) New() T {
	return m.tableFunc()
}

func (m *templateRepo[T]) RepoName() string {
	return m.repoName
}

func (m *templateRepo[T]) Create(item T) error {
	return m.db.Model(item).Create(item).Error
}
func (m *templateRepo[T]) CreateInBatch(items []T) (int64, error) {
	db := m.db.Model(&items).CreateInBatches(&items, m.batchSize)
	return db.RowsAffected, db.Error
}
func (m *templateRepo[T]) Updates(item any, wheres ...gormx.Option) (int64, error) {
	db := gormx.Options(wheres).Build(m.db.Model(item).Table(m.New().TableName())).Updates(item)
	return db.RowsAffected, db.Error
}
func (m *templateRepo[T]) Update(item T, wheres ...gormx.Option) error {
	return gormx.Options(wheres).Build(m.db.Model(item)).Updates(item).Error
}
func (m *templateRepo[T]) Delete(wheres ...gormx.Option) (int64, error) {
	db := gormx.Options(wheres).Build(m.db.Model(m.New())).Delete(m.New())
	return db.RowsAffected, db.Error
}
func (m *templateRepo[T]) First(item T, wheres ...gormx.Option) error {
	return gormx.Options(wheres).Build(m.db.Model(item)).First(item).Error
}
func (m *templateRepo[T]) Find(wheres ...gormx.Option) (items []T, err error) {
	err = gormx.Options(wheres).Build(m.db.Model(items)).Find(items).Error
	return
}
func (m *templateRepo[T]) Count(wheres ...gormx.Option) (c int64, err error) {
	err = gormx.Options(wheres).Build(m.db.Model(m.New())).Count(&c).Error
	return
}
