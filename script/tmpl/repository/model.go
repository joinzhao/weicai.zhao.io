package repository

import (
	"gorm.io/gorm"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/gormx/option"
	"weicai.zhao.io/script/tmpl/po"
)

type ModelRepo struct {
	repo *gormx.Repository
}

func NewModelRepo(db *gorm.DB, batchSize int) *ModelRepo {
	return &ModelRepo{
		repo: &gormx.Repository{
			DB:        db,
			BatchSize: batchSize,
		},
	}
}

func (repo *ModelRepo) GetBy(id int64) (item po.Model, err error) {
	return item, repo.repo.First(&item, option.Where{
		Column: "id",
		Expr:   option.Equal,
		Value:  []interface{}{id},
	})
}

func (repo *ModelRepo) Find(item po.Model) (items []po.Model, err error) {
	return items, repo.repo.Find(&items)
}

func (repo *ModelRepo) Cnt(item po.Model) (c int64, err error) {
	return repo.repo.Count(option.Table(item.TableName()))
}

func (repo *ModelRepo) Create(item *po.Model) error {
	return repo.repo.Create(item)
}

func (repo *ModelRepo) Update(id int64, item *po.Model) error {
	return repo.repo.Updates(item, option.Where{
		Column: "id",
		Expr:   option.Equal,
		Value:  []interface{}{id},
	})
}

func (repo *ModelRepo) Delete(id int64) error {
	return repo.repo.Delete(&po.Model{}, option.Where{
		Column: "id",
		Expr:   option.Equal,
		Value:  []interface{}{id},
	})
}
