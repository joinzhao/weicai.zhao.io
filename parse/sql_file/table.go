package sql_file

type Table struct {
	Name         string   `json:"name" gorm:"column:TABLE_NAME;"`
	TableComment string   `json:"tableComment" gorm:"column:TABLE_COMMENT;"`
	Column       []Column `json:"fields" gorm:"foreignKey:Name;references:Name;"`
}

func (t *Table) TableName() string {
	return "TABLES"
}
