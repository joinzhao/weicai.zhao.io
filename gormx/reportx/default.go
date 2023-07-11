package reportx

import (
	"gorm.io/gorm"
	"weicai.zhao.io/gormx/conditionx"
	"weicai.zhao.io/gormx/tablex"
)

type operatr struct {
	db *gorm.DB
}

func NewReport(db *gorm.DB) Report {
	return &operatr{db: db}
}

func NewReportReader(db *gorm.DB) ReportReader {
	return &operatr{db: db}
}

func NewReportWriter(db *gorm.DB) ReportWriter {
	return &operatr{db: db}
}

func (op *operatr) First(table tablex.Tabler, wheres ...conditionx.Where) error {
	db := op.db.Model(table)
	op.where(db, wheres...)
	return db.First(table).Error
}
func (op *operatr) Create(table tablex.Tabler) error {
	return op.db.Model(table).Create(table).Error
}

func (op *operatr) CreateInBatches(tablers ...tablex.Tabler) error {
	if len(tablers) > 0 {
		return op.db.Model(&tablers).CreateInBatches(&tablers, 500).Error

	}
	return nil
}

func (op *operatr) Update(table tablex.Tabler, wheres ...conditionx.Where) error {
	db := op.db.Model(table)
	op.where(db, wheres...)
	return db.Updates(table).Error
}

func (op *operatr) Delete(table tablex.Tabler, wheres ...conditionx.Where) error {
	db := op.db.Model(table)
	op.where(db, wheres...)
	return db.Delete(table).Error
}

func (op *operatr) where(db *gorm.DB, wheres ...conditionx.Where) {
	for _, where := range wheres {
		db = db.Where(where.Condition(), where.Args()...)
	}
}
