package sql_file

const (
	IsNullableYes = "YES"
	IsNullableNo  = "NO"
)

type Column struct {
	Column        string `json:"columnName" gorm:"column:COLUMN_NAME;"`
	ColumnDefault string `json:"columnDefault" gorm:"column:COLUMN_DEFAULT;"`
	IsNullable    string `json:"isNullable" gorm:"column:IS_NULLABLE;"`
	Name          string `json:"name" gorm:"column:TABLE_NAME;"`
	DataType      string `json:"dataType" gorm:"column:DATA_TYPE;"`
	ColumnType    string `json:"columnType" gorm:"column:COLUMN_TYPE;"`
	ColumnKey     string `json:"columnKey" gorm:"column:COLUMN_KEY;"`
	ColumnComment string `json:"columnComment" gorm:"column:COLUMN_COMMENT;"`
	Length        string `json:"length" gorm:"column:CHARACTER_MAXIMUM_LENGTH;"`
}

func (c Column) TableName() string {
	return "COLUMNS"
}

var dataTypeRef = map[string]string{
	"bigint":     "int64",
	"varchar":    "string",
	"timestamp":  "time.Time",
	"int":        "int64",
	"longtext":   "string",
	"enum":       "string",
	"text":       "string",
	"mediumtext": "string",
	"json":       "[]byte",
	"datetime":   "time.Time",
	"set":        "[]byte",
	"binary":     "[]byte",
	"char":       "string",
	"varbinary":  "[]byte",
	"tinyint":    "uint8",
	"blob":       "string",
	"double":     "float64",
	"decimal":    "float64",
	"longblob":   "[]byte",
	"smallint":   "int64",
	"mediumblob": "[]byte",
	"time":       "time.Time",
	"float":      "float64",
}

func RegisterDataType(k, v string) {
	dataTypeRef[k] = v
}

func TransferSqlType(dataType string) string {
	return dataTypeRef[dataType]
}
