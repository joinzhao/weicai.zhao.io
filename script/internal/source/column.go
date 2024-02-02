package source

import (
	"fmt"
	"gorm.io/gorm"
	"weicai.zhao.io/script/internal"
)

// New return result object implement internal.TableColumnSource
func New(db *gorm.DB) *tableColumnSource {
	return &tableColumnSource{
		db: db,
	}
}

// NewUnq return result object internal.TableColumnSource
func NewUnq(db *gorm.DB) *tableUnqColumnSource {
	return &tableUnqColumnSource{
		db: db,
	}
}

type tableColumnSource struct {
	db *gorm.DB
}

func (s *tableColumnSource) Find(dbName, table string) (columns []internal.TableColumn, err error) {
	columns = make([]internal.TableColumn, 0)
	var sql = fmt.Sprintf("SELECT `COLUMN_NAME`,`IS_NULLABLE`,`COLUMN_DEFAULT`,`DATA_TYPE`,`COLUMN_TYPE`,`COLUMN_KEY`,`COLUMN_COMMENT` "+
		"FROM `information_schema`.`COLUMNS` "+
		"WHERE `TABLE_SCHEMA` = '%s' AND `TABLE_NAME` = '%s' "+
		"ORDER BY `ORDINAL_POSITION` ASC;", dbName, table)

	return columns, s.db.Raw(sql).Scan(&columns).Error
}

type tableUnqColumnSource struct {
	db *gorm.DB
}

func (s *tableUnqColumnSource) Find(dbName, table string) (columns []internal.TableColumn, err error) {
	columns = make([]internal.TableColumn, 0)
	var sql = fmt.Sprintf("SELECT `COLUMN_NAME`,`IS_NULLABLE`,`COLUMN_DEFAULT`,`DATA_TYPE`,`COLUMN_TYPE`,`COLUMN_KEY`,`COLUMN_COMMENT` "+
		"FROM `information_schema`.`COLUMNS` "+
		"WHERE `TABLE_SCHEMA` = '%s' AND `TABLE_NAME` = '%s' AND COLUMN_KEY = '%s' "+
		"ORDER BY `ORDINAL_POSITION` ASC;", dbName, table, "UNI")

	return columns, s.db.Raw(sql).Scan(&columns).Error
}
