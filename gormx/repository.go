package gormx

import (
	"gorm.io/gorm"
	"weicai.zhao.io/gormx/option"
)

type Repository struct {
	DB        *gorm.DB
	BatchSize int
}

func (repo *Repository) First(result interface{}, ops ...option.Option) error {
	return option.Apply(repo.DB.Model(result), ops...).First(result).Error
}

func (repo *Repository) Find(result interface{}, ops ...option.Option) error {
	return option.Apply(repo.DB.Model(result), ops...).Find(result).Error
}

func (repo *Repository) Count(ops ...option.Option) (c int64, err error) {
	return c, option.Apply(repo.DB, ops...).Count(&c).Error
}

func (repo *Repository) Create(result interface{}) error {
	return repo.DB.Model(result).Create(result).Error
}

func (repo *Repository) CreateInBatch(result interface{}) error {
	if repo.BatchSize == 0 {
		repo.BatchSize = 500
	}
	return repo.DB.CreateInBatches(result, repo.BatchSize).Error
}

func (repo *Repository) Updates(result interface{}, ops ...option.Option) error {
	return option.Apply(repo.DB, ops...).Updates(result).Error
}

func (repo *Repository) UpdateColumn(column string, value interface{}, ops ...option.Option) error {
	return option.Apply(repo.DB, ops...).UpdateColumn(column, value).Error
}

func (repo *Repository) Delete(result interface{}, ops ...option.Option) error {
	return option.Apply(repo.DB, ops...).Delete(result).Error
}
