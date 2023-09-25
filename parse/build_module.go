package parse

import (
	"fmt"
	"path"
	"weicai.zhao.io/parse/go_file"
	"weicai.zhao.io/parse/sql_file"
	"weicai.zhao.io/tools"
)

type Module struct {
	Path string
	Name string
}

func (m *Module) Parse(tables []sql_file.Table) {
	var tableNameFile = go_file.File{
		Path:     path.Join(m.Path, baseDir, poPathDir),
		FileName: "global_table_name.go",
		Content: go_file.Content{
			Consts: go_file.Consts{},
		},
	}
	for i := 0; i < len(tables); i++ {
		m.buildStructFromTable(tables[i])
		tableNameFile.Content.Consts = append(tableNameFile.Content.Consts, go_file.Const{
			Name:  tools.UnderlineToUpperCamelCase(tables[i].Name) + "TableName",
			Value: fmt.Sprintf(`"%s"`, tables[i].Name),
		})
	}
	// create internal/module/po/global_table_name.go
	tableNameFile.Create()

}

// create internal/module/po/table_name.go
// create internal/module/po/table_name_get.go
// create internal/module/po/table_name_set.go
func (m *Module) buildStructFromTable(table sql_file.Table) {
	var (
		str = m.tableToStruct(table)
	)
	m.structTableNameFunc(&str)

	tableFile := go_file.File{
		Path:     path.Join(m.Path, baseDir, poPathDir),
		FileName: fmt.Sprintf("%s.go", table.Name),
		Content: go_file.Content{
			Structs: go_file.Structs{str},
		},
	}
	tableFile.Create()

	tableGetFile := go_file.File{
		Path:     path.Join(m.Path, baseDir, poPathDir),
		FileName: fmt.Sprintf("%s_get.go", table.Name),
		Content: go_file.Content{
			Functions: m.structGetFunc(str),
		},
	}
	tableGetFile.Create()

	tableSetFile := go_file.File{
		Path:     path.Join(m.Path, baseDir, poPathDir),
		FileName: fmt.Sprintf("%s_set.go", table.Name),
		Content: go_file.Content{
			Functions: m.structSetFunc(str),
		},
	}
	tableSetFile.Create()
}

func (m *Module) structGetFunc(str go_file.Struct) go_file.Functions {
	var res = go_file.Functions{}
	if str.Fields == nil || len(str.Fields) == 0 {
		return res
	}

	for i := 0; i < len(str.Fields); i++ {

		res = append(res, go_file.Function{
			Imports:      go_file.Imports{str.Fields[i].GetImport()},
			FunctionName: fmt.Sprintf("Get%s", str.Fields[i].FieldName),
			Params:       nil,
			Owner: go_file.Owner{
				Usage: str.Usage,
				Name:  fmt.Sprintf("*%s", str.StructName),
			},
			Returns: go_file.Returns{
				{
					Type: str.Fields[i].Type,
				},
			},
			Value: fmt.Sprintf("    return %s.%s", str.Usage, str.Fields[i].FieldName),
		})
	}
	return res
}

func (m *Module) structSetFunc(str go_file.Struct) go_file.Functions {
	var res = go_file.Functions{}
	if str.Fields == nil || len(str.Fields) == 0 {
		return res
	}

	for i := 0; i < len(str.Fields); i++ {
		res = append(res, go_file.Function{
			Imports:      go_file.Imports{str.Fields[i].GetImport()},
			FunctionName: fmt.Sprintf("Set%s", str.Fields[i].FieldName),
			Params: go_file.Params{
				{
					Names: []string{"v"},
					Type:  str.Fields[i].Type,
				},
			},
			Owner: go_file.Owner{
				Usage: str.Usage,
				Name:  fmt.Sprintf("*%s", str.StructName),
			},
			Returns: go_file.Returns{},
			Value:   fmt.Sprintf("    %s.%s = v", str.Usage, str.Fields[i].FieldName),
		})
	}
	return res
}

func (m *Module) structTableNameFunc(str *go_file.Struct) {
	str.Functions = append(str.Functions, go_file.Function{
		Imports:      nil,
		FunctionName: "TableName",
		Params:       nil,
		Owner: go_file.Owner{
			Usage: "m",
			Name:  fmt.Sprintf("*%s", str.StructName),
		},
		Returns: go_file.Returns{
			{
				Type: "string",
			},
		},
		Value: fmt.Sprintf(`    return %sTableName`, str.StructName),
	})
}

func (m *Module) tableToStruct(table sql_file.Table) go_file.Struct {
	var (
		str = go_file.Struct{
			StructName: tools.UnderlineToUpperCamelCase(table.Name),
			Usage:      "m",
			Comment:    table.TableComment,
		}
	)

	if table.Column != nil && len(table.Column) > 0 {
		str.Fields = go_file.Fields{}
		for i := 0; i < len(table.Column); i++ {
			str.Fields = append(str.Fields, m.columnToField(table.Column[i]))
		}
	}

	return str
}

func (m *Module) columnToField(column sql_file.Column) go_file.Field {
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

const (
	baseDir           = "internal"
	poPathDir         = "po"
	repoPathDir       = "repo"
	dependencyPathDir = "dependency"
)
