package parse

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"strings"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/parse/go_file"
	"weicai.zhao.io/parse/sql_file"
	"weicai.zhao.io/tools"
)

func Generate(conf gormx.Config) {
	tables := getTables(conf)
	var _path = "../po"

	packageName := getPackageName(_path)

	for i := 0; i < len(tables); i++ {
		var f = go_file.File{
			Path:     _path,
			FileName: fmt.Sprintf("%s.go", tables[i].Name),
			Content:  go_file.Content{},
		}
		f.Content = tableToContent(tables[i])
		f.Content.PackageName = packageName

		f.Create()
	}
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getPackageName(_path string) string {
	log.Println("path: ", _path)
	list := strings.Split(_path, "/")
	if list != nil && len(list) > 0 {
		packageName := list[len(list)-1]
		switch packageName {
		case ".":
			return getPackageName(getCurrentDirectory())
		case "..":
			_path = getPackageName(getParentDirectory(getCurrentDirectory()))
		}
		return packageName
	}

	return getPackageName(getCurrentDirectory())
}

func tableToContent(table *sql_file.Table) go_file.Content {
	var sts = go_file.Struct{
		StructName: tools.UnderlineToUpperCamelCase(table.Name),
		Usage:      "m",
		Comment:    table.TableComment,
		Fields:     go_file.Fields{},
		Functions: go_file.Functions{
			tableNameFunc(table.Name),
		},
	}

	if table.Column != nil {
		var fields = go_file.Fields{}
		for i := 0; i < len(table.Column); i++ {
			fields = append(fields, columnToField(table.Column[i]))
		}
		sts.Fields = fields
	}

	cts := setStructFunction(table.Name, &sts)

	return go_file.Content{
		Consts:  cts,
		Structs: go_file.Structs{sts},
	}
}

func setStructFunction(tableName string, sts *go_file.Struct) go_file.Consts {
	var constName = fmt.Sprintf("%sTableName", sts.StructName)
	var cts = go_file.Consts{
		{
			Name:  constName,
			Value: fmt.Sprintf(`"%s"`, tableName),
		},
	}

	sts.Functions = go_file.Functions{
		{
			FunctionName: "TableName",
			Owner: go_file.Owner{
				Usage: sts.Usage,
				Name:  fmt.Sprintf("*%s", sts.StructName),
			},
			Returns: go_file.Returns{
				{
					Type: "string",
				},
			},
			Value: fmt.Sprintf(`    return %s`, constName),
		},
	}

	if sts.Fields != nil && len(sts.Fields) > 0 {
		for i := 0; i < len(sts.Fields); i++ {
			sts.Functions = append(sts.Functions, go_file.Function{
				Imports:      nil,
				FunctionName: fmt.Sprintf("Get%s", sts.Fields[i].FieldName),
				Params:       nil,
				Owner: go_file.Owner{
					Usage: sts.Usage,
					Name:  fmt.Sprintf("*%s", sts.StructName),
				},
				Returns: go_file.Returns{
					{
						Type: sts.Fields[i].Type,
					},
				},
				Value: fmt.Sprintf("    return %s.%s", sts.Usage, sts.Fields[i].FieldName),
			}, go_file.Function{
				Imports:      nil,
				FunctionName: fmt.Sprintf("Set%s", sts.Fields[i].FieldName),
				Params: go_file.Params{
					{
						Names: []string{"v"},
						Type:  sts.Fields[i].Type,
					},
				},
				Owner: go_file.Owner{
					Usage: sts.Usage,
					Name:  fmt.Sprintf("*%s", sts.StructName),
				},
				Returns: go_file.Returns{},
				Value:   fmt.Sprintf("    %s.%s = v", sts.Usage, sts.Fields[i].FieldName),
			})
		}
	}

	return cts
}

func tableNameFunc(tableName string) go_file.Function {
	return go_file.Function{
		FunctionName: "TableName",
		Owner: go_file.Owner{
			Usage: "m",
			Name:  fmt.Sprintf("*%s", tools.UnderlineToUpperCamelCase(tableName)),
		},
		Returns: go_file.Returns{
			{
				Type: "string",
			},
		},
		Value: fmt.Sprintf(`    return "%s"`, tableName),
	}
}

func columnToField(column *sql_file.Column) go_file.Field {
	var field = go_file.Field{
		Import:    go_file.Import{},
		FieldName: tools.UnderlineToUpperCamelCase(column.Column),
		Comment:   column.ColumnComment,
	}

	field.Type = sql_file.TransferSqlType(column.DataType)

	im := go_file.TransferType(field.Type)
	if im != "" {
		field.Import = go_file.Import{Value: im}
	}

	var gormTag = go_file.GormTag{
		Value: []string{
			fmt.Sprintf("column:%s", column.Column),
			fmt.Sprintf("default:%s", column.ColumnDefault),
			fmt.Sprintf("comment:%s", column.ColumnComment),
			fmt.Sprintf("type:%s", column.ColumnType),
		},
	}
	if column.IsNullable == sql_file.IsNullableNo {
		gormTag.Value = append(gormTag.Value, "not null")
	} else {
		gormTag.Value = append(gormTag.Value, "null")
	}
	if column.ColumnKey == "PRI" {
		gormTag.Value = append(gormTag.Value, "primaryKey")
	}

	field.Tags = go_file.Tags{
		go_file.JsonTag{
			Key:   "",
			Value: []string{column.Column},
		},
		gormTag,
	}

	return field
}

func getTables(conf gormx.Config) []*sql_file.Table {
	var databaseName = conf.Database
	conf.Database = "information_schema"
	manager := gormx.New([]*gormx.Config{&conf})

	var items = make([]*sql_file.Table, 0)
	err := manager.Default().
		Model(&items).
		Where("TABLE_SCHEMA = ?", databaseName).
		Preload("Column", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("*").Where("TABLE_SCHEMA = ?", databaseName).Order("ORDINAL_POSITION ASC")
		}).
		Find(&items).
		Error
	if err != nil {
		panic(err)
	}

	return items
}
