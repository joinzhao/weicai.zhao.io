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

func NewRepo[T schema.Tabler](tableName string, db *gorm.DB, batchSize int, f NewTabler[T]) ReaderWriterRepo[T] {
	return &templateRepo[T]{
		repoName:  tableName,
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
func (m *templateRepo[T]) Updates(item any, wheres ...gormx.Where) (int64, error) {
	db := gormx.Wheres(wheres).Build(m.db.Model(item).Table(m.New().TableName())).Updates(item)
	return db.RowsAffected, db.Error
}
func (m *templateRepo[T]) Update(item T, wheres ...gormx.Where) error {
	return gormx.Wheres(wheres).Build(m.db.Model(item)).Updates(item).Error
}
func (m *templateRepo[T]) Delete(wheres ...gormx.Where) (int64, error) {
	db := gormx.Wheres(wheres).Build(m.db.Model(m.New())).Delete(m.New())
	return db.RowsAffected, db.Error
}
func (m *templateRepo[T]) First(item T, wheres ...gormx.Where) error {
	return gormx.Wheres(wheres).Build(m.db.Model(item)).First(item).Error
}
func (m *templateRepo[T]) Find(limit int, offset int, wheres ...gormx.Where) (items []T, err error) {
	err = gormx.Wheres(wheres).Build(m.db.Model(items)).Limit(limit).Offset(offset).Find(items).Error
	return
}
func (m *templateRepo[T]) Count(wheres ...gormx.Where) (c int64, err error) {
	err = gormx.Wheres(wheres).Build(m.db.Model(m.New())).Count(&c).Error
	return
}

func (m *templateRepo[T]) Transaction(f func(repo ReaderWriterRepo[T]) error) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		tmp := &templateRepo[T]{
			repoName:  m.repoName,
			db:        tx,
			batchSize: m.batchSize,
			tableFunc: m.tableFunc,
		}
		return f(tmp)
	})
}
