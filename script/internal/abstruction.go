package internal

// Cmd 命令
// GetName 命令名称
// Do 处理命令
type Cmd interface {
	GetName() string
	Do() error
}

type TableColumn struct {
	COLUMN_NAME    string
	IS_NULLABLE    string
	COLUMN_DEFAULT int
	DATA_TYPE      string
	COLUMN_TYPE    string
	COLUMN_KEY     string
	COLUMN_COMMENT string
}

type TableColumnSource interface {
	Find(dbName, table string) ([]TableColumn, error)
}
